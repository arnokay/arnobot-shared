package service

import (
	"context"
	"log/slog"

	"github.com/nats-io/nats.go"

	"arnobot-shared/applog"
	"arnobot-shared/data"
	"arnobot-shared/mbtypes"
	"arnobot-shared/pkg/errs"
	"arnobot-shared/topics"
	"arnobot-shared/trace"
)

type AuthModuleService struct {
	mb     *nats.Conn
	logger *slog.Logger
}

func NewAuthModuleService(mb *nats.Conn) *AuthModuleService {
	logger := applog.NewServiceLogger("auth-module-service")

	return &AuthModuleService{
		mb:     mb,
		logger: logger,
	}
}

func (s *AuthModuleService) UserSessionValidate(ctx context.Context, token string) (bool, error) {
	req := mbtypes.AuthSessionTokenRequest{
		Data:    token,
		TraceID: trace.FromContext(ctx),
	}

	reqBytes, _ := req.Encode()

	msg, err := s.mb.RequestWithContext(ctx, topics.AuthSessionTokenValidate, reqBytes)
	if err != nil {
		s.logger.ErrorContext(ctx, "cannot validate user session", "err", err)
		return false, errs.ErrInternal
	}

	var response mbtypes.AuthSessionTokenValidateResponse

	err = response.Decode(msg.Data)
	if err != nil {
		s.logger.ErrorContext(ctx, "cannot decode response", "err", err)
		return false, errs.ErrInternal
	}

	if !response.Success {
		return false, errs.New(response.Code, response.Error, nil)
	}

	return true, nil
}

func (s *AuthModuleService) AuthSessionExchange(ctx context.Context, token string) (*data.User, error) {
	req := mbtypes.AuthSessionTokenRequest{
		TraceID: trace.FromContext(ctx),
		Data:    token,
	}

	b, _ := req.Encode()

	msg, err := s.mb.RequestWithContext(ctx, topics.AuthSessionTokenExchange, b)
	if err != nil {
		s.logger.ErrorContext(ctx, "cannot request user session exchange", "err", err)
		return nil, errs.ErrInternal
	}

	var res mbtypes.AuthSessionTokenExchangeResponse

	err = res.Decode(msg.Data)
	if err != nil {
		s.logger.ErrorContext(ctx, "cannot decode response", "err", err)
		return nil, errs.ErrInternal
	}

	if !res.Success {
		return nil, errs.New(res.Code, res.Error, nil)
	}

	return res.Data, nil
}

func (s *AuthModuleService) AuthProviderGet(ctx context.Context, data data.AuthProviderGet) (*data.AuthProvider, error) {
	req := mbtypes.AuthProviderGetRequest{
		TraceID: trace.FromContext(ctx),
		Data:    data,
	}

	b, _ := req.Encode()

	msg, err := s.mb.RequestWithContext(ctx, topics.AuthProviderTokenGet, b)
	if err != nil {
		s.logger.ErrorContext(ctx, "cannot request provider", "err", err)
		return nil, errs.ErrInternal
	}

	var res mbtypes.AuthProviderGetResponse

	err = res.Decode(msg.Data)
	if err != nil {
		s.logger.ErrorContext(ctx, "cannot decode provider response", "err", err)
		return nil, errs.ErrInternal
	}

	if !res.Success {
		return nil, errs.New(res.Code, res.Error, nil)
	}

	return res.Data, nil
}

func (s *AuthModuleService) AuthProviderUpdateTokens(ctx context.Context, data data.AuthProviderUpdateTokens) error {
	req := mbtypes.AuthProviderUpdateTokensRequest{
		TraceID: trace.FromContext(ctx),
		Data:    data,
	}

	b, _ := req.Encode()

	msg, err := s.mb.RequestWithContext(ctx, topics.AuthProviderTokenUpdateTokens, b)
	if err != nil {
		s.logger.ErrorContext(ctx, "cannot request provider tokens update", "err", err)
		return errs.ErrInternal
	}

	var res mbtypes.AuthProviderUpdateTokensResponse

	err = res.Decode(msg.Data)
	if err != nil {
		s.logger.ErrorContext(ctx, "cannot decode provider response", "err", err)
		return errs.ErrInternal
	}

	if !res.Success {
		return errs.New(res.Code, res.Error, nil)
	}

	return nil
}
