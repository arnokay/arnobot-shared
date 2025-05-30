package service

import (
	"context"
	"log/slog"

	"github.com/nats-io/nats.go"

	"arnobot-shared/applog"
	"arnobot-shared/events"
	"arnobot-shared/mbtypes"
	"arnobot-shared/pkg/errs"
	"arnobot-shared/topics"
	"arnobot-shared/trace"
)

type TwitchModuleService struct {
	mb     *nats.Conn
	logger *slog.Logger
}

func NewTwitchModuleService(mb *nats.Conn) *TwitchModuleService {
	logger := applog.NewServiceLogger("twitch-module-service")

	return &TwitchModuleService{
		mb:     mb,
		logger: logger,
	}
}

func (s *TwitchModuleService) ChatSendMessage(ctx context.Context, arg events.MessageSend) error {
	payload := mbtypes.PlatformChatMessageSend{
		Data:    arg,
		TraceID: trace.FromContext(ctx),
	}

	payloadBytes, _ := payload.Encode()

	err := s.mb.Publish(topics.TwitchChatMessageSend, payloadBytes)
	if err != nil {
		s.logger.ErrorContext(ctx, "cannot send chat message", "err", err)
		return errs.ErrInternal
	}

	return nil
}
