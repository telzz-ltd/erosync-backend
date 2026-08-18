package api

import (
	"encoding/json"
	"net/http"
)

type Map map[string]any

const (
	CodeOK           = "000000"
	CodeUnknownError = "000099"
	CodeInvalidBody  = "000016"
)

type APIResponse struct {
	Message string `json:"message"`
	Code    string `json:"code"`
	Data    any    `json:"data"`
}

func OK(w http.ResponseWriter, data any) {
	writeJSON(w, http.StatusOK, APIResponse{
		Message: "success",
		Code:    CodeOK,
		Data:    data,
	})
}

func BadRequest(w http.ResponseWriter, message string) {
	writeJSON(w, http.StatusBadRequest, APIResponse{
		Message: message,
		Code:    CodeInvalidBody,
		Data:    nil,
	})
}

func ValidationError(w http.ResponseWriter, object any) {
	writeJSON(w, http.StatusBadRequest, APIResponse{
		Message: "validation error",
		Code:    CodeInvalidBody,
		Data:    object,
	})
}

func ServerError(w http.ResponseWriter, message string) {
	writeJSON(w, 500, APIResponse{
		Message: message,
		Code:    CodeUnknownError,
		Data:    nil,
	})
}

func Error(w http.ResponseWriter, status int, message string, code string) {
	if code == "" {
		code = CodeUnknownError
	}

	writeJSON(w, status, APIResponse{
		Message: message,
		Code:    code,
		Data:    nil,
	})
}

func writeJSON(w http.ResponseWriter, statusCode int, response APIResponse) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	_ = json.NewEncoder(w).Encode(response)
}
