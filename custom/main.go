package main

import (
	"context"

	"github.com/ServiceWeaver/weaver-kube/tool"
	"go.opentelemetry.io/otel/exporters/jaeger"
	"go.opentelemetry.io/otel/sdk/trace"
)

func main() {
	const jaegerURL = "http://jaeger:14268/api/traces"
	endpoint := jaeger.WithCollectorEndpoint(jaeger.WithEndpoint(jaegerURL))
	traceExporter, err := jaeger.New(endpoint)
	if err != nil {
		panic(err)
	}

	tool.Run("deploy", tool.Plugins{
		HandleTraceSpans: func(ctx context.Context, spans []trace.ReadOnlySpan) error {
			return traceExporter.ExportSpans(ctx, spans)
		},
	})
}
