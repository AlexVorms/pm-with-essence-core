package app_errors

import (
	"fmt"
	"net/http"
)

type HttpError struct {
	Code    int
	Message string `json:"message"`
	Err     error
}

func (e *HttpError) Unwrap() error {
	return e.Err
}
func (e *HttpError) Error() string {
	return fmt.Sprintf("[%d] %v (blame: %s)", e.Code, e.Err)
}

func NewBadRequestError(message string) *HttpError {
	return &HttpError{
		Code:    http.StatusBadRequest,
		Message: message,
	}
}
func NewPostgresWriteError(err error, message string) *HttpError {
	return &HttpError{
		Err:     err,
		Code:    http.StatusInternalServerError,
		Message: message,
	}
}

func NewPostgresReadError(err error, message string) *HttpError {
	return &HttpError{
		Err:     err,
		Code:    http.StatusInternalServerError,
		Message: message,
	}
}

func NewPostgresDuplicatedKeyError(err error) *HttpError {
	return &HttpError{
		Err:     err,
		Code:    http.StatusInternalServerError,
		Message: "this key already exists",
	}
}

func NewNotFoundError(err error, message string) *HttpError {
	return &HttpError{
		Err:     err,
		Code:    http.StatusNotFound,
		Message: message,
	}
}

func NewUnauthorizedError(err error) *HttpError {
	return &HttpError{
		Err:     err,
		Code:    http.StatusUnauthorized,
		Message: "unauthorized user",
	}
}

func NewParseEnumError(err error) *HttpError {
	return &HttpError{
		Err:     err,
		Code:    http.StatusInternalServerError,
		Message: "failed to parse enum",
	}
}

func NewHttpServerConnectError(err error) *HttpError {
	return &HttpError{
		Err:     err,
		Code:    http.StatusInternalServerError,
		Message: "failed connect http server",
	}
}

func NewHttpServerRequestError(err error, code int) *HttpError {
	return &HttpError{
		Err:     err,
		Code:    code,
		Message: "failed to request",
	}
}

func NewReadByteError(err error) *HttpError {
	return &HttpError{
		Err:     err,
		Code:    http.StatusInternalServerError,
		Message: "failed read byte",
	}
}

func NewJsonUnmarshalError(err error) *HttpError {
	return &HttpError{
		Err:     err,
		Code:    http.StatusInternalServerError,
		Message: "failed unmarshal json",
	}
}

func NewJsonMarshalError(err error) *HttpError {
	return &HttpError{
		Err:     err,
		Code:    http.StatusInternalServerError,
		Message: "failed marshal json",
	}
}

func NewLoginError() *HttpError {
	return &HttpError{
		Code:    http.StatusUnauthorized,
		Message: "Invalid email or password",
	}
}

func NewNotSessionError(err error) *HttpError {
	return &HttpError{
		Err:     err,
		Code:    http.StatusBadRequest,
		Message: "not found session",
	}
}

func NewExpiredDate(err error) *HttpError {
	return &HttpError{
		Err:     err,
		Code:    http.StatusUnauthorized,
		Message: "refresh token expired",
	}
}

func NewCreateJWTError(err error) *HttpError {
	return &HttpError{
		Err:     err,
		Code:    http.StatusInternalServerError,
		Message: "failed build jwt",
	}
}
func NewPasswordHashError(err error, message string) *HttpError {
	return &HttpError{
		Err:     err,
		Code:    http.StatusInternalServerError,
		Message: message,
	}
}
