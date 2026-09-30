package middleware

import (
	"context"
	"log"
	"net/http"
	"strings"

	"github.com/telzz/erosync-api/internal/service"
	"github.com/telzz/erosync-api/pkg/response"
)

var UserIDKey = "userID"

func Auth(next http.Handler, jwt *service.JwtService) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := strings.TrimSpace(r.Header.Get("Authorization"))
		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer") {
			response.Error(w, 401, "absent or invalid auth token", nil)
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		userId, err := jwt.ValidateToken(r.Context(), tokenString)
		if err != nil {
			log.Println("JWT verification failed:", err)
			response.Error(w, 401, "absent or invalid auth token", nil)
			return
		}

		ctx := context.WithValue(r.Context(), UserIDKey, userId)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
