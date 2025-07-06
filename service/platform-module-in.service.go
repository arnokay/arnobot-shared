package service

import (
	"context"

	"github.com/nats-io/nats.go"

	"github.com/arnokay/arnobot-shared/applog"
	"github.com/arnokay/arnobot-shared/data"
	"github.com/arnokay/arnobot-shared/events"
	"github.com/arnokay/arnobot-shared/topics"
)

type PlatformModuleIn struct {
	mb     *nats.Conn
	logger applog.Logger
}

func NewPlatformModuleIn(mb *nats.Conn) *PlatformModuleIn {
	logger := applog.NewServiceLogger("platform-module-service")

	return &PlatformModuleIn{
		mb:     mb,
		logger: logger,
	}
}

func (s *PlatformModuleIn) GetBot(ctx context.Context, arg data.PlatformBotGet) (data.PlatformBot, error) {
	topicBuilder := topics.TopicBuilder(topics.PlatformGetBot)
	topicBuilder.Platform(arg.Platform)
	topic := topicBuilder.Build()

	return HandleRequest[data.PlatformBot](ctx, s.mb, s.logger, topic, arg)
}

func (s *PlatformModuleIn) StartBot(ctx context.Context, arg data.PlatformBotToggle) error {
	topicBulder := topics.TopicBuilder(topics.PlatformStartBot)
	topicBulder.Platform(arg.Platform)
	topic := topicBulder.Build()

	return HandlePublish(ctx, s.mb, s.logger, topic, arg)
}

func (s *PlatformModuleIn) StopBot(ctx context.Context, arg data.PlatformBotToggle) error {
	topicBulder := topics.TopicBuilder(topics.PlatformStopBot)
	topicBulder.Platform(arg.Platform)
	topic := topicBulder.Build()

	return HandlePublish(ctx, s.mb, s.logger, topic, arg)
}

func (s *PlatformModuleIn) ChatSendMessage(ctx context.Context, arg events.MessageSend) error {
	topicBulder := topics.TopicBuilder(topics.PlatformBroadcasterChatMessageSend)
	topicBulder.Platform(arg.Platform)
	topicBulder.BroadcasterID(arg.BroadcasterID)
	topic := topicBulder.Build()

	return HandlePublish(ctx, s.mb, s.logger, topic, arg)
}
