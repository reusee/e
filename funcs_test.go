package e

import (
	"strings"
	"testing"
)

func TestFunc(t *testing.T) {
	me, _, _, _ := New(Default.WithStack().WithName("e_test"))
	str := me(nil, "foo").Error()
	if !strings.HasPrefix(str, "e_test: foo (") {
		t.Fatal()
	}
}
