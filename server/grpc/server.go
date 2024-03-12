package grpc

import (
	"context"
	"errors"
	"net"
	"reflect"
	"time"

	"google.golang.org/grpc"

	"github.com/starudream/go-lib/core/v2/slog"
	"github.com/starudream/go-lib/core/v2/utils/optionutil"
	"github.com/starudream/go-lib/server/v2"
)

type Server struct {
	srv *grpc.Server

	srvOpts []grpc.ServerOption
}

func NewServer(options ...Option) *Server {
	s := optionutil.Build(&Server{
		srvOpts: []grpc.ServerOption{
			grpc.MaxRecvMsgSize(64 * 1024 * 1024),
		},
	}, options...)
	s.srv = grpc.NewServer(s.srvOpts...)
	return s
}

var _ server.Server = (*Server)(nil)

func (s *Server) Start(ln net.Listener) error {
	slog.Info("grpc server started, listening on %s", ln.Addr())
	return s.srv.Serve(ln)
}

func (s *Server) Stop(timeout time.Duration) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	go func() {
		s.srv.GracefulStop()
		cancel()
	}()
	<-ctx.Done()
	if err := ctx.Err(); errors.Is(err, context.DeadlineExceeded) {
		slog.Warn("grpc server shutdown timeout")
	}
	slog.Info("grpc server stopped")
}

func (s *Server) RegisterServer(fn, svc any) {
	reflect.ValueOf(fn).Call([]reflect.Value{reflect.ValueOf(s.srv), reflect.ValueOf(svc)})
}
