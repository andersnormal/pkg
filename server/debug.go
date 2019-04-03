package server

import (
	"net/http"

	"net/http/pprof"
)

var _ Listener = (*debug)(nil)

type debug struct {
	opts        *DebugOpts
	mux         *http.ServeMux
	handler     *http.Server
	promHandler http.Handler
	proof       bool
}

// DebugOpt ...
type DebugOpt func(*DebugOpts)

// DebugOpts ...
type DebugOpts struct {
	Addr string

	Proof       bool
	PromHandler http.Handler
}

// NewDebugListener ...
func NewDebugListener(opts ...DebugOpt) Listener {
	options := &DebugOpts{}

	d := new(debug)
	d.opts = options
	d.mux = http.NewServeMux()

	configureDebug(d, opts...)
	configureMux(d)

	d.handler = new(http.Server)
	d.handler.Addr = d.opts.Addr
	d.handler.Handler = d.mux

	return d
}

// Start ...
func (d *debug) Start() func() error {
	return func() error {

		err := d.handler.ListenAndServe()

		return err
	}
}

// Stop ...
func (d *debug) Stop() error {
	if d.handler == nil {
		return nil
	}

	return d.handler.Close()
}

// WithStatusAddr is adding this status addr as an option.
func WithStatusAddr(addr string) func(o *DebugOpts) {
	return func(o *DebugOpts) {
		o.Addr = addr
	}
}

// WithProof ...
func WithProof() func(o *DebugOpts) {
	return func(o *DebugOpts) {
		o.Proof = true
	}
}

// WithPrometheusHandler is adding this prometheus http handler as an option.
func WithPrometheusHandler(handler http.Handler) func(o *DebugOpts) {
	return func(o *DebugOpts) {
		o.PromHandler = handler
	}
}

func configureMux(d *debug) error {
	if d.proof {
		d.mux.Handle("/debug/pprof/", http.HandlerFunc(pprof.Index))
		d.mux.Handle("/debug/pprof/cmdline", http.HandlerFunc(pprof.Cmdline))
		d.mux.Handle("/debug/pprof/profile", http.HandlerFunc(pprof.Profile))
		d.mux.Handle("/debug/pprof/symbol", http.HandlerFunc(pprof.Symbol))
		d.mux.Handle("/debug/pprof/trace", http.HandlerFunc(pprof.Trace))
	}

	if d.promHandler != nil {
		d.mux.Handle("/debug/metrics", d.promHandler)
	}

	return nil
}

func configureDebug(d *debug, opts ...DebugOpt) error {
	for _, o := range opts {
		o(d.opts)
	}

	d.proof = d.opts.Proof
	d.promHandler = d.opts.PromHandler

	return nil
}
