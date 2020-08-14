package e

import (
	"errors"
	"fmt"
	"os"
	"strings"
	"testing"
)

func TestFunc(t *testing.T) {
	me := Default.WithStack().WithName("e_test")
	str := me(nil, "foo").Error()
	if !strings.HasPrefix(str, "e_test: foo (") {
		t.Fatal()
	}
}

func TestErrIs(t *testing.T) {
	me := Default
	err := fmt.Errorf("foo")
	e := me(err, "wrap")
	e = me(e, "wrap")
	e = me(e, "wrap")
	if !errors.Is(e, err) {
		t.Fatal()
	}
}

func TestErrAs(t *testing.T) {
	pathErr := &os.PathError{
		Path: "foo",
	}
	me := Default
	e := me(pathErr, "wrap")
	e = me(e, "wrap")
	e = me(e, "wrap")
	var err *os.PathError
	if !errors.As(e, &err) {
		t.Fatal()
	}
	if err.Path != "foo" {
		t.Fatal()
	}
}
