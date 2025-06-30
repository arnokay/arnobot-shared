package service

import (
	"context"
	"log/slog"

	"github.com/nats-io/nats.go"

	"github.com/arnokay/arnobot-shared/applog"
	"github.com/arnokay/arnobot-shared/data"
	"github.com/arnokay/arnobot-shared/topics"
)

type AuthModule struct {
	mb     *nats.Conn
	logger *slog.Logger
}

func NewAuthModule(mb *nats.Conn) *AuthModule {
	logger := applog.NewServiceLogger("auth-module-service")

	return &AuthModule{
		mb:     mb,
		logger: logger,
	}
}

func (s *AuthModule) AuthSessionValidate(ctx context.Context, token string) (bool, error) {
  return HandleRequest[bool](ctx, s.mb, s.logger, topics.AuthSessionTokenValidate, token)
}

func (s *AuthModule) AuthSessionGetOwner(ctx context.Context, token string) (*data.User, error) {
  return HandleRequest[*data.User](ctx, s.mb, s.logger, topics.AuthSessionTokenExchange, token)
}

func (s *AuthModule) AuthProviderGet(ctx context.Context, arg data.AuthProviderGet) (*data.AuthProvider, error) {
  return HandleRequest[*data.AuthProvider](ctx, s.mb, s.logger, topics.AuthProviderTokenGet, arg)
}

func (s *AuthModule) AuthProviderUpdateTokens(ctx context.Context, arg data.AuthProviderUpdateTokens) error {
  return HandlePublish(ctx, s.mb, s.logger, topics.AuthProviderTokenUpdateTokens, arg)
}
