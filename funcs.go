package e

import (
	"fmt"

	"errors"
)

type Err struct {
	Pkg  string
	Info string
	Prev error
}

func (e Err) Error() string {
	if e.Prev == nil {
		return fmt.Sprintf("%s: %s", e.Pkg, e.Info)
	}
	return fmt.Sprintf("%s: %s\n%v", e.Pkg, e.Info, e.Prev)
}

var _ errors.Wrapper = Err{}

func (e Err) Unwrap() error {
	return e.Prev
}

var _ errors.Formatter = Err{}

func (e Err) FormatError(p errors.Printer) error {
	p.Printf("%s: %s", e.Pkg, e.Info)
	return e.Prev
}

func WithPackage(
	pkg string,
) ErrFunc {
	return func(err error, args ...interface{}) error {
		format := args[0].(string)
		args = args[1:]
		if len(args) > 0 {
			return Err{
				Pkg:  pkg,
				Info: fmt.Sprintf(format, args...),
				Prev: err,
			}
		}
		return Err{
			Pkg:  pkg,
			Info: format,
			Prev: err,
		}
	}
}
