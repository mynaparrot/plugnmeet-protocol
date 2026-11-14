package connections

import "testing"

func TestValidateDatabaseInfo(t *testing.T) {
	tt := []struct {
		name    string
		info    *DatabaseInfo
		wantErr string
	}{
		{
			name: "valid",
			info: &DatabaseInfo{Host: "localhost", Port: 3306, DBName: "plugnmeet"},
		},
		{
			name:    "missing host",
			info:    &DatabaseInfo{Port: 3306, DBName: "plugnmeet"},
			wantErr: "database_info.host: required",
		},
		{
			name:    "blank host",
			info:    &DatabaseInfo{Host: "  ", Port: 3306, DBName: "plugnmeet"},
			wantErr: "database_info.host: required",
		},
		{
			name:    "missing port",
			info:    &DatabaseInfo{Host: "localhost", DBName: "plugnmeet"},
			wantErr: "database_info.port: port 0 must be between 1 and 65535",
		},
		{
			name:    "invalid port",
			info:    &DatabaseInfo{Host: "localhost", Port: 70000, DBName: "plugnmeet"},
			wantErr: "database_info.port: port 70000 must be between 1 and 65535",
		},
		{
			name:    "missing db",
			info:    &DatabaseInfo{Host: "localhost", Port: 3306},
			wantErr: "database_info.db: required",
		},
		{
			name: "valid replica with inherited port",
			info: &DatabaseInfo{
				Host: "localhost", Port: 3306, DBName: "plugnmeet",
				Replicas: []ReplicaDBInfo{{Host: "replica.local"}, {Host: "replica2.local", Port: 3307}},
			},
		},
		{
			name: "replica missing host",
			info: &DatabaseInfo{
				Host: "localhost", Port: 3306, DBName: "plugnmeet",
				Replicas: []ReplicaDBInfo{{Host: "replica.local"}, {Port: 3307}},
			},
			wantErr: "database_info.replicas[1].host: required",
		},
		{
			name: "replica invalid port",
			info: &DatabaseInfo{
				Host: "localhost", Port: 3306, DBName: "plugnmeet",
				Replicas: []ReplicaDBInfo{{Host: "replica.local", Port: -1}},
			},
			wantErr: "database_info.replicas[0].port: port -1 must be between 1 and 65535",
		},
		{
			name: "valid failover hosts",
			info: &DatabaseInfo{
				Host: "localhost", Port: 3306, DBName: "plugnmeet",
				Failover: &FailoverSettings{Hosts: []ReplicaDBInfo{{Host: "db2.local"}, {Host: "db3.local", Port: 3308}}},
			},
		},
		{
			name: "failover host missing host",
			info: &DatabaseInfo{
				Host: "localhost", Port: 3306, DBName: "plugnmeet",
				Failover: &FailoverSettings{Hosts: []ReplicaDBInfo{{Port: 3307}}},
			},
			wantErr: "database_info.failover.hosts[0].host: required",
		},
		{
			name: "failover host invalid port",
			info: &DatabaseInfo{
				Host: "localhost", Port: 3306, DBName: "plugnmeet",
				Failover: &FailoverSettings{Hosts: []ReplicaDBInfo{{Host: "db2.local", Port: 66000}}},
			},
			wantErr: "database_info.failover.hosts[0].port: port 66000 must be between 1 and 65535",
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			err := ValidateDatabaseInfo(tc.info)
			assertErr(t, err, tc.wantErr)
			if err2 := tc.info.Validate(); (err2 == nil) != (err == nil) {
				t.Fatalf("Validate() and ValidateDatabaseInfo() disagree: %v vs %v", err2, err)
			}
		})
	}
}

func TestValidateRedisInfo(t *testing.T) {
	tt := []struct {
		name    string
		info    *RedisInfo
		wantErr string
	}{
		{
			name: "valid standalone",
			info: &RedisInfo{Host: "localhost:6379"},
		},
		{
			name:    "standalone missing port",
			info:    &RedisInfo{Host: "localhost"},
			wantErr: `redis_info.host: "localhost" must be in host:port format`,
		},
		{
			name:    "standalone empty host",
			info:    &RedisInfo{},
			wantErr: `redis_info.host: "" must be in host:port format`,
		},
		{
			name:    "standalone invalid port",
			info:    &RedisInfo{Host: "localhost:notaport"},
			wantErr: `redis_info.host: "localhost:notaport" has an invalid port`,
		},
		{
			name:    "standalone port out of range",
			info:    &RedisInfo{Host: "localhost:70000"},
			wantErr: "redis_info.host: port 70000 must be between 1 and 65535",
		},
		{
			name: "valid sentinel",
			info: &RedisInfo{MasterName: "mymaster", SentinelAddresses: []string{"10.0.0.1:26379", "10.0.0.2:26379"}},
		},
		{
			name:    "sentinel without master name",
			info:    &RedisInfo{SentinelAddresses: []string{"10.0.0.1:26379"}},
			wantErr: "redis_info.sentinel_master_name: required when sentinel_addresses is set",
		},
		{
			name:    "sentinel invalid address",
			info:    &RedisInfo{MasterName: "mymaster", SentinelAddresses: []string{"10.0.0.1"}},
			wantErr: `redis_info.sentinel_addresses[0]: "10.0.0.1" must be in host:port format`,
		},
		{
			name: "valid cluster",
			info: &RedisInfo{ClusterAddresses: []string{"10.0.0.1:6379", "10.0.0.2:6379"}},
		},
		{
			name:    "cluster invalid address",
			info:    &RedisInfo{ClusterAddresses: []string{"10.0.0.1:6379", "10.0.0.2:abc"}},
			wantErr: `redis_info.cluster_addresses[1]: "10.0.0.2:abc" has an invalid port`,
		},
		{
			name:    "cluster and sentinel conflict",
			info:    &RedisInfo{MasterName: "mymaster", ClusterAddresses: []string{"10.0.0.1:6379"}, SentinelAddresses: []string{"10.0.0.1:26379"}},
			wantErr: "redis_info: cannot set both cluster_addresses and sentinel_addresses",
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			err := ValidateRedisInfo(tc.info)
			assertErr(t, err, tc.wantErr)
			if err2 := tc.info.Validate(); (err2 == nil) != (err == nil) {
				t.Fatalf("Validate() and ValidateRedisInfo() disagree: %v vs %v", err2, err)
			}
		})
	}
}

func assertErr(t *testing.T, err error, want string) {
	t.Helper()

	if want == "" {
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		return
	}
	if err == nil {
		t.Fatalf("expected error %q, got nil", want)
	}
	if err.Error() != want {
		t.Fatalf("expected error %q, got %q", want, err.Error())
	}
}
