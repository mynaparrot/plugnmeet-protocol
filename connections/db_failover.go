package connections

import (
	"context"
	"database/sql/driver"
	"errors"
	"fmt"
	"io"
	"net"
	"strconv"
	"strings"
	"sync"
	"sync/atomic"
	"time"

	"github.com/go-sql-driver/mysql"
	log "github.com/sirupsen/logrus"
)

const (
	defaultConnectTimeout      = 3 * time.Second
	defaultFailoverCooldown    = 30 * time.Second
	defaultFailbackDelay       = 60 * time.Second
	defaultHealthCheckInterval = 5 * time.Second

	galeraStatusQuery = "SHOW GLOBAL STATUS WHERE Variable_name IN ('wsrep_ready','wsrep_local_state')"
)

var (
	_ io.Closer                 = (*failoverConnector)(nil)
	_ driver.Conn               = (*failoverConn)(nil)
	_ driver.Validator          = (*failoverConn)(nil)
	_ driver.Pinger             = (*failoverConn)(nil)
	_ driver.Execer             = (*failoverConn)(nil)
	_ driver.ExecerContext      = (*failoverConn)(nil)
	_ driver.Queryer            = (*failoverConn)(nil)
	_ driver.QueryerContext     = (*failoverConn)(nil)
	_ driver.ConnPrepareContext = (*failoverConn)(nil)
	_ driver.ConnBeginTx        = (*failoverConn)(nil)
	_ driver.SessionResetter    = (*failoverConn)(nil)
	_ driver.NamedValueChecker  = (*failoverConn)(nil)
)

// failoverHost is one node of the priority list (index 0 is the primary).
type failoverHost struct {
	name         string
	connector    driver.Connector
	downUntil    time.Time
	down         bool
	healthySince time.Time
}

// failoverConnector routes all connections to one active host, chosen by priority.
// A failed host is skipped for the cooldown; failover to a lower-priority host is
// immediate, while the background monitor only fails back to a higher-priority host
// once it stayed healthy for failbackDelay.
// On a switch, pooled connections to the old host are dropped when returned or reused,
// but in-flight transactions finish on their host, so a conflicting write there may
// fail with a Galera certification/deadlock error during that short window.
type failoverConnector struct {
	hosts          []*failoverHost
	cooldown       time.Duration
	failback       bool
	failbackDelay  time.Duration
	attemptTimeout time.Duration
	now            func() time.Time
	healthCheck    func(ctx context.Context, conn driver.Conn) error
	logger         *log.Logger
	active         atomic.Int32

	mu      sync.Mutex // guards hosts' down/downUntil/healthySince, allDown and stop
	allDown bool
	stop    chan struct{}
	closed  sync.Once
	wg      sync.WaitGroup
}

// newFailoverConnector builds a connector for the primary plus the failover hosts.
func newFailoverConnector(info *DatabaseInfo, failover *FailoverSettings, primaryDSN string, timeout time.Duration, ll *log.Logger) (driver.Connector, error) {
	base, err := mysql.ParseDSN(primaryDSN)
	if err != nil {
		return nil, err
	}

	newHost := func(host string, port int32, username, password string) (*failoverHost, error) {
		cnf := base.Clone()
		cnf.Addr = net.JoinHostPort(host, strconv.Itoa(int(port)))
		cnf.Timeout = timeout
		if username != "" {
			cnf.User = username
		}
		if password != "" {
			cnf.Passwd = password
		}
		connector, err := mysql.NewConnector(cnf)
		if err != nil {
			return nil, err
		}
		return &failoverHost{name: cnf.Addr, connector: connector}, nil
	}

	hosts := make([]*failoverHost, 0, len(failover.Hosts)+1)
	primary, err := newHost(info.Host, info.Port, "", "")
	if err != nil {
		return nil, err
	}
	hosts = append(hosts, primary)

	for _, f := range failover.Hosts {
		if f.Port == 0 {
			f.Port = info.Port
		}
		host, err := newHost(f.Host, f.Port, f.Username, f.Password)
		if err != nil {
			return nil, err
		}
		hosts = append(hosts, host)
	}

	healthCheck := noHealthCheck
	if failover.GaleraCheck == nil || *failover.GaleraCheck {
		healthCheck = galeraHealthCheck
	}

	c := newFailoverConnectorWithHosts(hosts, failoverCooldown(failover), healthCheck)
	if ll != nil {
		c.logger = ll
	}
	c.failback = failbackEnabled(failover)
	c.failbackDelay = failbackDelay(failover)
	c.attemptTimeout = 2 * timeout
	return c, nil
}

// newFailoverConnectorWithHosts builds a connector around an existing host list.
func newFailoverConnectorWithHosts(hosts []*failoverHost, cooldown time.Duration, healthCheck func(context.Context, driver.Conn) error) *failoverConnector {
	c := &failoverConnector{
		hosts:         hosts,
		cooldown:      cooldown,
		failback:      true,
		failbackDelay: defaultFailbackDelay,
		now:           time.Now,
		healthCheck:   healthCheck,
		logger:        log.StandardLogger(),
	}
	c.active.Store(-1)
	return c
}

func noHealthCheck(context.Context, driver.Conn) error { return nil }

func (c *failoverConnector) Connect(ctx context.Context) (driver.Conn, error) {
	conn, idx, skipped, err := c.connectPass(ctx, true)
	if err == nil {
		c.setAllDown(false, nil)
		return c.track(conn, idx), nil
	}
	if !skipped {
		c.setAllDown(true, err)
		return nil, err
	}

	// every host is cooling down, so ignore the cooldown as a last resort
	conn, idx, _, err = c.connectPass(ctx, false)
	if err != nil {
		c.setAllDown(true, err)
		return nil, err
	}
	c.setAllDown(false, nil)
	return c.track(conn, idx), nil
}

func (c *failoverConnector) Driver() driver.Driver {
	return &mysql.MySQLDriver{}
}

// connectPass uses the active host when it is healthy and otherwise takes the
// first healthy host in priority order, so failover away is immediate while the
// active host never moves down to a higher priority one.
func (c *failoverConnector) connectPass(ctx context.Context, respectCooldown bool) (driver.Conn, int, bool, error) {
	var lastErr error
	skipped := false
	tried := -1

	if active := int(c.active.Load()); active >= 0 {
		if respectCooldown && c.isDown(active) {
			skipped = true
		} else {
			tried = active
			conn, err := c.attempt(ctx, active)
			if err == nil {
				c.activate(active)
				return conn, active, skipped, nil
			}
			lastErr = err
			c.markDown(active, err)
		}
	}

	for idx := range c.hosts {
		if idx == tried {
			continue
		}
		if respectCooldown && c.isDown(idx) {
			skipped = true
			continue
		}

		conn, err := c.attempt(ctx, idx)
		if err == nil {
			c.activate(idx)
			return conn, idx, skipped, nil
		}

		lastErr = err
		c.markDown(idx, err)
	}

	if lastErr == nil {
		lastErr = errors.New("no database host available")
	}
	return nil, -1, skipped, lastErr
}

// attempt connects to one host and health checks it, bounded by attemptTimeout.
func (c *failoverConnector) attempt(ctx context.Context, idx int) (driver.Conn, error) {
	attemptCtx := ctx
	if c.attemptTimeout > 0 {
		var cancel context.CancelFunc
		attemptCtx, cancel = context.WithTimeout(ctx, c.attemptTimeout)
		defer cancel()
	}

	conn, err := c.hosts[idx].connector.Connect(attemptCtx)
	if err != nil {
		return nil, fmt.Errorf("connect: %w", err)
	}
	if err := c.healthCheck(attemptCtx, conn); err != nil {
		conn.Close()
		return nil, fmt.Errorf("health check: %w", err)
	}
	return conn, nil
}

// checkOnce probes the hosts, switching away from a dead active host and, when
// failback is enabled, back to a higher-priority host that stays healthy.
func (c *failoverConnector) checkOnce(ctx context.Context) {
	conn, _, _, err := c.connectPass(ctx, true)
	if err != nil {
		c.setAllDown(true, err)
		return
	}
	c.setAllDown(false, nil)
	if err := conn.Close(); err != nil {
		c.logger.WithError(err).Debug("closing health check connection failed")
	}
	c.failbackPass(ctx)
}

// failbackPass probes higher-priority hosts up to the first healthy one and
// switches once it has been healthy for failbackDelay.
func (c *failoverConnector) failbackPass(ctx context.Context) {
	active := int(c.active.Load())
	if !c.failback || active <= 0 {
		return
	}

	for idx := 0; idx < active; idx++ {
		if c.isDown(idx) {
			continue
		}

		conn, err := c.attempt(ctx, idx)
		if err != nil {
			c.markDown(idx, err) // also resets healthySince
			continue
		}
		if err := conn.Close(); err != nil {
			c.logger.WithError(err).Debug("closing health check connection failed")
		}
		c.noteHealthy(idx)
		return
	}
}

// noteHealthy starts or continues the failback wait of a healthy host.
func (c *failoverConnector) noteHealthy(idx int) {
	c.mu.Lock()
	first := c.hosts[idx].healthySince.IsZero()
	if first {
		c.hosts[idx].healthySince = c.now()
	}
	since := c.hosts[idx].healthySince
	name := c.hosts[idx].name
	c.mu.Unlock()

	if first {
		c.logger.Infof("database host %s is healthy, failing back after %s", name, c.failbackDelay)
	}

	if waited := c.now().Sub(since); waited < c.failbackDelay {
		c.logger.Debugf("host %s healthy, failback in %s", name, c.failbackDelay-waited)
		return
	}
	c.activate(idx)
}

// setAllDown records whether no host is usable, logging only on transitions.
func (c *failoverConnector) setAllDown(down bool, err error) {
	c.mu.Lock()
	changed := c.allDown != down
	c.allDown = down
	c.mu.Unlock()

	switch {
	case changed && down:
		c.logger.WithError(err).Warn("no healthy database host available")
	case changed:
		c.logger.Info("database host available again")
	case down:
		c.logger.WithError(err).Debug("no healthy database host available")
	}
}

// start launches the background health monitor, 0 disables it.
func (c *failoverConnector) start(interval time.Duration) {
	if interval <= 0 {
		return
	}

	c.mu.Lock()
	if c.stop != nil {
		c.mu.Unlock()
		return
	}
	stop := make(chan struct{})
	c.stop = stop
	c.mu.Unlock()

	c.wg.Add(1)
	go func() {
		defer c.wg.Done()

		ticker := time.NewTicker(interval)
		defer ticker.Stop()
		for {
			select {
			case <-ticker.C:
				c.checkOnce(context.Background())
			case <-stop:
				return
			}
		}
	}()
}

// Close stops the background health monitor; it is safe to call more than once.
func (c *failoverConnector) Close() error {
	c.closed.Do(func() {
		c.mu.Lock()
		if c.stop != nil {
			close(c.stop)
		}
		c.mu.Unlock()
		c.wg.Wait()
	})
	return nil
}

func (c *failoverConnector) track(conn driver.Conn, idx int) driver.Conn {
	return &failoverConn{Conn: conn, connector: c, host: idx}
}

func (c *failoverConnector) activate(idx int) {
	c.mu.Lock()
	recovered := c.hosts[idx].down
	c.hosts[idx].down = false
	c.hosts[idx].downUntil = time.Time{}
	c.hosts[idx].healthySince = time.Time{}
	c.mu.Unlock()

	if recovered {
		c.logger.Infof("database host %s recovered", c.hosts[idx].name)
	}

	prev := int(c.active.Swap(int32(idx)))
	switch {
	case prev == -1:
		c.logger.Infof("connected to DB %s", c.hosts[idx].name)
	case prev != idx:
		c.logger.Warnf("switched DB from %s to %s", c.hosts[prev].name, c.hosts[idx].name)
	}
}

func (c *failoverConnector) markDown(idx int, err error) {
	c.mu.Lock()
	first := !c.hosts[idx].down
	c.hosts[idx].down = true
	c.hosts[idx].downUntil = c.now().Add(c.cooldown)
	c.hosts[idx].healthySince = time.Time{}
	name := c.hosts[idx].name
	c.mu.Unlock()

	if first {
		c.logger.WithError(err).Warnf("database host %s failed, skipping it for %s", name, c.cooldown)
		return
	}
	c.logger.WithError(err).Debugf("database host %s failed, skipping it for %s", name, c.cooldown)
}

func (c *failoverConnector) isDown(idx int) bool {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.now().Before(c.hosts[idx].downUntil)
}

// galeraHealthCheck returns an error when the node is not a synced Galera member.
func galeraHealthCheck(ctx context.Context, conn driver.Conn) error {
	queryer, ok := conn.(driver.QueryerContext)
	if !ok {
		return nil
	}

	rows, err := queryer.QueryContext(ctx, galeraStatusQuery, nil)
	if err != nil {
		return err
	}
	defer rows.Close()

	var ready, state string
	for {
		values := make([]driver.Value, 2)
		if err := rows.Next(values); err != nil {
			if errors.Is(err, io.EOF) {
				break
			}
			return err
		}
		switch strings.ToLower(valueString(values[0])) {
		case "wsrep_ready":
			ready = valueString(values[1])
		case "wsrep_local_state":
			state = valueString(values[1])
		}
	}

	if ready == "" || state == "" {
		return errors.New("missing Galera wsrep status")
	}
	if !strings.EqualFold(ready, "ON") {
		return fmt.Errorf("wsrep_ready is %q", ready)
	}
	if state != "4" {
		return fmt.Errorf("wsrep_local_state is %q", state)
	}
	return nil
}

func valueString(v driver.Value) string {
	switch s := v.(type) {
	case string:
		return s
	case []byte:
		return string(s)
	case nil:
		return ""
	}
	return fmt.Sprint(v)
}

// failoverConn remembers its host, and is invalid once another host is active.
type failoverConn struct {
	driver.Conn
	connector *failoverConnector
	host      int
}

func (c *failoverConn) IsValid() bool {
	if int(c.connector.active.Load()) != c.host {
		return false
	}
	if v, ok := c.Conn.(driver.Validator); ok {
		return v.IsValid()
	}
	return true
}

func (c *failoverConn) Ping(ctx context.Context) error {
	if int(c.connector.active.Load()) != c.host {
		return driver.ErrBadConn
	}
	if p, ok := c.Conn.(driver.Pinger); ok {
		return p.Ping(ctx)
	}
	return driver.ErrSkip
}

func (c *failoverConn) Exec(query string, args []driver.Value) (driver.Result, error) {
	if e, ok := c.Conn.(driver.Execer); ok {
		return e.Exec(query, args)
	}
	return nil, driver.ErrSkip
}

func (c *failoverConn) ExecContext(ctx context.Context, query string, args []driver.NamedValue) (driver.Result, error) {
	if e, ok := c.Conn.(driver.ExecerContext); ok {
		return e.ExecContext(ctx, query, args)
	}
	return nil, driver.ErrSkip
}

func (c *failoverConn) Query(query string, args []driver.Value) (driver.Rows, error) {
	if q, ok := c.Conn.(driver.Queryer); ok {
		return q.Query(query, args)
	}
	return nil, driver.ErrSkip
}

func (c *failoverConn) QueryContext(ctx context.Context, query string, args []driver.NamedValue) (driver.Rows, error) {
	if q, ok := c.Conn.(driver.QueryerContext); ok {
		return q.QueryContext(ctx, query, args)
	}
	return nil, driver.ErrSkip
}

func (c *failoverConn) PrepareContext(ctx context.Context, query string) (driver.Stmt, error) {
	if p, ok := c.Conn.(driver.ConnPrepareContext); ok {
		return p.PrepareContext(ctx, query)
	}
	return nil, driver.ErrSkip
}

func (c *failoverConn) BeginTx(ctx context.Context, opts driver.TxOptions) (driver.Tx, error) {
	if b, ok := c.Conn.(driver.ConnBeginTx); ok {
		return b.BeginTx(ctx, opts)
	}
	return nil, driver.ErrSkip
}

func (c *failoverConn) ResetSession(ctx context.Context) error {
	// database/sql only validates connections on putConn, so pooled idle
	// connections to a host that is no longer active get reused once unless
	// they are rejected here as well.
	if int(c.connector.active.Load()) != c.host {
		return driver.ErrBadConn
	}
	if r, ok := c.Conn.(driver.SessionResetter); ok {
		return r.ResetSession(ctx)
	}
	return nil
}

func (c *failoverConn) CheckNamedValue(nv *driver.NamedValue) error {
	if n, ok := c.Conn.(driver.NamedValueChecker); ok {
		return n.CheckNamedValue(nv)
	}
	return driver.ErrSkip
}
