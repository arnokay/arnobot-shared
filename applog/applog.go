package applog

import (
	"context"
	"io"
	"log/slog"

	"github.com/arnokay/arnobot-shared/consts"
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

func Init(
	appName string,
	w io.Writer,
	logLevel int,
) *slog.Logger {
	baseHandler := slog.NewTextHandler(w, &slog.HandlerOptions{
		Level: slog.Level(logLevel),
	})

	customHandler := &appHandler{
		Handler: baseHandler,
  }

	logger := slog.New(customHandler).With(
		slog.String(consts.LOGGER_APP_KEY, appName),
	)

	slog.SetDefault(logger)

	return logger
}

func NewWithAttrs(attrs ...slog.Attr) *slog.Logger {
	logger := slog.New(
		slog.Default().Handler().WithAttrs(
			attrs,
		),
	)

	return logger
}

func NewServiceLogger(serviceName string, attributes ...slog.Attr) *slog.Logger {
	attrs := []slog.Attr{
		slog.String(consts.LOGGER_SERVICE_KEY, serviceName),
	}
	attrs = append(attrs, attributes...)

	return NewWithAttrs(attrs...)
}
