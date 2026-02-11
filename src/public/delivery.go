package public

import (
	"encoding/json"
	"net/http"
)

// HealthHandler exposes HTTP endpoints backed by an abstract HealthUseCase.

type HealthHandler struct {
	Svc HealthUseCase
}

func NewHealthHandler(svc HealthUseCase) *HealthHandler {
	return &HealthHandler{Svc: svc}
}

func (h *HealthHandler) Health(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	status, app, env := h.Svc.Health()
	_ = json.NewEncoder(w).Encode(HealthResponse{Status: status, App: app, Env: env})
}

func (h *HealthHandler) HealthDB(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	status, app, env := h.Svc.HealthDB()
	_ = json.NewEncoder(w).Encode(HealthDBResponse{Status: status, App: app, Env: env})
}
