package app

import (
	"context"
	"encoding/json"
	"net/http"
)

func ShouldBindJSON(r *http.Request, t any) error {
	if err := json.NewDecoder(r.Body).Decode(&t); err != nil {
		return err
	}

	if err := Validate.Struct(t); err != nil {
		return err
	}

	return nil
}

func GetValue(r *http.Request, key string) any {
	return r.Context().Value(key)
}

func SetValue(r *http.Request, key string, value any) {
	r.WithContext(context.WithValue(r.Context(), key, value))
}
