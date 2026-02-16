package app

import (
	"github.com/rs/zerolog"
	"github.com/xeventa/base-service/core/contract"
	"github.com/xeventa/base-service/core/db"
	"github.com/xeventa/base-service/core/environment"
	"github.com/xeventa/base-service/core/logger"
	"github.com/xeventa/base-service/src/public"
)

// Provider functions (shared by Wire injector)
func ProvideConfig() (*environment.Config, error) {
	return environment.ProvideConfig()
}
func ProvideLogger(config *environment.Config) zerolog.Logger {
	return logger.New(config.AppEnv)
}

func ProvideMySQL(config *environment.Config) *db.MySQL {
	return db.NewMySQL(config)
}

func ProvideRoutes(
	public *public.Route,
) []contract.IRoute {
	return []contract.IRoute{
		public,
		// ... add more routes here
	}
}
