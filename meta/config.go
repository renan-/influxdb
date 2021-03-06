package meta

import (
	"time"

	"github.com/influxdb/influxdb/toml"
)

const (
	// DefaultBindAddress is the default address to bind to.
	DefaultBindAddress = ":8088"

	// DefaultHeartbeatTimeout is the default heartbeat timeout for the store.
	DefaultHeartbeatTimeout = 1000 * time.Millisecond

	// DefaultElectionTimeout is the default election timeout for the store.
	DefaultElectionTimeout = 1000 * time.Millisecond

	// DefaultLeaderLeaseTimeout is the default leader lease for the store.
	DefaultLeaderLeaseTimeout = 500 * time.Millisecond

	// DefaultCommitTimeout is the default commit timeout for the store.
	DefaultCommitTimeout = 50 * time.Millisecond
)

// Config represents the meta configuration.
type Config struct {
	Hostname           string        `toml:"-"`
	Dir                string        `toml:"dir"`
	BindAddress        string        `toml:"bind-address"`
	ElectionTimeout    toml.Duration `toml:"election-timeout"`
	HeartbeatTimeout   toml.Duration `toml:"heartbeat-timeout"`
	LeaderLeaseTimeout toml.Duration `toml:"leader-lease-timeout"`
	CommitTimeout      toml.Duration `toml:"commit-timeout"`
}

func NewConfig() Config {
	return Config{
		BindAddress:        DefaultBindAddress,
		ElectionTimeout:    toml.Duration(DefaultElectionTimeout),
		HeartbeatTimeout:   toml.Duration(DefaultHeartbeatTimeout),
		LeaderLeaseTimeout: toml.Duration(DefaultLeaderLeaseTimeout),
		CommitTimeout:      toml.Duration(DefaultCommitTimeout),
	}
}
