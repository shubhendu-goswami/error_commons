package utils

import (
	"github.com/shubhendu-goswami/error_commons/model"
)

type HttpErrorUtils struct {
}

func (_ HttpErrorUtils) BadRequest() model.HttpError {
	return model.NewHttpError("Bad Request", 400, "bad_request")
}

func (_ HttpErrorUtils) UnprocessableEntity() model.HttpError {
	return model.NewHttpError("Unprocessable Entity", 422, "unprocessable_entity")
}

func (_ HttpErrorUtils) NotFound() model.HttpError {
	return model.NewHttpError("Not Found", 404, "page_not_found")
}

func (_ HttpErrorUtils) Forbidden() model.HttpError {
	return model.NewHttpError("Forbidden", 404, "forbidden")
}

func (_ HttpErrorUtils) Unauthorised() model.HttpError {
	return model.NewHttpError("Unauthorised", 401, "unauthorised")
}

func (_ HttpErrorUtils) GatewayTimeout() model.HttpError {
	return model.NewHttpError("Gateway Timeout", 504, "gateway_timeout")
}

func (_ HttpErrorUtils) InternalServerError(message string) model.HttpError {
	return model.NewHttpError("Internal Server Error", 500, message)
}

func (_ HttpErrorUtils) NotImplemented() model.HttpError {
	return model.NewHttpError("Not Implemented", 501, "not_implemented")
}

func (_ HttpErrorUtils) BadGateway() model.HttpError {
	return model.NewHttpError("Bad Gateway", 502, "bad_gateway")
}

func (_ HttpErrorUtils) ServiceUnavailable() model.HttpError {
	return model.NewHttpError("Service Unavailable", 503, "service_unavailable")
}

func HttpError() HttpErrorUtils {
	return HttpErrorUtils{}
}
