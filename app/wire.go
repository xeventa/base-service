//go:build wireinject
// +build wireinject

package app

import (
	"github.com/google/wire"
	"github.com/xeventa/base-service/src/public"
)

var (
	DomainSet = wire.NewSet(
		// surrounding domain

		// public domain
		public.NewService,
		wire.Bind(new(public.IService), new(*public.Service)),
		public.NewDelivery,
		public.NewRoute,

		// route provider
		ProvideRoutes,
	)
)
