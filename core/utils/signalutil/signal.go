package signalutil

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"reflect"
	"runtime"
	"sync"
	"syscall"
	"time"

	"github.com/starudream/go-lib/core/v2/utils/osutil"
	"github.com/starudream/go-lib/core/v2/utils/sliceutil"

	"github.com/starudream/go-lib/core/v2/internal/logs"
)

type Context struct {
	once sync.Once

	ctx    context.Context
	cancel context.CancelFunc
	done   chan struct{}

	sig os.Signal

	fns sliceutil.SyncSlice[func()]
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

func (c *Context) Ctx() context.Context {
	return c.ctx
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
				logs.D("user cancel, the process will exit")
			}
			// add interval between Defer() and Done()
			time.Sleep(50 * time.Millisecond)
			// close done after all defer functions are executed
			c.wait()
			// all done
			close(c.done)
			signal.Stop(ch)
			// force exit
			go func() {
				time.Sleep(time.Second)
				logs.D("something still running after 1s, force exit")
				if osutil.ArgTest() {
					os.Exit(1)
				} else {
					os.Exit(0)
				}
			}()
		}()
	})
	return c
}

func (c *Context) wait() {
	wg := sync.WaitGroup{}
	wg.Add(1)
	c.fns.Range(func(i int, fn func()) bool {
		wg.Add(1)
		go func() {
			defer func(start time.Time) {
				defer wg.Done()
				logs.D("graceful shutdown", "func", runtime.FuncForPC(reflect.ValueOf(fn).Pointer()).Name(), "took", time.Since(start))
			}(time.Now())
			fn()
		}()
		return true
	})
	wg.Done()
	wg.Wait()
}

func (c *Context) Defer(fn func()) *Context {
	if fn == nil {
		return c
	}
	c.fns.Append(fn)
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
