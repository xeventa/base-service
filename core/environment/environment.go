package environment

import (
	"os"
)

// Config holds general environment parameters for the service.
// Values are loaded from OS env vars with sensible defaults.
// Adjust as needed for your deployment.
type Config struct {
	AppName     string
	Env         string // e.g., development, staging, production
	HTTPPort    string // e.g., ":8080"
	DatabaseURL string // PostgreSQL DSN
}

// Load reads environment variables and returns a Config.
// Defaults:
// - AppName: BASE_SERVICE
// - Env: development
// - HTTPPort: :8080
// - DatabaseURL: empty (optional)
func Load() Config {
	cfg := Config{
		AppName:     getEnv("APP_NAME", "BASE_SERVICE"),
		Env:         getEnv("APP_ENV", "development"),
		HTTPPort:    getEnv("HTTP_PORT", ":8080"),
		DatabaseURL: getEnv("DATABASE_URL", "root:GgSfGkOmQ~HyemuLzvJZM~THJ0joWfk9@tcp(tramway.proxy.rlwy.net:25766)/?parseTime=true&charset=utf8mb4&loc=Local"),
	}
	return cfg
}

func getEnv(key, def string) string {
	v := os.Getenv(key)
	if v == "" {
		return def
	}
	return v
}
