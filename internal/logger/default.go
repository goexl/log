package logger

import (
	"fmt"
	"path/filepath"
	"runtime"
	"sort"
	"strings"

	"github.com/goexl/gox"
	"github.com/goexl/gox/field"
	"github.com/goexl/log"
	"github.com/goexl/simaqian/internal/core"
	"github.com/goexl/simaqian/internal/param"
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

func (d *Default) Enable(lvl core.Level) {
	d.params.Level = lvl
}

func (d *Default) Enabled(lvl core.Level) bool {
	return d.params.Level.Rank() >= lvl.Rank()
}

func (d *Default) Debug(msg string, fields ...gox.Field[any]) {
	if d.params.Level.Rank() <= core.LevelDebug.Rank() {
		d.addCaller(&fields)
		d.executor.Debug(msg, fields...)
	}
}

func (d *Default) Info(msg string, fields ...gox.Field[any]) {
	if d.params.Level.Rank() <= log.LevelInfo.Rank() {
		d.addCaller(&fields)
		d.executor.Info(msg, fields...)
	}
}

func (d *Default) Warn(msg string, fields ...gox.Field[any]) {
	if d.params.Level.Rank() <= log.LevelWarn.Rank() {
		d.addCaller(&fields)
		d.executor.Warn(msg, fields...)
	}
}

func (d *Default) Error(msg string, fields ...gox.Field[any]) {
	if d.params.Level.Rank() <= log.LevelError.Rank() {
		d.addCaller(&fields)
		d.executor.Error(msg, fields...)
	}
}

func (d *Default) Panic(msg string, fields ...gox.Field[any]) {
	if d.params.Level.Rank() > log.LevelPanic.Rank() {
		return
	}

	d.addCaller(&fields)
	d.addStacks(&fields)
	d.executor.Panic(msg, fields...)
}

func (d *Default) Fatal(msg string, fields ...gox.Field[any]) {
	if d.params.Level.Rank() > log.LevelFatal.Rank() {
		return
	}

	d.addCaller(&fields)
	d.addStacks(&fields)
	d.executor.Fatal(msg, fields...)
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
