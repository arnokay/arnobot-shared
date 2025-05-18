package middlewares

import (
	"log/slog"

	"arnobot-shared/applog"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

func TraceIDMiddleware() echo.MiddlewareFunc {
  return func(next echo.HandlerFunc) echo.HandlerFunc {
    return func(c echo.Context) error {
      traceID := uuid.New().String()
     
      c.Set("trace_id", traceID)

      logger := applog.NewWithAttrs(slog.String("trace_id", traceID))

      c.Set("logger", logger)
      c.Response().Header().Set("X-Trace-ID", traceID)

      return next(c)
    }
  }
}
