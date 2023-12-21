package hsrv

import (
	"context"
	"errors"
	"net"
	"net/http"
	"sync"
	"time"

	"github.com/starudream/go-lib/core/v2/slog"
	"github.com/starudream/go-lib/core/v2/utils/signalutil"
)

func Default() Router {
	r := NewRouter()
	return r
}

var (
	_r     *Mux
	_rOnce sync.Once
)

func Root() Router {
	_rOnce.Do(func() {
		_r = Default().(*Mux)
	})
	return _r
}

func Run(options ...Option) (err error) {
	opts := newOptions(options...)

	srv := &http.Server{Addr: opts.Addr, Handler: opts.Handler}

	ln, err := net.Listen("tcp", srv.Addr)
	if err != nil {
		return err
	}

	slog.Info("http server listening on %s", ln.Addr())

	var (
		ech   = make(chan struct{}, 1)
		start = func() {
			slog.Info("http server started")
			err = srv.Serve(ln)
			if err != nil && !errors.Is(err, http.ErrServerClosed) {
				close(ech)
			}
		}
		stop = func() {
			ctx, cancel := context.WithTimeout(context.Background(), time.Second)
			defer cancel()
			err = srv.Shutdown(ctx)
			if err != nil {
				if errors.Is(err, context.DeadlineExceeded) {
					slog.Warn("http server shutdown forced")
				} else {
					slog.Error("http server shutdown error: %v", err)
				}
			}
			slog.Info("http server stopped")
		}
	)

	go start()

	select {
	case <-ech:
		return err
	case <-signalutil.Defer(stop).Done():
	}

	return nil
}
