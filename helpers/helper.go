package helpers

import (
	"github.com/go-playground/validator/v10"
)

type Response struct {
	Status  string      `json:"status"`
	Message any         `json:"message"`
	Data    interface{} `json:"data"`
}

type ResponseFailed struct {
	Status  string `json:"status"`
	Message any    `json:"message"`
}

func APIResponse(status string, message any, data interface{}) Response {
	jsonResponse := Response{
		Status:  status,
		Message: message,
		Data:    data,
	}
	return jsonResponse
}

func APIResponseFailed(status string, message any) ResponseFailed {
	jsonResponse := ResponseFailed{
		Status:  status,
		Message: message,
	}
	return jsonResponse
}

func FormatValidationError(err error) []string {
	var errors []string

	for _, e := range err.(validator.ValidationErrors) {
		errors = append(errors, e.Error())
	}
	return errors
}
