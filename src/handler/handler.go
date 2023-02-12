package handler

import "net/http"

func NewBadRequestErrorResponse(err string) *ErrorResponse {
	return NewErrorResponse(err, http.StatusBadRequest)
}

func NewForbiddenErrorResponse(err string) *ErrorResponse {
	return NewErrorResponse(err, http.StatusForbidden)
}

func NewNotFoundErrorResponse(err string) *ErrorResponse {
	return NewErrorResponse(err, http.StatusNotFound)
}

func NewErrorResponse(err string, code int) *ErrorResponse {
	return &ErrorResponse{
		Error: err,
		Code:  code,
	}
}

type ErrorResponse struct {
	Error string `json:"error"`
	Code  int    `json:"-"`
}
