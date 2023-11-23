package signalutil

import (
	"context"
	"os"
)

var _c = NewContext()

func Ctx() context.Context {
	return _c.Ctx()
}

func Defer(fn func()) *Context {
	return _c.Defer(fn)
}

// Done must be called last
func Done() <-chan struct{} {
	return _c.Done()
}

func Cancel() {
	_c.Cancel()
}

func Signal() os.Signal {
	return _c.Signal()
}
