package server

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"time"

	"github.com/andersnormal/pkg/config"
)

// Env ...
type Env string

const (
	Development Env = "development"
	Production  Env = "production"
)

// ServerError ...
type ServerError struct {
	Err error
}

// Error ...
func (s *ServerError) Error() string { return fmt.Sprintf("server error: %s", s.Err) }

// Unwrap ...
func (s *ServerError) Unwrap() error { return s.Err }

// NewServerError ...
func NewServerError(err error) *ServerError {
	return &ServerError{Err: err}
}

// Server is the interface to be implemented
// to run the server.
//
// 	s, ctx := WithContext(context.Background())
//	s.Listen(listener, false)
//
//	if err := s.Wait(); err != nil {
//		panic(err)
//	}
type Server interface {
	// Run is running a new go routine
	Listen(listener Listener, ready bool)
	// Name
	Name() string
	// Env
	Env() Env
	// Waits for the server to fail,
	// or gracefully shutdown if context is canceled
	Wait() error
}

// Listener is the interface to a listener,
// so starting and shutdown of a listener,
// or any routine.
type Listener interface {
	// Start is being called on the listener
	Start(ctx context.Context, ready func()) func() error
}

// Opt ...
type Opt func(*Opts)

// Opts ...
type Opts struct {
	Name   string
	Env    Env
	Config *config.Config
}

type listeners map[Listener]bool

// server holds the instance info of the server
type server struct {
	ctx    context.Context
	cancel context.CancelFunc

	wg      sync.WaitGroup
	errOnce sync.Once
	err     error

	listeners map[Listener]bool

	ready chan bool
	sys   chan os.Signal

	opts *Opts
}

// WithContext ...
func WithContext(ctx context.Context, opts ...Opt) (Server, context.Context) {
	ctx, cancel := context.WithCancel(ctx)

	// new server
	s := newServer(ctx, opts...)
	s.cancel = cancel
	s.ctx = ctx

	return s, ctx
}

// Name ...
func Name(name string) func(o *Opts) {
	return func(o *Opts) {
		o.Name = name
	}
}

// WithConfig ...
func WithConfig(cfg *config.Config) func(o *Opts) {
	return func(o *Opts) {
		o.Config = cfg
	}
}

func newServer(ctx context.Context, opts ...Opt) *server {
	options := new(Opts)
	options.Config = config.New()
	options.Env = Development

	s := new(server)
	s.opts = options

	ctx, cancel := context.WithCancel(ctx)
	s.cancel = cancel

	s.listeners = make(listeners)
	s.ready = make(chan bool)

	configure(s, opts...)
	configureSignals(s)

	return s
}

// Listen ...
func (s *server) Listen(listener Listener, ready bool) {
	// if we found the listener already, we simply reject to be added
	if _, found := s.listeners[listener]; found {
		return
	}

	// add to listeners
	s.listeners[listener] = ready
}

// Env ...
func (s *server) Env() Env {
	return s.opts.Env
}

// Name ...
func (s *server) Name() string {
	return s.opts.Name
}

// Wait ...
func (s *server) Wait() error {
	// create ticker for interrupt signals
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

OUTTER:
	// start all listeners in order
	for l, ready := range s.listeners {
		fn := func() {
			r := ready

			var readyOnce sync.Once
			readyOnce.Do(func() {
				if r {
					s.ready <- true
				}
			})
		}

		// schedule to routines
		s.run(l.Start(s.ctx, fn))

		// this blocks until ready is called
		if ready {
			select {
			case <-s.ready:
				continue
			case <-s.ctx.Done():
				break OUTTER
			}
		}
	}

	// this is the main loop
	for {
		select {
		case <-ticker.C:
		case <-s.sys:
			// if there is sys interupt
			// cancel the context of the routines
			s.cancel()
		case <-s.ctx.Done():
			if err := s.ctx.Err(); err != nil {
				return NewServerError(s.err)
			}

			return nil
		}
	}
}

func (s *server) run(f func() error) {
	s.wg.Add(1)

	fn := func() {
		defer s.wg.Done()

		if err := f(); err != nil {
			s.errOnce.Do(func() {
				s.err = err
				if s.cancel != nil {
					s.cancel()
				}
			})
		}
	}

	go fn()
}

func configureSignals(s *server) {
	s.sys = make(chan os.Signal, 1)
	signal.Notify(s.sys, s.opts.Config.TermSignal)
}

func configure(s *server, opts ...Opt) error {
	for _, o := range opts {
		o(s.opts)
	}

	return nil
}
