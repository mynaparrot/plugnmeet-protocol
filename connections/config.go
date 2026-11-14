package connections

import "time"

type DatabaseInfo struct {
	DriverName      string            `yaml:"driver_name"`
	Host            string            `yaml:"host"`
	Port            int32             `yaml:"port"`
	Username        string            `yaml:"username"`
	Password        string            `yaml:"password"`
	DBName          string            `yaml:"db"`
	Prefix          string            `yaml:"prefix"`
	Charset         *string           `yaml:"charset"`
	Loc             *string           `yaml:"loc"`
	ConnMaxLifetime *time.Duration    `yaml:"conn_max_lifetime"`
	MaxOpenConns    *int              `yaml:"max_open_conns"`
	Replicas        []ReplicaDBInfo   `yaml:"replicas"`
	ConnectTimeout  *time.Duration    `yaml:"connect_timeout"`
	ReadTimeout     *time.Duration    `yaml:"read_timeout"`
	WriteTimeout    *time.Duration    `yaml:"write_timeout"`
	Failover        *FailoverSettings `yaml:"failover"`
}

type FailoverSettings struct {
	Hosts               []ReplicaDBInfo `yaml:"hosts"`
	Cooldown            *time.Duration  `yaml:"cooldown"`
	HealthCheckInterval *time.Duration  `yaml:"health_check_interval"`
	Failback            *bool           `yaml:"failback"`
	FailbackDelay       *time.Duration  `yaml:"failback_delay"`
	GaleraCheck         *bool           `yaml:"galera_check"`
}

// ReplicaDBInfo holds connection details for a read replica or failover host.
type ReplicaDBInfo struct {
	Host     string `yaml:"host"`
	Port     int32  `yaml:"port"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
}

type RedisInfo struct {
	Host              string         `yaml:"host"`
	Username          string         `yaml:"username"`
	Password          string         `yaml:"password"`
	DBName            int            `yaml:"db"`
	DefaultTTL        *time.Duration `yaml:"default_ttl"`
	UseTLS            bool           `yaml:"use_tls"`
	MasterName        string         `yaml:"sentinel_master_name"`
	SentinelUsername  string         `yaml:"sentinel_username"`
	SentinelPassword  string         `yaml:"sentinel_password"`
	SentinelAddresses []string       `yaml:"sentinel_addresses"`
	ClusterAddresses  []string       `yaml:"cluster_addresses"`
}
