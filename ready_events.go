package pkg

import (
	"errors"
	"reflect"
	"time"
)

var _ ReadyEvents = (*readyEvents)(nil)

var (
	ErrReadyEventsTimeout = errors.New("pkg: ready events timeout")
	ErrReadyEventsDone    = errors.New("pkg: ready events not done")
)

type ReadyEvents interface {
	// Register is registering an event to occur upon readyness
	// of an app.
	Register(event interface{}) error
	// Wait blocks until all events have occured or when
	// the timeout is hit.
	Wait() error
	// Ready is submitting an event to be done
}

func NewReadyEvents(d time.Duration) *readyEvents {
	return &readyEvents{
		timeout: d,
	}
}

type readyEvents struct {
	events  []interface{}
	ready   chan interface{}
	done    chan int
	timeout time.Duration
}

func (r *readyEvents) Register(event interface{}) error {
	r.events = append(r.events, event)

	// noop
	return nil
}

// Ready allows to send in an event to be marked as ready.
func (r *readyEvents) Ready(event interface{}) {
	r.ready <- event
}

// Wait blocks until all events have occured or when
// the timeout is hit.
func (r *readyEvents) Wait() error {
	for {
		select {
		case event := <-r.ready:
			for i := range r.events {
				if reflect.TypeOf(r.events[i]) == reflect.TypeOf(event) {
					r.events = append(r.events[:i], r.events[i+1:]...)
				}
			}

			if len(r.events) == 0 {
				r.done <- 1
			}
		case done := <-r.done:
			if done != 0 {
				return ErrReadyEventsDone
			}

			return nil
		case <-time.After(r.timeout):
			return ErrReadyEventsTimeout
		}
	}
}
