package otel

import (
	"context"

	"go.opentelemetry.io/otel/trace"
)

type SpanContext = trace.SpanContext

var (
	ContextWithSpan              = trace.ContextWithSpan
	ContextWithSpanContext       = trace.ContextWithSpanContext
	ContextWithRemoteSpanContext = trace.ContextWithRemoteSpanContext

	SpanFromContext        = trace.SpanFromContext
	SpanContextFromContext = trace.SpanContextFromContext

	SpanIDFromHex  = trace.SpanIDFromHex
	TraceIDFromHex = trace.TraceIDFromHex
)

func SpanID(ctx context.Context) string {
	return SpanContextFromContext(ctx).SpanID().String()
}

func TraceID(ctx context.Context) string {
	return SpanContextFromContext(ctx).TraceID().String()
}
