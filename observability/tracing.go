package observability

import "context"

var tracer Tracer

func SetTracer(t Tracer) {
	if t != nil {
		tracer = t
	}
}

func Start(ctx context.Context, name string) (context.Context, Span) {
	return tracer.Start(ctx, name)
}
