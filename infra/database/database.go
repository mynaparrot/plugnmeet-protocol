package factory

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/plugin/dbresolver"
)

type DatabaseInfo struct {
	DriverName      string          `yaml:"driver_name"`
	Host            string          `yaml:"host"`
	Port            int32           `yaml:"port"`
	Username        string          `yaml:"username"`
	Password        string          `yaml:"password"`
	DBName          string          `yaml:"db"`
	Prefix          string          `yaml:"prefix"`
	Charset         *string         `yaml:"charset"`
	Loc             *string         `yaml:"loc"`
	ConnMaxLifetime *time.Duration  `yaml:"conn_max_lifetime"`
	MaxOpenConns    *int            `yaml:"max_open_conns"`
	Replicas        []ReplicaDBInfo `yaml:"replicas"`
}

// ReplicaDBInfo holds connection details for a read replica database.
type ReplicaDBInfo struct {
	Host     string `yaml:"host"`
	Port     int32  `yaml:"port"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
}

func NewDatabaseConnection(ctx context.Context, info *DatabaseInfo, debug bool, log *logrus.Logger) (*gorm.DB, error) {
	charset := "utf8mb4"
	loc := "UTC"
	connMaxLifetime := time.Minute * 4
	maxOpenConns := 10

	if info.Charset != nil && *info.Charset != "" {
		charset = *info.Charset
	}
	if info.Loc != nil && *info.Loc != "" {
		loc = strings.ReplaceAll(*info.Loc, "/", "%2F")
	}
	if info.ConnMaxLifetime != nil && *info.ConnMaxLifetime > 0 {
		connMaxLifetime = *info.ConnMaxLifetime
	}
	if info.MaxOpenConns != nil && *info.MaxOpenConns > 0 {
		maxOpenConns = *info.MaxOpenConns
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=True&loc=%s", info.Username, info.Password, info.Host, info.Port, info.DBName, charset, loc)

	mysqlCnf := mysql.Config{
		DSN: dsn, // data source name
	}
	cnf := &gorm.Config{}

	loggerCnf := logger.Config{
		SlowThreshold:             time.Second, // Slow SQL threshold
		LogLevel:                  logger.Info,
		IgnoreRecordNotFoundError: true,
		ParameterizedQueries:      false,
		Colorful:                  true,
	}

	if !debug {
		loggerCnf.LogLevel = logger.Warn
		cnf.Logger = logger.New(log, loggerCnf)
	} else {
		cnf.Logger = logger.New(log, loggerCnf)
	}

	db, err := gorm.Open(mysql.New(mysqlCnf), cnf)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	// If read replicas are configured, set up the dbresolver.
	if len(info.Replicas) > 0 {
		log.Infof("found %d read replicas, configuring dbresolver", len(info.Replicas))
		var replicaDialectors []gorm.Dialector

		for _, r := range info.Replicas {
			// Use primary's settings as default for replicas if not specified.
			if r.Username == "" {
				r.Username = info.Username
			}
			if r.Password == "" {
				r.Password = info.Password
			}
			if r.Port == 0 {
				r.Port = info.Port
			}

			replicaDsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=True&loc=%s", r.Username, r.Password, r.Host, r.Port, info.DBName, charset, loc)
			replicaDialectors = append(replicaDialectors, mysql.Open(replicaDsn))
		}
		resolverCnf := dbresolver.Config{
			Replicas: replicaDialectors,
			Policy:   dbresolver.RandomPolicy{}, // Use random policy to distribute read load.
		}
		if debug {
			resolverCnf.TraceResolverMode = true
		}

		err = db.Use(dbresolver.Register(resolverCnf).
			SetConnMaxLifetime(connMaxLifetime).
			SetMaxOpenConns(maxOpenConns).
			SetMaxIdleConns(maxOpenConns))
		if err != nil {
			return nil, fmt.Errorf("failed to configure dbresolver: %w", err)
		}
	}

	d, err := db.DB()
	if err != nil {
		return nil, fmt.Errorf("failed to get database connection: %w", err)
	}

	// https://github.com/go-sql-driver/mysql?tab=readme-ov-file#important-settings
	d.SetConnMaxLifetime(connMaxLifetime)
	d.SetMaxOpenConns(maxOpenConns)
	d.SetMaxIdleConns(maxOpenConns)

	err = d.PingContext(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	dbVersion := ""
	db.Raw("SELECT VERSION()").Scan(&dbVersion)
	log.WithField("version", dbVersion).Info("successfully connected to database")

	return db, nil
}
