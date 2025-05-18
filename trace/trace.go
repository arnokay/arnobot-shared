package trace

import (
	"context"

	"github.com/google/uuid"
)

type traceIDKey string

func (t traceIDKey) String() string {
  return string(t)
}

const TraceIDKey traceIDKey = "trace_id"

func New() string {
  return uuid.NewString()
}

func FromContext(ctx context.Context) string {
	v := ctx.Value(TraceIDKey)

	if traceID, ok := v.(string); ok {
		return traceID
	}

	return ""
}

func Context(ctx context.Context, traceID string) context.Context {
  if traceID == "" {
    return ctx
  }

  return context.WithValue(ctx, TraceIDKey, traceID)
}
