package executor

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/goexl/gox"
	"github.com/goexl/log/internal/core"
)

var _ core.Executor = (*Builtin)(nil)

type Builtin struct {
	logger *log.Logger
}

func NewBuiltin() *Builtin {
	return &Builtin{
		logger: log.Default(),
	}
}

func (b *Builtin) Debug(msg string, required gox.Field[any], optionals ...gox.Field[any]) {
	b.logger.Println(b.parse(core.LevelDebug, msg, required, optionals...))
}

func (b *Builtin) Info(msg string, required gox.Field[any], optionals ...gox.Field[any]) {
	b.logger.Println(b.parse(core.LevelInfo, msg, required, optionals...))
}

func (b *Builtin) Warn(msg string, required gox.Field[any], optionals ...gox.Field[any]) {
	b.logger.Println(b.parse(core.LevelWarn, msg, required, optionals...))
}

func (b *Builtin) Error(msg string, required gox.Field[any], optionals ...gox.Field[any]) {
	b.logger.Println(b.parse(core.LevelError, msg, required, optionals...))
}

func (b *Builtin) Panic(msg string, required gox.Field[any], optionals ...gox.Field[any]) {
	b.logger.Println(b.parse(core.LevelPanic, msg, required, optionals...))
}

func (b *Builtin) Fatal(msg string, required gox.Field[any], optionals ...gox.Field[any]) {
	b.logger.Println(b.parse(core.LevelFatal, msg, required, optionals...))
}

func (b *Builtin) Sync() (err error) {
	return
}

func (b *Builtin) parse(level core.Level, msg string, required gox.Field[any], optionals ...gox.Field[any]) (args any) { // nolint:lll
	data := make(map[string]any, len(optionals)+3)
	data[required.Key()] = required.Value()
	for _, field := range optionals {
		data[field.Key()] = field.Value()
	}
	if bytes, me := json.Marshal(data); nil != me {
		args = b.string(level, msg, required, optionals...)
	} else {
		args = gox.StringBuilder(`{"level":`, level, `,"message":"`, msg, `",`, string(bytes[1:])).String()
	}

	return
}

func (b *Builtin) string(level core.Level, msg string, required gox.Field[any], optionals ...gox.Field[any]) (args []any) { // nolint:lll
	args = make([]any, 0, len(optionals)+2)
	args = append(args, level)
	args = append(args, msg)

	fields := append([]gox.Field[any]{required}, optionals...)
	if 0 != len(fields) {
		args = append(args, "[")
		for _, field := range fields {
			args = append(args, fmt.Sprintf("{%s = %v}", field.Key(), field.Value()))
		}
		args = append(args, "]")
	}

	return
}
