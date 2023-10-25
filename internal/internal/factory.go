package internal

import (
	"github.com/goexl/simaqian/internal/core"
	"github.com/goexl/simaqian/internal/executor"
)

type Factory struct{}

func NewFactory() *Factory {
	return new(Factory)
}

func (f *Factory) New() (exec core.Executor, err error) {
	exec = executor.NewBuiltin()

	return
}
