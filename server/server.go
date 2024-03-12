package server

import (
	"net"
	"sync"
	"time"

	"github.com/soheilhy/cmux"
	"golang.org/x/sync/errgroup"

	"github.com/starudream/go-lib/core/v2/gh"
	"github.com/starudream/go-lib/core/v2/slog"
	"github.com/starudream/go-lib/core/v2/utils/signalutil"
)

type Server interface {
	Start(ln net.Listener) error
	Stop(timeout time.Duration)
}

func Run(address string, options ...Option) error {
	opts := newOptions(options...)

	ln, err := net.Listen("tcp", address)
	if err != nil {
		return err
	}

	var (
		cm    = cmux.New(ln)
		eg, _ = errgroup.WithContext(signalutil.Ctx())
		start = func() {
			if opts.gs != nil {
				gl := cm.MatchWithWriters(cmux.HTTP2MatchHeaderFieldSendSettings("content-type", "application/grpc"))
				eg.Go(func() error { return opts.gs.Start(gl) })
			}
			if opts.hs != nil {
				hl := cm.Match(cmux.HTTP1Fast("PATCH"))
				eg.Go(func() error { return opts.hs.Start(hl) })
			}
			eg.Go(func() error { return cm.Serve() })
			if e := eg.Wait(); e != nil {
				slog.Fatal("start server error: %v", e)
				signalutil.Cancel()
			}
		}
		stop = func() {
			wg := sync.WaitGroup{}
			wg.Add(2)
			go func() {
				defer wg.Done()
				if opts.gs != nil {
					opts.gs.Stop(opts.timeout)
				}
			}()
			go func() {
				defer wg.Done()
				if opts.hs != nil {
					opts.hs.Stop(opts.timeout)
				}
			}()
			wg.Wait()
			gh.Close(cm, ln)
		}
	)

	go start()

	<-signalutil.Defer(stop).Done()

	return nil
}
