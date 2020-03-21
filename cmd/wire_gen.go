// Code generated by Wire. DO NOT EDIT.

//go:generate wire
//+build !wireinject

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
	authService, err := service.ProvideAuthService(configConfig)
	if err != nil {
		return nil, err
	}
	db, err := repository.ProvideDB(configConfig)
	if err != nil {
		return nil, err
	}
	userQueryRepositoryImpl := &repository.UserQueryRepositoryImpl{
		DB: db,
	}
	authorizeWrapper := middleware.NewAuthorizeWrapper(configConfig, authService, userQueryRepositoryImpl)
	dao := repository.DAO{
		DB_: db,
	}
	postCommandRepositoryImpl := &repository.PostCommandRepositoryImpl{
		DAO: dao,
	}
	hashtagCommandRepositoryImpl := &repository.HashtagCommandRepositoryImpl{
		DAO: dao,
	}
	wordpress := configConfig.Wordpress
	wordpressQueryRepository := repository.NewWordpressQueryRepositoryImpl(wordpress)
	categoryQueryRepositoryImpl := &repository.CategoryQueryRepositoryImpl{
		DB: db,
	}
	hashtagQueryRepositoryImpl := &repository.HashtagQueryRepositoryImpl{
		DB: db,
	}
	hashtagCommandServiceImpl := &service.HashtagCommandServiceImpl{
		HashtagQueryRepository:   hashtagQueryRepositoryImpl,
		HashtagCommandRepository: hashtagCommandRepositoryImpl,
	}
	wordpressServiceImpl := &service.WordpressServiceImpl{
		WordpressQueryRepository: wordpressQueryRepository,
		UserQueryRepository:      userQueryRepositoryImpl,
		CategoryQueryRepository:  categoryQueryRepositoryImpl,
		HashtagCommandService:    hashtagCommandServiceImpl,
	}
	transactionServiceImpl := &repository.TransactionServiceImpl{
		DB: db,
	}
	postCommandServiceImpl := &service.PostCommandServiceImpl{
		PostCommandRepository:    postCommandRepositoryImpl,
		HashtagCommandRepository: hashtagCommandRepositoryImpl,
		WordpressQueryRepository: wordpressQueryRepository,
		WordpressService:         wordpressServiceImpl,
		TransactionService:       transactionServiceImpl,
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
	postFavoriteCommandRepositoryImpl := &repository.PostFavoriteCommandRepositoryImpl{
		DAO: dao,
	}
	postFavoriteQueryRepositoryImpl := &repository.PostFavoriteQueryRepositoryImpl{
		DB: db,
	}
	postFavoriteCommandServiceImpl := &service.PostFavoriteCommandServiceImpl{
		PostFavoriteCommandRepository: postFavoriteCommandRepositoryImpl,
		PostFavoriteQueryRepository:   postFavoriteQueryRepositoryImpl,
		PostQueryRepository:           postQueryRepositoryImpl,
		PostCommandRepository:         postCommandRepositoryImpl,
		TransactionService:            transactionServiceImpl,
	}
	postFavoriteCommandController := api.PostFavoriteCommandController{
		PostQueryService:           postQueryServiceImpl,
		PostFavoriteCommandService: postFavoriteCommandServiceImpl,
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
	session, err := repository.ProvideAWSSession(configConfig)
	if err != nil {
		return nil, err
	}
	aws := configConfig.AWS
	reviewCommandRepositoryImpl := &repository.ReviewCommandRepositoryImpl{
		DAO:        dao,
		AWSSession: session,
		AWSConfig:  aws,
	}
	touristSpotCommandRepositoryImpl := &repository.TouristSpotCommandRepositoryImpl{
		DAO: dao,
	}
	reviewCommandServiceImpl := &service.ReviewCommandServiceImpl{
		ReviewQueryRepository:        reviewQueryRepositoryImpl,
		ReviewCommandRepository:      reviewCommandRepositoryImpl,
		HashtagCommandRepository:     hashtagCommandRepositoryImpl,
		CategoryQueryRepository:      categoryQueryRepositoryImpl,
		InnQueryRepository:           innQueryRepositoryImpl,
		TouristSpotCommandRepository: touristSpotCommandRepositoryImpl,
		TransactionService:           transactionServiceImpl,
	}
	reviewCommandScenarioImpl := &scenario.ReviewCommandScenarioImpl{
		ReviewCommandService:  reviewCommandServiceImpl,
		HashtagCommandService: hashtagCommandServiceImpl,
	}
	reviewCommandController := api.ReviewCommandController{
		ReviewCommandScenario: reviewCommandScenarioImpl,
		ReviewCommandService:  reviewCommandServiceImpl,
	}
	reviewFavoriteCommandRepositoryImpl := &repository.ReviewFavoriteCommandRepositoryImpl{
		DAO: dao,
	}
	reviewFavoriteQueryRepositoryImpl := &repository.ReviewFavoriteQueryRepositoryImpl{
		DB: db,
	}
	reviewFavoriteCommandServiceImpl := &service.ReviewFavoriteCommandServiceImpl{
		ReviewFavoriteCommandRepository: reviewFavoriteCommandRepositoryImpl,
		ReviewFavoriteQueryRepository:   reviewFavoriteQueryRepositoryImpl,
		ReviewQueryRepository:           reviewQueryRepositoryImpl,
		ReviewCommandRepository:         reviewCommandRepositoryImpl,
		TransactionService:              transactionServiceImpl,
	}
	reviewFavoriteCommandController := api.ReviewFavoriteCommandController{
		ReviewQueryService:           reviewQueryServiceImpl,
		ReviewFavoriteCommandService: reviewFavoriteCommandServiceImpl,
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
	uploader := repository.ProvideS3Uploader(session)
	userCommandRepositoryImpl := &repository.UserCommandRepositoryImpl{
		DB:            db,
		MediaUploader: uploader,
		AWSConfig:     aws,
	}
	userCommandServiceImpl := &service.UserCommandServiceImpl{
		UserCommandRepository: userCommandRepositoryImpl,
		UserQueryRepository:   userQueryRepositoryImpl,
	}
	userCommandController := api.UserCommandController{
		UserCommandService: userCommandServiceImpl,
		AuthService:        authService,
	}
	healthCheckRepositoryImpl := &repository.HealthCheckRepositoryImpl{
		DB: db,
	}
	healthCheckController := api.HealthCheckController{
		HealthCheckRepository: healthCheckRepositoryImpl,
	}
	categoryCommandRepositoryImpl := &repository.CategoryCommandRepositoryImpl{
		DB: db,
	}
	categoryCommandServiceImpl := &service.CategoryCommandServiceImpl{
		CategoryCommandRepository: categoryCommandRepositoryImpl,
		WordpressQueryRepository:  wordpressQueryRepository,
		WordpressService:          wordpressServiceImpl,
	}
	comicCommandRepositoryImpl := &repository.ComicCommandRepositoryImpl{
		DB: db,
	}
	comicCommandServiceImpl := &service.ComicCommandServiceImpl{
		ComicCommandRepository:   comicCommandRepositoryImpl,
		WordpressQueryRepository: wordpressQueryRepository,
		WordpressService:         wordpressServiceImpl,
	}
	featureCommandRepositoryImpl := &repository.FeatureCommandRepositoryImpl{
		DB: db,
	}
	featureCommandServiceImpl := &service.FeatureCommandServiceImpl{
		FeatureCommandRepository: featureCommandRepositoryImpl,
		WordpressQueryRepository: wordpressQueryRepository,
		WordpressService:         wordpressServiceImpl,
	}
	lcategoryCommandRepositoryImpl := &repository.LcategoryCommandRepositoryImpl{
		DB: db,
	}
	lcategoryCommandServiceImpl := &service.LcategoryCommandServiceImpl{
		LcategoryCommandRepository: lcategoryCommandRepositoryImpl,
		WordpressQueryRepository:   wordpressQueryRepository,
		WordpressService:           wordpressServiceImpl,
	}
	touristSpotCommandServiceImpl := &service.TouristSpotCommandServiceImpl{
		TouristSpotCommandRepository: touristSpotCommandRepositoryImpl,
		WordpressQueryRepository:     wordpressQueryRepository,
		WordpressService:             wordpressServiceImpl,
	}
	vlogCommandRepositoryImpl := &repository.VlogCommandRepositoryImpl{
		DB: db,
	}
	vlogCommandServiceImpl := &service.VlogCommandServiceImpl{
		VlogCommandRepository:    vlogCommandRepositoryImpl,
		WordpressQueryRepository: wordpressQueryRepository,
		WordpressService:         wordpressServiceImpl,
	}
	wordpressCallbackServiceImpl := &service.WordpressCallbackServiceImpl{
		CategoryCommandService:    categoryCommandServiceImpl,
		ComicCommandService:       comicCommandServiceImpl,
		FeatureCommandService:     featureCommandServiceImpl,
		LcategoryCommandService:   lcategoryCommandServiceImpl,
		PostCommandService:        postCommandServiceImpl,
		TouristSpotCommandService: touristSpotCommandServiceImpl,
		VlogCommandService:        vlogCommandServiceImpl,
	}
	wordpressCallbackController := api.WordpressCallbackController{
		WordpressCallbackService: wordpressCallbackServiceImpl,
	}
	s3SignatureFactory := factory.S3SignatureFactory{
		Session:   session,
		AWSConfig: aws,
	}
	s3CommandServiceImpl := &service.S3CommandServiceImpl{
		S3SignatureFactory: s3SignatureFactory,
	}
	s3CommandController := api.S3CommandController{
		S3CommandService: s3CommandServiceImpl,
	}
	touristSpotQueryServiceImpl := &service.TouristSpotQueryServiceImpl{
		TouristSpotQueryRepository: touristSpotQueryRepositoryImpl,
	}
	touristSpotQueryController := api.TouristSpotQueryController{
		TouristSpotQueryService: touristSpotQueryServiceImpl,
	}
	interestQueryRepositoryImpl := &repository.InterestQueryRepositoryImpl{
		DB: db,
	}
	interestQueryServiceImpl := &service.InterestQueryServiceImpl{
		InterestQueryRepository: interestQueryRepositoryImpl,
	}
	interestQueryController := api.InterestQueryController{
		InterestQueryService: interestQueryServiceImpl,
	}
	app := &App{
		Config:                          configConfig,
		Echo:                            echoEcho,
		AuthorizeWrapper:                authorizeWrapper,
		PostCommandController:           postCommandController,
		PostQueryController:             postQueryController,
		PostFavoriteCommandController:   postFavoriteCommandController,
		CategoryQueryController:         categoryQueryController,
		ComicQueryController:            comicQueryController,
		ReviewQueryController:           reviewQueryController,
		ReviewCommandController:         reviewCommandController,
		ReviewFavoriteCommandController: reviewFavoriteCommandController,
		HashtagQueryController:          hashtagQueryController,
		SearchQueryController:           searchQueryController,
		FeatureQueryController:          featureQueryController,
		VlogQueryController:             vlogQueryController,
		UserQueryController:             userQueryController,
		UserCommandController:           userCommandController,
		HealthCheckController:           healthCheckController,
		WordpressCallbackController:     wordpressCallbackController,
		S3CommandController:             s3CommandController,
		TouristSpotQueryController:      touristSpotQueryController,
		InteresetQueryController:        interestQueryController,
	}
	return app, nil
}

var (
	_wireConfigValue = &client.Config{}
)

// wire.go:

var controllerSet = wire.NewSet(api.PostQueryControllerSet, api.PostCommandControllerSet, api.PostFavoriteCommandControllerSet, api.CategoryQueryControllerSet, api.ComicQueryControllerSet, api.ReviewQueryControllerSet, api.ReviewCommandControllerSet, api.ReviewFavoriteCommandControllerSet, api.TouristSpotQeuryControllerSet, api.SearchQueryControllerSet, api.FeatureQueryControllerSet, api.VlogQueryControllerSet, api.HashtagQueryControllerSet, api.UserQueryControllerSet, api.UserCommandControllerSet, api.HealthCheckControllerSet, api.WordpressCallbackControllerSet, api.S3CommandControllerSet, api.InterestQueryControllerSet)

var scenarioSet = wire.NewSet(scenario.ReviewCommandScenarioSet)

var serviceSet = wire.NewSet(service.PostQueryServiceSet, service.PostCommandServiceSet, service.PostFavoriteCommandServiceSet, service.CategoryQueryServiceSet, service.CategoryCommandServiceSet, service.ComicQueryServiceSet, service.ComicCommandServiceSet, service.ReviewQueryServiceSet, service.ReviewCommandServiceSet, service.ReviewFavoriteCommandServiceSet, service.WordpressServiceSet, service.TouristSpotQueryServiceSet, service.SearchQueryServiceSet, service.FeatureQueryServiceSet, service.FeatureCommandServiceSet, service.VlogQueryServiceSet, service.VlogCommandServiceSet, service.HashtagQueryServiceSet, service.HashtagCommandServiceSet, service.TouristSpotCommandServiceSet, service.LcategoryCommandServiceSet, service.WordpressCallbackServiceSet, service.UserQueryServiceSet, service.UserCommandServiceSet, service.S3CommandServiceSet, service.ProvideAuthService, service.InterestQueryServiceSet)

var factorySet = wire.NewSet(factory.S3SignatureFactorySet)
