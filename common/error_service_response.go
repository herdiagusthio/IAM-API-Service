package common

import (
	"net/http"

	"iam-api-service/service"
)

type errorServiceResponse string

const (
	errInternalServerError errorServiceResponse = "500"
	errNotFound            errorServiceResponse = "404"
	errBadRequest          errorServiceResponse = "400"
	errConflict            errorServiceResponse = "409"
)

//BusinessResponse default payload response
type ServiceResponse struct {
	Code    errorServiceResponse `json:"code"`
	Message string               `json:"message"`
	Data    interface{}          `json:"data"`
}

//NewErrorBusinessResponse Response return choosen http status like 400 bad request 422 unprocessable entity, ETC, based on responseCode
func NewErrorBusinessResponse(err error) (int, ServiceResponse) {
	return errorMapping(err)
}

//errorMapping error for missing header key with given value
func errorMapping(err error) (int, ServiceResponse) {
	switch err {
	default:
		return newInternalServerErrorResponse()
	case service.ErrNotFound:
		return newNotFoundResponse()
	case service.ErrInvalidData:
		return newValidationResponse(err.Error())
	case service.ErrLogin:
		return newErrLoginResponse(err.Error())
	case service.ErrRegister:
		return newErrRegisterResponse(err.Error())
	}
}

//newInternalServerErrorResponse default internal server error response
func newInternalServerErrorResponse() (int, ServiceResponse) {
	return http.StatusInternalServerError, ServiceResponse{
		errInternalServerError,
		"Internal server error",
		map[string]interface{}{},
	}
}

//newNotFoundResponse default not found error response
func newNotFoundResponse() (int, ServiceResponse) {
	return http.StatusNotFound, ServiceResponse{
		errNotFound,
		"Data Not found",
		map[string]interface{}{},
	}
}

//newValidationResponse failed to validate request payload
func newValidationResponse(message string) (int, ServiceResponse) {
	return http.StatusBadRequest, ServiceResponse{
		errBadRequest,
		"Validation failed, " + message,
		map[string]interface{}{},
	}
}

//newErrLoginResponse email or password is incorrect
func newErrLoginResponse(message string) (int, ServiceResponse) {
	return http.StatusBadRequest, ServiceResponse{
		errBadRequest,
		message,
		map[string]interface{}{},
	}
}

//newErrRegisterResponse email or password already registered
func newErrRegisterResponse(message string) (int, ServiceResponse) {
	return http.StatusConflict, ServiceResponse{
		errConflict,
		message,
		map[string]interface{}{},
	}
}
