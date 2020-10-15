//+build wireinject

package main

import (
	"github.com/google/wire"
	"github.com/labstack/echo/v4"
	"github.com/uma-co82/shupple2-api/pkg/adaptor/api"
	"github.com/uma-co82/shupple2-api/pkg/adaptor/infrastructure/repository"
	"github.com/uma-co82/shupple2-api/pkg/config"
)

var controllerSet = wire.NewSet(
	api.HealthCheckControllerSet,
)

func InitializeApp(configFilePath config.FilePath) (*App, error) {
	wire.Build(
		echo.New,
		controllerSet,
		wire.Struct(new(App), "*"),
		config.GetConfig,
		repository.RepositoriesSet,
	)

	return new(App), nil
}
