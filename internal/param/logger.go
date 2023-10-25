package param

import (
	"github.com/goexl/simaqian/internal/core"
	"github.com/goexl/simaqian/internal/internal"
	"github.com/goexl/simaqian/internal/internal/constant"
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
