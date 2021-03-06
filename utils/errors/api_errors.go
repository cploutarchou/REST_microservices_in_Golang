package errors

import (
	"net/http"
)

type APIError struct {
	Message string `json:"message"`
	Status  int    `json:"code"`
	Error   string `json:"error"`
}

func NewBadRequest(message string) *APIError {
	return &APIError{
		Message: message,
		Status:  http.StatusBadRequest,
		Error:   "bad_request",
	}
}
func NewNotFound(message string) *APIError {
	return &APIError{
		Message: message,
		Status:  http.StatusNotFound,
		Error:   "bad_request",
	}
}

func Conflict(message string) *APIError {
	return &APIError{
		Message: message,
		Status:  http.StatusConflict,
		Error:   "conflict",
	}
}
