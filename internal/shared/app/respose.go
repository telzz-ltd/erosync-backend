package app

import (
	"encoding/json"
	"net/http"
)

type H map[string]any

func JSON(w http.ResponseWriter, status int, data any) {
	dataByte, _ := json.Marshal(data)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(dataByte)
}
