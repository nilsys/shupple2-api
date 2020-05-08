//+build wireinject

package main

import (
	"github.com/google/wire"
	"github.com/labstack/echo/v4"
	"github.com/stayway-corp/stayway-media-api/pkg/adaptor/api"
	"github.com/stayway-corp/stayway-media-api/pkg/adaptor/api/converter"
	"github.com/stayway-corp/stayway-media-api/pkg/adaptor/api/middleware"
	"github.com/stayway-corp/stayway-media-api/pkg/adaptor/infrastructure/client"
	"github.com/stayway-corp/stayway-media-api/pkg/adaptor/infrastructure/repository"
	"github.com/stayway-corp/stayway-media-api/pkg/application/scenario"
	"github.com/stayway-corp/stayway-media-api/pkg/application/service"
	"github.com/stayway-corp/stayway-media-api/pkg/config"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/factory"
	domain_service "github.com/stayway-corp/stayway-media-api/pkg/domain/service"
)

var controllerSet = wire.NewSet(
	converter.ConvertersSet,
	api.PostQueryControllerSet,
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
	api.ThemeQueryControllerSet,
	api.WordpressCallbackControllerSet,
	api.SitemapControllerSet,
	api.S3CommandControllerSet,
	api.InterestQueryControllerSet,
	api.AreaQueryControllerSet,
	api.InnQueryControllerSet,
	api.NoticeQueryControllerSet,
	api.ReportCommandControllerSet,
)

var scenarioSet = wire.NewSet(
	scenario.ReviewCommandScenarioSet,
)

var domainServiceSet = wire.NewSet(
	domain_service.NoticeDomainServiceSet,
	domain_service.TaggedUserDomainServiceSet,
)

var serviceSet = wire.NewSet(
	service.PostQueryServiceSet,
	service.PostCommandServiceSet,
	service.PostFavoriteCommandServiceSet,
	service.CategoryQueryServiceSet,
	service.CategoryCommandServiceSet,
	service.AreaCategoryQueryServiceSet,
	service.AreaCategoryCommandServiceSet,
	service.ThemeCategoryQueryServiceSet,
	service.ThemeCategoryCommandServiceSet,
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
	service.SpotCategoryCommandServiceSet,
	service.SitemapServiceSet,
	service.WordpressCallbackServiceSet,
	service.UserQueryServiceSet,
	service.UserCommandServiceSet,
	service.S3CommandServiceSet,
	service.ProvideAuthService,
	service.InterestQueryServiceSet,
	service.InnQueryServiceSet,
	service.NoticeQueryServiceSet,
	service.ReportCommandServiceSet,
)

var factorySet = wire.NewSet(
	factory.S3SignatureFactorySet,
)

func InitializeApp(configFilePath config.FilePath) (*App, error) {
	wire.Build(
		echo.New,
		wire.Struct(new(App), "*"),
		config.GetConfig,
		wire.FieldsOf(new(config.Stayway), "Metasearch", "Media"),
		client.NewClient,
		wire.Value(&client.Config{}),
		wire.FieldsOf(new(*config.Config), "Wordpress", "Stayway", "AWS", "Slack", "Env"),
		middleware.AuthorizeSet,
		controllerSet,
		scenarioSet,
		domainServiceSet,
		serviceSet,
		factorySet,
		repository.RepositoriesSet,
		repository.ProvideS3Uploader,
	)

	return new(App), nil
}
