package builder

import (
	"github.com/goexl/simaqian/internal/core"
	"github.com/goexl/simaqian/internal/logger"
	"github.com/goexl/simaqian/internal/param"
)

type Logger struct {
	params *param.Logger
}

func NewLogger() *Logger {
	return &Logger{
		params: param.NewLogger(),
	}
}

func (l *Logger) Debug() *Logger {
	return l.Level(core.LevelDebug)
}

func (l *Logger) Info() *Logger {
	return l.Level(core.LevelInfo)
}

func (l *Logger) Warn() *Logger {
	return l.Level(core.LevelWarn)
}

func (l *Logger) Error() *Logger {
	return l.Level(core.LevelError)
}

func (l *Logger) Fatal() *Logger {
	return l.Level(core.LevelFatal)
}

func (l *Logger) Level(lvl core.Level) (logger *Logger) {
	l.params.Level = lvl
	logger = l

	return l
}

func (l *Logger) Skip(skip int) (logger *Logger) {
	l.params.Skip = skip
	logger = l

	return l
}

func (l *Logger) Stacktrace(stacktrace int) (logger *Logger) {
	l.params.Stacktrace = stacktrace
	logger = l

	return l
}

func (l *Logger) Executor(executor core.Executor) (logger *Logger) {
	l.params.Executor = executor
	logger = l

	return l
}

func (l *Logger) Build() core.Logger {
	return logger.NewDefault(l.params)
}
