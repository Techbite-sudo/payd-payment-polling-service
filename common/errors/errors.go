package errors

import "fmt"

type AppError struct {
	Code    int
	Message string
	Err     error
}

func (e *AppError) Error() string {
	if e.Err != nil {
		return fmt.Sprintf("AppError: Code=%d, Message=%s, Error=%v", e.Code, e.Message, e.Err)
	}
	return fmt.Sprintf("AppError: Code=%d, Message=%s", e.Code, e.Message)
}

func NewAppError(code int, message string, err error) *AppError {
	return &AppError{
		Code:    code,
		Message: message,
		Err:     err,
	}
}
