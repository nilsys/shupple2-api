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
)

var controllerSet = wire.NewSet(
	api.PostQueryControllerSet,
	api.PostCommandControllerSet,
	api.CategoryQueryControllerSet,
	api.ComicQueryControllerSet,
	api.ReviewQueryControllerSet,
	api.SearchQueryControllerSet,
	api.FeatureQueryControllerSet,
	api.VlogQueryControllerSet,
	api.HashtagQueryControllerSet,
	api.UserQueryControllerSet,
	api.HealthCheckControllerSet,
)

var serviceSet = wire.NewSet(
	service.PostQueryServiceSet,
	service.PostCommandServiceSet,
	service.CategoryQueryServiceSet,
	service.ComicQueryServiceSet,
	service.ReviewQueryServiceSet,
	service.WordpressServiceSet,
	service.SearchQueryServiceSet,
	service.FeatureQueryServiceSet,
	service.VlogQueryServiceSet,
	service.HashtagQueryServiceSet,
	service.UserQueryServiceSet,
)

var configSet = wire.FieldsOf(new(*config.Config), "Stayway")

type App struct {
	Config                  *config.Config
	Echo                    *echo.Echo
	PostCommandController   api.PostCommandController
	PostQueryController     api.PostQueryController
	CategoryQueryController api.CategoryQueryController
	ComicQueryController    api.ComicQueryController
	ReviewQueryController   api.ReviewQueryController
	HashtagQueryController  api.HashtagQueryController
	SearchQueryController   api.SearchQueryController
	FeatureQueryController  api.FeatureQueryController
	VlogQueryController     api.VlogQueryController
	UserQueryController     api.UserQueryController
	HealthCheckController   api.HealthCheckController
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
		repository.RepositoriesSet,
	)

	return new(App), nil
}
