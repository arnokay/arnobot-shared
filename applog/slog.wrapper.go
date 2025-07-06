package applog

import (
	"context"
	"io"
	"log/slog"

	"github.com/arnokay/arnobot-shared/trace"
)

type appHandler struct {
	slog.Handler
}

func (h *appHandler) Handle(ctx context.Context, r slog.Record) error {
	if traceID, ok := ctx.Value(trace.TraceIDKey).(string); ok {
		r.AddAttrs(slog.String("trace_id", traceID))
	}

	return h.Handler.Handle(ctx, r)
}

func (h *appHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	return &appHandler{
		Handler: h.Handler.WithAttrs(attrs),
	}
}

func (h *appHandler) WithGroup(name string) slog.Handler {
	return &appHandler{
		Handler: h.Handler.WithGroup(name),
	}
}

type SlogLogger struct {
	*slog.Logger
}

func (l *SlogLogger) With(args ...any) Logger {
	return &SlogLogger{l.Logger.With(args...)}
}

func (l *SlogLogger) WithGroup(name string) Logger {
	return &SlogLogger{l.Logger.WithGroup(name)}
}

func NewSlogLogger(
	w io.Writer,
  appName string,
	logLevel int,
  handler slog.Handler,
) Logger {
  if handler == nil {
    handler = slog.NewTextHandler(w, &slog.HandlerOptions{
      Level: slog.Level(logLevel),
    })
  }

	customHandler := &appHandler{
		Handler: handler,
	}

	sl := slog.New(customHandler)

	if appName != "" {
		sl = sl.With(
			slog.String(LoggerAppKey, appName),
		)
	}
	slog.SetDefault(sl)

	return &SlogLogger{sl}
}
