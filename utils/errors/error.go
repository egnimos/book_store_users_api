package errors

import "net/http"

//RestErr : "common"
type RestErr struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
	Code    string `json:"code"`
}

//NewBadRequestError : check the error is for bad request status
func NewBadRequestError(message string) *RestErr {
	return &RestErr{
		Message: message,
		Status:  http.StatusBadRequest,
		Code:    "bad_request",
	}
}

//NewNotFoundError : check the error is for bad request status
func NewNotFoundError(message string) *RestErr {
	return &RestErr{
		Message: message,
		Status:  http.StatusNotFound,
		Code:    "not_found",
	}
}

//NewInternalServerError : check the error is caused by the server
func NewInternalServerError(message string) *RestErr {
	return &RestErr{
		Message: message,
		Status:  http.StatusInternalServerError,
		Code:    "internal_server_error",
	}
}
