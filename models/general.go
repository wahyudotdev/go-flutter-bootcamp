package models

type GeneralError struct {
	Message string `json:"message"`
	Error   string `json:"error"`
}

type GeneralResponse struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
	Token   string      `json:"token,omitempty"`
}
