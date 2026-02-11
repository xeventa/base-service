package public

// HealthResponse is the public response for /health.
type HealthResponse struct {
	Status string `json:"status"`
	App    string `json:"app"`
	Env    string `json:"env"`
}

// HealthDBResponse is the public response for /health/db.
type HealthDBResponse struct {
	Status string `json:"status"` // ok | error | skipped
	App    string `json:"app"`
	Env    string `json:"env"`
}
