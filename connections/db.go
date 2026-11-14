package connections

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"gorm.io/plugin/dbresolver"
)

// DBOption customizes the GORM connection.
type DBOption func(*dbOptions)

type dbOptions struct {
	debug           bool
	singularTable   bool
	defaultTimezone string
}

// WithDebug enables SQL tracing, including dbresolver routing.
func WithDebug(debug bool) DBOption {
	return func(o *dbOptions) { o.debug = debug }
}

// WithSingularTable uses singular table names and applies info.Prefix through
// GORM's naming strategy. Callers that already prefix names in TableName() must
// not enable it.
func WithSingularTable(singular bool) DBOption {
	return func(o *dbOptions) { o.singularTable = singular }
}

// WithDefaultTimezone sets the DSN timezone, info.Loc still wins.
func WithDefaultTimezone(tz string) DBOption {
	return func(o *dbOptions) { o.defaultTimezone = tz }
}

// NewDatabaseConnection opens the primary connection, wiring replicas and the
// priority-based failover on top when they are configured.
func NewDatabaseConnection(ctx context.Context, info *DatabaseInfo, ll *logrus.Logger, opts ...DBOption) (*gorm.DB, error) {
	if ll == nil {
		ll = logrus.StandardLogger()
	}
	log := ll.WithField("method", "NewDatabaseConnection")

	o := &dbOptions{defaultTimezone: "UTC"}
	for _, opt := range opts {
		opt(o)
	}

	charset := "utf8mb4"
	if info.Charset != nil && *info.Charset != "" {
		charset = *info.Charset
	}
	timezone := o.defaultTimezone
	if info.Loc != nil && *info.Loc != "" {
		timezone = *info.Loc
	}
	loc := strings.ReplaceAll(timezone, "/", "%2F")

	connectTimeout := defaultConnectTimeout
	if info.ConnectTimeout != nil && *info.ConnectTimeout > 0 {
		connectTimeout = *info.ConnectTimeout
	}

	connMaxLifetime := 4 * time.Minute
	if info.ConnMaxLifetime != nil && *info.ConnMaxLifetime > 0 {
		connMaxLifetime = *info.ConnMaxLifetime
	}
	maxOpenConns := 10
	if info.MaxOpenConns != nil && *info.MaxOpenConns > 0 {
		maxOpenConns = *info.MaxOpenConns
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=True&loc=%s&timeout=%s%s",
		info.Username, info.Password, info.Host, info.Port, info.DBName, charset, loc, connectTimeout, timeoutParams(info))

	logLevel := logger.Warn
	if o.debug {
		logLevel = logger.Info
		if lvl := logger.LogLevel(ll.GetLevel()); lvl > logLevel {
			logLevel = lvl
		}
	}

	cnf := &gorm.Config{
		Logger: logger.New(ll, logger.Config{
			SlowThreshold:             time.Second,
			LogLevel:                  logLevel,
			IgnoreRecordNotFoundError: true,
			ParameterizedQueries:      false,
			Colorful:                  true,
		}),
	}
	if ns := namingStrategy(o, info); ns != nil {
		cnf.NamingStrategy = *ns
	}

	var db *gorm.DB
	var err error
	var failover *failoverConnector
	var sqlDB *sql.DB

	// Nothing is returned on failure, so the monitor and the pooled connection
	// must be released to not leak the health check goroutine.
	ok := false
	defer func() {
		if ok {
			return
		}
		if failover != nil {
			_ = failover.Close()
		}
		if sqlDB != nil {
			_ = sqlDB.Close()
		}
	}()

	if info.Failover != nil && len(info.Failover.Hosts) > 0 {
		log.Infof("found %d failover hosts, using priority-based DB failover", len(info.Failover.Hosts))

		connector, cerr := newFailoverConnector(info, info.Failover, dsn, connectTimeout, ll)
		if cerr != nil {
			return nil, cerr
		}
		failover, _ = connector.(*failoverConnector)
		sqlDB = sql.OpenDB(connector)
		db, err = gorm.Open(mysql.New(mysql.Config{Conn: sqlDB}), cnf)
	} else {
		db, err = gorm.Open(mysql.New(mysql.Config{DSN: dsn}), cnf)
	}
	if err != nil {
		return nil, err
	}

	if failover != nil {
		if interval := healthCheckInterval(info.Failover); interval > 0 {
			log.Infof("starting DB health monitor every %s", interval)
			failover.start(interval)
		} else {
			log.Info("DB health monitor disabled")
		}
	}

	if len(info.Replicas) > 0 {
		log.Infof("found %d read replicas, configuring dbresolver", len(info.Replicas))
		replicaDialectors := make([]gorm.Dialector, 0, len(info.Replicas))
		for _, r := range info.Replicas {
			// inherit the primary settings that are not set per replica
			if r.Username == "" {
				r.Username = info.Username
			}
			if r.Password == "" {
				r.Password = info.Password
			}
			if r.Port == 0 {
				r.Port = info.Port
			}

			replicaDsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=True&loc=%s&timeout=%s%s",
				r.Username, r.Password, r.Host, r.Port, info.DBName, charset, loc, connectTimeout, timeoutParams(info))
			replicaDialectors = append(replicaDialectors, mysql.Open(replicaDsn))
		}

		resolverCnf := dbresolver.Config{
			Replicas: replicaDialectors,
			Policy:   dbresolver.RandomPolicy{}, // distribute read load
		}
		if o.debug {
			resolverCnf.TraceResolverMode = true
		}

		err = db.Use(dbresolver.Register(resolverCnf).
			SetConnMaxLifetime(connMaxLifetime).
			SetMaxOpenConns(maxOpenConns).
			SetMaxIdleConns(maxOpenConns))
		if err != nil {
			return nil, err
		}
	}

	d, err := db.DB()
	if err != nil {
		return nil, err
	}

	// https://github.com/go-sql-driver/mysql?tab=readme-ov-file#important-settings
	d.SetConnMaxLifetime(connMaxLifetime)
	d.SetMaxOpenConns(maxOpenConns)
	d.SetMaxIdleConns(maxOpenConns)

	if err := d.PingContext(ctx); err != nil {
		return nil, err
	}

	var version string
	if err := d.QueryRowContext(ctx, "SELECT VERSION()").Scan(&version); err != nil {
		log.WithError(err).Debug("failed to read database version")
	}
	log.WithField("version", version).Info("Successfully connected to database")

	ok = true
	return db, nil
}

// namingStrategy returns nil unless singular tables are requested, so callers
// that prefix table names in their own TableName() methods keep GORM's default
// behaviour and do not apply the prefix twice.
func namingStrategy(o *dbOptions, info *DatabaseInfo) *schema.NamingStrategy {
	if !o.singularTable {
		return nil
	}
	return &schema.NamingStrategy{TablePrefix: info.Prefix, SingularTable: true}
}

// timeoutParams returns the optional DSN read/write timeout parameters.
func timeoutParams(info *DatabaseInfo) string {
	var params string
	if info.ReadTimeout != nil && *info.ReadTimeout > 0 {
		params += fmt.Sprintf("&readTimeout=%s", info.ReadTimeout.String())
	}
	if info.WriteTimeout != nil && *info.WriteTimeout > 0 {
		params += fmt.Sprintf("&writeTimeout=%s", info.WriteTimeout.String())
	}
	return params
}

// failoverCooldown returns how long a failed host is skipped, nil means the default.
func failoverCooldown(f *FailoverSettings) time.Duration {
	if f.Cooldown == nil || *f.Cooldown <= 0 {
		return defaultFailoverCooldown
	}
	return *f.Cooldown
}

// healthCheckInterval returns the monitor interval, nil means the default.
func healthCheckInterval(f *FailoverSettings) time.Duration {
	if f.HealthCheckInterval == nil {
		return defaultHealthCheckInterval
	}
	return *f.HealthCheckInterval
}

// failbackEnabled returns whether traffic may return to a higher-priority host.
func failbackEnabled(f *FailoverSettings) bool {
	return f.Failback == nil || *f.Failback
}

// failbackDelay returns how long a higher-priority host must stay healthy, nil or
// a non-positive value means the default.
func failbackDelay(f *FailoverSettings) time.Duration {
	if f.FailbackDelay == nil || *f.FailbackDelay <= 0 {
		return defaultFailbackDelay
	}
	return *f.FailbackDelay
}
