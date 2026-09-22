package app

import (
	"encoding/json"
	"net/http"
)

type Map map[string]any

func JSON(w http.ResponseWriter, status int, data any) {
	dataByte, _ := json.Marshal(data)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(dataByte)
}

func Error(w http.ResponseWriter, err error) {
	JSON(w, 500, Map{"message": err.Error()})
}
