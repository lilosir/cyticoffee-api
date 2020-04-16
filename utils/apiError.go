package utils

import "net/http"

//APIError struct
type APIError struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

var (
	Success          = NewAPIError(http.StatusOK, "success", nil)
	ServerError      = NewAPIError(http.StatusInternalServerError, "server error", nil)
	NotFound         = NewAPIError(http.StatusNotFound, "not found", nil)
	PasswordNotMatch = NewAPIError(http.StatusUnauthorized, "Email and password do not match", nil)
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
