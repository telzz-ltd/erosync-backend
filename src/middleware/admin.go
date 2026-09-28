package middleware

import (
	"erosync/internal/domain"
	"erosync/pkg/response"
	"net/http"
	"slices"
	"strings"
)

func Admin(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		acceptedRoles := []domain.UserRole{
			domain.UserRoleAdmin,
			domain.UserRoleModerator,
		}

		role, ok := r.Context().Value(UserRoleKey).(string)
		if !ok || strings.TrimSpace(role) == "" {
			response.Error(w, 401, "unauthorized", nil)
			return
		}

		if !slices.Contains(acceptedRoles, domain.UserRole(role)) {
			response.Error(w, 403, response.MsgForbidden, nil)
			return
		}

		next.ServeHTTP(w, r)
	})
}
