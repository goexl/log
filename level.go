package log

import (
	"github.com/goexl/simaqian/internal/core"
)

const (
	LevelDebug = core.LevelDebug
	LevelInfo  = core.LevelInfo
	LevelWarn  = core.LevelWarn
	LevelError = core.LevelError
	LevelPanic = core.LevelPanic
	LevelFatal = core.LevelFatal
)

// Level 日志级别
type Level = core.Level

// ParseLevel 解析日志级别
func ParseLevel(level string) core.Level {
	return core.ParseLevel(level)
}
