package e

import "math/rand"

type ErrFunc func(error, ...interface{}) error

type thrownError struct {
	err error
	sig int64
}

func (t thrownError) String() string {
	return t.err.Error()
}

func New(
	newErr ErrFunc,
) (
	makeErr func(prev error, args ...interface{}) error,
	wrapErr func(err error, args ...interface{}) error,
	check func(err error),
	handle func(errp *error),
) {

	sig := rand.Int63()

	check = func(err error) {
		if err != nil {
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

	makeErr = func(prev error, args ...interface{}) error {
		return newErr(prev, args...)
	}

	wrapErr = func(err error, args ...interface{}) error {
		if err == nil {
			return nil
		}
		return newErr(err, args...)
	}

	return
}
