package hggw

import (
	"google.golang.org/grpc"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"

	"github.com/starudream/go-lib/core/v2/utils/optionutil"
)

type Option = optionutil.I[Server]

func WithMuxOptions(muxOpts ...runtime.ServeMuxOption) Option {
	return optionutil.New(func(opts *Server) {
		opts.muxOpts = append(opts.muxOpts, muxOpts...)
	})
}

func WithDialOptions(dialOpts ...grpc.DialOption) Option {
	return optionutil.New(func(opts *Server) {
		opts.dialOpts = append(opts.dialOpts, dialOpts...)
	})
}
