package helpers

import (
	"encoding/json"
	"net/http"

	"github.com/agusheryanto182/go-schedule/models/web"
)

func ReadFromRequestBody(req *http.Request, result interface{}) {
	decoder := json.NewDecoder(req.Body)
	err := decoder.Decode(result)
	PanicIfError(err)
}

func WriteToResponseBody(res http.ResponseWriter, result interface{}) {
	res.Header().Add("Content-Type", "application/json")
	encoder := json.NewEncoder(res)
	err := encoder.Encode(result)
	PanicIfError(err)
}

func WriteToResponseBodyError(res http.ResponseWriter, code int, err string) {
	res.WriteHeader(code)
	webResponse := web.Response{
		Status:  http.StatusText(code),
		Message: err,
	}

	WriteToResponseBody(res, webResponse)
}
