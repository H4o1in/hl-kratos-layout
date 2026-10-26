//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"log/slog"

	"gin-gorm-kratos-layout/internal/biz"
	"gin-gorm-kratos-layout/internal/conf"
	"gin-gorm-kratos-layout/internal/data"
	"gin-gorm-kratos-layout/internal/server"
	"gin-gorm-kratos-layout/internal/service"

	"github.com/go-kratos/kratos/v3"
	"github.com/google/wire"
)

// wireApp init kratos application.
func wireApp(*conf.Server, *conf.Data, *slog.Logger) (*kratos.App, func(), error) {
	panic(wire.Build(server.ProviderSet, data.ProviderSet, biz.ProviderSet, service.ProviderSet, newApp))
}
