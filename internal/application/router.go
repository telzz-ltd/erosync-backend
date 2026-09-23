package application

import (
	"erosync/pkg/response"
	"net/http"
)

func (app *Application) RegisterRouter() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /health", func(w http.ResponseWriter, r *http.Request) {
		response.JSON(w, 200, response.Map{"message": "app working fine"})
	})

	return mux
}
