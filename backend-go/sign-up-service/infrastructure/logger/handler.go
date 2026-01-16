package infrastructure_logger

import (
	"context"
	"log/slog"
)

// ContextHandler wraps a standard slog.Handler to inject context values
type ContextHandler struct {
	slog.Handler
}

// Handle intercepts the log record and injects the Trace ID from context
func (h *ContextHandler) Handle(ctx context.Context, r slog.Record) error {
	if traceID, ok := ctx.Value(TraceIDKey).(string); ok {
		// Add the trace_id attribute to the record automatically
		r.AddAttrs(slog.String("trace_id", traceID))
	}
	return h.Handler.Handle(ctx, r)
}
