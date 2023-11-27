package helpers

import (
	"github.com/go-playground/validator/v10"
)

type Response struct {
	Message string      `json:"message"`
	Status  string      `json:"status"`
	Data    interface{} `json:"data"`
}

type ResponseFailed struct {
	Message string `json:"message"`
	Status  string `json:"status"`
}

func APIResponse(message string, status string, data interface{}) Response {
	jsonResponse := Response{
		Message: message,
		Status:  status,
		Data:    data,
	}
	return jsonResponse
}

func APIResponseFailed(message string, status string) ResponseFailed {
	jsonResponse := ResponseFailed{
		Message: message,
		Status:  status,
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
