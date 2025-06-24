package middlewares

import (
	"context"
	"errors"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"

	"github.com/arnokay/arnobot-shared/apperror"
	"github.com/arnokay/arnobot-shared/applog"
	"github.com/arnokay/arnobot-shared/trace"
)

func AttachTraceID(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		traceID := trace.New()

		c.Set(trace.TraceIDKey.String(), traceID)
		ctx2 := context.WithValue(c.Request().Context(), trace.TraceIDKey, traceID)
		c.SetRequest(c.Request().WithContext(ctx2))
		c.Response().Header().Add("X-Trace-ID", traceID)

		return next(c)
	}
}

func HTTPLogger(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		req := c.Request()
		res := c.Response()
		start := time.Now()

		logger := applog.NewServiceLogger("http")

		logger.InfoContext(
			req.Context(),
			"HTTP Request",
			"method", req.Method,
			"uri", req.RequestURI,
			"ip", c.RealIP(),
		)

		err := next(c)

		latency := time.Since(start)

		logger.InfoContext(req.Context(), "HTTP Response",
			"method", req.Method,
			"uri", req.RequestURI,
			"status", res.Status,
			"latency_ms", latency.Milliseconds(),
			"remote_ip", c.RealIP(),
		)

		return err
	}
}


func ErrHandler(err error, c echo.Context) {
	logger := applog.NewServiceLogger("err-handler")
	if c.Response().Committed {
		return
	}
	status := http.StatusInternalServerError
	responseErr := apperror.ErrInternal

	var appErr apperror.AppError
	if errors.As(err, &appErr) {
		status = apperror.ToHTTPStatus(appErr)
		responseErr = appErr
	} else {
		if he, ok := err.(*echo.HTTPError); ok {
			status = he.Code
			responseErr = apperror.New(apperror.CodeHTTP, he.Error(), he)
		}
	}

  logger.DebugContext(c.Request().Context(), "sending error", "status", status, "err", responseErr)

	c.JSON(status, responseErr)
}
