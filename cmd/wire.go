//+build wireinject

package main

import (
	"github.com/google/wire"
	"github.com/labstack/echo/v4"
	"github.com/stayway-corp/stayway-media-api/pkg/adaptor/api"
	"github.com/stayway-corp/stayway-media-api/pkg/adaptor/infrastructure/client"
	"github.com/stayway-corp/stayway-media-api/pkg/adaptor/infrastructure/repository"
	"github.com/stayway-corp/stayway-media-api/pkg/application/scenario"
	"github.com/stayway-corp/stayway-media-api/pkg/application/service"
	"github.com/stayway-corp/stayway-media-api/pkg/config"
)

var controllerSet = wire.NewSet(
	api.PostQueryControllerSet,
	api.PostCommandControllerSet,
	api.CategoryQueryControllerSet,
	api.ComicQueryControllerSet,
	api.ReviewQueryControllerSet,
	api.ReviewCommandControllerSet,
	api.TouristSpotQeuryControllerSet,
	api.SearchQueryControllerSet,
	api.FeatureQueryControllerSet,
	api.VlogQueryControllerSet,
	api.HashtagQueryControllerSet,
	api.UserQueryControllerSet,
	api.HealthCheckControllerSet,
	api.WordpressCallbackControllerSet,
	api.InterestQueryControllerSet,
)

var scenarioSet = wire.NewSet(
	scenario.ReviewCommandScenarioSet,
)

var serviceSet = wire.NewSet(
	service.PostQueryServiceSet,
	service.PostCommandServiceSet,
	service.CategoryQueryServiceSet,
	service.CategoryCommandServiceSet,
	service.ComicQueryServiceSet,
	service.ComicCommandServiceSet,
	service.ReviewQueryServiceSet,
	service.ReviewCommandServiceSet,
	service.WordpressServiceSet,
	service.TouristSpotQueryServiceSet,
	service.SearchQueryServiceSet,
	service.FeatureQueryServiceSet,
	service.FeatureCommandServiceSet,
	service.VlogQueryServiceSet,
	service.VlogCommandServiceSet,
	service.HashtagQueryServiceSet,
	service.HashtagCommandServiceSet,
	service.TouristSpotCommandServiceSet,
	service.LcategoryCommandServiceSet,
	service.WordpressCallbackServiceSet,
	service.UserQueryServiceSet,
	service.InterestQueryServiceSet,
)

var configSet = wire.FieldsOf(new(*config.Config), "Stayway")

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
		scenarioSet,
		serviceSet,
		repository.RepositoriesSet,
	)

	return new(App), nil
}
