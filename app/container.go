package app

import (
	"database/sql"
	"net/http"

	"github.com/rs/zerolog"

	"github.com/xeventa/base-service/core/db"
	"github.com/xeventa/base-service/core/environment"
	"github.com/xeventa/base-service/core/logger"
	httpPkg "github.com/xeventa/base-service/src/public"
)

// Container wires dependencies for the application following clean architecture boundaries.
// It exposes the HTTP handler for the public API.

type Container struct {
	Config  environment.Config
	Logger  zerolog.Logger
	DB      *sql.DB
	Handler http.Handler
}

func New() (*Container, error) {
	// Load environment
	cfg := environment.Load()
	// Logger
	log := logger.New(cfg.Env)
	// Optional DB (MySQL/MariaDB)
	database, err := db.NewMySQL(cfg.DatabaseURL)
	if err != nil {
		log.Error().Err(err).Msg("failed to init database; continuing without DB")
		database = nil
	}
	// Layered wiring within http package
	dbp := httpPkg.SQLDBPinger{DB: database}
	svc := httpPkg.NewHealthService(log, dbp, cfg.AppName, cfg.Env)
	h := httpPkg.NewHealthHandler(svc)
	// Router
	r := httpPkg.Router(h)
	return &Container{
		Config:  cfg,
		Logger:  log,
		DB:      database,
		Handler: r,
	}, nil
}
