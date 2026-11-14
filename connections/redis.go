package connections

import (
	"context"
	"crypto/tls"
	"strings"

	"github.com/redis/go-redis/v9"
	log "github.com/sirupsen/logrus"
)

// NewRedisConnection opens a cluster, sentinel or standalone Redis client
// depending on the configured addresses.
func NewRedisConnection(ctx context.Context, rf *RedisInfo, ll *log.Logger) (redis.UniversalClient, error) {
	if ll == nil {
		ll = log.StandardLogger()
	}
	logger := ll.WithField("method", "NewRedisConnection")

	rdb := newRedisClient(rf)

	if _, err := rdb.Ping(ctx).Result(); err != nil {
		_ = rdb.Close()
		return nil, err
	}

	version := ""
	if info, err := rdb.Info(ctx, "server").Result(); err == nil {
		version = redisVersion(info)
	}
	logger.WithField("version", version).Info("Successfully connected to Redis")

	return rdb, nil
}

// newRedisClient builds the client for the configured deployment mode.
func newRedisClient(rf *RedisInfo) redis.UniversalClient {
	var tlsConfig *tls.Config
	if rf.UseTLS {
		tlsConfig = &tls.Config{MinVersion: tls.VersionTLS12}
	}

	switch {
	case len(rf.ClusterAddresses) > 0:
		return redis.NewClusterClient(&redis.ClusterOptions{
			Addrs:     rf.ClusterAddresses,
			Username:  rf.Username,
			Password:  rf.Password,
			TLSConfig: tlsConfig,
		})
	case len(rf.SentinelAddresses) > 0:
		return redis.NewFailoverClient(&redis.FailoverOptions{
			SentinelAddrs:    rf.SentinelAddresses,
			SentinelUsername: rf.SentinelUsername,
			SentinelPassword: rf.SentinelPassword,
			MasterName:       rf.MasterName,
			Username:         rf.Username,
			Password:         rf.Password,
			DB:               rf.DBName,
			TLSConfig:        tlsConfig,
		})
	default:
		return redis.NewClient(&redis.Options{
			Addr:      rf.Host,
			Username:  rf.Username,
			Password:  rf.Password,
			DB:        rf.DBName,
			TLSConfig: tlsConfig,
		})
	}
}

// redisVersion extracts redis_version from the INFO output.
func redisVersion(info string) string {
	for _, line := range strings.Split(info, "\n") {
		if v, found := strings.CutPrefix(strings.TrimSpace(line), "redis_version:"); found {
			return strings.TrimSpace(v)
		}
	}
	return ""
}
