package group

import (
	"context"
	"sync"
)

// GroupRun is an abstraction on WaitGroup to run multiple functions in parallel.
// It mimics 'errgroup' to extend structs with functions to run in parallel.
type GroupRun struct {
	wg sync.WaitGroup
}

// ResolverFunc ...
type GroupFunc func(ctx context.Context) error

// Run is creating a new go routine to run a function in parallel.
func (g *GroupRun) Run(ctx context.Context, fn GroupFunc) {
	g.wg.Add(1)

	go func() {
		defer g.wg.Done()

		err := fn(ctx)
		if err != nil {
			return // ignoring errors for now
		}
	}()
}

// Wait is waiting for all go routines to finish.
func (g *GroupRun) Wait() {
	g.wg.Wait()
}
