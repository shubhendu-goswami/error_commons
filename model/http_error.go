package model

type HttpError struct {
	Message string `json:"message"`
	Code    int64  `json:"code"`
	Status  string `json:"status"`
}

func NewHttpError(message string, code int64, status string) HttpError {
	return HttpError{
		Message: message,
		Code:    code,
		Status:  status,
	}
}

func (e HttpError) GetMessage() string {
	return e.Message
}

func (e HttpError) GetCode() int64 {
	return e.Code
}

func (e HttpError) GetStatus() string {
	return e.Status
}
