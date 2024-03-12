package hggw

import (
	"google.golang.org/grpc"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"

	"github.com/starudream/go-lib/core/v2/utils/optionutil"
)

type Option = optionutil.I[Server]

func WithMuxOptions(muxOpts ...runtime.ServeMuxOption) Option {
	return optionutil.New(func(s *Server) {
		s.muxOpts = append(s.muxOpts, muxOpts...)
	})
}

func WithDialOptions(dialOpts ...grpc.DialOption) Option {
	return optionutil.New(func(s *Server) {
		s.dialOpts = append(s.dialOpts, dialOpts...)
	})
}
