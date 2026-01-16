package infrastructure_logger

import (
	"context"

	"github.com/google/uuid"
)

type contextKey string

const TraceIDKey contextKey = "trace_id"

// FromContext retrieves the trace ID from the context
func FromContext(ctx context.Context) string {
	if s, ok := ctx.Value(TraceIDKey).(string); ok {
		return s
	}
	return "unknown"
}

// WithTraceID adds a new UUID to a context
func WithTraceID(ctx context.Context) context.Context {
	return context.WithValue(ctx, TraceIDKey, uuid.New().String())
}
