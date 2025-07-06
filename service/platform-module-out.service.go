package service

import (
	"context"

	"github.com/nats-io/nats.go"

	"github.com/arnokay/arnobot-shared/applog"
	"github.com/arnokay/arnobot-shared/events"
	"github.com/arnokay/arnobot-shared/topics"
)

type PlatformModuleOut struct {
	mb     *nats.Conn
	logger applog.Logger
}

func NewPlatformModuleOut(mb *nats.Conn) *PlatformModuleOut {
	logger := applog.NewServiceLogger("platform-module-out")

	return &PlatformModuleOut{
		mb:     mb,
		logger: logger,
	}
}

func (s *PlatformModuleOut) ChatMessageNotify(ctx context.Context, arg events.Message) error {
	topicBulder := topics.TopicBuilder(topics.PlatformBroadcasterChatMessageNotify)
	topicBulder.Platform(arg.Platform)
  topicBulder.BroadcasterID(arg.BroadcasterID)
	topic := topicBulder.Build()

  return HandlePublish(ctx, s.mb, s.logger, topic, arg)
}
