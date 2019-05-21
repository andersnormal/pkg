package watchcat

import (
	"context"
	"time"

	ws "github.com/gorilla/websocket"
)

var _ Watchcat = (*watchcat)(nil)

// New ...
func New(opts ...Opt) Watchcat {
	options := new(Opts)

	w := new(watchcat)
	w.opts = options

	configure(w, opts...)

	return w
}

// Start ...
func (w *watchcat) Start(ctx context.Context, ready func()) func() error {
	return func() error {
		if err := w.http.ListenAndServe(); err != nil {
			return err
		}

		return nil
	}
}

// Stop ...
func (w *watchcat) Stop() error {
	if w.http != nil {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		if err := w.http.Shutdown(ctx); err != nil {
			return err
		}
	}

	return nil
}

var upgrader = ws.Upgrader{}

func configure(w *watchcat, opts ...Opt) error {
	for _, o := range opts {
		o(w.opts)
	}

	w.addr = w.opts.Addr

	return nil
}
