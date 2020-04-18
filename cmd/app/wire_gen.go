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
	service2 "github.com/stayway-corp/stayway-media-api/pkg/domain/service"
)

import (
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// Injectors from wire.go:

func InitializeApp(configFilePath config.FilePath) (*App, error) {
	configConfig, err := config.GetConfig(configFilePath)
	if err != nil {
		return nil, err
	}
	db, err := repository.ProvideDB(configConfig)
	if err != nil {
		return nil, err
	}
	echoEcho := echo.New()
	authService, err := service.ProvideAuthService(configConfig)
	if err != nil {
		return nil, err
	}
	userQueryRepositoryImpl := &repository.UserQueryRepositoryImpl{
		DB: db,
	}
	authorize := middleware.Authorize{
		AuthService: authService,
		UserRepo:    userQueryRepositoryImpl,
	}
	dao := repository.DAO{
		UnderlyingDB: db,
	}
	postCommandRepositoryImpl := &repository.PostCommandRepositoryImpl{
		DAO: dao,
	}
	hashtagCommandRepositoryImpl := &repository.HashtagCommandRepositoryImpl{
		DAO: dao,
	}
	wordpress := configConfig.Wordpress
	stayway := configConfig.Stayway
	staywayMedia := stayway.Media
	wordpressQueryRepositoryImpl := repository.NewWordpressQueryRepositoryImpl(wordpress, staywayMedia)
	areaCategoryQueryRepositoryImpl := &repository.AreaCategoryQueryRepositoryImpl{
		DB: db,
	}
	themeCategoryQueryRepositoryImpl := &repository.ThemeCategoryQueryRepositoryImpl{
		DB: db,
	}
	spotCategoryQueryRepositoryImpl := &repository.SpotCategoryQueryRepositoryImpl{
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
		WordpressQueryRepository:     wordpressQueryRepositoryImpl,
		UserQueryRepository:          userQueryRepositoryImpl,
		AreaCategoryQueryRepository:  areaCategoryQueryRepositoryImpl,
		ThemeCategoryQueryRepository: themeCategoryQueryRepositoryImpl,
		SpotCategoryQueryRepository:  spotCategoryQueryRepositoryImpl,
		HashtagCommandService:        hashtagCommandServiceImpl,
	}
	transactionServiceImpl := &repository.TransactionServiceImpl{
		DB: db,
	}
	postCommandServiceImpl := &service.PostCommandServiceImpl{
		PostCommandRepository:    postCommandRepositoryImpl,
		HashtagCommandRepository: hashtagCommandRepositoryImpl,
		WordpressQueryRepository: wordpressQueryRepositoryImpl,
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
	noticeCommandRepositoryImpl := &repository.NoticeCommandRepositoryImpl{
		DAO: dao,
	}
	taggedUserDomainServiceImpl := service2.TaggedUserDomainServiceImpl{
		UserQueryRepository: userQueryRepositoryImpl,
	}
	noticeDomainServiceImpl := &service2.NoticeDomainServiceImpl{
		NoticeCommandRepository: noticeCommandRepositoryImpl,
		TaggedUserDomainService: taggedUserDomainServiceImpl,
	}
	postFavoriteCommandServiceImpl := &service.PostFavoriteCommandServiceImpl{
		PostFavoriteCommandRepository: postFavoriteCommandRepositoryImpl,
		PostFavoriteQueryRepository:   postFavoriteQueryRepositoryImpl,
		PostQueryRepository:           postQueryRepositoryImpl,
		PostCommandRepository:         postCommandRepositoryImpl,
		NoticeDomainService:           noticeDomainServiceImpl,
		TransactionService:            transactionServiceImpl,
	}
	postFavoriteCommandController := api.PostFavoriteCommandController{
		PostFavoriteCommandService: postFavoriteCommandServiceImpl,
	}
	categoryQueryServiceImpl := &service.CategoryQueryServiceImpl{
		AreaCategoryRepository:  areaCategoryQueryRepositoryImpl,
		ThemeCategoryRepository: themeCategoryQueryRepositoryImpl,
		SpotCategoryRepository:  spotCategoryQueryRepositoryImpl,
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
	staywayMetasearch := stayway.Metasearch
	clientConfig := _wireConfigValue
	clientClient := client.NewClient(clientConfig)
	innQueryRepositoryImpl := &repository.InnQueryRepositoryImpl{
		MetasearchConfig: staywayMetasearch,
		Client:           clientClient,
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
		InnQueryRepository:           innQueryRepositoryImpl,
		TouristSpotCommandRepository: touristSpotCommandRepositoryImpl,
		NoticeDomainService:          noticeDomainServiceImpl,
		TransactionService:           transactionServiceImpl,
	}
	reviewCommandScenarioImpl := &scenario.ReviewCommandScenarioImpl{
		ReviewQueryService:    reviewQueryServiceImpl,
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
		NoticeDomainService:             noticeDomainServiceImpl,
		TransactionService:              transactionServiceImpl,
	}
	reviewFavoriteCommandController := api.ReviewFavoriteCommandController{
		ReviewFavoriteCommandService: reviewFavoriteCommandServiceImpl,
	}
	hashtagQueryServiceImpl := &service.HashtagQueryServiceImpl{
		HashtagQueryRepository: hashtagQueryRepositoryImpl,
	}
	hashtagQueryController := api.HashtagQueryController{
		HashtagQueryService: hashtagQueryServiceImpl,
	}
	hashtagCommandController := api.HashtagCommandController{
		HashtagCommandService: hashtagCommandServiceImpl,
	}
	touristSpotQueryRepositoryImpl := &repository.TouristSpotQueryRepositoryImpl{
		DB: db,
	}
	searchQueryServiceImpl := &service.SearchQueryServiceImpl{
		AreaCategoryQueryRepository:  areaCategoryQueryRepositoryImpl,
		ThemeCategoryQueryRepository: themeCategoryQueryRepositoryImpl,
		TouristSpotQueryRepository:   touristSpotQueryRepositoryImpl,
		HashtagQueryRepository:       hashtagQueryRepositoryImpl,
		UserQueryRepository:          userQueryRepositoryImpl,
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
		DAO:           dao,
		MediaUploader: uploader,
		AWSConfig:     aws,
		AWSSession:    session,
	}
	userCommandServiceImpl := &service.UserCommandServiceImpl{
		UserCommandRepository:    userCommandRepositoryImpl,
		UserQueryRepository:      userQueryRepositoryImpl,
		WordpressQueryRepository: wordpressQueryRepositoryImpl,
		AuthService:              authService,
		NoticeDomainService:      noticeDomainServiceImpl,
		TransactionService:       transactionServiceImpl,
	}
	userCommandController := api.UserCommandController{
		UserCommandService: userCommandServiceImpl,
	}
	healthCheckRepositoryImpl := &repository.HealthCheckRepositoryImpl{
		DB: db,
	}
	healthCheckController := api.HealthCheckController{
		HealthCheckRepository: healthCheckRepositoryImpl,
		Config:                configConfig,
	}
	areaCategoryCommandRepositoryImpl := &repository.AreaCategoryCommandRepositoryImpl{
		DAO: dao,
	}
	areaCategoryCommandServiceImpl := &service.AreaCategoryCommandServiceImpl{
		AreaCategoryCommandRepository: areaCategoryCommandRepositoryImpl,
		AreaCategoryQueryRepository:   areaCategoryQueryRepositoryImpl,
		WordpressQueryRepository:      wordpressQueryRepositoryImpl,
		WordpressService:              wordpressServiceImpl,
		TransactionService:            transactionServiceImpl,
	}
	themeCategoryCommandRepositoryImpl := &repository.ThemeCategoryCommandRepositoryImpl{
		DAO: dao,
	}
	themeCategoryCommandServiceImpl := &service.ThemeCategoryCommandServiceImpl{
		ThemeCategoryCommandRepository: themeCategoryCommandRepositoryImpl,
		ThemeCategoryQueryRepository:   themeCategoryQueryRepositoryImpl,
		WordpressQueryRepository:       wordpressQueryRepositoryImpl,
		WordpressService:               wordpressServiceImpl,
		TransactionService:             transactionServiceImpl,
	}
	categoryCommandServiceImpl := &service.CategoryCommandServiceImpl{
		AreaCategoryCommandService:    areaCategoryCommandServiceImpl,
		ThemeCategoryCommandService:   themeCategoryCommandServiceImpl,
		AreaCategoryCommandRepository: areaCategoryCommandRepositoryImpl,
		WordpressQueryRepository:      wordpressQueryRepositoryImpl,
	}
	comicCommandRepositoryImpl := &repository.ComicCommandRepositoryImpl{
		DAO: dao,
	}
	comicCommandServiceImpl := &service.ComicCommandServiceImpl{
		ComicCommandRepository:   comicCommandRepositoryImpl,
		WordpressQueryRepository: wordpressQueryRepositoryImpl,
		WordpressService:         wordpressServiceImpl,
		TransactionService:       transactionServiceImpl,
	}
	featureCommandRepositoryImpl := &repository.FeatureCommandRepositoryImpl{
		DAO: dao,
	}
	featureCommandServiceImpl := &service.FeatureCommandServiceImpl{
		FeatureCommandRepository: featureCommandRepositoryImpl,
		WordpressQueryRepository: wordpressQueryRepositoryImpl,
		WordpressService:         wordpressServiceImpl,
		TransactionService:       transactionServiceImpl,
	}
	spotCategoryCommandRepositoryImpl := &repository.SpotCategoryCommandRepositoryImpl{
		DAO: dao,
	}
	spotCategoryCommandServiceImpl := &service.SpotCategoryCommandServiceImpl{
		SpotCategoryCommandRepository: spotCategoryCommandRepositoryImpl,
		WordpressQueryRepository:      wordpressQueryRepositoryImpl,
		WordpressService:              wordpressServiceImpl,
		TransactionService:            transactionServiceImpl,
	}
	touristSpotCommandServiceImpl := &service.TouristSpotCommandServiceImpl{
		TouristSpotCommandRepository: touristSpotCommandRepositoryImpl,
		WordpressQueryRepository:     wordpressQueryRepositoryImpl,
		WordpressService:             wordpressServiceImpl,
		TransactionService:           transactionServiceImpl,
	}
	vlogCommandRepositoryImpl := &repository.VlogCommandRepositoryImpl{
		DAO: dao,
	}
	vlogCommandServiceImpl := &service.VlogCommandServiceImpl{
		VlogCommandRepository:    vlogCommandRepositoryImpl,
		WordpressQueryRepository: wordpressQueryRepositoryImpl,
		WordpressService:         wordpressServiceImpl,
		TransactionService:       transactionServiceImpl,
	}
	wordpressCallbackServiceImpl := &service.WordpressCallbackServiceImpl{
		UserCommandService:         userCommandServiceImpl,
		CategoryCommandService:     categoryCommandServiceImpl,
		ComicCommandService:        comicCommandServiceImpl,
		FeatureCommandService:      featureCommandServiceImpl,
		SpotCategoryCommandService: spotCategoryCommandServiceImpl,
		PostCommandService:         postCommandServiceImpl,
		TouristSpotCommandService:  touristSpotCommandServiceImpl,
		VlogCommandService:         vlogCommandServiceImpl,
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
	themeCategoryQueryServiceImpl := &service.ThemeCategoryQueryServiceImpl{
		ThemeCategoryQueryRepository: themeCategoryQueryRepositoryImpl,
	}
	themeQueryController := api.ThemeQueryController{
		ThemeCategoryQueryService: themeCategoryQueryServiceImpl,
	}
	areaQueryServiceImpl := &service.AreaQueryServiceImpl{
		Repository: areaCategoryQueryRepositoryImpl,
	}
	areaQueryController := api.AreaQueryController{
		AreaQueryService: areaQueryServiceImpl,
	}
	innQueryServiceImpl := &service.InnQueryServiceImpl{
		InnQueryRepository:          innQueryRepositoryImpl,
		AreaCategoryQueryRepository: areaCategoryQueryRepositoryImpl,
		TouristSpotQueryRepository:  touristSpotQueryRepositoryImpl,
	}
	innQueryController := api.InnQueryController{
		InnQueryService: innQueryServiceImpl,
	}
	slack := configConfig.Slack
	env := configConfig.Env
	slackRepositoryImpl := &repository.SlackRepositoryImpl{
		SlackConfig: slack,
		EnvConfig:   env,
		Client:      clientClient,
	}
	reportCommandRepositoryImpl := &repository.ReportCommandRepositoryImpl{
		DAO:         dao,
		SlackConfig: slack,
		Client:      clientClient,
	}
	reportQueryRepositoryImpl := &repository.ReportQueryRepositoryImpl{
		DB: db,
	}
	reportCommandServiceImpl := &service.ReportCommandServiceImpl{
		ReviewQueryRepository:   reviewQueryRepositoryImpl,
		ReviewCommandRepository: reviewCommandRepositoryImpl,
		SlackRepository:         slackRepositoryImpl,
		ReportCommandRepository: reportCommandRepositoryImpl,
		ReportQueryRepository:   reportQueryRepositoryImpl,
		TransactionService:      transactionServiceImpl,
	}
	reportCommandController := api.ReportCommandController{
		ReportCommandService: reportCommandServiceImpl,
	}
	app := &App{
		Config:                          configConfig,
		DB:                              db,
		Echo:                            echoEcho,
		AuthorizeWrapper:                authorize,
		PostCommandController:           postCommandController,
		PostQueryController:             postQueryController,
		PostFavoriteCommandController:   postFavoriteCommandController,
		CategoryQueryController:         categoryQueryController,
		ComicQueryController:            comicQueryController,
		ReviewQueryController:           reviewQueryController,
		ReviewCommandController:         reviewCommandController,
		ReviewFavoriteCommandController: reviewFavoriteCommandController,
		HashtagQueryController:          hashtagQueryController,
		HashtagCommandController:        hashtagCommandController,
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
		ThemeQueryController:            themeQueryController,
		AreaQueryController:             areaQueryController,
		InnQueryController:              innQueryController,
		ReportCommandController:         reportCommandController,
	}
	return app, nil
}

var (
	_wireConfigValue = &client.Config{}
)

// wire.go:

var controllerSet = wire.NewSet(api.PostQueryControllerSet, api.PostCommandControllerSet, api.PostFavoriteCommandControllerSet, api.CategoryQueryControllerSet, api.ComicQueryControllerSet, api.ReviewQueryControllerSet, api.ReviewCommandControllerSet, api.ReviewFavoriteCommandControllerSet, api.TouristSpotQeuryControllerSet, api.SearchQueryControllerSet, api.FeatureQueryControllerSet, api.VlogQueryControllerSet, api.HashtagQueryControllerSet, api.HashtagCommandControllerSet, api.UserQueryControllerSet, api.UserCommandControllerSet, api.HealthCheckControllerSet, api.ThemeQueryControllerSet, api.WordpressCallbackControllerSet, api.S3CommandControllerSet, api.InterestQueryControllerSet, api.AreaQueryControllerSet, api.InnQueryControllerSet, api.ReportCommandControllerSet)

var scenarioSet = wire.NewSet(scenario.ReviewCommandScenarioSet)

var domainServiceSet = wire.NewSet(service2.NoticeDomainServiceSet, service2.TaggedUserDomainServiceSet)

var serviceSet = wire.NewSet(service.PostQueryServiceSet, service.PostCommandServiceSet, service.PostFavoriteCommandServiceSet, service.CategoryQueryServiceSet, service.CategoryCommandServiceSet, service.AreaCategoryQueryServiceSet, service.AreaCategoryCommandServiceSet, service.ThemeCategoryQueryServiceSet, service.ThemeCategoryCommandServiceSet, service.ComicQueryServiceSet, service.ComicCommandServiceSet, service.ReviewQueryServiceSet, service.ReviewCommandServiceSet, service.ReviewFavoriteCommandServiceSet, service.WordpressServiceSet, service.TouristSpotQueryServiceSet, service.SearchQueryServiceSet, service.FeatureQueryServiceSet, service.FeatureCommandServiceSet, service.VlogQueryServiceSet, service.VlogCommandServiceSet, service.HashtagQueryServiceSet, service.HashtagCommandServiceSet, service.TouristSpotCommandServiceSet, service.SpotCategoryCommandServiceSet, service.WordpressCallbackServiceSet, service.UserQueryServiceSet, service.UserCommandServiceSet, service.S3CommandServiceSet, service.ProvideAuthService, service.InterestQueryServiceSet, service.InnQueryServiceSet, service.ReportCommandServiceSet)

var factorySet = wire.NewSet(factory.S3SignatureFactorySet)
