package app

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

func ShouldBindJSON[T any](r *http.Request) (t T, err error) {
	if err := json.NewDecoder(r.Body).Decode(&t); err != nil {
		return t, err
	}
	fmt.Println("Req: ", t)

	if err := Validate.Struct(&t); err != nil {
		return t, err
	}

	return t, nil
}

func GetValue(r *http.Request, key string) any {
	return r.Context().Value(key)
}

func SetValue(r *http.Request, key string, value any) {
	r.WithContext(context.WithValue(r.Context(), key, value))
}
