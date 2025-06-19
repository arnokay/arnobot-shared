package service

import (
	"context"
	"log/slog"

	"github.com/nats-io/nats.go"

	"github.com/arnokay/arnobot-shared/apperror"
	"github.com/arnokay/arnobot-shared/applog"
	"github.com/arnokay/arnobot-shared/apptype"
	"github.com/arnokay/arnobot-shared/events"
	"github.com/arnokay/arnobot-shared/topics"
	"github.com/arnokay/arnobot-shared/trace"
)

type PlatformModuleService struct {
	mb     *nats.Conn
	logger *slog.Logger
}

func NewPlatformModuleService(mb *nats.Conn) *PlatformModuleService {
	logger := applog.NewServiceLogger("platform-module-service")

	return &PlatformModuleService{
		mb:     mb,
		logger: logger,
	}
}

func (s *PlatformModuleService) ChatSendMessage(ctx context.Context, arg events.MessageSend) error {
	payload := apptype.PlatformChatMessageSend{
		Data:    arg,
		TraceID: trace.FromContext(ctx),
	}

	payloadBytes, _ := payload.Encode()

	topic := topics.PlatformBroadcasterChatMessageSend.Build(arg.Platform, arg.BroadcasterID)

	err := s.mb.Publish(topic, payloadBytes)
	if err != nil {
		s.logger.ErrorContext(ctx, "cannot send chat message", "err", err, "platform", arg.Platform)
		return apperror.ErrInternal
	}

	return nil
}
