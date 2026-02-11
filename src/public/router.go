package public

import (
	chi "github.com/go-chi/chi/v5"
)

// Router aggregates public handlers.
// Additional routes can be mounted here following clean architecture boundaries.
func Router(healthHandler *HealthHandler) *chi.Mux {
	r := chi.NewRouter()
	// Public endpoints
	r.Get(RouteHealth, healthHandler.Health)
	r.Get(RouteHealthDB, healthHandler.HealthDB)
	return r
}
