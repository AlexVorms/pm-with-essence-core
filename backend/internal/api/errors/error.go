package errors

import (
	"fmt"
	"net/http"
)

type HttpError struct {
	Code    int
	Message string `json:"message"`
	Err     error
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

func NewLoginError(err error) *HttpError {
	return &HttpError{
		Err:     err,
		Code:    http.StatusUnauthorized,
		Message: "failed checks email or password",
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
func (e *HttpError) Error() string {
	return fmt.Sprintf("[%d] %v (blame: %s)", e.Code, e.Err)
}
