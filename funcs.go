package e

import (
	"fmt"
	"runtime"
	"strconv"
	"strings"
)

type Err struct {
	Pkg  string
	Info string
	Prev error
	File string
	Line int
}

func (e Err) Error() string {
	var b strings.Builder
	if e.Pkg != "" {
		b.WriteString(e.Pkg)
		b.WriteString(":")
	}
	if e.Info != "" {
		b.WriteString(" ")
		b.WriteString(e.Info)
	}
	if e.File != "" {
		b.WriteString(" (")
		b.WriteString(e.File)
		b.WriteString(":")
		b.WriteString(strconv.Itoa(e.Line))
		b.WriteString(")")
	}
	if e.Prev != nil {
		b.WriteString("\n")
		b.WriteString(e.Prev.Error())
	}
	return b.String()
}

type Func func(error, ...interface{}) error

var Default = Func(func(err error, args ...interface{}) error {
	format := args[0].(string)
	args = args[1:]
	if len(args) > 0 {
		return Err{
			Info: fmt.Sprintf(format, args...),
			Prev: err,
		}
	}
	return Err{
		Info: format,
		Prev: err,
	}
})

func (f Func) WithName(name string) Func {
	return func(err error, args ...interface{}) error {
		e := f(err, args...)
		er := e.(Err)
		er.Pkg = name
		return er
	}
}

func (f Func) WithStack() Func {
	return func(err error, args ...interface{}) error {
		e := f(err, args...)
		er := e.(Err)
		pcs := make([]uintptr, 8)
		skip := 1
	loop_callers:
		for {
			n := runtime.Callers(skip, pcs)
			if n == 0 {
				break
			}
			frames := runtime.CallersFrames(pcs[:n])
			for {
				frame, more := frames.Next()
				if !strings.HasSuffix(frame.File, "funcs.go") &&
					!strings.HasSuffix(frame.File, "error.go") {
					er.File = frame.File
					er.Line = frame.Line
					break loop_callers
				}
				if !more {
					break
				}
			}
			skip += n
		}
		return er
	}
}

func WithPackage(
	pkg string,
) Func {
	return Default.WithName(pkg).WithStack()
}
