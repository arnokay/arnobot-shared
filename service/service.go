package service

import (
	"context"

	"github.com/nats-io/nats.go"

	"github.com/arnokay/arnobot-shared/apperror"
	"github.com/arnokay/arnobot-shared/applog"
	"github.com/arnokay/arnobot-shared/apptype"
	"github.com/arnokay/arnobot-shared/trace"
)

func HandleRequest[TRes any](
	ctx context.Context,
	mb *nats.Conn,
	logger applog.Logger,
	topic string,
	data any,
) (TRes, error) {
	req := apptype.Request[any]{
		TraceID: trace.FromContext(ctx),
		Data:    data,
	}
	var res apptype.Response[TRes]

	b, _ := req.Encode()

	msg, err := mb.RequestWithContext(ctx, topic, b)
	if err != nil {
		logger.ErrorContext(ctx, "cannot request topic: "+topic, "err", err)
		return *new(TRes), apperror.ErrInternal
	}

	res.Decode(msg.Data)
	if !res.Success {
		return *new(TRes), apperror.New(res.Code, res.Error, nil)
	}

	return res.Data, nil
}

func HandlePublish[TReq any](
	ctx context.Context,
	mb *nats.Conn,
	logger applog.Logger,
	topic string,
	data TReq,
) error {
	req := apptype.Request[TReq]{
		TraceID: trace.FromContext(ctx),
		Data:    data,
	}

	b, _ := req.Encode()

	err := mb.Publish(topic, b)
	if err != nil {
		logger.ErrorContext(ctx, "cannot publish topic: "+topic, "err", err)
		return apperror.ErrInternal
	}

	return nil
}
