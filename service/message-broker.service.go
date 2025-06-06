package service

import (
	"context"
	"encoding/json"
	"log/slog"

	"github.com/nats-io/nats.go"

	"arnobot-shared/applog"
	"arnobot-shared/pkg/assert"
	"arnobot-shared/apperror"
)

type MessageBrokerService struct {
	mb     *nats.Conn
	logger *slog.Logger
}

func NewMessageBrokerService(messageBroker *nats.Conn) *MessageBrokerService {
	logger := applog.NewServiceLogger("MessageBrokerService")
	return &MessageBrokerService{
		mb:     messageBroker,
		logger: logger,
	}
}

func (s *MessageBrokerService) Request(ctx context.Context, topic string, data any, resp any) error {
	b, err := json.Marshal(data)
	assert.NoError(err, "#mbs: cannot marshal data")

	msg, err := s.mb.RequestWithContext(ctx, string(topic), b)
	if err != nil {
    s.logger.Error("cannot request", "topic", topic, "data", data)
		return apperror.ErrInternal
	}

  err = json.Unmarshal(msg.Data, &resp)
  assert.NoError(err, "#mbs: cannot unmarshal in response")

  return nil
}

// func (s *MessageBrokerService) Subscribe(topic topics.Topic) error {
//   s.mb.Subscribe(string(topic), handler)
//   return nil
// }
