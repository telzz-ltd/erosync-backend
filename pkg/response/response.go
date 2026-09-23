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

func JSON(w http.ResponseWriter, status int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if data != nil {
		_ = json.NewEncoder(w).Encode(data)
	}
}

func Error(w http.ResponseWriter, status int, message string, errors any) {
	payload := Map{"error": message}
	if errors != nil {
		payload["details"] = errors
	}
	JSON(w, status, payload)
}
