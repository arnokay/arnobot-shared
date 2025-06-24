package middlewares

import (
	"log/slog"
	"strings"

	"github.com/labstack/echo/v4"

	"github.com/arnokay/arnobot-shared/appctx"
	"github.com/arnokay/arnobot-shared/apperror"
	"github.com/arnokay/arnobot-shared/applog"
	"github.com/arnokay/arnobot-shared/service"
)

type AuthMiddlewares struct {
	logger     *slog.Logger
	authModule *service.AuthModule
}

func NewAuthMiddleware(
	authModule *service.AuthModule,
) *AuthMiddlewares {
	logger := applog.NewServiceLogger("auth-middleware")

	return &AuthMiddlewares{
		logger:     logger,
		authModule: authModule,
	}
}

func (m *AuthMiddlewares) UserSessionGuard(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		header := c.Request().Header.Get("Authorization")

		if !strings.HasPrefix(header, "Session") {
			m.logger.DebugContext(c.Request().Context(), "header has no session prefix")
			return apperror.ErrUnauthorized
		}

		parts := strings.SplitN(header, " ", 2)
		if len(parts) != 2 {
			m.logger.DebugContext(c.Request().Context(), "header has no token")
			return apperror.ErrUnauthorized
		}

		sessionToken := parts[1]

		valid, err := m.authModule.AuthSessionValidate(c.Request().Context(), sessionToken)
		if err != nil {
			m.logger.DebugContext(c.Request().Context(), "cannot get auth session validate", "err", err)
			return apperror.ErrUnauthorized
		}

		if !valid {
			m.logger.DebugContext(c.Request().Context(), "token is not valid")
			return apperror.ErrUnauthorized
		}

		return next(c)
	}
}

func (m *AuthMiddlewares) SessionGetOwner(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		header := c.Request().Header.Get("Authorization")

		if !strings.HasPrefix(header, "Session") {
			return next(c)
		}

		parts := strings.SplitN(header, " ", 2)
		if len(parts) != 2 {
			return next(c)
		}

		sessionToken := parts[1]

		user, err := m.authModule.AuthSessionGetOwner(c.Request().Context(), sessionToken)
		if err != nil {
			return next(c)
		}

		ctx := appctx.SetUser(c.Request().Context(), user)
		c.SetRequest(c.Request().WithContext(ctx))

		return next(c)
	}
}
