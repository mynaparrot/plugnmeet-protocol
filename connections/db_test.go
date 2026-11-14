package connections

import (
	"context"
	"testing"
	"time"

	log "github.com/sirupsen/logrus"
)

func TestDBOptions(t *testing.T) {
	o := &dbOptions{defaultTimezone: "UTC"}
	WithDebug(true)(o)
	WithSingularTable(true)(o)
	WithDefaultTimezone("Asia/Dhaka")(o)

	if !o.debug || !o.singularTable {
		t.Fatalf("expected debug and singular table to be enabled: %+v", o)
	}
	if o.defaultTimezone != "Asia/Dhaka" {
		t.Fatalf("unexpected timezone %q", o.defaultTimezone)
	}
}

func TestNamingStrategy(t *testing.T) {
	info := &DatabaseInfo{Prefix: "pnm_"}

	// Callers that prefix names in TableName() must not get a second prefix.
	if ns := namingStrategy(&dbOptions{}, info); ns != nil {
		t.Fatalf("expected no naming strategy without singular table, got %+v", ns)
	}

	ns := namingStrategy(&dbOptions{singularTable: true}, info)
	if ns == nil || ns.TablePrefix != "pnm_" || !ns.SingularTable {
		t.Fatalf("unexpected naming strategy %+v", ns)
	}
}

func TestTimeoutParams(t *testing.T) {
	read := 5 * time.Second
	write := 7 * time.Second

	if got := timeoutParams(&DatabaseInfo{}); got != "" {
		t.Fatalf("expected empty params, got %q", got)
	}
	if got, want := timeoutParams(&DatabaseInfo{ReadTimeout: &read}), "&readTimeout=5s"; got != want {
		t.Fatalf("expected %q, got %q", want, got)
	}
	if got, want := timeoutParams(&DatabaseInfo{ReadTimeout: &read, WriteTimeout: &write}), "&readTimeout=5s&writeTimeout=7s"; got != want {
		t.Fatalf("expected %q, got %q", want, got)
	}
}

func TestFailoverDefaults(t *testing.T) {
	f := &FailoverSettings{}
	if got := failoverCooldown(f); got != defaultFailoverCooldown {
		t.Fatalf("unexpected cooldown %s", got)
	}
	if got := healthCheckInterval(f); got != defaultHealthCheckInterval {
		t.Fatalf("unexpected interval %s", got)
	}
	if got := failbackDelay(f); got != defaultFailbackDelay {
		t.Fatalf("unexpected delay %s", got)
	}
	if !failbackEnabled(f) {
		t.Fatal("expected failback to be enabled by default")
	}

	cooldown := time.Minute
	interval := 0 * time.Second
	delay := 2 * time.Minute
	disabled := false
	f = &FailoverSettings{Cooldown: &cooldown, HealthCheckInterval: &interval, FailbackDelay: &delay, Failback: &disabled}

	if got := failoverCooldown(f); got != cooldown {
		t.Fatalf("unexpected cooldown %s", got)
	}
	if got := healthCheckInterval(f); got != interval {
		t.Fatalf("unexpected interval %s", got)
	}
	if got := failbackDelay(f); got != delay {
		t.Fatalf("unexpected delay %s", got)
	}
	if failbackEnabled(f) {
		t.Fatal("expected failback to be disabled")
	}
}

func TestNewDatabaseConnectionUnreachable(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	timeout := 500 * time.Millisecond
	info := &DatabaseInfo{
		Host: "127.0.0.1", Port: 1, Username: "root", DBName: "plugnmeet",
		ConnectTimeout: &timeout,
	}

	if _, err := NewDatabaseConnection(ctx, info, log.StandardLogger()); err == nil {
		t.Fatal("expected an error for an unreachable database")
	}
}

func TestNewDatabaseConnectionFailoverUnreachable(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	timeout := 300 * time.Millisecond
	interval := time.Hour
	info := &DatabaseInfo{
		Host: "127.0.0.1", Port: 1, Username: "root", DBName: "plugnmeet",
		ConnectTimeout: &timeout,
		Failover: &FailoverSettings{
			HealthCheckInterval: &interval,
			Hosts:               []ReplicaDBInfo{{Host: "127.0.0.1", Port: 2}},
		},
	}

	if _, err := NewDatabaseConnection(ctx, info, log.StandardLogger(), WithDebug(true), WithSingularTable(true)); err == nil {
		t.Fatal("expected an error when no failover host is reachable")
	}
}
