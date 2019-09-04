package config

import (
	"syscall"
)

// Config ...
type Config struct {
	// Verbose toggles the verbosity
	Verbose bool
	// LogLevel is the level with with to log for this config
	LogLevel string `mapstructure:"log_level"`
	// LogFormat is the format that is used for logging
	LogFormat string `mapstructure:"log_format"`
	// ReloadSignal ...
	ReloadSignal syscall.Signal
	// TermSignal ...
	TermSignal syscall.Signal
	// KillSignal ...
	KillSignal syscall.Signal
}

// Configer ...
type Configer interface {
	Configure(opts ...Opt)
}

// Opt ...
type Opt func(*Config)

// Configure ...
func (c *Config) Configure(opts ...Opt) {
	for _, o := range opts {
		o(c)
	}
}

const (
	// DefaultLogLevel is the default logging level.
	DefaultLogLevel = "warn"
	// DefaultLogFormat is the default format of the logger
	DefaultLogFormat = "text"
	// DefaultTermSignal is the signal to term the agent.
	DefaultTermSignal = syscall.SIGTERM
	// DefaultReloadSignal is the default signal for reload.
	DefaultReloadSignal = syscall.SIGHUP
	// DefaultKillSignal is the default signal for termination.
	DefaultKillSignal = syscall.SIGINT
	// DefaultVerbose is the default verbuse status.
	DefaultVerbose = false
)

// New returns a new Config
func New(opts ...Opt) *Config {
	c := &Config{
		LogLevel:     DefaultLogLevel,
		LogFormat:    DefaultLogFormat,
		ReloadSignal: DefaultReloadSignal,
		TermSignal:   DefaultTermSignal,
		KillSignal:   DefaultKillSignal,
		Verbose:      DefaultVerbose,
	}

	c.Configure(opts...)

	return c
}
