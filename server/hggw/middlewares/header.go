package middlewares

import (
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
)

func WithIncomingHeaderMatcher() runtime.ServeMuxOption {
	fn := func(key string) (string, bool) {
		return key, true
	}
	return runtime.WithIncomingHeaderMatcher(fn)
}

func WithOutgoingHeaderMatcher() runtime.ServeMuxOption {
	fn := func(key string) (string, bool) {
		return key, true
	}
	return runtime.WithOutgoingHeaderMatcher(fn)
}
