//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"assignment03/internal/biz"
	"assignment03/internal/conf"
	"assignment03/internal/data"
	"assignment03/internal/server"
	"assignment03/internal/service"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

// initApp init kratos application.
func initApp(*conf.Server, *conf.Data, log.Logger) (*kratos.App, func(), error) {
	panic(wire.Build(server.ProviderSet, data.ProviderSet, biz.ProviderSet, service.ProviderSet, newApp))
}
