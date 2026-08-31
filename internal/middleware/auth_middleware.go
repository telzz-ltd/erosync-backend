package middleware

import (
	"erosync/internal/shared/app"
	"erosync/internal/shared/security"
	"log"
	"net/http"
	"strings"
)

func Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("authorization")
		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer") {
			app.JSON(w, 401, app.H{"message": "unauthorized"})
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		if tokenString == "" {
			app.JSON(w, 401, app.H{"message": "unauthorized"})
			return
		}

		userId, role, err := security.VerifyToken(tokenString)
		if err != nil {
			log.Println(err)
			app.JSON(w, 401, app.H{"message": "unauthorized"})
			return
		}

		app.SetValue(r, "userId", userId)
		app.SetValue(r, "role", role)
		next.ServeHTTP(w, r)
	})
}
