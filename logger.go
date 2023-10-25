package log

import (
	"github.com/goexl/log/internal/builder"
	"github.com/goexl/log/internal/core"
)

// New 创建
func New() *builder.Logger {
	return builder.NewLogger()
}

// Logger 日志接口
type Logger = core.Logger
