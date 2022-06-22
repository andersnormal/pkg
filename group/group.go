package group

import (
	"context"
	"sync"
)

// GroupRun is an abstraction on WaitGroup to run multiple functions concurrently.
// It mimics 'errgroup' to extend structs with functions to run concurrently with
// a root context.
type GroupRun struct {
	wg  sync.WaitGroup
	err *Error
	sync.Mutex
}

// GroupFunc ...
type GroupFunc func(ctx context.Context) error

// Run is creating a new go routine to run a function concurrently.
func (g *GroupRun) Run(ctx context.Context, fn GroupFunc) {
	g.wg.Add(1)

	go func() {
		defer g.wg.Done()

		err := fn(ctx)
		if err != nil {
			g.Lock()
			g.err = Append(g.err, err)
			g.Unlock()
		}
	}()
}

// Wait is waiting for all go routines to finish.
func (g *GroupRun) Wait() error {
	g.wg.Wait()
	g.Lock()
	defer g.Unlock()

	return g.err
}
