package e

import "math/rand"

type MakeErr = func(error, ...interface{}) error

type thrownError struct {
	err error
	sig int64
}

func (t thrownError) String() string {
	return t.err.Error()
}

func New(
	makeErr MakeErr,
) (
	check func(err error, args ...interface{}),
	handle func(errp *error),
) {

	sig := rand.Int63()

	check = func(err error, args ...interface{}) {
		if err != nil {
			if len(args) > 0 {
				err = makeErr(err, args...)
			}
			panic(thrownError{
				err: err,
				sig: sig,
			})
		}
	}

	handle = func(errp *error) {
		if errp == nil {
			return
		}
		if p := recover(); p != nil {
			if e, ok := p.(thrownError); ok && e.sig == sig {
				*errp = e.err
			} else {
				panic(p)
			}
		}
	}

	return
}
