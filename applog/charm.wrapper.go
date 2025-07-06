package applog

import (
	"context"
	"io"
	"log/slog"

	"github.com/charmbracelet/log"

	"github.com/arnokay/arnobot-shared/trace"
)

type CharmLogWrapper struct {
	logger *log.Logger
}

func NewCharmLogger(
	w io.Writer,
  appName string,
	logLevel int,
	options *log.Options,
) Logger {
	if options == nil {
		options = &log.Options{}
	}

	options.CallerOffset = 1
	options.Level = log.Level(logLevel)
	logger := log.NewWithOptions(w, *options)

  logger.SetPrefix(appName)

	log.SetDefault(logger)

	return &CharmLogWrapper{logger: logger}
}

func (c *CharmLogWrapper) Debug(msg string, args ...any) {
	c.logger.Debug(msg, args...)
}

func (c *CharmLogWrapper) Info(msg string, args ...any) {
	c.logger.Info(msg, args...)
}

func (c *CharmLogWrapper) Warn(msg string, args ...any) {
	c.logger.Warn(msg, args...)
}

func (c *CharmLogWrapper) Error(msg string, args ...any) {
	c.logger.Error(msg, args...)
}

func (c *CharmLogWrapper) DebugContext(ctx context.Context, msg string, args ...any) {
	if traceID, ok := ctx.Value(trace.TraceIDKey).(string); ok {
		args = append(args, "trace_id", traceID)
	}
	c.logger.Debug(msg, args...)
}

func (c *CharmLogWrapper) InfoContext(ctx context.Context, msg string, args ...any) {
	if traceID, ok := ctx.Value(trace.TraceIDKey).(string); ok {
		args = append(args, "trace_id", traceID)
	}
	c.logger.Info(msg, args...)
}

func (c *CharmLogWrapper) WarnContext(ctx context.Context, msg string, args ...any) {
	if traceID, ok := ctx.Value(trace.TraceIDKey).(string); ok {
		args = append(args, "trace_id", traceID)
	}
	c.logger.Warn(msg, args...)
}

func (c *CharmLogWrapper) ErrorContext(ctx context.Context, msg string, args ...any) {
	if traceID, ok := ctx.Value(trace.TraceIDKey).(string); ok {
		args = append(args, "trace_id", traceID)
	}
	c.logger.Error(msg, args...)
}

func (c *CharmLogWrapper) Log(ctx context.Context, level slog.Level, msg string, args ...any) {
	switch level {
	case slog.LevelDebug:
		c.logger.Debug(msg, args...)
	case slog.LevelInfo:
		c.logger.Info(msg, args...)
	case slog.LevelWarn:
		c.logger.Warn(msg, args...)
	case slog.LevelError:
		c.logger.Error(msg, args...)
	default:
		// For custom levels, use info as fallback
		c.logger.Info(msg, args...)
	}
}

func (c *CharmLogWrapper) Enabled(ctx context.Context, level slog.Level) bool {
	charmLevel := c.logger.GetLevel()

	// Convert slog.Level to charm log level for comparison
	switch level {
	case slog.LevelDebug:
		return charmLevel <= log.DebugLevel
	case slog.LevelInfo:
		return charmLevel <= log.InfoLevel
	case slog.LevelWarn:
		return charmLevel <= log.WarnLevel
	case slog.LevelError:
		return charmLevel <= log.ErrorLevel
	default:
		return true
	}
}

func (c *CharmLogWrapper) With(args ...any) Logger {
	newLogger := c.logger.With(args...)
	return &CharmLogWrapper{logger: newLogger}
}

func (c *CharmLogWrapper) WithGroup(name string) Logger {
	newLogger := c.logger.WithPrefix(name)
	return &CharmLogWrapper{logger: newLogger}
}

func (c *CharmLogWrapper) Handler() slog.Handler {
	return nil
}
