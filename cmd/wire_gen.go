// Code generated by Wire. DO NOT EDIT.

//go:generate wire
//+build !wireinject

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

import (
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// Injectors from wire.go:

func InitializeApp(configFilePath2 config.ConfigFilePath) (*App, error) {
	configConfig, err := config.GetConfig(configFilePath2)
	if err != nil {
		return nil, err
	}
	echoEcho := echo.New()
	db, err := repository.ProvideDB(configConfig)
	if err != nil {
		return nil, err
	}
	postCommandRepositoryImpl := &repository.PostCommandRepositoryImpl{
		DB: db,
	}
	wordpress := configConfig.Wordpress
	wordpressQueryRepository := repository.NewWordpressQueryRepositoryImpl(wordpress)
	userQueryRepositoryImpl := &repository.UserQueryRepositoryImpl{
		DB: db,
	}
	categoryQueryRepositoryImpl := &repository.CategoryQueryRepositoryImpl{
		DB: db,
	}
	wordpressServiceImpl := &service.WordpressServiceImpl{
		WordpressQueryRepository: wordpressQueryRepository,
		UserQueryRepository:      userQueryRepositoryImpl,
		CategoryQueryRepository:  categoryQueryRepositoryImpl,
	}
	postCommandServiceImpl := &service.PostCommandServiceImpl{
		PostCommandRepository:    postCommandRepositoryImpl,
		WordpressQueryRepository: wordpressQueryRepository,
		WordpressService:         wordpressServiceImpl,
	}
	postCommandController := api.PostCommandController{
		PostService: postCommandServiceImpl,
	}
	postQueryRepositoryImpl := &repository.PostQueryRepositoryImpl{
		DB: db,
	}
	postQueryServiceImpl := &service.PostQueryServiceImpl{
		PostQueryRepository: postQueryRepositoryImpl,
	}
	postQueryController := api.PostQueryController{
		PostService: postQueryServiceImpl,
	}
	categoryQueryServiceImpl := &service.CategoryQueryServiceImpl{
		Repository: categoryQueryRepositoryImpl,
	}
	categoryQueryController := api.CategoryQueryController{
		CategoryQueryService: categoryQueryServiceImpl,
	}
	comicQueryRepositoryImpl := &repository.ComicQueryRepositoryImpl{
		DB: db,
	}
	comicQueryServiceImpl := &service.ComicQueryServiceImpl{
		ComicQueryRepository: comicQueryRepositoryImpl,
		UserQueryRepository:  userQueryRepositoryImpl,
	}
	comicQueryController := api.ComicQueryController{
		ComicQueryService: comicQueryServiceImpl,
	}
	reviewQueryRepositoryImpl := &repository.ReviewQueryRepositoryImpl{
		DB: db,
	}
	staywayConfig := configConfig.Stayway
	clientConfig := _wireConfigValue
	clientClient := client.NewClient(clientConfig)
	innQueryRepositoryImpl := &repository.InnQueryRepositoryImpl{
		StaywayConfig: staywayConfig,
		Client:        clientClient,
	}
	reviewQueryServiceImpl := &service.ReviewQueryServiceImpl{
		ReviewQueryRepository: reviewQueryRepositoryImpl,
		InnQueryRepository:    innQueryRepositoryImpl,
	}
	reviewQueryController := api.ReviewQueryController{
		ReviewQueryService: reviewQueryServiceImpl,
	}
	hashtagQueryRepositoryImpl := &repository.HashtagQueryRepositoryImpl{
		DB: db,
	}
	hashtagQueryServiceImpl := &service.HashtagQueryServiceImpl{
		HashtagQueryRepository: hashtagQueryRepositoryImpl,
	}
	hashtagQueryController := api.HashtagQueryController{
		HashtagQueryService: hashtagQueryServiceImpl,
	}
	touristSpotQueryRepositoryImpl := &repository.TouristSpotQueryRepositoryImpl{
		DB: db,
	}
	searchQueryServiceImpl := &service.SearchQueryServiceImpl{
		CategoryQueryRepository:    categoryQueryRepositoryImpl,
		TouristSpotQueryRepository: touristSpotQueryRepositoryImpl,
		HashtagQueryRepository:     hashtagQueryRepositoryImpl,
		UserQueryRepository:        userQueryRepositoryImpl,
	}
	searchQueryController := api.SearchQueryController{
		SearchQueryService: searchQueryServiceImpl,
	}
	featureQueryRepositoryImpl := &repository.FeatureQueryRepositoryImpl{
		DB: db,
	}
	featureQueryServiceImpl := &service.FeatureQueryServiceImpl{
		FeatureQueryRepository: featureQueryRepositoryImpl,
	}
	featureQueryController := api.FeatureQueryController{
		FeatureQueryService: featureQueryServiceImpl,
	}
	vlogQueryRepositoryImpl := &repository.VlogQueryRepositoryImpl{
		DB: db,
	}
	vlogQueryServiceImpl := &service.VlogQueryServiceImpl{
		VlogQueryRepository: vlogQueryRepositoryImpl,
	}
	vlogQueryController := api.VlogQueryController{
		VlogQueryService: vlogQueryServiceImpl,
	}
	userQueryServiceImpl := &service.UserQueryServiceImpl{
		UserQueryRepository: userQueryRepositoryImpl,
	}
	userQueryController := api.UserQueryController{
		UserQueryService: userQueryServiceImpl,
	}
	healthCheckRepositoryImpl := &repository.HealthCheckRepositoryImpl{
		DB: db,
	}
	healthCheckController := api.HealthCheckController{
		HealthCheckRepository: healthCheckRepositoryImpl,
	}
	app := &App{
		Config:                  configConfig,
		Echo:                    echoEcho,
		PostCommandController:   postCommandController,
		PostQueryController:     postQueryController,
		CategoryQueryController: categoryQueryController,
		ComicQueryController:    comicQueryController,
		ReviewQueryController:   reviewQueryController,
		HashtagQueryController:  hashtagQueryController,
		SearchQueryController:   searchQueryController,
		FeatureQueryController:  featureQueryController,
		VlogQueryController:     vlogQueryController,
		UserQueryController:     userQueryController,
		HealthCheckController:   healthCheckController,
	}
	return app, nil
}

var (
	_wireConfigValue = &client.Config{}
)

// wire.go:

var controllerSet = wire.NewSet(api.PostQueryControllerSet, api.PostCommandControllerSet, api.CategoryQueryControllerSet, api.ComicQueryControllerSet, api.ReviewQueryControllerSet, api.SearchQueryControllerSet, api.FeatureQueryControllerSet, api.VlogQueryControllerSet, api.HashtagQueryControllerSet, api.UserQueryControllerSet, api.HealthCheckControllerSet)

var serviceSet = wire.NewSet(service.PostQueryServiceSet, service.PostCommandServiceSet, service.CategoryQueryServiceSet, service.ComicQueryServiceSet, service.ReviewQueryServiceSet, service.WordpressServiceSet, service.SearchQueryServiceSet, service.FeatureQueryServiceSet, service.VlogQueryServiceSet, service.HashtagQueryServiceSet, service.UserQueryServiceSet)

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
