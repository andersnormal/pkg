package pkg

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestReadyEvents(t *testing.T) {
	tests := []struct {
		desc    string
		events  []interface{}
		timeout time.Duration
		err     error
	}{
		{
			desc:    "should not return error with no events",
			events:  []interface{}{},
			timeout: 60 * time.Second,
			err:     nil,
		},
		{
			desc: "should not return error with events in time",
			events: []interface{}{
				struct {
					finish bool
				}{
					finish: true,
				},
			},
			timeout: 60 * time.Second,
			err:     nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			readyEvents := NewReadyEvents(tt.timeout)

			for i := range tt.events {
				readyEvents.Register(tt.events[i])
			}

			go func(events []interface{}) {
				for i := range events {
					readyEvents.Ready(tt.events[i])
				}
			}(tt.events)

			err := readyEvents.Wait()

			if tt.err != nil {
				return
			}

			assert.NoError(t, err)
		})
	}
}
