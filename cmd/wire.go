//+build wireinject

package main

import (
	"github.com/google/wire"
	"github.com/labstack/echo/v4"
	"github.com/stayway-corp/stayway-media-api/pkg/adaptor/api"
	"github.com/stayway-corp/stayway-media-api/pkg/adaptor/infrastructure/repository"
	"github.com/stayway-corp/stayway-media-api/pkg/application/service"
	"github.com/stayway-corp/stayway-media-api/pkg/config"
)

var controllerSet = wire.NewSet(
	api.PostQueryControllerSet,
	api.PostCommandControllerSet,
)

var repositorySet = wire.NewSet(
	repository.ProvideDB,
	repository.PostQueryRepositorySet,
	repository.PostCommandRepositorySet,
)

var serviceSet = wire.NewSet(
	service.PostQueryServiceSet,
	service.PostCommandServiceSet,
)

type App struct {
	Config                *config.Config
	Echo                  *echo.Echo
	PostCommandController api.PostCommandController
	PostQueryController   api.PostQueryController
}

func InitializeApp(configFilePath config.ConfigFilePath) (*App, error) {
	wire.Build(
		echo.New,
		wire.Struct(new(App), "*"),
		config.GetConfig,
		controllerSet,
		serviceSet,
		repositorySet,
	)

	return new(App), nil
}
