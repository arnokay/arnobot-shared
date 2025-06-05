package apperror

import (
	"errors"
	"net/http"
)

type ErrorCode string

func (ec ErrorCode) String() string {
	return string(ec)
}

const (
	CodeHTTP           ErrorCode = "http"
	CodeNotFound       ErrorCode = "not_found"
	CodeAlreadyExists  ErrorCode = "already_exists"
	CodeInvalidInput   ErrorCode = "invalid_input"
	CodeInternal       ErrorCode = "internal"
	CodeUnauthorized   ErrorCode = "unauthorized"
	CodeForbidden      ErrorCode = "forbidden"
	CodeExternal       ErrorCode = "external"
	CodeNotImplemented ErrorCode = "not_implemented"
)

var (
	ErrNotFound       = AppError{Code: CodeNotFound, Message: "resource not found"}
	ErrInvalidInput   = AppError{Code: CodeInvalidInput, Message: "invalid input provided"}
	ErrInternal       = AppError{Code: CodeInternal, Message: "internal server error"}
	ErrUnauthorized   = AppError{Code: CodeUnauthorized, Message: "unauthorized access"}
	ErrAlreadyExists  = AppError{Code: CodeAlreadyExists, Message: "resource already exists"}
	ErrHTTP           = AppError{Code: CodeHTTP, Message: "http error"}
	ErrExternal       = AppError{Code: CodeExternal, Message: "external error"}
	ErrForbidden      = AppError{Code: CodeForbidden, Message: "no access"}
	ErrNotImplemented = AppError{Code: CodeNotImplemented, Message: "not implemented"}
)

type AppError struct {
	Code    ErrorCode `json:"code"`
	Message string    `json:"message"`
	cause   error     `json:"-"`
}

func New(code ErrorCode, msg string, err error) AppError {
	return AppError{
		Code:    code,
		cause:   err,
		Message: msg,
	}
}

func (e AppError) Error() string {
	return e.Message
}

func (e AppError) Unwrap() error {
	return e.cause
}

func ToHTTPStatus(err error) int {
	var appErr AppError
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
		case CodeExternal:
			return http.StatusServiceUnavailable
		case CodeForbidden:
			return http.StatusForbidden
		case CodeNotImplemented:
			return http.StatusNotImplemented
		default:
			return http.StatusInternalServerError
		}
	}
	return http.StatusInternalServerError
}

func IsAppErr(err error) bool {
	var appErr AppError
	return errors.As(err, &appErr)
}
