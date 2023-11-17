package signalutil

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"reflect"
	"runtime"
	"sync"
	"sync/atomic"
	"syscall"
	"time"

	"github.com/starudream/go-lib/core/v2/internal/logs"
)

type Context struct {
	once sync.Once

	ctx    context.Context
	cancel context.CancelFunc

	sig os.Signal

	done    chan struct{}
	stopped atomic.Bool
	fns     []func()
}

func NewContext(ctx ...context.Context) *Context {
	if len(ctx) == 0 || ctx[0] == nil {
		ctx = []context.Context{context.Background()}
	}
	c := &Context{}
	c.ctx, c.cancel = context.WithCancel(ctx[0])
	c.done = make(chan struct{})
	return c
}

func (c *Context) init() *Context {
	c.once.Do(func() {
		ch := make(chan os.Signal, 1)
		signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
		go func() {
			select {
			case c.sig = <-ch:
				fmt.Printf("\n\n")
				logs.D("receive signal, the process will exit", "signal", c.sig)
				c.cancel()
			case <-c.ctx.Done():
			}
			c.stopped.Store(true)
			// close done after all defer functions are executed
			c.wait()
			close(c.done)
			signal.Stop(ch)
			// force exit because some goroutines may not exit
			os.Exit(0)
		}()
	})
	return c
}

func (c *Context) wait() {
	wg := sync.WaitGroup{}
	wg.Add(len(c.fns))
	for _, fn := range c.fns {
		go func(fn func()) {
			defer func(start time.Time) {
				defer wg.Done()
				took := time.Since(start)
				name := runtime.FuncForPC(reflect.ValueOf(fn).Pointer()).Name()
				logs.D("graceful shutdown", "func", name, "took", took)
			}(time.Now())
			fn()
		}(fn)
	}
	wg.Wait()
}

func (c *Context) Defer(fn func()) *Context {
	if fn == nil || c.stopped.Load() {
		return c
	}
	c.fns = append(c.fns, fn)
	return c
}

func (c *Context) Done() <-chan struct{} {
	return c.init().done
}

func (c *Context) Cancel() {
	c.init().cancel()
}

func (c *Context) Signal() os.Signal {
	return c.init().sig
}
