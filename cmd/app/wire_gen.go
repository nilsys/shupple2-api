// Code generated by Wire. DO NOT EDIT.

//go:generate wire
//+build !wireinject

package main

import (
	"github.com/google/wire"
	"github.com/labstack/echo/v4"
	"github.com/stayway-corp/stayway-media-api/pkg/adaptor/api"
	"github.com/stayway-corp/stayway-media-api/pkg/adaptor/api/converter"
	"github.com/stayway-corp/stayway-media-api/pkg/adaptor/api/middleware"
	"github.com/stayway-corp/stayway-media-api/pkg/adaptor/infrastructure/client"
	"github.com/stayway-corp/stayway-media-api/pkg/adaptor/infrastructure/repository"
	"github.com/stayway-corp/stayway-media-api/pkg/adaptor/infrastructure/repository/payjp"
	"github.com/stayway-corp/stayway-media-api/pkg/application/scenario"
	"github.com/stayway-corp/stayway-media-api/pkg/application/service"
	"github.com/stayway-corp/stayway-media-api/pkg/config"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/factory"
	service2 "github.com/stayway-corp/stayway-media-api/pkg/domain/service"

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
	shippingQueryRepositoryImpl := &repository.ShippingQueryRepositoryImpl{
		DAO: dao,
	}
	shippingQueryServiceImpl := &service.ShippingQueryServiceImpl{
		ShippingQueryRepository: shippingQueryRepositoryImpl,
	}
	converters := converter.Converters{
		Config: configConfig,
	}
	shippingQueryController := api.ShippingQueryController{
		ShippingQueryService: shippingQueryServiceImpl,
		Converters:           converters,
	}
	shippingCommandRepositoryImpl := &repository.ShippingCommandRepositoryImpl{
		DAO: dao,
	}
	shippingCommandServiceImpl := &service.ShippingCommandServiceImpl{
		ShippingCommandRepository: shippingCommandRepositoryImpl,
		ShippingQueryRepository:   shippingQueryRepositoryImpl,
	}
	shippingCommandController := api.ShippingCommandController{
		ShippingCommandService: shippingCommandServiceImpl,
		Converters:             converters,
	}
	payjpService := repository.ProvidePayjp(configConfig)
	cardQueryRepositoryImpl := &repository.CardQueryRepositoryImpl{
		DAO:         dao,
		PayjpClient: payjpService,
	}
	payjpCardQueryRepositoryImpl := &payjp.CardQueryRepositoryImpl{
		PayjpClient: payjpService,
	}
	cardQueryServiceImpl := &service.CardQueryServiceImpl{
		CardQueryRepository:      cardQueryRepositoryImpl,
		PayjpCardQueryRepository: payjpCardQueryRepositoryImpl,
	}
	cardQueryController := api.CardQueryController{
		CardQueryService: cardQueryServiceImpl,
		Converters:       converters,
	}
	paymentCommandRepositoryImpl := &repository.PaymentCommandRepositoryImpl{
		DAO: dao,
	}
	cfProjectQueryRepositoryImpl := &repository.CfProjectQueryRepositoryImpl{
		DAO: dao,
	}
	chargeCommandRepositoryImpl := &payjp.ChargeCommandRepositoryImpl{
		PayjpClient: payjpService,
	}
	returnGiftQueryRepositoryImpl := &repository.ReturnGiftQueryRepositoryImpl{
		DAO: dao,
	}
	transactionServiceImpl := &repository.TransactionServiceImpl{
		DB: db,
	}
	chargeCommandServiceImpl := &service.ChargeCommandServiceImpl{
		PaymentCommandRepository:  paymentCommandRepositoryImpl,
		CardQueryRepository:       cardQueryRepositoryImpl,
		CfProjectQueryRepository:  cfProjectQueryRepositoryImpl,
		ChargeCommandRepository:   chargeCommandRepositoryImpl,
		ReturnGiftQueryRepository: returnGiftQueryRepositoryImpl,
		ShippingQueryRepository:   shippingQueryRepositoryImpl,
		TransactionService:        transactionServiceImpl,
	}
	chargeCommandController := api.ChargeCommandController{
		ChargeCommandService: chargeCommandServiceImpl,
		Converters:           converters,
	}
	cardCommandRepositoryImpl := &repository.CardCommandRepositoryImpl{
		DAO:         dao,
		PayjpClient: payjpService,
	}
	payjpCardCommandRepositoryImpl := &payjp.CardCommandRepositoryImpl{
		PayjpClient: payjpService,
	}
	cardCommandServiceImpl := &service.CardCommandServiceImpl{
		CardCommandRepository:      cardCommandRepositoryImpl,
		PayjpCardCommandRepository: payjpCardCommandRepositoryImpl,
		TransactionService:         transactionServiceImpl,
	}
	cardCommandController := api.CardCommandController{
		CardCommandService: cardCommandServiceImpl,
		Converters:         converters,
	}
	areaCategoryQueryRepositoryImpl := &repository.AreaCategoryQueryRepositoryImpl{
		DB: db,
	}
	themeCategoryQueryRepositoryImpl := &repository.ThemeCategoryQueryRepositoryImpl{
		DB: db,
	}
	categoryIDMapFactory := factory.CategoryIDMapFactory{
		AreaCategoryQueryRepository:  areaCategoryQueryRepositoryImpl,
		ThemeCategoryQueryRepository: themeCategoryQueryRepositoryImpl,
	}
	postQueryRepositoryImpl := &repository.PostQueryRepositoryImpl{
		DB: db,
	}
	postQueryServiceImpl := &service.PostQueryServiceImpl{
		PostQueryRepository: postQueryRepositoryImpl,
	}
	postQueryScenarioImpl := &scenario.PostQueryScenarioImpl{
		CategoryIDMapFactory: categoryIDMapFactory,
		PostQueryService:     postQueryServiceImpl,
	}
	postQueryController := api.PostQueryController{
		Converters:        converters,
		PostQueryScenario: postQueryScenarioImpl,
	}
	postFavoriteCommandRepositoryImpl := &repository.PostFavoriteCommandRepositoryImpl{
		DAO: dao,
	}
	postFavoriteQueryRepositoryImpl := &repository.PostFavoriteQueryRepositoryImpl{
		DB: db,
	}
	postCommandRepositoryImpl := &repository.PostCommandRepositoryImpl{
		DAO: dao,
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
		Converters:                 converters,
		PostFavoriteCommandService: postFavoriteCommandServiceImpl,
	}
	spotCategoryQueryRepositoryImpl := &repository.SpotCategoryQueryRepositoryImpl{
		DB: db,
	}
	categoryQueryServiceImpl := &service.CategoryQueryServiceImpl{
		AreaCategoryRepository:  areaCategoryQueryRepositoryImpl,
		ThemeCategoryRepository: themeCategoryQueryRepositoryImpl,
		SpotCategoryRepository:  spotCategoryQueryRepositoryImpl,
	}
	categoryQueryController := api.CategoryQueryController{
		Converters:           converters,
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
		Converters:        converters,
		ComicQueryService: comicQueryServiceImpl,
	}
	reviewQueryRepositoryImpl := &repository.ReviewQueryRepositoryImpl{
		DB: db,
	}
	stayway := configConfig.Stayway
	staywayMetasearch := stayway.Metasearch
	clientConfig := _wireConfigValue
	clientClient := client.NewClient(clientConfig)
	innQueryRepositoryImpl := &repository.InnQueryRepositoryImpl{
		MetasearchConfig: staywayMetasearch,
		Client:           clientClient,
	}
	reviewQueryServiceImpl := &service.ReviewQueryServiceImpl{
		ReviewQueryRepository:       reviewQueryRepositoryImpl,
		InnQueryRepository:          innQueryRepositoryImpl,
		AreaCategoryQueryRepository: areaCategoryQueryRepositoryImpl,
	}
	reviewQueryController := api.ReviewQueryController{
		Converters:         converters,
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
	hashtagCommandRepositoryImpl := &repository.HashtagCommandRepositoryImpl{
		DAO: dao,
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
	hashtagQueryRepositoryImpl := &repository.HashtagQueryRepositoryImpl{
		DB: db,
	}
	hashtagCommandServiceImpl := &service.HashtagCommandServiceImpl{
		HashtagQueryRepository:   hashtagQueryRepositoryImpl,
		HashtagCommandRepository: hashtagCommandRepositoryImpl,
	}
	reviewCommandScenarioImpl := &scenario.ReviewCommandScenarioImpl{
		ReviewQueryService:    reviewQueryServiceImpl,
		ReviewCommandService:  reviewCommandServiceImpl,
		HashtagCommandService: hashtagCommandServiceImpl,
	}
	hashtagQueryServiceImpl := &service.HashtagQueryServiceImpl{
		HashtagQueryRepository: hashtagQueryRepositoryImpl,
	}
	reviewCommandController := api.ReviewCommandController{
		Converters:            converters,
		ReviewCommandScenario: reviewCommandScenarioImpl,
		ReviewCommandService:  reviewCommandServiceImpl,
		HashtagQueryService:   hashtagQueryServiceImpl,
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
		Converters:                   converters,
		ReviewFavoriteCommandService: reviewFavoriteCommandServiceImpl,
	}
	wordpressQueryRepositoryImpl := repository.NewWordpressQueryRepositoryImpl(configConfig)
	wordpress := configConfig.Wordpress
	staywayMedia := stayway.Media
	rssServiceImpl := &service.RSSServiceImpl{
		WordpressQueryRepository: wordpressQueryRepositoryImpl,
		WordpressConfig:          wordpress,
		MediaConfig:              staywayMedia,
	}
	rssController := api.RSSController{
		RSSService: rssServiceImpl,
	}
	hashtagQueryController := api.HashtagQueryController{
		Converters:          converters,
		HashtagQueryService: hashtagQueryServiceImpl,
	}
	hashtagCommandController := api.HashtagCommandController{
		Converters:            converters,
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
		Converters:         converters,
		SearchQueryService: searchQueryServiceImpl,
	}
	featureQueryRepositoryImpl := &repository.FeatureQueryRepositoryImpl{
		DB: db,
	}
	featureQueryServiceImpl := &service.FeatureQueryServiceImpl{
		FeatureQueryRepository: featureQueryRepositoryImpl,
		CategoryIDMapFactory:   categoryIDMapFactory,
	}
	featureQueryScenarioImpl := &scenario.FeatureQueryScenarioImpl{
		CategoryIDMapFactory: categoryIDMapFactory,
		FeatureQueryService:  featureQueryServiceImpl,
	}
	featureQueryController := api.FeatureQueryController{
		Converters:           converters,
		FeatureQueryScenario: featureQueryScenarioImpl,
		FeatureQueryService:  featureQueryServiceImpl,
	}
	vlogQueryRepositoryImpl := &repository.VlogQueryRepositoryImpl{
		DB: db,
	}
	vlogQueryServiceImpl := &service.VlogQueryServiceImpl{
		VlogQueryRepository:  vlogQueryRepositoryImpl,
		CategoryIDMapFactory: categoryIDMapFactory,
	}
	vlogQueryScenarioImpl := &scenario.VlogQueryScenarioImpl{
		VlogQueryService:     vlogQueryServiceImpl,
		CategoryIDMapFactory: categoryIDMapFactory,
	}
	vlogQueryController := api.VlogQueryController{
		Converters:        converters,
		VlogQueryScenario: vlogQueryScenarioImpl,
	}
	userQueryServiceImpl := &service.UserQueryServiceImpl{
		UserQueryRepository: userQueryRepositoryImpl,
	}
	userQueryController := api.UserQueryController{
		Converters:       converters,
		UserQueryService: userQueryServiceImpl,
	}
	uploader := repository.ProvideS3Uploader(session)
	userCommandRepositoryImpl := &repository.UserCommandRepositoryImpl{
		DAO:           dao,
		MediaUploader: uploader,
		AWSConfig:     aws,
		AWSSession:    session,
	}
	customerCommandRepositoryImpl := &payjp.CustomerCommandRepositoryImpl{
		PayjpClient: payjpService,
	}
	userCommandServiceImpl := &service.UserCommandServiceImpl{
		UserCommandRepository:     userCommandRepositoryImpl,
		UserQueryRepository:       userQueryRepositoryImpl,
		WordpressQueryRepository:  wordpressQueryRepositoryImpl,
		CustomerCommandRepository: customerCommandRepositoryImpl,
		AuthService:               authService,
		NoticeDomainService:       noticeDomainServiceImpl,
		TransactionService:        transactionServiceImpl,
	}
	userCommandController := api.UserCommandController{
		Converters:         converters,
		UserCommandService: userCommandServiceImpl,
	}
	healthCheckRepositoryImpl := &repository.HealthCheckRepositoryImpl{
		DB: db,
	}
	healthCheckController := api.HealthCheckController{
		Converters:            converters,
		HealthCheckRepository: healthCheckRepositoryImpl,
		Config:                configConfig,
	}
	areaCategoryCommandRepositoryImpl := &repository.AreaCategoryCommandRepositoryImpl{
		DAO: dao,
	}
	wordpressServiceImpl := &service.WordpressServiceImpl{
		WordpressQueryRepository:     wordpressQueryRepositoryImpl,
		UserQueryRepository:          userQueryRepositoryImpl,
		AreaCategoryQueryRepository:  areaCategoryQueryRepositoryImpl,
		ThemeCategoryQueryRepository: themeCategoryQueryRepositoryImpl,
		SpotCategoryQueryRepository:  spotCategoryQueryRepositoryImpl,
		HashtagCommandService:        hashtagCommandServiceImpl,
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
	postCommandServiceImpl := &service.PostCommandServiceImpl{
		PostCommandRepository:    postCommandRepositoryImpl,
		HashtagCommandRepository: hashtagCommandRepositoryImpl,
		WordpressQueryRepository: wordpressQueryRepositoryImpl,
		WordpressService:         wordpressServiceImpl,
		TransactionService:       transactionServiceImpl,
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
		Converters:               converters,
		WordpressCallbackService: wordpressCallbackServiceImpl,
	}
	sitemapServiceImpl := &service.SitemapServiceImpl{
		WordpressQueryRepository: wordpressQueryRepositoryImpl,
		WordpressConfig:          wordpress,
		MediaConfig:              staywayMedia,
	}
	sitemapController := api.SitemapController{
		Converters:     converters,
		SitemapService: sitemapServiceImpl,
	}
	s3SignatureFactory := factory.S3SignatureFactory{
		Session:   session,
		AWSConfig: aws,
	}
	s3CommandServiceImpl := &service.S3CommandServiceImpl{
		S3SignatureFactory: s3SignatureFactory,
	}
	s3CommandController := api.S3CommandController{
		Converters:       converters,
		S3CommandService: s3CommandServiceImpl,
	}
	touristSpotQueryServiceImpl := &service.TouristSpotQueryServiceImpl{
		TouristSpotQueryRepository: touristSpotQueryRepositoryImpl,
		CategoryIDMapFactory:       categoryIDMapFactory,
	}
	touristSpotQueryScenarioImpl := &scenario.TouristSpotQueryScenarioImpl{
		CategoryIDMapFactory:    categoryIDMapFactory,
		TouristSpotQueryService: touristSpotQueryServiceImpl,
	}
	touristSpotQueryController := api.TouristSpotQueryController{
		Converters:               converters,
		TouristSpotQueryScenario: touristSpotQueryScenarioImpl,
	}
	interestQueryRepositoryImpl := &repository.InterestQueryRepositoryImpl{
		DB: db,
	}
	interestQueryServiceImpl := &service.InterestQueryServiceImpl{
		InterestQueryRepository: interestQueryRepositoryImpl,
	}
	interestQueryController := api.InterestQueryController{
		Converters:           converters,
		InterestQueryService: interestQueryServiceImpl,
	}
	themeCategoryQueryServiceImpl := &service.ThemeCategoryQueryServiceImpl{
		ThemeCategoryQueryRepository: themeCategoryQueryRepositoryImpl,
	}
	themeQueryController := api.ThemeQueryController{
		Converters:                converters,
		ThemeCategoryQueryService: themeCategoryQueryServiceImpl,
	}
	areaQueryServiceImpl := &service.AreaQueryServiceImpl{
		Repository: areaCategoryQueryRepositoryImpl,
	}
	areaQueryController := api.AreaQueryController{
		Converters:       converters,
		AreaQueryService: areaQueryServiceImpl,
	}
	innQueryServiceImpl := &service.InnQueryServiceImpl{
		InnQueryRepository:          innQueryRepositoryImpl,
		AreaCategoryQueryRepository: areaCategoryQueryRepositoryImpl,
		TouristSpotQueryRepository:  touristSpotQueryRepositoryImpl,
	}
	innQueryController := api.InnQueryController{
		Converters:      converters,
		InnQueryService: innQueryServiceImpl,
	}
	noticeQueryRepositoryImpl := &repository.NoticeQueryRepositoryImpl{
		DB: db,
	}
	noticeQueryServiceImpl := &service.NoticeQueryServiceImpl{
		NoticeQueryRepository:   noticeQueryRepositoryImpl,
		NoticeCommandRepository: noticeCommandRepositoryImpl,
		TransactionService:      transactionServiceImpl,
	}
	noticeQueryController := api.NoticeQueryController{
		Converters:         converters,
		NoticeQueryService: noticeQueryServiceImpl,
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
		Converters:           converters,
		ReportCommandService: reportCommandServiceImpl,
	}
	comicFavoriteCommandRepositoryImpl := &repository.ComicFavoriteCommandRepositoryImpl{
		DAO: dao,
	}
	comicFavoriteQueryRepositoryImpl := &repository.ComicFavoriteQueryRepositoryImpl{
		DB: db,
	}
	comicFavoriteCommandServiceImpl := &service.ComicFavoriteCommandServiceImpl{
		ComicFavoriteCommandRepository: comicFavoriteCommandRepositoryImpl,
		ComicFavoriteQueryRepository:   comicFavoriteQueryRepositoryImpl,
		ComicCommandRepository:         comicCommandRepositoryImpl,
		ComicQueryRepository:           comicQueryRepositoryImpl,
		NoticeDomainService:            noticeDomainServiceImpl,
		TransactionService:             transactionServiceImpl,
	}
	comicFavoriteCommandController := api.ComicFavoriteCommandController{
		ComicFavoriteCommandService: comicFavoriteCommandServiceImpl,
	}
	vlogFavoriteCommandRepositoryImpl := &repository.VlogFavoriteCommandRepositoryImpl{
		DAO: dao,
	}
	vlogFavoriteQueryRepositoryImpl := &repository.VlogFavoriteQueryRepositoryImpl{
		DB: db,
	}
	vlogFavoriteCommandServiceImpl := &service.VlogFavoriteCommandServiceImpl{
		VlogFavoriteCommandRepository: vlogFavoriteCommandRepositoryImpl,
		VlogFavoriteQueryRepository:   vlogFavoriteQueryRepositoryImpl,
		VlogQueryRepository:           vlogQueryRepositoryImpl,
		VlogCommandRepository:         vlogCommandRepositoryImpl,
		NoticeDomainService:           noticeDomainServiceImpl,
		TransactionService:            transactionServiceImpl,
	}
	vlogFavoriteCommandController := api.VlogFavoriteCommandController{
		VlogFavoriteCommandService: vlogFavoriteCommandServiceImpl,
	}
	app := &App{
		Config:                          configConfig,
		DB:                              db,
		Echo:                            echoEcho,
		AuthorizeWrapper:                authorize,
		ShippingQueryController:         shippingQueryController,
		ShippingCommandController:       shippingCommandController,
		CardQueryController:             cardQueryController,
		ChargeCommandController:         chargeCommandController,
		CardCommandController:           cardCommandController,
		PostQueryController:             postQueryController,
		PostFavoriteCommandController:   postFavoriteCommandController,
		CategoryQueryController:         categoryQueryController,
		ComicQueryController:            comicQueryController,
		ReviewQueryController:           reviewQueryController,
		ReviewCommandController:         reviewCommandController,
		ReviewFavoriteCommandController: reviewFavoriteCommandController,
		RssController:                   rssController,
		HashtagQueryController:          hashtagQueryController,
		HashtagCommandController:        hashtagCommandController,
		SearchQueryController:           searchQueryController,
		FeatureQueryController:          featureQueryController,
		VlogQueryController:             vlogQueryController,
		UserQueryController:             userQueryController,
		UserCommandController:           userCommandController,
		HealthCheckController:           healthCheckController,
		WordpressCallbackController:     wordpressCallbackController,
		SitemapController:               sitemapController,
		S3CommandController:             s3CommandController,
		TouristSpotQueryController:      touristSpotQueryController,
		InteresetQueryController:        interestQueryController,
		ThemeQueryController:            themeQueryController,
		AreaQueryController:             areaQueryController,
		InnQueryController:              innQueryController,
		NoticeQueryController:           noticeQueryController,
		ReportCommandController:         reportCommandController,
		ComicFavoriteCommandController:  comicFavoriteCommandController,
		VlogFavoriteCommandController:   vlogFavoriteCommandController,
	}
	return app, nil
}

var (
	_wireConfigValue = &client.Config{}
)

// wire.go:

var controllerSet = wire.NewSet(converter.ConvertersSet, api.ShippingQueryControllerSet, api.ShippingCommandControllerSet, api.ChargeCommandControllerSet, api.CardQueryControllerSet, api.CardCommandControllerSet, api.PostQueryControllerSet, api.PostFavoriteCommandControllerSet, api.CategoryQueryControllerSet, api.ComicQueryControllerSet, api.ReviewQueryControllerSet, api.ReviewCommandControllerSet, api.ReviewFavoriteCommandControllerSet, api.RSSControllerSet, api.TouristSpotQeuryControllerSet, api.SearchQueryControllerSet, api.FeatureQueryControllerSet, api.VlogQueryControllerSet, api.HashtagQueryControllerSet, api.HashtagCommandControllerSet, api.UserQueryControllerSet, api.UserCommandControllerSet, api.HealthCheckControllerSet, api.ThemeQueryControllerSet, api.WordpressCallbackControllerSet, api.SitemapControllerSet, api.S3CommandControllerSet, api.InterestQueryControllerSet, api.AreaQueryControllerSet, api.InnQueryControllerSet, api.NoticeQueryControllerSet, api.ReportCommandControllerSet, api.ComicFavoriteCommandControllerSet, api.VlogFavoriteCommandControllerSet)

var scenarioSet = wire.NewSet(scenario.ReviewCommandScenarioSet, scenario.PostQueryScenarioSet, scenario.FeatureQueryScenarioSet, scenario.VlogQueryScenarioSet, scenario.TouristSpotQueryScenarioSet)

var domainServiceSet = wire.NewSet(service2.NoticeDomainServiceSet, service2.TaggedUserDomainServiceSet)

var serviceSet = wire.NewSet(service.ShippingQueryServiceSet, service.ShippingCommandServiceSet, service.ChargeCommandServiceSet, service.CardCommandServiceSet, service.CardQueryServiceSet, service.PostQueryServiceSet, service.PostCommandServiceSet, service.PostFavoriteCommandServiceSet, service.CategoryQueryServiceSet, service.CategoryCommandServiceSet, service.AreaCategoryQueryServiceSet, service.AreaCategoryCommandServiceSet, service.ThemeCategoryQueryServiceSet, service.ThemeCategoryCommandServiceSet, service.ComicQueryServiceSet, service.ComicCommandServiceSet, service.ReviewQueryServiceSet, service.ReviewCommandServiceSet, service.ReviewFavoriteCommandServiceSet, service.RssServiceSet, service.WordpressServiceSet, service.TouristSpotQueryServiceSet, service.SearchQueryServiceSet, service.FeatureQueryServiceSet, service.FeatureCommandServiceSet, service.VlogQueryServiceSet, service.VlogCommandServiceSet, service.HashtagQueryServiceSet, service.HashtagCommandServiceSet, service.TouristSpotCommandServiceSet, service.SpotCategoryCommandServiceSet, service.SitemapServiceSet, service.WordpressCallbackServiceSet, service.UserQueryServiceSet, service.UserCommandServiceSet, service.S3CommandServiceSet, service.ProvideAuthService, service.InterestQueryServiceSet, service.InnQueryServiceSet, service.NoticeQueryServiceSet, service.ReportCommandServiceSet, service.ComicFavoriteCommandServiceSet, service.VlogFavoriteCommandServiceSet)

var factorySet = wire.NewSet(factory.S3SignatureFactorySet, factory.CategoryIDMapFactorySet)
