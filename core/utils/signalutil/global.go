package signalutil

import (
	"os"
)

var _c = NewContext()

func Done() <-chan struct{} {
	return _c.ctx.Done()
}

func Cancel() {
	_c.cancel()
}

func Signal() os.Signal {
	return _c.sig
}
