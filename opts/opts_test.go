package opts

import (
	"syscall"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
)

func TestConfig_New(t *testing.T) {
	o := New()

	assert.Equal(t, o.Verbose, false)
	assert.Equal(t, o.Env, Development)
	assert.Equal(t, o.KillSignal, syscall.SIGINT)
	assert.Equal(t, o.ReloadSignal, syscall.SIGHUP)
	assert.Equal(t, o.TermSignal, syscall.SIGTERM)
	assert.Nil(t, o.Logger)
}

func TestConfig_WithLogger(t *testing.T) {
	logger, err := zap.NewProduction()
	defer func() { _ = logger.Sync() }()
	assert.NoError(t, err)

	o := New(WithLogger(logger))
	assert.NotNil(t, o.Logger)
}

func TestConfig_WithName(t *testing.T) {
	o := New(WithEnv(Env("fooBar")))
	assert.Equal(t, o.Env, Env("fooBar"))
}
