package grpc

import (
	"google.golang.org/grpc"

	"github.com/starudream/go-lib/core/v2/utils/optionutil"
)

type Option = optionutil.I[Server]

func WithServerOptions(opts ...grpc.ServerOption) Option {
	return optionutil.New(func(s *Server) {
		s.srvOpts = append(s.srvOpts, opts...)
	})
}

func WithUnaryInterceptor(ints ...grpc.UnaryServerInterceptor) Option {
	return optionutil.New(func(s *Server) {
		s.uInts = append(s.uInts, ints...)
	})
}

func WithStreamInterceptor(ints ...grpc.StreamServerInterceptor) Option {
	return optionutil.New(func(s *Server) {
		s.sInts = append(s.sInts, ints...)
	})
}

func WithReflection(t bool) Option {
	return optionutil.New(func(s *Server) {
		s.reflection = t
	})
}
