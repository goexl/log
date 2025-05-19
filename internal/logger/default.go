package logger

import (
	"fmt"
	"path/filepath"
	"runtime"
	"sort"
	"strings"

	"github.com/goexl/gox"
	"github.com/goexl/gox/field"
	"github.com/goexl/log/internal/core"
	"github.com/goexl/log/internal/param"
)

var _ core.Logger = (*Default)(nil)

type Default struct {
	params   *param.Logger
	executor core.Executor
}

func NewDefault(params *param.Logger, executor core.Executor) *Default {
	return &Default{
		params:   params,
		executor: executor,
	}
}

func (d *Default) Level() core.Level {
	return d.params.Level
}

func (d *Default) Enable(level core.Level) {
	d.params.Level = level
}

func (d *Default) Enabled(level core.Level) bool {
	return d.params.Level.Rank() >= level.Rank()
}

func (d *Default) Debug(msg string, required gox.Field[any], optionals ...gox.Field[any]) {
	if d.params.Level.Rank() <= core.LevelDebug.Rank() {
		d.addCaller(&optionals)
		d.executor.Debug(msg, required, optionals...)
	}
}

func (d *Default) Info(msg string, required gox.Field[any], optionals ...gox.Field[any]) {
	if d.params.Level.Rank() <= core.LevelInfo.Rank() {
		d.addCaller(&optionals)
		d.executor.Info(msg, required, optionals...)
	}
}

func (d *Default) Warn(msg string, required gox.Field[any], optionals ...gox.Field[any]) {
	if d.params.Level.Rank() <= core.LevelWarn.Rank() {
		d.addCaller(&optionals)
		d.executor.Warn(msg, required, optionals...)
	}
}

func (d *Default) Error(msg string, required gox.Field[any], optionals ...gox.Field[any]) {
	if d.params.Level.Rank() <= core.LevelError.Rank() {
		d.addCaller(&optionals)
		d.executor.Error(msg, required, optionals...)
	}
}

func (d *Default) Panic(msg string, required gox.Field[any], optionals ...gox.Field[any]) {
	if d.params.Level.Rank() > core.LevelPanic.Rank() {
		return
	}

	d.addCaller(&optionals)
	d.addStacks(&optionals)
	d.executor.Panic(msg, required, optionals...)
}

func (d *Default) Fatal(msg string, required gox.Field[any], optionals ...gox.Field[any]) {
	if d.params.Level.Rank() > core.LevelFatal.Rank() {
		return
	}

	d.addCaller(&optionals)
	d.addStacks(&optionals)
	d.executor.Fatal(msg, required, optionals...)
}

func (d *Default) Sync() error {
	return d.executor.Sync()
}

func (d *Default) addCaller(fields *[]gox.Field[any]) {
	if _, file, no, ok := runtime.Caller(2 + d.params.Skip); ok { // ! 默认封闭的时候加了2层调用栈
		caller := fmt.Sprintf("%s:%d", filepath.Base(file), no)
		*fields = append(*fields, field.New("caller", caller))
	}
}

func (d *Default) addStacks(fields *[]gox.Field[any]) {
	callers := make([]uintptr, d.params.Stacktrace)
	count := runtime.Callers(2+d.params.Skip, callers) // ! 默认封闭的时候加了2层调用栈
	frames := runtime.CallersFrames(callers[:count])

	stacks := make([]string, 0)
	for {
		frame, more := frames.Next()
		stacks = append(stacks, fmt.Sprintf("%s[%s]:%d", filepath.Base(frame.File), frame.Function, frame.Line))
		if !more {
			break
		}
	}
	sort.SliceStable(stacks, d.swap)
	*fields = append(*fields, field.New("stacks", strings.Join(stacks, " -> ")))
}

func (d *Default) swap(_, _ int) bool {
	return true
}
