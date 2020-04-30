package utils

import "net/http"

//APIError struct
type APIError struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

var (
	Success            = NewAPIError(http.StatusOK, "success", nil)
	ServerError        = NewAPIError(http.StatusInternalServerError, "server error", nil)
	NotFound           = NewAPIError(http.StatusNotFound, "not found", nil)
	NotFoundInDatabase = NewAPIError(http.StatusNoContent, "there is no result", nil)
	PasswordNotMatch   = NewAPIError(http.StatusUnauthorized, "Email and password do not match", nil)
	Unauthenticated    = NewAPIError(http.StatusUnauthorized, "Validation failed, please log in again.", nil)
)

// NewAPIError creates a new response body
func NewAPIError(code int, message string, data interface{}) *APIError {
	return &APIError{
		Code:    code,
		Message: message,
		Data:    data,
	}
}

func (e *APIError) Error() string {
	return e.Message
}
