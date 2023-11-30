package helpers

import (
	"strings"

	"github.com/agusheryanto182/go-schedule/models/web"
	"github.com/go-playground/validator/v10"
)

type Response struct {
	Status  string      `json:"status"`
	Message any         `json:"message"`
	Data    interface{} `json:"data"`
}

type ResponseGetAll struct {
	Status  string                            `json:"status"`
	Message any                               `json:"message"`
	Data    map[string][]web.ScheduleResponse `json:"data"`
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

func APIResponseGetData(status string, message any, data map[string][]web.ScheduleResponse) ResponseGetAll {
	jsonResponse := ResponseGetAll{
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

func IsValidEmail(email string) bool {
	// Pengecekan akhiran email
	// if !strings.HasSuffix(email, ".com") {
	// 	return false
	// }

	// Pengecekan apakah email mengandung karakter "@"
	if !strings.Contains(email, "@") {
		return false
	}

	// Jika melewati kedua pengecekan di atas, email dianggap valid
	return true
}

func IsDayValid(day string) bool {
	if !strings.EqualFold(day, "monday") && !strings.EqualFold(day, "tuesday") && !strings.EqualFold(day, "wednesday") && !strings.EqualFold(day, "thursday") && !strings.EqualFold(day, "friday") {
		return false
	}

	return true
}
