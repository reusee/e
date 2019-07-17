package e

import (
	"fmt"
	"testing"
)

func BenchmarkOverhead(b *testing.B) {
	ce, he := New(Default)
	e := fmt.Errorf("foo")
	fn := func() (err error) {
		defer he(&err)
		ce(e)
		return
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		fn()
	}
}

func BenchmarkOverheadNoError(b *testing.B) {
	ce, he := New(Default)
	fn := func() (err error) {
		defer he(&err)
		ce(nil)
		return
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		fn()
	}
}
