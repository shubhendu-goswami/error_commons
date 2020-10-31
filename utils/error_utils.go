package utils

import "github.com/shubhendu-goswami/error_commons/model"

type ErrorUtils struct {
}

func (_ ErrorUtils) CustomError(message string, code int64) model.Error {
	return model.NewError(message, code)
}

func Error() ErrorUtils {
	return ErrorUtils{}
}
