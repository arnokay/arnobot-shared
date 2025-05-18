package errs

import (
	"errors"
	"net/http"
)

type ErrorCode string

func (ec ErrorCode) String() string {
	return string(ec)
}

const (
	CodeHTTP          ErrorCode = "http"
	CodeNotFound      ErrorCode = "not_found"
	CodeAlreadyExists ErrorCode = "already_exists"
	CodeInvalidInput  ErrorCode = "invalid_input"
	CodeInternal      ErrorCode = "internal"
	CodeUnauthorized  ErrorCode = "unauthorized"
)

var (
	ErrNotFound      = AppError{Code: CodeNotFound, Message: "resource not found"}
	ErrInvalidInput  = AppError{Code: CodeInvalidInput, Message: "invalid input provided"}
	ErrInternal      = AppError{Code: CodeInternal, Message: "internal server error"}
	ErrUnauthorized  = AppError{Code: CodeUnauthorized, Message: "unauthorized access"}
	ErrAlreadyExists = AppError{Code: CodeAlreadyExists, Message: "resource already exists"}
	ErrHTTP          = AppError{Code: CodeHTTP, Message: "http error"}
)

type AppError struct {
	Code    ErrorCode
	Message string
	Cause   error
}

func New(code ErrorCode, msg string, err error) AppError {
	return AppError{
		Code:    code,
		Cause:   err,
		Message: msg,
	}
}

func (e AppError) Error() string {
	return e.Message
}

func (e AppError) Unwrap() error {
	return e.Cause
}

func FromCode(code string) AppError {
	switch code {
	case CodeAlreadyExists.String():
		return ErrAlreadyExists
	case CodeNotFound.String():
		return ErrNotFound
	case CodeInvalidInput.String():
		return ErrInvalidInput
	case CodeInternal.String():
		return ErrInternal
	case CodeUnauthorized.String():
		return ErrUnauthorized
	default:
		return ErrInternal
	}
}

func ToHTTPStatus(err error) int {
	var appErr *AppError
	if errors.As(err, &appErr) {
		switch appErr.Code {
		case CodeAlreadyExists:
			return http.StatusConflict
		case CodeNotFound:
			return http.StatusNotFound
		case CodeInvalidInput:
			return http.StatusBadRequest
		case CodeUnauthorized:
			return http.StatusUnauthorized
		case CodeInternal:
			return http.StatusInternalServerError
		default:
			return http.StatusInternalServerError
		}
	}
	return http.StatusInternalServerError
}
