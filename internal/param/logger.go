package param

import (
	"github.com/goexl/simaqian/internal/core"
	"github.com/goexl/simaqian/internal/executor"
	"github.com/goexl/simaqian/internal/internal/constant"
)

type Logger struct {
	Level      core.Level
	Skip       int
	Stacktrace int
	Executor   core.Executor
}

func NewLogger() *Logger {
	return &Logger{
		Level:      core.LevelInfo,
		Stacktrace: constant.DefaultStacktrace,
		Executor:   executor.NewBuiltin(),
	}
}
