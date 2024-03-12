package http

import (
	"context"
	"errors"
	"net"
	"net/http"
	"time"

	"github.com/starudream/go-lib/core/v2/slog"
	"github.com/starudream/go-lib/core/v2/utils/optionutil"
	"github.com/starudream/go-lib/server/v2"
)

type Server struct {
	*Mux
	srv *http.Server
}

func NewServer(options ...Option) *Server {
	s := optionutil.Build(&Server{
		Mux: NewMux(),
	}, options...)
	s.srv = &http.Server{Handler: s}
	return s
}

var _ server.Server = (*Server)(nil)

func (s *Server) Start(ln net.Listener) error {
	slog.Info("http server started, listening on %s", ln.Addr())
	return s.srv.Serve(ln)
}

func (s *Server) Stop(timeout time.Duration) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	err := s.srv.Shutdown(ctx)
	if err != nil {
		if errors.Is(err, context.DeadlineExceeded) {
			slog.Warn("http server shutdown timeout")
		} else {
			slog.Error("http server shutdown error: %v", err)
		}
	}
	slog.Info("http server stopped")
}
