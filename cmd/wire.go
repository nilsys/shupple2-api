//+build wireinject

package main

import (
	"github.com/google/wire"
	"github.com/labstack/echo/v4"
	"github.com/stayway-corp/stayway-media-api/pkg/adaptor/api"
	"github.com/stayway-corp/stayway-media-api/pkg/adaptor/infrastructure/client"
	"github.com/stayway-corp/stayway-media-api/pkg/adaptor/infrastructure/repository"
	"github.com/stayway-corp/stayway-media-api/pkg/application/service"
	"github.com/stayway-corp/stayway-media-api/pkg/config"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/factory"
)

var controllerSet = wire.NewSet(
	api.PostQueryControllerSet,
	api.PostCommandControllerSet,
	api.ReviewQueryControllerSet,
)

var serviceSet = wire.NewSet(
	service.PostQueryServiceSet,
	service.PostCommandServiceSet,
	service.ReviewQueryServiceSet,
	service.WordpressServiceSet,
)

var factorySet = wire.NewSet(
	factory.PostDetailFactorySet,
)

var configSet = wire.FieldsOf(new(*config.Config), "Stayway")

type App struct {
	Config                *config.Config
	Echo                  *echo.Echo
	PostCommandController api.PostCommandController
	PostQueryController   api.PostQueryController
	ReviewQueryController api.ReviewQueryController
}

func InitializeApp(configFilePath config.ConfigFilePath) (*App, error) {
	wire.Build(
		echo.New,
		wire.Struct(new(App), "*"),
		config.GetConfig,
		configSet,
		client.NewClient,
		wire.Value(&client.Config{}),
		wire.FieldsOf(new(*config.Config), "Wordpress"),
		controllerSet,
		serviceSet,
		factorySet,
		repository.RepositoriesSet,
	)

	return new(App), nil
}
