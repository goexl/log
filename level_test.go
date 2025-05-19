package log_test

import (
	"testing"

	"github.com/goexl/log"
	"github.com/goexl/log/internal/core"
	"github.com/stretchr/testify/assert"
)

func TestLevel(t *testing.T) {
	assert.Equal(t, log.LevelDebug, core.LevelDebug)
	assert.Equal(t, log.LevelInfo, core.LevelInfo)
	assert.Equal(t, log.LevelWarn, core.LevelWarn)
	assert.Equal(t, log.LevelError, core.LevelError)
	assert.Equal(t, log.LevelPanic, core.LevelPanic)
	assert.Equal(t, log.LevelFatal, core.LevelFatal)
}

func TestParseLevel(t *testing.T) {
	testcases := []struct {
		in       string
		expected log.Level
	}{
		{in: "debug", expected: log.LevelDebug},
		{in: "info", expected: log.LevelInfo},
		{in: "warn", expected: log.LevelWarn},
		{in: "error", expected: log.LevelError},
		{in: "panic", expected: log.LevelPanic},
		{in: "fatal", expected: log.LevelFatal},
	}

	for _, test := range testcases {
		got := log.ParseLevel(test.in)
		assert.Equal(t, test.expected, got)
	}
}
