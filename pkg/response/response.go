package response

import (
	"encoding/json"
	"net/http"
)

type Map map[string]any

var (
	MsgInvalidBody = "invalid request body"
	MsgUnknown     = "an unknown error occurred"
)

func JSON(w http.ResponseWriter, statusCode int, data any) {
	w.Header().Set("Content-Type", "application/json")

	dataByte, err := json.Marshal(data)
	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte(`{"message": "internal server error"}`))
		return
	}

	w.WriteHeader(statusCode)
	w.Write(dataByte)
}

func Success(w http.ResponseWriter, data any) {
	JSON(w, 200, Map{"message": "success", "data": data})
}

func Error(w http.ResponseWriter, statusCode int, message string, errData any) {
	if errData != nil {
		JSON(w, statusCode, Map{"message": message, "errors": errData})
	} else {
		JSON(w, statusCode, Map{"message": message})
	}
}

func Unauthenticated(w http.ResponseWriter) {
	Error(w, 401, "Unauthenticated", nil)
}

func Denied(w http.ResponseWriter) {
	Error(w, 403, "Insufficient Permission", nil)
}

func ServerError(w http.ResponseWriter, message string) {
	Error(w, 500, message, nil)
}
