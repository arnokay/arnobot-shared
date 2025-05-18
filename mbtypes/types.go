package mbtypes

import (
	"encoding/json"
	"fmt"

	"arnobot-shared/pkg/errs"
)

type Request[T any] struct {
	TraceID string `json:"traceId"`
	Data    T      `json:"data"`
}

func (r Request[T]) Encode() ([]byte, error) {
	b, err := r.JSON()

	return b, err
}

func (r *Request[T]) Decode(b []byte) error {
	err := r.DecodeJSON(b)

	return err
}

func (r Request[T]) JSON() ([]byte, error) {
	b, err := json.Marshal(r)
	if err != nil {
		return nil, fmt.Errorf("cannot marshal request data: %v+, error: %w", r.Data, err)
	}

	return b, nil
}

func (r *Request[T]) DecodeJSON(b []byte) error {
	err := json.Unmarshal(b, r)
	if err != nil {
		return fmt.Errorf("cannot unmarshal request data: %s, error: %w", string(b), err)
	}

	return nil
}

type Response[T any] struct {
	Success bool           `json:"success"`
	Error   string         `json:"error"`
	Code    errs.ErrorCode `json:"code"`
	TraceID string         `json:"traceId"`
	Data    T              `json:"data"`
}

func (r *Response[T]) ToSuccess(data T) {
	r.Success = true
	r.Data = data
}

func (r *Response[T]) ToFail(code errs.ErrorCode, reason string) {
	r.Success = false
	r.Error = reason
	r.Code = code
}

func (r *Response[T]) ToFailErr(err error) {
	r.Success = false
	if appErr, ok := err.(errs.AppError); ok {
		r.Code = appErr.Code
		r.Error = appErr.Message
		return
	}

	r.Code = errs.CodeInternal
	r.Error = err.Error()
}

func (r Response[T]) Encode() ([]byte, error) {
	b, err := r.JSON()

	return b, err
}

func (r *Response[T]) Decode(b []byte) error {
	err := r.DecodeJSON(b)

	return err
}

func (r Response[T]) JSON() ([]byte, error) {
	b, err := json.Marshal(r)
	if err != nil {
		return nil, fmt.Errorf("cannot marshal response data: %v+, error: %w", r.Data, err)
	}

	return b, nil
}

func (r *Response[T]) DecodeJSON(b []byte) error {
	err := json.Unmarshal(b, r)
	if err != nil {
		return fmt.Errorf("cannot unmarshal response data: %s, error: %w", string(b), err)
	}

	return nil
}
