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
	payload := mbtypes.CoreChatMessageNotify{
		Data:    arg,
		TraceID: trace.FromContext(ctx),
	}

	payloadBytes, _ := payload.Encode()

	err := s.mb.Publish(topics.CoreChatMessageNotify, payloadBytes)
	if err != nil {
		s.logger.ErrorContext(ctx, "cannot notify about new message", "err", err)
		return errs.ErrInternal
	}

	return nil
}
