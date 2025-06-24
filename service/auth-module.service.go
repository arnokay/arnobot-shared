package service

import (
	"context"
	"log/slog"

	"github.com/nats-io/nats.go"

	"github.com/arnokay/arnobot-shared/applog"
	"github.com/arnokay/arnobot-shared/data"
	"github.com/arnokay/arnobot-shared/apptype"
	"github.com/arnokay/arnobot-shared/apperror"
	"github.com/arnokay/arnobot-shared/topics"
	"github.com/arnokay/arnobot-shared/trace"
)

type AuthModule struct {
	mb     *nats.Conn
	logger *slog.Logger
}

func NewAuthModule(mb *nats.Conn) *AuthModule {
	logger := applog.NewServiceLogger("auth-module-service")

	return &AuthModule{
		mb:     mb,
		logger: logger,
	}
}

func (s *AuthModule) AuthSessionValidate(ctx context.Context, token string) (bool, error) {
	req := apptype.AuthSessionTokenRequest{
		Data:    token,
		TraceID: trace.FromContext(ctx),
	}

	reqBytes, _ := req.Encode()

	msg, err := s.mb.RequestWithContext(ctx, topics.AuthSessionTokenValidate, reqBytes)
	if err != nil {
		s.logger.ErrorContext(ctx, "cannot validate user session", "err", err)
		return false, apperror.ErrInternal
	}

	var response apptype.AuthSessionTokenValidateResponse

	err = response.Decode(msg.Data)
	if err != nil {
		s.logger.ErrorContext(ctx, "cannot decode response", "err", err)
		return false, apperror.ErrInternal
	}

	if !response.Success {
		return false, apperror.New(response.Code, response.Error, nil)
	}

	return true, nil
}

func (s *AuthModule) AuthSessionGetOwner(ctx context.Context, token string) (*data.User, error) {
	req := apptype.AuthSessionTokenRequest{
		TraceID: trace.FromContext(ctx),
		Data:    token,
	}

	b, _ := req.Encode()

	msg, err := s.mb.RequestWithContext(ctx, topics.AuthSessionTokenExchange, b)
	if err != nil {
		s.logger.ErrorContext(ctx, "cannot request user session exchange", "err", err)
		return nil, apperror.ErrInternal
	}

	var res apptype.AuthSessionTokenExchangeResponse

	err = res.Decode(msg.Data)
	if err != nil {
		s.logger.ErrorContext(ctx, "cannot decode response", "err", err)
		return nil, apperror.ErrInternal
	}

	if !res.Success {
		return nil, apperror.New(res.Code, res.Error, nil)
	}

	return res.Data, nil
}

func (s *AuthModule) AuthProviderGet(ctx context.Context, data data.AuthProviderGet) (*data.AuthProvider, error) {
	req := apptype.AuthProviderGetRequest{
		TraceID: trace.FromContext(ctx),
		Data:    data,
	}

	b, _ := req.Encode()

	msg, err := s.mb.RequestWithContext(ctx, topics.AuthProviderTokenGet, b)
	if err != nil {
		s.logger.ErrorContext(ctx, "cannot request provider", "err", err)
		return nil, apperror.ErrInternal
	}

	var res apptype.AuthProviderGetResponse

	err = res.Decode(msg.Data)
	if err != nil {
		s.logger.ErrorContext(ctx, "cannot decode provider response", "err", err)
		return nil, apperror.ErrInternal
	}

	if !res.Success {
    s.logger.DebugContext(ctx, "unsuccessful request for auth provider", "err", res.Error)
		return nil, apperror.New(res.Code, res.Error, nil)
	}

	return res.Data, nil
}

func (s *AuthModule) AuthProviderUpdateTokens(ctx context.Context, id int32, data data.AuthProviderUpdateTokens) error {
	req := apptype.AuthProviderUpdateTokensRequest{
		TraceID: trace.FromContext(ctx),
		Data: apptype.AuthProviderUpdateTokensPayload{
			ID:   id,
			Data: data,
		},
	}

	b, _ := req.Encode()

	msg, err := s.mb.RequestWithContext(ctx, topics.AuthProviderTokenUpdateTokens, b)
	if err != nil {
		s.logger.ErrorContext(ctx, "cannot request provider tokens update", "err", err)
		return apperror.ErrInternal
	}

	var res apptype.AuthProviderUpdateTokensResponse

	err = res.Decode(msg.Data)
	if err != nil {
		s.logger.ErrorContext(ctx, "cannot decode provider response", "err", err)
		return apperror.ErrInternal
	}

	if !res.Success {
		return apperror.New(res.Code, res.Error, nil)
	}

	return nil
}
