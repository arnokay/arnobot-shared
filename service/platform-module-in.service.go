package service

import (
	"context"
	"log/slog"

	"github.com/nats-io/nats.go"

	"github.com/arnokay/arnobot-shared/apperror"
	"github.com/arnokay/arnobot-shared/applog"
	"github.com/arnokay/arnobot-shared/apptype"
	"github.com/arnokay/arnobot-shared/data"
	"github.com/arnokay/arnobot-shared/events"
	"github.com/arnokay/arnobot-shared/topics"
	"github.com/arnokay/arnobot-shared/trace"
)

type PlatformModuleIn struct {
	mb     *nats.Conn
	logger *slog.Logger
}

func NewPlatformModuleIn(mb *nats.Conn) *PlatformModuleIn {
	logger := applog.NewServiceLogger("platform-module-service")

	return &PlatformModuleIn{
		mb:     mb,
		logger: logger,
	}
}

func (s *PlatformModuleIn) GetBot(ctx context.Context, arg data.PlatformBotGet) (data.PlatformBot, error) {
	req := apptype.PlatformBotGetRequest{
		Data:    arg,
		TraceID: trace.FromContext(ctx),
	}

	b, _ := req.Encode()

	topicBuilder := topics.TopicBuilder(topics.PlatformGetBot)
	topicBuilder.Platform(arg.Platform)
	topic := topicBuilder.Build()

	msg, err := s.mb.RequestWithContext(ctx, topic, b)
	if err != nil {
		s.logger.ErrorContext(ctx, "cannot request bot", "err", err)
		return data.PlatformBot{}, apperror.ErrInternal
	}

	var resp apptype.PlatformBotGetResponse
	resp.Decode(msg.Data)

	if !resp.Success {
		return data.PlatformBot{}, apperror.New(resp.Code, resp.Error, nil)
	}

	return resp.Data, nil
}

func (s *PlatformModuleIn) StartBot(ctx context.Context, arg data.PlatformBotToggle) error {
	req := apptype.PlatformStartBot{
		Data:    arg,
		TraceID: trace.FromContext(ctx),
	}

	reqBytes, _ := req.Encode()

	topicBulder := topics.TopicBuilder(topics.PlatformStartBot)
	topicBulder.Platform(arg.Platform)
	topic := topicBulder.Build()

	msg, err := s.mb.RequestWithContext(ctx, topic, reqBytes)
	if err != nil {
		s.logger.ErrorContext(ctx, "cannot start bot", "err", err, "platform", arg.Platform)
		return apperror.ErrInternal
	}

	var resp apptype.EmptyResponse
	resp.Decode(msg.Data)

	if !resp.Success {
		return apperror.New(resp.Code, resp.Error, nil)
	}

	return nil
}

func (s *PlatformModuleIn) StopBot(ctx context.Context, arg data.PlatformBotToggle) error {
	req := apptype.PlatformStopBot{
		Data:    arg,
		TraceID: trace.FromContext(ctx),
	}

	reqBytes, _ := req.Encode()

	topicBulder := topics.TopicBuilder(topics.PlatformStopBot)
	topicBulder.Platform(arg.Platform)
	topic := topicBulder.Build()

	msg, err := s.mb.RequestWithContext(ctx, topic, reqBytes)
	if err != nil {
		s.logger.ErrorContext(ctx, "cannot stop bot", "err", err, "platform", arg.Platform)
		return apperror.ErrInternal
	}

	var resp apptype.EmptyResponse
	resp.Decode(msg.Data)

	if !resp.Success {
		return apperror.New(resp.Code, resp.Error, nil)
	}

	return nil
}

func (s *PlatformModuleIn) ChatSendMessage(ctx context.Context, arg events.MessageSend) error {
	payload := apptype.PlatformChatMessageSend{
		Data:    arg,
		TraceID: trace.FromContext(ctx),
	}

	payloadBytes, _ := payload.Encode()

	topicBulder := topics.TopicBuilder(topics.PlatformBroadcasterChatMessageSend)
	topicBulder.Platform(arg.Platform)
	topicBulder.BroadcasterID(arg.BroadcasterID)
	topic := topicBulder.Build()

	err := s.mb.Publish(topic, payloadBytes)
	if err != nil {
		s.logger.ErrorContext(ctx, "cannot send chat message", "err", err, "platform", arg.Platform)
		return apperror.ErrInternal
	}

	return nil
}
