package utils

//RespMes struct
type RespMes struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

// NewRespMes creates a new response body
func NewRespMes(message string, data interface{}) *RespMes {
	return &RespMes{
		Message: message,
		Data:    data,
	}
}
