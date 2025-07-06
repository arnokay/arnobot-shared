package service

import (
	"context"

	"github.com/google/uuid"
	"github.com/nats-io/nats.go"

	"github.com/arnokay/arnobot-shared/applog"
	"github.com/arnokay/arnobot-shared/data"
	"github.com/arnokay/arnobot-shared/topics"
)

type UserCommandModule struct {
	mb     *nats.Conn
	logger applog.Logger
}

func NewUserCommandModule(mb *nats.Conn) *UserCommandModule {
	logger := applog.NewServiceLogger("user-command-module-service")

	return &UserCommandModule{
		mb:     mb,
		logger: logger,
	}
}

func (s *UserCommandModule) GetByUserID(ctx context.Context, userID uuid.UUID) ([]data.UserCommand, error) {
	return HandleRequest[[]data.UserCommand](
		ctx,
		s.mb,
		s.logger,
		topics.CoreUserCommandGetByUserID,
		userID,
	)
}

func (s *UserCommandModule) GetOne(ctx context.Context, arg data.UserCommandGetOne) (data.UserCommand, error) {
	return HandleRequest[data.UserCommand](
		ctx,
		s.mb,
		s.logger,
		topics.CoreUserCommandGetOne,
		arg,
	)
}

func (s *UserCommandModule) Create(ctx context.Context, arg data.UserCommandCreate) (data.UserCommand, error) {
	return HandleRequest[data.UserCommand](
		ctx,
		s.mb,
		s.logger,
		topics.CoreUserCommandCreate,
		arg,
	)
}

func (s *UserCommandModule) Update(ctx context.Context, arg data.UserCommandUpdate) (data.UserCommand, error) {
	return HandleRequest[data.UserCommand](
		ctx,
		s.mb,
		s.logger,
		topics.CoreUserCommandCreate,
		arg,
	)
}

func (s *UserCommandModule) Delete(ctx context.Context, arg data.UserCommandDelete) (data.UserCommand, error) {
	return HandleRequest[data.UserCommand](
		ctx,
		s.mb,
		s.logger,
		topics.CoreUserCommandDelete,
		arg,
	)
}
