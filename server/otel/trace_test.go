package otel

import (
	"context"
	"testing"

	"go.opentelemetry.io/otel/trace"

	"github.com/starudream/go-lib/core/v2/utils/testutil"
)

func Test(t *testing.T) {
	spanId, err := SpanIDFromHex("00f067aa0ba902b7")
	testutil.LogNoErr(t, err, spanId)

	traceId, err := TraceIDFromHex("4bf92f3577b34da6a3ce929d0e0e4736")
	testutil.LogNoErr(t, err, traceId)

	spanCtx := trace.NewSpanContext(trace.SpanContextConfig{TraceID: traceId, SpanID: spanId})

	ctx := ContextWithSpanContext(context.Background(), spanCtx)
	testutil.Log(t, spanCtx, SpanID(ctx), TraceID(ctx))
}
