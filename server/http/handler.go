package http

import (
	"fmt"
	"net/http"

	"github.com/starudream/go-lib/core/v2/slog"
	"github.com/starudream/go-lib/core/v2/utils/osutil"
	"github.com/starudream/go-lib/server/v2/ierr"
)

type (
	Request        = http.Request
	ResponseWriter = http.ResponseWriter

	Handler     = http.Handler
	HandlerFunc = http.HandlerFunc

	Middleware = func(next Handler) Handler
)

type HandlerCtx func(c *Context) error

func (hc HandlerCtx) ServeHTTP(w ResponseWriter, r *Request) {
	c := NewContext(w, r)
	defer func() { PanicHandler(c, recover()) }()
	ErrorHandler(c, hc(c))
}

var ErrorHandler = func(c *Context, err error) {
	if err == nil {
		return
	}
	e := ierr.FromError(err)
	_ = c.JSON(e.Status(), e)
}

var PanicHandler = func(c *Context, rec any) {
	if rec == nil {
		return
	}
	slog.Error("[%s] %v", osutil.CallerString(3), rec, slog.String("stack", osutil.Stack(2)), slog.GetAttrs(c))
	err, ok := rec.(error)
	if !ok {
		err = fmt.Errorf("%v", rec)
	}
	ErrorHandler(c, err)
}
