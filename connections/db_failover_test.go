package connections

import (
	"context"
	"database/sql/driver"
	"errors"
	"fmt"
	"io"
	"strings"
	"sync"
	"sync/atomic"
	"testing"
	"time"

	"github.com/go-sql-driver/mysql"
	log "github.com/sirupsen/logrus"
)

// logHook captures logrus records so tests can assert on log transitions.
type logHook struct {
	mu      sync.Mutex
	records []logRecord
}

type logRecord struct {
	level log.Level
	msg   string
}

func (h *logHook) Levels() []log.Level { return log.AllLevels }

func (h *logHook) Fire(entry *log.Entry) error {
	h.mu.Lock()
	h.records = append(h.records, logRecord{level: entry.Level, msg: entry.Message})
	h.mu.Unlock()
	return nil
}

func (h *logHook) count(level log.Level, msg string) int {
	h.mu.Lock()
	defer h.mu.Unlock()

	n := 0
	for _, r := range h.records {
		if r.level == level && r.msg == msg {
			n++
		}
	}
	return n
}

// captureLogs redirects the standard logger to a hook until the test ends.
func captureLogs(t *testing.T) *logHook {
	t.Helper()

	logger := log.StandardLogger()
	oldLevel := logger.GetLevel()
	oldHooks := logger.ReplaceHooks(make(log.LevelHooks))
	hook := &logHook{}
	logger.AddHook(hook)
	logger.SetLevel(log.DebugLevel)

	t.Cleanup(func() {
		logger.ReplaceHooks(oldHooks)
		logger.SetLevel(oldLevel)
	})
	return hook
}

type fakeClock struct {
	t time.Time
}

func (c *fakeClock) now() time.Time { return c.t }

func (c *fakeClock) advance(d time.Duration) { c.t = c.t.Add(d) }

type fakeConn struct {
	name  string
	valid bool
	pings int
}

func (c *fakeConn) Prepare(string) (driver.Stmt, error) { return nil, errors.New("unsupported") }
func (c *fakeConn) Close() error                        { return nil }
func (c *fakeConn) Begin() (driver.Tx, error)           { return nil, errors.New("unsupported") }
func (c *fakeConn) IsValid() bool                       { return c.valid }

func (c *fakeConn) Ping(context.Context) error {
	c.pings++
	return nil
}

// galeraConn answers the wsrep status query with fixed rows.
type galeraConn struct {
	fakeConn
	rows [][]driver.Value
}

func (c *galeraConn) QueryContext(context.Context, string, []driver.NamedValue) (driver.Rows, error) {
	return &fakeRows{values: c.rows}, nil
}

type fakeRows struct {
	values [][]driver.Value
	idx    int
}

func (r *fakeRows) Columns() []string { return []string{"Variable_name", "Value"} }
func (r *fakeRows) Close() error      { return nil }

func (r *fakeRows) Next(dest []driver.Value) error {
	if r.idx >= len(r.values) {
		return io.EOF
	}
	copy(dest, r.values[r.idx])
	r.idx++
	return nil
}

type fakeConnector struct {
	name     string
	attempts int
	open     func() (driver.Conn, error)
}

func (c *fakeConnector) Connect(context.Context) (driver.Conn, error) {
	c.attempts++
	return c.open()
}

func (c *fakeConnector) Driver() driver.Driver { return &mysql.MySQLDriver{} }

func healthy(context.Context, driver.Conn) error { return nil }

type failoverSetup struct {
	connector *failoverConnector
	clock     *fakeClock
	up        map[string]bool
	hosts     map[string]*fakeConnector
}

func newFailoverSetup(cooldown time.Duration, healthCheck func(context.Context, driver.Conn) error, names ...string) *failoverSetup {
	if len(names) == 0 {
		names = []string{"primary", "backup"}
	}

	s := &failoverSetup{
		clock: &fakeClock{t: time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC)},
		up:    map[string]bool{},
		hosts: map[string]*fakeConnector{},
	}

	hosts := make([]*failoverHost, 0, len(names))
	for _, name := range names {
		s.up[name] = true
		connector := &fakeConnector{name: name}
		connector.open = func() (driver.Conn, error) {
			if !s.up[name] {
				return nil, fmt.Errorf("%s down", name)
			}
			return &fakeConn{name: name, valid: true}, nil
		}
		s.hosts[name] = connector
		hosts = append(hosts, &failoverHost{name: name, connector: connector})
	}

	s.connector = newFailoverConnectorWithHosts(hosts, cooldown, healthCheck)
	s.connector.now = s.clock.now
	return s
}

func (s *failoverSetup) activeIndex() int { return int(s.connector.active.Load()) }

func (s *failoverSetup) connect(t *testing.T) (*failoverConn, *fakeConn) {
	t.Helper()

	conn, err := s.connector.Connect(context.Background())
	if err != nil {
		t.Fatalf("unexpected connect error: %v", err)
	}
	fc, ok := conn.(*failoverConn)
	if !ok {
		t.Fatalf("expected *failoverConn, got %T", conn)
	}
	return fc, fc.Conn.(*fakeConn)
}

func TestFailoverPrimaryDownAtStart(t *testing.T) {
	s := newFailoverSetup(30*time.Second, healthy)
	s.up["primary"] = false

	conn, inner := s.connect(t)
	if conn.host != 1 || inner.name != "backup" {
		t.Fatalf("expected backup conn, got host=%d name=%s", conn.host, inner.name)
	}
	if s.activeIndex() != 1 {
		t.Fatalf("expected active index 1, got %d", s.activeIndex())
	}
	if !conn.IsValid() {
		t.Fatal("expected active conn to be valid")
	}
}

func TestFailoverSwitchesWhileRunning(t *testing.T) {
	s := newFailoverSetup(30*time.Second, healthy)

	first, _ := s.connect(t)
	if first.host != 0 {
		t.Fatalf("expected primary, got host %d", first.host)
	}
	if !first.IsValid() {
		t.Fatal("expected primary conn to be valid")
	}

	s.up["primary"] = false
	second, inner := s.connect(t)
	if second.host != 1 || inner.name != "backup" {
		t.Fatalf("expected backup, got host=%d name=%s", second.host, inner.name)
	}
	if first.IsValid() {
		t.Fatal("expected conn on non-active host to be invalid")
	}
	if !second.IsValid() {
		t.Fatal("expected conn on active host to be valid")
	}
}

func TestFailoverSkipsUnhealthyHost(t *testing.T) {
	healthCheck := func(_ context.Context, conn driver.Conn) error {
		if conn.(*fakeConn).name == "primary" {
			return errors.New("wsrep_ready is OFF")
		}
		return nil
	}
	s := newFailoverSetup(30*time.Second, healthCheck)

	conn, inner := s.connect(t)
	if conn.host != 1 || inner.name != "backup" {
		t.Fatalf("expected backup, got host=%d name=%s", conn.host, inner.name)
	}
	if !s.connector.isDown(0) {
		t.Fatal("expected unhealthy host to be marked down")
	}
}

func TestFailoverLiveConnectStaysOnActiveBackup(t *testing.T) {
	s := newFailoverSetup(30*time.Second, healthy)
	s.up["primary"] = false

	if _, inner := s.connect(t); inner.name != "backup" {
		t.Fatalf("expected backup, got %s", inner.name)
	}
	if s.hosts["primary"].attempts != 1 {
		t.Fatalf("expected 1 primary attempt, got %d", s.hosts["primary"].attempts)
	}

	// primary recovered, but live connects must stay on the healthy active host
	s.up["primary"] = true
	s.clock.advance(31 * time.Second)
	for i := 0; i < 3; i++ {
		if _, inner := s.connect(t); inner.name != "backup" {
			t.Fatalf("expected backup on live connect, got %s", inner.name)
		}
	}
	if s.hosts["primary"].attempts != 1 {
		t.Fatalf("primary must not be probed by live connects, attempts=%d", s.hosts["primary"].attempts)
	}
	if s.activeIndex() != 1 {
		t.Fatalf("expected active index 1, got %d", s.activeIndex())
	}
}

func TestFailoverConnIsValidDelegates(t *testing.T) {
	s := newFailoverSetup(30*time.Second, healthy)
	s.up["primary"] = false
	s.connect(t)

	stale := &failoverConn{Conn: &fakeConn{name: "primary", valid: true}, connector: s.connector, host: 0}
	if stale.IsValid() {
		t.Fatal("expected conn on non-active host to be invalid")
	}

	inner := &failoverConn{Conn: &fakeConn{name: "backup", valid: false}, connector: s.connector, host: 1}
	if inner.IsValid() {
		t.Fatal("expected inner validator to be honoured")
	}
}

func TestFailoverConnResetSessionRejectsStaleHost(t *testing.T) {
	s := newFailoverSetup(30*time.Second, healthy)
	s.up["primary"] = false

	stale, inner := s.connect(t)
	if stale.host != 1 || inner.name != "backup" {
		t.Fatalf("expected backup, got host=%d name=%s", stale.host, inner.name)
	}

	// primary recovers and becomes active again, so the backup conn is stale
	s.up["primary"] = true
	s.clock.advance(31 * time.Second)
	s.connector.checkOnce(context.Background())
	s.clock.advance(60 * time.Second)
	s.connector.checkOnce(context.Background())

	active, activeInner := s.connect(t)
	if active.host != 0 || activeInner.name != "primary" {
		t.Fatalf("expected failback to primary, got host=%d name=%s", active.host, activeInner.name)
	}

	if err := stale.ResetSession(context.Background()); !errors.Is(err, driver.ErrBadConn) {
		t.Fatalf("expected driver.ErrBadConn for conn on non-active host, got %v", err)
	}
	if err := active.ResetSession(context.Background()); err != nil {
		t.Fatalf("expected nil for conn on active host, got %v", err)
	}
}

func TestFailoverAllHostsDown(t *testing.T) {
	s := newFailoverSetup(30*time.Second, healthy)
	s.up["primary"] = false
	s.up["backup"] = false

	if _, err := s.connector.Connect(context.Background()); err == nil {
		t.Fatal("expected error when all hosts are down")
	}
	for idx := range s.connector.hosts {
		if !s.connector.isDown(idx) {
			t.Fatalf("expected host %d to be marked down", idx)
		}
	}
}

func TestFailoverLastResortIgnoresCooldown(t *testing.T) {
	s := newFailoverSetup(30*time.Second, healthy, "primary", "backup", "third")
	s.up["primary"] = false
	s.up["backup"] = false
	s.up["third"] = false

	if _, err := s.connector.Connect(context.Background()); err == nil {
		t.Fatal("expected error when all hosts are down")
	}

	// still within cooldown, but every host is skipped, so retry them all
	s.up["backup"] = true
	conn, inner := s.connect(t)
	if inner.name != "backup" || conn.host != 1 {
		t.Fatalf("expected backup via last-resort pass, got host=%d name=%s", conn.host, inner.name)
	}
}

func TestFailoverConnectorDriver(t *testing.T) {
	s := newFailoverSetup(30*time.Second, healthy)
	if _, ok := s.connector.Driver().(*mysql.MySQLDriver); !ok {
		t.Fatalf("expected *mysql.MySQLDriver, got %T", s.connector.Driver())
	}
}

func TestFailoverActiveIndexBeforeFirstConnect(t *testing.T) {
	s := newFailoverSetup(30*time.Second, healthy)
	if s.activeIndex() != -1 {
		t.Fatalf("expected active index -1 before first connect, got %d", s.activeIndex())
	}

	conn, _ := s.connect(t)
	if conn.host != 0 {
		t.Fatalf("expected primary conn, got host %d", conn.host)
	}
	if s.activeIndex() != 0 {
		t.Fatalf("expected active index 0 after connect, got %d", s.activeIndex())
	}
}

func TestFailoverConnPing(t *testing.T) {
	s := newFailoverSetup(30*time.Second, healthy)
	active, _ := s.connect(t)

	inner := active.Conn.(*fakeConn)
	if err := active.Ping(context.Background()); err != nil {
		t.Fatalf("expected nil ping on active host, got %v", err)
	}
	if inner.pings != 1 {
		t.Fatalf("expected ping delegated to inner conn, pings=%d", inner.pings)
	}

	stale := &failoverConn{Conn: &fakeConn{name: "backup", valid: true}, connector: s.connector, host: 1}
	if err := stale.Ping(context.Background()); !errors.Is(err, driver.ErrBadConn) {
		t.Fatalf("expected driver.ErrBadConn for stale host, got %v", err)
	}
}

func wsrepRows(ready, state string) [][]driver.Value {
	var rows [][]driver.Value
	if ready != "" {
		rows = append(rows, []driver.Value{"wsrep_ready", ready})
	}
	if state != "" {
		rows = append(rows, []driver.Value{"wsrep_local_state", state})
	}
	return rows
}

func TestGaleraHealthCheck(t *testing.T) {
	tests := []struct {
		name    string
		rows    [][]driver.Value
		wantErr string
	}{
		{name: "missing both", rows: nil, wantErr: "missing Galera wsrep status"},
		{name: "missing state", rows: wsrepRows("ON", ""), wantErr: "missing Galera wsrep status"},
		{name: "missing ready", rows: wsrepRows("", "4"), wantErr: "missing Galera wsrep status"},
		{name: "not ready", rows: wsrepRows("OFF", "4"), wantErr: `wsrep_ready is "OFF"`},
		{name: "not synced", rows: wsrepRows("ON", "2"), wantErr: `wsrep_local_state is "2"`},
		{name: "synced", rows: wsrepRows("ON", "4")},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := galeraHealthCheck(context.Background(), &galeraConn{rows: tt.rows})
			if tt.wantErr == "" {
				if err != nil {
					t.Fatalf("expected nil error, got %v", err)
				}
				return
			}
			if err == nil || !strings.Contains(err.Error(), tt.wantErr) {
				t.Fatalf("expected error containing %q, got %v", tt.wantErr, err)
			}
		})
	}
}

func TestFailoverGaleraCheckConfig(t *testing.T) {
	const dsn = "user:pass@tcp(127.0.0.1:3306)/db"
	info := &DatabaseInfo{Host: "127.0.0.1", Port: 3306, Failover: &FailoverSettings{}}

	connector, err := newFailoverConnector(info, info.Failover, dsn, time.Second, log.StandardLogger())
	if err != nil {
		t.Fatalf("unexpected connector error: %v", err)
	}
	c := connector.(*failoverConnector)
	if got := c.active.Load(); got != -1 {
		t.Fatalf("expected active -1 after construction, got %d", got)
	}
	if err := c.healthCheck(context.Background(), &galeraConn{}); err == nil {
		t.Fatal("expected default Galera check to reject node without wsrep status")
	}

	disabled := false
	info.Failover.GaleraCheck = &disabled
	connector, err = newFailoverConnector(info, info.Failover, dsn, time.Second, log.StandardLogger())
	if err != nil {
		t.Fatalf("unexpected connector error: %v", err)
	}
	if err := connector.(*failoverConnector).healthCheck(context.Background(), &galeraConn{}); err != nil {
		t.Fatalf("expected node to be accepted when galera_check is false, got %v", err)
	}
}

func TestFailoverCheckOnceSwitchesFromFailingPrimary(t *testing.T) {
	unhealthy := map[string]bool{}
	healthCheck := func(_ context.Context, conn driver.Conn) error {
		if unhealthy[conn.(*fakeConn).name] {
			return errors.New("wsrep_ready is OFF")
		}
		return nil
	}
	s := newFailoverSetup(30*time.Second, healthCheck)

	primary, _ := s.connect(t)
	if s.activeIndex() != 0 {
		t.Fatalf("expected active index 0, got %d", s.activeIndex())
	}

	unhealthy["primary"] = true
	s.connector.checkOnce(context.Background())

	if s.activeIndex() != 1 {
		t.Fatalf("expected active index 1, got %d", s.activeIndex())
	}
	if primary.IsValid() {
		t.Fatal("expected conn on failing primary to be invalid")
	}
}

func TestFailoverCheckOnceKeepsBackupDuringCooldown(t *testing.T) {
	s := newFailoverSetup(30*time.Second, healthy)
	s.up["primary"] = false

	if _, inner := s.connect(t); inner.name != "backup" {
		t.Fatalf("expected backup, got %s", inner.name)
	}
	primaryAttempts := s.hosts["primary"].attempts

	// primary recovered but is still cooling down: no retry, stay on backup
	s.up["primary"] = true
	s.connector.checkOnce(context.Background())
	if s.activeIndex() != 1 {
		t.Fatalf("expected active index 1 during cooldown, got %d", s.activeIndex())
	}
	if s.hosts["primary"].attempts != primaryAttempts {
		t.Fatalf("primary must not be retried during cooldown, attempts=%d", s.hosts["primary"].attempts)
	}
}

func TestFailoverCheckOnceFailsBackAfterDelay(t *testing.T) {
	hook := captureLogs(t)
	s := newFailoverSetup(30*time.Second, healthy)
	s.up["primary"] = false

	if _, inner := s.connect(t); inner.name != "backup" {
		t.Fatalf("expected backup, got %s", inner.name)
	}

	// cooldown expired: the primary is probed and starts its failback wait
	s.up["primary"] = true
	s.clock.advance(31 * time.Second)
	s.connector.checkOnce(context.Background())
	if s.activeIndex() != 1 {
		t.Fatalf("expected active index 1 when the wait starts, got %d", s.activeIndex())
	}
	const started = "database host primary is healthy, failing back after 1m0s"
	if got := hook.count(log.InfoLevel, started); got != 1 {
		t.Fatalf("expected 1 failback-wait log, got %d", got)
	}

	// still inside the delay: no switch and no repeated start log
	s.clock.advance(30 * time.Second)
	s.connector.checkOnce(context.Background())
	if s.activeIndex() != 1 {
		t.Fatalf("expected active index 1 inside the delay, got %d", s.activeIndex())
	}
	if got := hook.count(log.InfoLevel, started); got != 1 {
		t.Fatalf("expected the start log only once, got %d", got)
	}
	if got := hook.count(log.DebugLevel, "host primary healthy, failback in 30s"); got != 1 {
		t.Fatalf("expected 1 remaining-delay debug log, got %d", got)
	}

	// delay elapsed: fail back to the primary
	s.clock.advance(30 * time.Second)
	s.connector.checkOnce(context.Background())
	if s.activeIndex() != 0 {
		t.Fatalf("expected failback to primary, got active index %d", s.activeIndex())
	}
	if s.connector.isDown(0) {
		t.Fatal("expected down mark to be cleared after failback")
	}
	if got := hook.count(log.WarnLevel, "switched DB from backup to primary"); got != 1 {
		t.Fatalf("expected 1 switch log, got %d", got)
	}
}

func TestFailoverFailbackWaitRestartsAfterProbeFailure(t *testing.T) {
	unhealthy := map[string]bool{}
	healthCheck := func(_ context.Context, conn driver.Conn) error {
		if unhealthy[conn.(*fakeConn).name] {
			return errors.New("wsrep_ready is OFF")
		}
		return nil
	}
	s := newFailoverSetup(30*time.Second, healthCheck)
	s.up["primary"] = false

	if _, inner := s.connect(t); inner.name != "backup" {
		t.Fatalf("expected backup, got %s", inner.name)
	}

	s.up["primary"] = true
	s.clock.advance(31 * time.Second)
	s.connector.checkOnce(context.Background())
	if s.activeIndex() != 1 {
		t.Fatalf("expected active index 1 when the wait starts, got %d", s.activeIndex())
	}

	// the primary fails a probe, which restarts its wait after the cooldown
	unhealthy["primary"] = true
	s.clock.advance(30 * time.Second)
	s.connector.checkOnce(context.Background())
	if s.activeIndex() != 1 || !s.connector.isDown(0) {
		t.Fatalf("expected backup with primary marked down, active=%d down=%v", s.activeIndex(), s.connector.isDown(0))
	}

	// healthy again, but the delay restarts from scratch
	unhealthy["primary"] = false
	s.clock.advance(61 * time.Second)
	s.connector.checkOnce(context.Background())
	if s.activeIndex() != 1 {
		t.Fatalf("expected backup after the restarted wait, got %d", s.activeIndex())
	}

	s.clock.advance(60 * time.Second)
	s.connector.checkOnce(context.Background())
	if s.activeIndex() != 0 {
		t.Fatalf("expected failback after the restarted delay, got %d", s.activeIndex())
	}
}

func TestFailoverFailbackDisabled(t *testing.T) {
	s := newFailoverSetup(30*time.Second, healthy)
	s.connector.failback = false
	s.up["primary"] = false

	if _, inner := s.connect(t); inner.name != "backup" {
		t.Fatalf("expected backup, got %s", inner.name)
	}

	// the healthy primary is never probed or re-activated by the monitor
	s.up["primary"] = true
	s.clock.advance(time.Hour)
	s.connector.checkOnce(context.Background())
	if s.activeIndex() != 1 {
		t.Fatalf("expected active index 1 with failback disabled, got %d", s.activeIndex())
	}
	if s.hosts["primary"].attempts != 1 {
		t.Fatalf("primary must not be probed with failback disabled, attempts=%d", s.hosts["primary"].attempts)
	}

	// a failure of the active host still moves to the healthy primary immediately
	s.up["backup"] = false
	if _, inner := s.connect(t); inner.name != "primary" {
		t.Fatalf("expected primary after backup failure, got %s", inner.name)
	}
	if s.activeIndex() != 0 {
		t.Fatalf("expected active index 0, got %d", s.activeIndex())
	}
}

func TestFailoverStartupUsesPriorityOrder(t *testing.T) {
	s := newFailoverSetup(30*time.Second, healthy, "primary", "backup", "third")
	s.up["primary"] = false

	if _, inner := s.connect(t); inner.name != "backup" {
		t.Fatalf("expected first healthy host in priority order, got %s", inner.name)
	}
	if s.activeIndex() != 1 {
		t.Fatalf("expected active index 1, got %d", s.activeIndex())
	}
	if s.hosts["third"].attempts != 0 {
		t.Fatalf("lower-priority host must not be tried, attempts=%d", s.hosts["third"].attempts)
	}
}

// hungConnector blocks in Connect until the attempt context is done.
type hungConnector struct {
	attempts int32
}

func (c *hungConnector) Connect(ctx context.Context) (driver.Conn, error) {
	atomic.AddInt32(&c.attempts, 1)
	<-ctx.Done()
	return nil, ctx.Err()
}

func (c *hungConnector) Driver() driver.Driver { return &mysql.MySQLDriver{} }

func TestFailoverAttemptTimeoutHungHost(t *testing.T) {
	backup := &fakeConnector{name: "backup"}
	backup.open = func() (driver.Conn, error) { return &fakeConn{name: "backup", valid: true}, nil }

	t.Run("hung connect", func(t *testing.T) {
		hung := &hungConnector{}
		s := &failoverSetup{
			clock: &fakeClock{t: time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC)},
			up:    map[string]bool{},
			hosts: map[string]*fakeConnector{"backup": backup},
		}
		s.connector = newFailoverConnectorWithHosts([]*failoverHost{
			{name: "hung", connector: hung},
			{name: "backup", connector: backup},
		}, time.Minute, healthy)
		s.connector.now = s.clock.now
		s.connector.attemptTimeout = 20 * time.Millisecond

		conn, idx, _, err := s.connector.connectPass(context.Background(), true)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		defer conn.Close()
		if idx != 1 {
			t.Fatalf("expected backup index 1, got %d", idx)
		}
		if atomic.LoadInt32(&hung.attempts) != 1 {
			t.Fatalf("expected single hung attempt, got %d", hung.attempts)
		}
	})

	t.Run("hung health check", func(t *testing.T) {
		hungCheck := func(ctx context.Context, conn driver.Conn) error {
			if conn.(*fakeConn).name == "hung" {
				<-ctx.Done()
				return ctx.Err()
			}
			return nil
		}
		hung := &fakeConnector{name: "hung"}
		hung.open = func() (driver.Conn, error) { return &fakeConn{name: "hung", valid: true}, nil }

		s := &failoverSetup{
			clock: &fakeClock{t: time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC)},
			up:    map[string]bool{},
			hosts: map[string]*fakeConnector{"backup": backup},
		}
		s.connector = newFailoverConnectorWithHosts([]*failoverHost{
			{name: "hung", connector: hung},
			{name: "backup", connector: backup},
		}, time.Minute, hungCheck)
		s.connector.now = s.clock.now
		s.connector.attemptTimeout = 20 * time.Millisecond

		conn, idx, _, err := s.connector.connectPass(context.Background(), true)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		defer conn.Close()
		if idx != 1 {
			t.Fatalf("expected backup index 1, got %d", idx)
		}
	})
}

func TestFailoverCheckOnceAllHostsDown(t *testing.T) {
	s := newFailoverSetup(30*time.Second, healthy)
	s.connect(t)

	s.up["primary"] = false
	s.up["backup"] = false
	s.connector.checkOnce(context.Background())

	if s.activeIndex() != 0 {
		t.Fatalf("expected active index to stay 0, got %d", s.activeIndex())
	}
}

func TestFailoverMonitorStartStop(t *testing.T) {
	var attempts atomic.Int32
	primary := &fakeConnector{name: "primary"}
	primary.open = func() (driver.Conn, error) {
		attempts.Add(1)
		return &fakeConn{name: "primary", valid: true}, nil
	}

	c := newFailoverConnectorWithHosts([]*failoverHost{{name: "primary", connector: primary}}, time.Minute, healthy)
	c.attemptTimeout = time.Second
	c.start(5 * time.Millisecond)

	deadline := time.Now().Add(5 * time.Second)
	for attempts.Load() == 0 && time.Now().Before(deadline) {
		time.Sleep(time.Millisecond)
	}
	if attempts.Load() == 0 {
		t.Fatal("expected the monitor to probe the primary")
	}

	if err := c.Close(); err != nil {
		t.Fatalf("unexpected close error: %v", err)
	}
	before := attempts.Load()
	time.Sleep(30 * time.Millisecond)
	if got := attempts.Load(); got != before {
		t.Fatalf("monitor still running after close, attempts %d -> %d", before, got)
	}
	if c.active.Load() != 0 {
		t.Fatalf("expected active index 0, got %d", c.active.Load())
	}
	if err := c.Close(); err != nil {
		t.Fatalf("expected idempotent close, got %v", err)
	}
}

func TestFailoverMarkDownLogsTransitionOnly(t *testing.T) {
	hook := captureLogs(t)
	s := newFailoverSetup(30*time.Second, healthy)

	failed := fmt.Sprintf("database host %s failed, skipping it for %s", "primary", 30*time.Second)
	err := errors.New("connect: primary down")

	s.connector.markDown(0, err)
	s.connector.markDown(0, err)
	s.connector.markDown(0, err)

	if got := hook.count(log.WarnLevel, failed); got != 1 {
		t.Fatalf("expected 1 warn for repeated failures, got %d", got)
	}
	if got := hook.count(log.DebugLevel, failed); got != 2 {
		t.Fatalf("expected 2 debug lines for repeated failures, got %d", got)
	}

	recovered := "database host primary recovered"
	s.connector.activate(0)
	if got := hook.count(log.InfoLevel, recovered); got != 1 {
		t.Fatalf("expected 1 recovery log, got %d", got)
	}

	// a host that keeps working must not log its recovery again
	s.connector.activate(0)
	if got := hook.count(log.InfoLevel, recovered); got != 1 {
		t.Fatalf("expected recovery log only once, got %d", got)
	}
}

func TestFailoverAllDownLogsTransitionOnly(t *testing.T) {
	hook := captureLogs(t)
	s := newFailoverSetup(30*time.Second, healthy)
	s.up["primary"] = false
	s.up["backup"] = false

	const unavailable = "no healthy database host available"

	s.connector.checkOnce(context.Background())
	s.connector.checkOnce(context.Background())

	if got := hook.count(log.WarnLevel, unavailable); got != 1 {
		t.Fatalf("expected 1 warn while all hosts are down, got %d", got)
	}
	if got := hook.count(log.DebugLevel, unavailable); got != 1 {
		t.Fatalf("expected 1 debug while all hosts stay down, got %d", got)
	}

	s.up["backup"] = true
	s.clock.advance(31 * time.Second)
	s.connector.checkOnce(context.Background())

	const available = "database host available again"
	if got := hook.count(log.InfoLevel, available); got != 1 {
		t.Fatalf("expected 1 recovery log, got %d", got)
	}
	if s.activeIndex() != 1 {
		t.Fatalf("expected active index 1, got %d", s.activeIndex())
	}

	s.connector.checkOnce(context.Background())
	if got := hook.count(log.InfoLevel, available); got != 1 {
		t.Fatalf("expected recovery log only once, got %d", got)
	}
	if got := hook.count(log.WarnLevel, unavailable); got != 1 {
		t.Fatalf("expected no extra warn after recovery, got %d", got)
	}
}

func TestFailoverErrorStages(t *testing.T) {
	s := newFailoverSetup(30*time.Second, healthy)
	s.up["primary"] = false
	s.up["backup"] = false

	_, err := s.connector.Connect(context.Background())
	if err == nil || !strings.Contains(err.Error(), "connect:") {
		t.Fatalf("expected connect stage prefix, got %v", err)
	}

	unhealthy := func(context.Context, driver.Conn) error { return errors.New("wsrep_ready is OFF") }
	s = newFailoverSetup(30*time.Second, unhealthy)

	_, err = s.connector.Connect(context.Background())
	if err == nil || !strings.Contains(err.Error(), "health check:") {
		t.Fatalf("expected health check stage prefix, got %v", err)
	}
	if !strings.Contains(err.Error(), "wsrep_ready is OFF") {
		t.Fatalf("expected wrapped cause in error, got %v", err)
	}
}
