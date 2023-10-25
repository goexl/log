package builder

import (
	"github.com/goexl/log/internal/core"
	"github.com/goexl/log/internal/logger"
	"github.com/goexl/log/internal/param"
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

func (l *Logger) Factory(factory core.Factory) (logger *Logger) {
	l.params.Factory = factory
	logger = l

	return l
}

func (l *Logger) Build() (log core.Logger, err error) {
	if executor, ne := l.params.Factory.New(); nil != ne {
		err = ne
	} else {
		log = logger.NewDefault(l.params, executor)
	}

	return
}

func (l *Logger) Apply() (logger core.Logger) {
	if created, be := l.Build(); nil != be {
		panic(be)
	} else {
		logger = created
	}

	return
}
