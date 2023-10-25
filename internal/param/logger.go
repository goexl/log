package param

import (
	"github.com/goexl/log/internal/core"
	"github.com/goexl/log/internal/internal"
	"github.com/goexl/log/internal/internal/constant"
)

type Logger struct {
	Level      core.Level
	Skip       int
	Stacktrace int
	Factory    core.Factory
}

func NewLogger() *Logger {
	return &Logger{
		Level:      core.LevelInfo,
		Stacktrace: constant.DefaultStacktrace,
		Factory:    internal.NewFactory(),
	}
}
