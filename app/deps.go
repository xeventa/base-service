package app

import (
	"database/sql"
	"net/http"

	"github.com/rs/zerolog"

	"github.com/xeventa/base-service/core/environment"
)

// Deps aggregates app-level dependencies wired via Wire.
type Deps struct {
	Config  environment.Config
	Logger  zerolog.Logger
	DB      *sql.DB
	Handler http.Handler
}
