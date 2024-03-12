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
