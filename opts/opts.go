package opts

import (
	"syscall"

	"go.uber.org/zap"
)

const (
	// DefaultVerbose ...
	DefaultVerbose = false
	// DefaultTermSignal is the signal to term the agent.
	DefaultTermSignal = syscall.SIGTERM
	// DefaultReloadSignal is the default signal for reload.
	DefaultReloadSignal = syscall.SIGHUP
	// DefaultKillSignal is the default signal for termination.
	DefaultKillSignal = syscall.SIGINT
)

// Env is a string that identifies an environment.
type Env string

const (
	// Development symbolizes a development environment this program runs in.
	Development Env = "development"
	// Production symbolizes a production environment this program runs in.
	Production Env = "production"
)

// Opts ...
type Opts struct {
	// Verbose toggles verbosity
	Verbose bool
	// ReloadSignal ...
	ReloadSignal syscall.Signal
	// TermSignal ...
	TermSignal syscall.Signal
	// KillSignal ...
	KillSignal syscall.Signal
	// Logger ...
	Logger *zap.Logger
	// Env ...
	Env Env
}

// Opt is an option
type Opt func(*Opts)

// New returns a new instance of the options.
func New(opts ...Opt) *Opts {
	o := NewDefaultOpts()
	o.Configure(opts...)

	return o
}

// NewDefaultOpts returns options with a default configuration.
func NewDefaultOpts() *Opts {
	return &Opts{
		Verbose:      DefaultVerbose,
		ReloadSignal: DefaultReloadSignal,
		TermSignal:   DefaultTermSignal,
		KillSignal:   DefaultKillSignal,
		Env:          Development,
	}
}

// WithLogger is setting a logger for options
func WithLogger(logger *zap.Logger) Opt {
	return func(opts *Opts) {
		opts.Logger = logger
	}
}

// WithEnv configures a new environment.
func WithEnv(env Env) Opt {
	return func(opts *Opts) {
		opts.Env = env
	}
}

// Configure os configuring the options.
func (s *Opts) Configure(opts ...Opt) {
	for _, o := range opts {
		o(s)
	}
}
