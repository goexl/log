package core

import (
	"github.com/goexl/gox"
)

type Executor interface {
	Debug(msg string, required gox.Field[any], optionals ...gox.Field[any])

	Info(msg string, required gox.Field[any], optionals ...gox.Field[any])

	Warn(msg string, required gox.Field[any], optionals ...gox.Field[any])

	Error(msg string, required gox.Field[any], optionals ...gox.Field[any])

	Panic(msg string, required gox.Field[any], optionals ...gox.Field[any])

	Fatal(msg string, required gox.Field[any], optionals ...gox.Field[any])

	Sync() error
}
