package connections

import (
	"fmt"
	"net"
	"strconv"
	"strings"
)

// ValidateDatabaseInfo validates the primary, replica and failover settings.
func ValidateDatabaseInfo(d *DatabaseInfo) error {
	return d.Validate()
}

func (d *DatabaseInfo) Validate() error {
	if strings.TrimSpace(d.Host) == "" {
		return fmt.Errorf("database_info.host: required")
	}
	if err := validatePort("database_info.port", int(d.Port)); err != nil {
		return err
	}
	if strings.TrimSpace(d.DBName) == "" {
		return fmt.Errorf("database_info.db: required")
	}
	for i, r := range d.Replicas {
		if strings.TrimSpace(r.Host) == "" {
			return fmt.Errorf("database_info.replicas[%d].host: required", i)
		}
		if r.Port != 0 {
			if err := validatePort(fmt.Sprintf("database_info.replicas[%d].port", i), int(r.Port)); err != nil {
				return err
			}
		}
	}
	if d.Failover == nil {
		return nil
	}
	for i, h := range d.Failover.Hosts {
		if strings.TrimSpace(h.Host) == "" {
			return fmt.Errorf("database_info.failover.hosts[%d].host: required", i)
		}
		if h.Port != 0 {
			if err := validatePort(fmt.Sprintf("database_info.failover.hosts[%d].port", i), int(h.Port)); err != nil {
				return err
			}
		}
	}
	return nil
}

// ValidateRedisInfo validates standalone, sentinel and cluster settings.
func ValidateRedisInfo(r *RedisInfo) error {
	return r.Validate()
}

func (r *RedisInfo) Validate() error {
	if len(r.ClusterAddresses) > 0 {
		if len(r.SentinelAddresses) > 0 {
			return fmt.Errorf("redis_info: cannot set both cluster_addresses and sentinel_addresses")
		}
		for i, addr := range r.ClusterAddresses {
			if err := validateHostPort(fmt.Sprintf("redis_info.cluster_addresses[%d]", i), addr); err != nil {
				return err
			}
		}
		return nil
	}
	if len(r.SentinelAddresses) > 0 {
		if strings.TrimSpace(r.MasterName) == "" {
			return fmt.Errorf("redis_info.sentinel_master_name: required when sentinel_addresses is set")
		}
		for i, addr := range r.SentinelAddresses {
			if err := validateHostPort(fmt.Sprintf("redis_info.sentinel_addresses[%d]", i), addr); err != nil {
				return err
			}
		}
		return nil
	}
	return validateHostPort("redis_info.host", r.Host)
}

// validateHostPort ensures raw is in "host:port" format with a valid port.
func validateHostPort(field, raw string) error {
	host, portStr, err := net.SplitHostPort(strings.TrimSpace(raw))
	if err != nil || host == "" {
		return fmt.Errorf("%s: %q must be in host:port format", field, raw)
	}
	port, err := strconv.Atoi(portStr)
	if err != nil {
		return fmt.Errorf("%s: %q has an invalid port", field, raw)
	}
	return validatePort(field, port)
}

// validatePort ensures port is within the valid TCP range.
func validatePort(field string, port int) error {
	if port < 1 || port > 65535 {
		return fmt.Errorf("%s: port %d must be between 1 and 65535", field, port)
	}
	return nil
}
