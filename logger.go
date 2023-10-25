package log

import (
	"github.com/goexl/simaqian/internal/builder"
	"github.com/goexl/simaqian/internal/core"
)

// New 创建
func New() *builder.Logger {
	return builder.NewLogger()
}

// Logger 日志接口
type Logger = core.Logger
