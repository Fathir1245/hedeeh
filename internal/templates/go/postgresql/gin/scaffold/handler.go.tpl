package handler

import (
	"encoding/json"
	"net/http"

	"{{.ModuleName}}/internal/service"
)

type HealthHandler struct {
	service *service.HealthService
}

func NewHealthHandler() *HealthHandler {
	return &HealthHandler{
		service: service.NewHealthService(),
	}
}

func (h *HealthHandler) Check(w http.ResponseWriter, r *http.Request) {
	response := map[string]string{
		"status": h.service.Status(),
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(response)
}
