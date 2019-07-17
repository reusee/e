package e

import "math/rand"

type MakeErr = func(error, ...interface{}) error

type thrownError struct {
	err error
	sig int64
}

func (t *thrownError) String() string {
	return t.err.Error()
}

func New(
	makeErr MakeErr,
) (
	check func(err error, args ...interface{}),
	catch func(errp *error, args ...interface{}),
) {

	sig := rand.Int63()

	check = func(err error, args ...interface{}) {
		if err != nil {
			if len(args) > 0 {
				err = makeErr(err, args...)
			}
			panic(&thrownError{
				err: err,
				sig: sig,
			})
		}
	}

	catch = func(errp *error, args ...interface{}) {
		if errp == nil {
			return
		}
		if p := recover(); p != nil {
			if e, ok := p.(*thrownError); ok && e.sig == sig {
				if len(args) > 0 {
					e.err = makeErr(e.err, args...)
				}
				*errp = e.err
			} else {
				panic(p)
			}
		}
	}

	return
}
