package middlewares

import (
	"context"
	"net/http"

	"google.golang.org/grpc/metadata"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
)

func WithMetadata() runtime.ServeMuxOption {
	fn := func(ctx context.Context, req *http.Request) metadata.MD {
		md := metadata.New(map[string]string{
			gwMetadataPrefix + "method":    req.Method,
			gwMetadataPrefix + "raw-query": req.URL.RawQuery,
		})
		for k, vs := range req.URL.Query() {
			md.Set(gwMetadataPrefix+"query-"+k, vs...)
		}
		return md
	}
	return runtime.WithMetadata(fn)
}

func WithMarshalerOption() runtime.ServeMuxOption {
	opt := &runtime.HTTPBodyMarshaler{
		Marshaler: &runtime.JSONPb{
			MarshalOptions: protojson.MarshalOptions{
				UseEnumNumbers:  true,
				EmitUnpopulated: true,
			},
			UnmarshalOptions: protojson.UnmarshalOptions{
				DiscardUnknown: true,
			},
		},
	}
	return runtime.WithMarshalerOption(runtime.MIMEWildcard, opt)
}

func WithForwardResponseOption() runtime.ServeMuxOption {
	fn := func(ctx context.Context, w http.ResponseWriter, m proto.Message) error {
		return nil
	}
	return runtime.WithForwardResponseOption(fn)
}
