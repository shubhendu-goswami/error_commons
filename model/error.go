package model

type Error struct {
	Message string
	Code    int64
}

func NewError(message string, code int64) Error {
	return Error{
		Message: message,
		Code:    code,
	}
}

func (e Error) GetMessage() string {
	return e.Message
}

func (e Error) GetCode() int64 {
	return e.Code
}
