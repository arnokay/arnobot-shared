package middlewares

import (
	"arnobot-shared/pkg/errs"
	"arnobot-shared/trace"
	"context"
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"
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

func ErrHandler(err error, c echo.Context) {
	if c.Response().Committed {
		return
	}
	status := http.StatusInternalServerError
  responseErr := errs.ErrInternal

	var appErr *errs.AppError
	if errors.As(err, &appErr) {
		status = errs.ToHTTPStatus(appErr)
    responseErr = *appErr
	} else {
		if he, ok := err.(*echo.HTTPError); ok {
			status = he.Code
      responseErr = errs.New(errs.CodeHTTP, he.Error(), he)
		}
	}

	c.JSON(status, responseErr)
}
