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

// Env ...
type Env string

const (
	Development Env = "development"
	Production  Env = "production"
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

// Opt ...
type Opt func(*Opts)

// New ...
func New(opts ...Opt) *Opts {
	o := NewDefaultOpts()
	o.Configure(opts...)

	return o
}

// DefaultOpts ...
func NewDefaultOpts() *Opts {
	return &Opts{
		Verbose:      DefaultVerbose,
		ReloadSignal: DefaultReloadSignal,
		TermSignal:   DefaultTermSignal,
		KillSignal:   DefaultKillSignal,
		Env:          Development,
	}
}

// WithLogger ...
func WithLogger(logger *zap.Logger) Opt {
	return func(opts *Opts) {
		opts.Logger = logger
	}
}

// WithEnv ...
func WithEnv(env Env) Opt {
	return func(opts *Opts) {
		opts.Env = env
	}
}

// Configure ...
func (s *Opts) Configure(opts ...Opt) error {
	for _, o := range opts {
		o(s)
	}

	return nil
}
