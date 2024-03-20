package middlewares

import (
	"context"
	"errors"
	"io"
	"net/http"
	"net/textproto"
	"strings"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"

	"github.com/starudream/go-lib/core/v2/slog"
	"github.com/starudream/go-lib/server/v2/ierr"
)

func WithErrorHandler() runtime.ServeMuxOption {
	return runtime.WithErrorHandler(ErrorHandler)
}

func ErrorHandler(ctx context.Context, _ *runtime.ServeMux, marshaler runtime.Marshaler, w http.ResponseWriter, r *http.Request, err error) {
	const fallback = `{"code":9999,"message":"failed to marshal error message"}`

	var customStatus *runtime.HTTPStatusError
	if errors.As(err, &customStatus) {
		err = customStatus.Err
	}

	err = ierr.FromError(err)

	w.Header().Del("Trailer")
	w.Header().Del("Transfer-Encoding")

	contentType := marshaler.ContentType(err)
	w.Header().Set("Content-Type", contentType)

	buf, me := marshaler.Marshal(err)
	if me != nil {
		slog.Warn("marshal error: %v", me)
		w.WriteHeader(http.StatusInternalServerError)
		if _, we := io.WriteString(w, fallback); we != nil {
			slog.Warn("write response body error: %v", we)
		}
		return
	}

	md, _ := runtime.ServerMetadataFromContext(ctx)

	handleForwardResponseServerMetadata(w, md)

	doForwardTrailers := requestAcceptsTrailers(r)

	if doForwardTrailers {
		handleForwardResponseTrailerHeader(w, md)
		w.Header().Set("Transfer-Encoding", "chunked")
	}

	status := ierr.Status(err)
	if customStatus != nil {
		status = customStatus.HTTPStatus
	}

	w.WriteHeader(status)
	if _, we := w.Write(buf); we != nil {
		slog.Warn("write response error: %v", we)
	}

	if doForwardTrailers {
		handleForwardResponseTrailer(w, md)
	}
}

func requestAcceptsTrailers(req *http.Request) bool {
	te := req.Header.Get("TE")
	return strings.Contains(strings.ToLower(te), "trailers")
}

func handleForwardResponseServerMetadata(w http.ResponseWriter, md runtime.ServerMetadata) {
	for k, vs := range md.HeaderMD {
		if h, ok := OutgoingHeaderMatcher(k); ok {
			for _, v := range vs {
				w.Header().Add(h, v)
			}
		}
	}
}

func handleForwardResponseTrailerHeader(w http.ResponseWriter, md runtime.ServerMetadata) {
	for k := range md.TrailerMD {
		if h, ok := OutgoingTrailerMatcher(k); ok {
			w.Header().Add("Trailer", textproto.CanonicalMIMEHeaderKey(h))
		}
	}
}

func handleForwardResponseTrailer(w http.ResponseWriter, md runtime.ServerMetadata) {
	for k, vs := range md.TrailerMD {
		if h, ok := OutgoingTrailerMatcher(k); ok {
			for _, v := range vs {
				w.Header().Add(h, v)
			}
		}
	}
}
