package middlewares

import (
	"fmt"
	"log/slog"
	"strings"

	"github.com/labstack/echo/v4"

	"arnobot-shared/appctx"
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
    fmt.Println("kek2")

		if !strings.HasPrefix(header, "Session") {
      m.logger.DebugContext(c.Request().Context(), "header has no session prefix")
			return errs.ErrUnauthorized
		}

		parts := strings.SplitN(header, " ", 2)
		if len(parts) != 2 {
      m.logger.DebugContext(c.Request().Context(), "header has no token")
			return errs.ErrUnauthorized
		}

		sessionToken := parts[1]

		valid, err := m.authModuleService.AuthSessionValidate(c.Request().Context(), sessionToken)
		if err != nil {
      m.logger.DebugContext(c.Request().Context(), "cannot get auth session validate", "err", err)
			return errs.ErrUnauthorized
		}

		if !valid {
      m.logger.DebugContext(c.Request().Context(), "token is not valid")
			return errs.ErrUnauthorized
		}

		return next(c)
	}
}

func (m *AuthMiddlewares) SessionGetOwner(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		header := c.Request().Header.Get("Authorization")
    fmt.Println("kek3")

		if !strings.HasPrefix(header, "Session") {
			return next(c)
		}

		parts := strings.SplitN(header, " ", 2)
		if len(parts) != 2 {
			return next(c)
		}

		sessionToken := parts[1]

		user, err := m.authModuleService.AuthSessionGetOwner(c.Request().Context(), sessionToken)
		if err != nil {
			return next(c)
		}

		ctx := appctx.SetUser(c.Request().Context(), user)
		c.SetRequest(c.Request().WithContext(ctx))

		return next(c)
	}
}
