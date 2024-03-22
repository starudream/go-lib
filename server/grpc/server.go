package grpc

import (
	"context"
	"errors"
	"net"
	"reflect"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/starudream/go-lib/core/v2/slog"
	"github.com/starudream/go-lib/core/v2/utils/optionutil"
	"github.com/starudream/go-lib/server/v2"
	"github.com/starudream/go-lib/server/v2/otel/otelgrpc"

	"github.com/starudream/go-lib/server/v2/grpc/middlewares/auth"
	"github.com/starudream/go-lib/server/v2/grpc/middlewares/logger"
	"github.com/starudream/go-lib/server/v2/grpc/middlewares/prepare"
	"github.com/starudream/go-lib/server/v2/grpc/middlewares/recovery"
	"github.com/starudream/go-lib/server/v2/grpc/middlewares/validator"
)

type Server struct {
	srv *grpc.Server

	srvOpts []grpc.ServerOption

	uInts []grpc.UnaryServerInterceptor
	sInts []grpc.StreamServerInterceptor

	reflection bool
}

func NewServer(options ...Option) *Server {
	s := optionutil.Build(&Server{
		srvOpts: []grpc.ServerOption{
			grpc.MaxRecvMsgSize(64 * 1024 * 1024),
		},
		uInts: []grpc.UnaryServerInterceptor{
			prepare.Unary(),
			recovery.Unary(),
			validator.Unary(),
			auth.Unary(),
			logger.Unary(),
		},
		reflection: true,
	}, options...)
	s.srvOpts = append(s.srvOpts,
		grpc.ChainUnaryInterceptor(s.uInts...),
		grpc.ChainStreamInterceptor(s.sInts...),
		grpc.StatsHandler(otelgrpc.NewServerHandler()),
	)
	s.srv = grpc.NewServer(s.srvOpts...)
	if s.reflection {
		reflection.Register(s.srv)
	}
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
	err := ctx.Err()
	if err != nil && !errors.Is(err, context.Canceled) {
		if errors.Is(err, context.DeadlineExceeded) {
			slog.Warn("grpc server shutdown timeout")
		} else {
			slog.Error("grpc server shutdown error: %v", err)
		}
	}
	slog.Info("grpc server stopped")
}

func (s *Server) RegisterServer(fn, svc any) {
	reflect.ValueOf(fn).Call([]reflect.Value{reflect.ValueOf(s.srv), reflect.ValueOf(svc)})
}
