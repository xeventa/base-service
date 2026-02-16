//go:build wireinject
// +build wireinject

package app

import (
	"github.com/google/wire"
	"github.com/xeventa/base-service/core/contract"
	"github.com/xeventa/base-service/core/db"
	"github.com/xeventa/base-service/core/environment"
)

var (
	AppModule = wire.NewSet(
		ProvideConfig,
		ProvideLogger,
		ProvideMySQL,
	)
)

func InjectAppConfig() (*environment.Config, error) {
	panic(wire.Build(AppModule))
}

func InjectLogger() (*environment.Config, error) {
	panic(wire.Build(AppModule))
}

func InjectMySQL() (*db.MySQL, error) {
	panic(wire.Build(AppModule))
}

func InjectRoutes() ([]contract.IRoute, error) {
	panic(wire.Build(AppModule, DomainSet))
}
