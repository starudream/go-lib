package middlewares

import (
	"context"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
)

func WithErrorHandler() runtime.ServeMuxOption {
	fn := func(ctx context.Context, mux *runtime.ServeMux, marshal runtime.Marshaler, w http.ResponseWriter, r *http.Request, err error) {
	}
	return runtime.WithErrorHandler(fn)
}
