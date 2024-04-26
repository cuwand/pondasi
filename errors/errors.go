package errors

import (
	"net/http"
)

type ErrorString struct {
	StatusCode   int
	ResponseCode string
	MessageId    string
	MessageEn    string
}

func (e ErrorString) Error() string {
	return e.MessageEn
}

// Error
func Code(errorCode string) error {
	errorInfo := GetErrorInfo(errorCode)

	return &ErrorString{
		StatusCode:   errorInfo.StatusCode,
		ResponseCode: errorInfo.ErrorCode,
		MessageId:    errorInfo.MsgId,
		MessageEn:    errorInfo.MsgEn,
	}
}

// Error
func CodeMsg(errorCode, msg string) error {
	return &ErrorString{
		ResponseCode: errorCode,
		MessageEn:    msg,
		MessageId:    msg,
	}
}

// BadRequest will throw if the given request-body or params is not valid
func BadRequest(msg string) error {
	return &ErrorString{
		StatusCode: http.StatusBadRequest,
		MessageEn:  msg,
		MessageId:  msg,
	}
}

// NotFound will throw if the requested item is not exists
func NotFound(msg string) error {
	return &ErrorString{
		StatusCode: http.StatusNotFound,
		MessageEn:  msg,
		MessageId:  msg,
	}
}

// NotFound will throw if the requested item is not exists
func MethodNotAllowed(msg string) error {
	return &ErrorString{
		StatusCode: http.StatusMethodNotAllowed,
		MessageEn:  msg,
		MessageId:  msg,
	}
}

// Conflict will throw if the current action already exists
func Conflict(msg string) error {
	return &ErrorString{
		StatusCode: http.StatusConflict,
		MessageEn:  msg,
		MessageId:  msg,
	}
}

// InternalServerError will throw if any the Internal Server Error happen,
// Database, Third Party etc.
func InternalServerError(msg string) error {
	return &ErrorString{
		StatusCode: http.StatusInternalServerError,
		MessageEn:  msg,
		MessageId:  msg,
	}
}

func UnauthorizedError(msg string) error {
	return &ErrorString{
		StatusCode: http.StatusUnauthorized,
		MessageEn:  msg,
		MessageId:  msg,
	}
}

func ForbiddenError(msg string) error {
	return &ErrorString{
		StatusCode: http.StatusForbidden,
		MessageEn:  msg,
		MessageId:  msg,
	}
}

func Error(errorString ErrorString) error {
	return &errorString
}
