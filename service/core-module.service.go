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

type CoreModuleService struct {
	mb     *nats.Conn
	logger *slog.Logger
}

func NewCoreModuleService(mb *nats.Conn) *CoreModuleService {
	logger := applog.NewServiceLogger("core-module-service")

	return &CoreModuleService{
		mb:     mb,
		logger: logger,
	}
}

func (s *CoreModuleService) ChatMessageNotify(ctx context.Context, arg events.Message) error {
	payload := apptype.CoreChatMessageNotify{
		Data:    arg,
		TraceID: trace.FromContext(ctx),
	}

	payloadBytes, _ := payload.Encode()

	topic := topics.PlatformBroadcasterChatMessageNotify.Build(arg.Platform, arg.BroadcasterID)

	err := s.mb.Publish(topic, payloadBytes)
	if err != nil {
		s.logger.ErrorContext(ctx, "cannot notify about new message", "err", err)
		return apperror.ErrInternal
	}

	return nil
}
