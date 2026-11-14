package connections

import (
	"context"
	"testing"
	"time"

	"github.com/redis/go-redis/v9"
)

func TestNewRedisClientStandalone(t *testing.T) {
	rdb := newRedisClient(&RedisInfo{Host: "localhost:6379", DBName: 2, Username: "u", Password: "p"})
	defer rdb.Close()

	client, ok := rdb.(*redis.Client)
	if !ok {
		t.Fatalf("expected *redis.Client, got %T", rdb)
	}
	opt := client.Options()
	if opt.Addr != "localhost:6379" {
		t.Fatalf("unexpected addr %q", opt.Addr)
	}
	if opt.DB != 2 || opt.Username != "u" || opt.Password != "p" {
		t.Fatalf("credentials not applied: %+v", opt)
	}
	if opt.TLSConfig != nil {
		t.Fatal("expected no TLS config")
	}
}

func TestNewRedisClientSentinel(t *testing.T) {
	rdb := newRedisClient(&RedisInfo{
		MasterName:        "mymaster",
		SentinelAddresses: []string{"10.0.0.1:26379"},
		SentinelUsername:  "su",
		SentinelPassword:  "sp",
		DBName:            3,
	})
	defer rdb.Close()

	client, ok := rdb.(*redis.Client)
	if !ok {
		t.Fatalf("expected *redis.Client for sentinel mode, got %T", rdb)
	}
	// go-redis resolves the master through the sentinels at runtime, so the
	// static address is just a marker instead of a real host:port
	if addr := client.Options().Addr; addr != "FailoverClient" {
		t.Fatalf("expected the failover marker addr, got %q", addr)
	}
	if client.Options().DB != 3 {
		t.Fatalf("unexpected db %d", client.Options().DB)
	}
}

func TestNewRedisClientCluster(t *testing.T) {
	rdb := newRedisClient(&RedisInfo{ClusterAddresses: []string{"10.0.0.1:6379", "10.0.0.2:6379"}})
	defer rdb.Close()

	if _, ok := rdb.(*redis.ClusterClient); !ok {
		t.Fatalf("expected *redis.ClusterClient, got %T", rdb)
	}
}

func TestNewRedisClientTLS(t *testing.T) {
	rdb := newRedisClient(&RedisInfo{Host: "localhost:6380", UseTLS: true})
	defer rdb.Close()

	client := rdb.(*redis.Client)
	if client.Options().TLSConfig == nil {
		t.Fatal("expected TLS config to be set")
	}
}

func TestNewRedisConnectionUnreachable(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	// nothing listens on port 1, so the ping must fail and the client is closed
	if _, err := NewRedisConnection(ctx, &RedisInfo{Host: "127.0.0.1:1"}, nil); err == nil {
		t.Fatal("expected an error for an unreachable Redis host")
	}
}

func TestRedisVersion(t *testing.T) {
	info := "# Server\r\nredis_version:7.2.4\r\nredis_git_sha1:00000000\r\n"
	if got := redisVersion(info); got != "7.2.4" {
		t.Fatalf("expected 7.2.4, got %q", got)
	}
	if got := redisVersion("# Server\r\nredis_mode:standalone\r\n"); got != "" {
		t.Fatalf("expected empty version, got %q", got)
	}
}
