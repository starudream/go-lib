package otel

import (
	"log/slog"

	"go.opentelemetry.io/contrib/propagators/b3"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/propagation"

	"github.com/go-logr/logr"
)

func init() {
	otel.SetLogger(logr.FromSlogHandler(slog.Default().Handler()))

	otel.SetTextMapPropagator(
		propagation.NewCompositeTextMapPropagator(
			propagation.TraceContext{},
			b3.New(b3.WithInjectEncoding(b3.B3MultipleHeader|b3.B3SingleHeader)),
		),
	)
}
