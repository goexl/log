package core

import (
	"github.com/goexl/gox"
)

// Logger 日志接口
type Logger interface {
	// Level 现在的日志等级
	Level() Level

	// Enable 开启日志级别
	Enable(level Level)

	// Enabled 日志等级是否开启
	Enabled(level Level) bool

	// Debug 记录调试日志
	Debug(msg string, required gox.Field[any], optionals ...gox.Field[any])

	// Info 记录普通信息日志
	Info(msg string, required gox.Field[any], optionals ...gox.Field[any])

	// Warn 记录警告日志
	Warn(msg string, required gox.Field[any], optionals ...gox.Field[any])

	// Error 记录错误日志
	Error(msg string, required gox.Field[any], optionals ...gox.Field[any])

	// Panic 记录异常日志，程序会退出，可以使用recover机制来阻止程序退出
	Panic(msg string, required gox.Field[any], optionals ...gox.Field[any])

	// Fatal 记录致命错误日志，程序会退出
	Fatal(msg string, required gox.Field[any], optionals ...gox.Field[any])

	// Sync 同步
	Sync() error
}
