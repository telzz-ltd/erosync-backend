package handler

import (
	"encoding/json"
	"erosync/internal/schema"
	"erosync/pkg/response"
	"net/http"
)

func (h *Handler) CreateBrand(w http.ResponseWriter, r *http.Request) {
	var req schema.CreateBrandRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.Error(w, 400, response.MsgInvalidBody, nil)
		return
	}

	if err := h.app.Validator.ValidateStruct(&req); err != nil {
		response.Error(w, 400, response.MsgInvalidBody, err)
		return
	}

	brand, err := h.app.Brands.Create(r.Context(), req)
	if err != nil {
		response.Error(w, 500, err.Error(), nil)
		return
	}

	response.JSON(w, 201, brand)
}

func (h *Handler) GetBrandCategories(w http.ResponseWriter, r *http.Request) {
	params := map[string]any{}

	if name := r.URL.Query().Get("name"); name != "" {
		params["name"] = name
	}

	categories, err := h.app.Brands.GetCategories(params)
	if err != nil {
		response.Error(w, 500, err.Error(), nil)
		return
	}

	response.OK(w, categories)
}
