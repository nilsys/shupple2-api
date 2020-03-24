//+build wireinject

package main

import (
	"github.com/google/wire"
	"github.com/labstack/echo/v4"
	"github.com/stayway-corp/stayway-media-api/pkg/adaptor/api"
	"github.com/stayway-corp/stayway-media-api/pkg/adaptor/api/middleware"
	"github.com/stayway-corp/stayway-media-api/pkg/adaptor/infrastructure/client"
	"github.com/stayway-corp/stayway-media-api/pkg/adaptor/infrastructure/repository"
	"github.com/stayway-corp/stayway-media-api/pkg/application/scenario"
	"github.com/stayway-corp/stayway-media-api/pkg/application/service"
	"github.com/stayway-corp/stayway-media-api/pkg/config"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/factory"
)

var controllerSet = wire.NewSet(
	api.PostQueryControllerSet,
	api.PostCommandControllerSet,
	api.PostFavoriteCommandControllerSet,
	api.CategoryQueryControllerSet,
	api.ComicQueryControllerSet,
	api.ReviewQueryControllerSet,
	api.ReviewCommandControllerSet,
	api.ReviewFavoriteCommandControllerSet,
	api.TouristSpotQeuryControllerSet,
	api.SearchQueryControllerSet,
	api.FeatureQueryControllerSet,
	api.VlogQueryControllerSet,
	api.HashtagQueryControllerSet,
	api.HashtagCommandControllerSet,
	api.UserQueryControllerSet,
	api.UserCommandControllerSet,
	api.HealthCheckControllerSet,
	api.WordpressCallbackControllerSet,
	api.S3CommandControllerSet,
	api.InterestQueryControllerSet,
	api.AreaQueryControllerSet,
	api.InnQueryControllerSet,
)

var scenarioSet = wire.NewSet(
	scenario.ReviewCommandScenarioSet,
)

var serviceSet = wire.NewSet(
	service.PostQueryServiceSet,
	service.PostCommandServiceSet,
	service.PostFavoriteCommandServiceSet,
	service.CategoryQueryServiceSet,
	service.CategoryCommandServiceSet,
	service.ComicQueryServiceSet,
	service.ComicCommandServiceSet,
	service.ReviewQueryServiceSet,
	service.ReviewCommandServiceSet,
	service.ReviewFavoriteCommandServiceSet,
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
	service.UserCommandServiceSet,
	service.S3CommandServiceSet,
	service.ProvideAuthService,
	service.InterestQueryServiceSet,
	service.AreaQueryServiceSet,
	service.InnQueryServiceSet,
)

var factorySet = wire.NewSet(
	factory.S3SignatureFactorySet,
)

func InitializeApp(configFilePath config.ConfigFilePath) (*App, error) {
	wire.Build(
		echo.New,
		wire.Struct(new(App), "*"),
		config.GetConfig,
		wire.FieldsOf(new(config.Stayway), "Metasearch", "Media"),
		client.NewClient,
		wire.Value(&client.Config{}),
		wire.FieldsOf(new(*config.Config), "Wordpress", "Stayway", "AWS"),
		middleware.NewAuthorizeWrapper,
		controllerSet,
		scenarioSet,
		serviceSet,
		factorySet,
		repository.RepositoriesSet,
		repository.ProvideS3Uploader,
	)

	return new(App), nil
}
