package model

type Error struct {
	Message string
	Code    int64
}

func (e Error) Error(message string, code int64) {
	e.Code = code
	e.Message = message
}

func (e Error) GetMessage() string {
	return e.Message
}

func (e Error) GetCode() int64 {
	return e.Code
}
