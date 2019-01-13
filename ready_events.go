package pkg

import (
	"errors"
	"reflect"
	"sync"
	"time"
)

var _ ReadyEvents = (*readyEvents)(nil)

var (
	ErrReadyEventsTimeout = errors.New("pkg: ready events timeout")
	ErrReadyEventsDone    = errors.New("pkg: ready events not done")
)

// ReadyEvent is an event that should be done
// before proceeding.
//
//	type CustomReadyEvent  {
//		*ReadyEvent
//	}
//
type ReadyEvent struct{}

type ReadyEvents interface {
	// Register is registering an event to occur upon readyness
	// of an app.
	Register(event interface{}) error
	// Wait blocks until all events have occured or when
	// the timeout is hit.
	Wait() error
	// Ready is submitting an event to be done
}

//  NewReadyEvents returns a new ready events factory
func NewReadyEvents(d time.Duration) *readyEvents {
	return &readyEvents{
		timeout: d,
		ready:   make(chan interface{}),
	}
}

type readyEvents struct {
	events  []interface{}
	ready   chan interface{}
	timeout time.Duration

	sync.RWMutex
}

// Register is pushing a new event to the splice of ready events
func (r *readyEvents) Register(event interface{}) error {
	r.Lock()
	defer r.Unlock()

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
	if len(r.events) == 0 {
		return nil
	}

	for {
		select {
		case event := <-r.ready:
			for i := range r.events {
				if reflect.TypeOf(r.events[i]) == reflect.TypeOf(event) {
					r.events = append(r.events[:i], r.events[i+1:]...)
				}
			}

			if len(r.events) == 0 {
				return nil
			}

		case <-time.After(time.Second * 3):
			return ErrReadyEventsTimeout
		}
	}
}
