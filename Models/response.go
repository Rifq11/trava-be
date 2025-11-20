package models

type SuccessResponse struct {
	Status  string `json:"status"`
	Message string `json:"message,omitempty"`
	Data    any    `json:"data,omitempty"`
}

type ErrorResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}
