package middlewares

import (
	"fmt"
	"net/textproto"
	"strings"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
)

var (
	gwMetadataPrefix = "gw-"
	gwHeaderPrefix   = textproto.CanonicalMIMEHeaderKey(gwMetadataPrefix)
)

func WithIncomingHeaderMatcher() runtime.ServeMuxOption {
	return runtime.WithIncomingHeaderMatcher(IncomingHeaderMatcher)
}

func IncomingHeaderMatcher(key string) (string, bool) {
	switch key = textproto.CanonicalMIMEHeaderKey(key); {
	case isForwardedHeader(key):
		return key, false
	case strings.HasPrefix(key, "X-"):
		return key, true
	case isPermanentHTTPHeader(key):
		return gwMetadataPrefix + key, true
	case strings.HasPrefix(key, gwHeaderPrefix):
		return key[len(gwHeaderPrefix):], true
	}
	return key, false
}

func WithOutgoingHeaderMatcher() runtime.ServeMuxOption {
	return runtime.WithOutgoingHeaderMatcher(OutgoingHeaderMatcher)
}

func OutgoingHeaderMatcher(key string) (string, bool) {
	key = textproto.CanonicalMIMEHeaderKey(key)
	if strings.HasPrefix(key, "X-") {
		return key, true
	}
	return key, false
}

func OutgoingTrailerMatcher(key string) (string, bool) {
	return fmt.Sprintf("%s%s", gwHeaderPrefix, key), true
}

func isForwardedHeader(hdr string) bool {
	switch hdr {
	case
		"X-Forwarded-For",
		"X-Forwarded-Host":
		return true
	}
	return false
}

func isPermanentHTTPHeader(hdr string) bool {
	switch hdr {
	case
		"Accept",
		"Accept-Charset",
		"Accept-Language",
		"Accept-Ranges",
		"Authorization",
		"Cache-Control",
		"Content-Type",
		"Cookie",
		"Date",
		"Expect",
		"From",
		"Host",
		"If-Match",
		"If-Modified-Since",
		"If-None-Match",
		"If-Schedule-Tag-Match",
		"If-Unmodified-Since",
		"Max-Forwards",
		"Origin",
		"Pragma",
		"Referer",
		"User-Agent",
		"Via",
		"Warning":
		return true
	}
	return false
}
