package e

import (
	"io"
	"os"
)

var (
	testMakeErr           = Default.WithStack().WithName("e_test")
	testCheck, testHandle = New(testMakeErr)
)

func ExampleCopyFile() {

	copyFile := func(src string, dest string) (err error) {
		defer testHandle(&err, "copy %s to %s", src, dest)

		r, err := os.Open(src)
		testCheck(err, "open source file: %s", src)
		defer r.Close()

		w, err := os.Create(dest)
		testCheck(err, "create dest file: %s", dest)
		defer func() {
			w.Close()
			if err != nil {
				os.Remove(dest)
			}
		}()

		_, err = io.Copy(w, r)
		testCheck(err, "copy")
		testCheck(w.Close())

		return nil
	}

	_ = copyFile

	// Output:
}
