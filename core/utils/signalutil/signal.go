package signalutil

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/starudream/go-lib/core/v2/internal/logs"
)

type Context struct {
	ctx    context.Context
	cancel context.CancelFunc

	sig os.Signal
}

func NewContext(ctx ...context.Context) *Context {
	if len(ctx) == 0 || ctx[0] == nil {
		ctx = []context.Context{context.Background()}
	}
	c := &Context{}
	c.ctx, c.cancel = context.WithCancel(ctx[0])
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		select {
		case c.sig = <-ch:
			logs.D("receive signal, the process will exit", "signal", c.sig)
			c.cancel()
		case <-c.ctx.Done():
		}
		signal.Stop(ch)
	}()
	return c
}

func (c *Context) Done() <-chan struct{} {
	return c.ctx.Done()
}

func (c *Context) Cancel() {
	c.cancel()
}

func (c *Context) Signal() os.Signal {
	return c.sig
}
