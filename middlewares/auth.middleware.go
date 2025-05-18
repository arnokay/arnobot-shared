package middlewares

import (
	"log/slog"
	"strings"

	"github.com/labstack/echo/v4"

	"arnobot-shared/applog"
	"arnobot-shared/pkg/errs"
	"arnobot-shared/service"
)

type AuthMiddlewares struct {
	logger            *slog.Logger
	authModuleService *service.AuthModuleService
}

func NewAuthMiddleware(
	authModuleService *service.AuthModuleService,
) *AuthMiddlewares {
	logger := applog.NewServiceLogger("auth-middleware")

	return &AuthMiddlewares{
		logger:            logger,
		authModuleService: authModuleService,
	}
}

func (m *AuthMiddlewares) UserSessionGuard(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		header := c.Request().Header.Get("Authorization")

		if !strings.HasPrefix(header, "Session") {
			return errs.ErrUnauthorized
		}

		parts := strings.SplitN(header, " ", 2)
		if len(parts) != 2 {
			return errs.ErrUnauthorized
		}

		sessionToken := parts[1]

		valid, err := m.authModuleService.UserSessionValidate(c.Request().Context(), sessionToken)
		if err != nil {
			m.logger.ErrorContext(c.Request().Context(), "cannot validate user session", "err", err)
			return errs.ErrUnauthorized
		}

		if !valid {
			return errs.ErrUnauthorized
		}

		user, err := m.authModuleService.AuthSessionExchange(c.Request().Context(), sessionToken)
		if err != nil {
			m.logger.ErrorContext(c.Request().Context(), "cannot exchange user session", "err", err)
			return errs.ErrUnauthorized
		}

		c.Set("user", user)

		return next(c)
	}
}
