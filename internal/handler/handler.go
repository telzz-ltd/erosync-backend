package handler

import (
	"erosync/internal/shared/app"
	"net/http"
)

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	app.JSON(w, 200, app.H{"message": "app working fine"})
}
