package core

const (
	LevelDebug Level = "debug"
	LevelInfo  Level = "info"
	LevelWarn  Level = "warn"
	LevelError Level = "error"
	LevelPanic Level = "panic"
	LevelFatal Level = "fatal"
)

type Level string

func ParseLevel(from string) (level Level) {
	switch Level(from) {
	case LevelDebug:
		level = LevelDebug
	case LevelInfo:
		level = LevelInfo
	case LevelWarn:
		level = LevelWarn
	case LevelError:
		level = LevelError
	case LevelPanic:
		level = LevelPanic
	case LevelFatal:
		level = LevelFatal
	}

	return
}

func (l Level) Rank() (rank int) {
	switch l {
	case LevelDebug:
		rank = 10
	case LevelInfo:
		rank = 20
	case LevelWarn:
		rank = 30
	case LevelError:
		rank = 40
	case LevelPanic:
		rank = 50
	case LevelFatal:
		rank = 60
	}

	return
}
