package errors

import "fmt"

func New(errorCode string, message string, data ...interface{}) *Error {
	var errorData interface{}
	if len(data) > 0 {
		errorData = data[0]
	}
	return &Error{
		Code:    errorCode,
		Message: message,
		Data:    errorData,
	}
}

type Error struct {
	Code    string
	Message string
	Data    interface{}
}

func (e Error) Error() string {
	return fmt.Sprintf("%s: %s - %v", e.Code, e.Message, e.Data)
}
