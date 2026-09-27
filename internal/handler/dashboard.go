package handler

import (
	"erosync/pkg/response"
	"net/http"
)

func (h *Handler) DashboardStats(w http.ResponseWriter, r *http.Request) {
	response.JSON(w, 200, response.Map{
		"products":       200,
		"quotes":         29,
		"orders":         45,
		"pendingQuotes":  20,
		"revenue":        245_000,
		"projects":       20,
		"pendingTickets": 10,

		"recentQuotes": []map[string]any{},
	})
}
