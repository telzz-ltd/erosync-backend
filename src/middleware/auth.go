package middleware

import (
	"context"
	"erosync/internal/service"
	"erosync/pkg/response"
	"net/http"
	"strings"
)

type contextKey string

const UserIDKey contextKey = "userID"
const UserRoleKey contextKey = "userRole"

func Authenticate(jwt *service.JwtService) Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			authHeader := r.Header.Get("Authorization")
			if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
				response.Error(w, http.StatusUnauthorized, "Missing or invalid authorization token", nil)
				return
			}

			token := strings.TrimPrefix(authHeader, "Bearer ")
			if token == "" { // simple check logic
				response.Error(w, http.StatusUnauthorized, "Invalid token", nil)
				return
			}

			// Validate token and extract user metadata...
			payload, err := jwt.VerifyToken(token)
			if err != nil {
				response.Error(w, http.StatusUnauthorized, "invalid authorization token", nil)
			}

			userID := payload.Sub
			userRole := payload.Role

			// Inject user info into request context
			ctx := context.WithValue(r.Context(), UserIDKey, userID)
			ctx = context.WithValue(ctx, UserRoleKey, userRole)

			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
