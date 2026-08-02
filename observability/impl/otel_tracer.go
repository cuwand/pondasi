package impl

import (
	"context"
	"fmt"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"

	"github.com/cuwand/pondasi/observability"
)

type otelTracer struct {
	tracer trace.Tracer
}

func NewOtelTracer(serviceName string) observability.Tracer {
	return &otelTracer{
		tracer: otel.Tracer(serviceName),
	}
}

func (o *otelTracer) Start(
	ctx context.Context,
	name string,
) (context.Context, observability.Span) {
	ctx, span := o.tracer.Start(ctx, name)
	return ctx, &otelSpan{span: span}
}

type otelSpan struct {
	span trace.Span
}

func (s *otelSpan) End() {
	s.span.End()
}

func (s *otelSpan) SetAttr(key string, value any) {
	s.span.SetAttributes(attribute.String(key, fmt.Sprint(value)))
}

func (s *otelSpan) RecordError(err error) {
	if err != nil {
		s.span.RecordError(err)
	}
}
