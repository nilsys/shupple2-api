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
	"github.com/stayway-corp/stayway-media-api/pkg/application/service/helper"
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
	aws := configConfig.AWS
	session, err := repository.ProvideAWSSession(configConfig)
	if err != nil {
		return nil, err
	}
	userQueryRepositoryImpl := &repository.UserQueryRepositoryImpl{
		DB:         db,
		AWSConfig:  aws,
		AWSSession: session,
	}
	authorize := middleware.Authorize{
		AuthService: authService,
		UserRepo:    userQueryRepositoryImpl,
	}
	converters := converter.Converters{
		Config: configConfig,
	}
	dao := repository.DAO{
		UnderlyingDB: db,
	}
	cfProjectQueryRepositoryImpl := &repository.CfProjectQueryRepositoryImpl{
		DAO: dao,
	}
	cfProjectQueryServiceImpl := &service.CfProjectQueryServiceImpl{
		CfProjectQueryRepository: cfProjectQueryRepositoryImpl,
	}
	cfProjectQueryScenarioImpl := &scenario.CfProjectQueryScenarioImpl{
		CfProjectQueryService:    cfProjectQueryServiceImpl,
		UserQueryRepository:      userQueryRepositoryImpl,
		CfProjectQueryRepository: cfProjectQueryRepositoryImpl,
	}
	cfProjectQueryController := api.CfProjectQueryController{
		Converters:             converters,
		CfProjectQueryScenario: cfProjectQueryScenarioImpl,
		CfProjectQueryService:  cfProjectQueryServiceImpl,
	}
	cfProjectCommandRepositoryImpl := &repository.CfProjectCommandRepositoryImpl{
		DAO: dao,
	}
	mailCommandRepository := repository.ProvideMailer(configConfig, session)
	wordpressQueryRepositoryImpl := repository.NewWordpressQueryRepositoryImpl(configConfig)
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
	hashtagCommandRepositoryImpl := &repository.HashtagCommandRepositoryImpl{
		DAO: dao,
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
	cfProjectCommandServiceImpl := &service.CfProjectCommandServiceImpl{
		CfProjectCommandRepository: cfProjectCommandRepositoryImpl,
		UserQueryRepository:        userQueryRepositoryImpl,
		MailCommandRepository:      mailCommandRepository,
		WordpressQueryRepository:   wordpressQueryRepositoryImpl,
		WordpressService:           wordpressServiceImpl,
		TransactionService:         transactionServiceImpl,
	}
	cfProjectCommandController := api.CfProjectCommandController{
		Converters:              converters,
		CfProjectCommandService: cfProjectCommandServiceImpl,
	}
	cfReturnGiftQueryRepositoryImpl := &repository.CfReturnGiftQueryRepositoryImpl{
		DAO: dao,
	}
	cfReturnGiftQueryServiceImpl := &service.CfReturnGiftQueryServiceImpl{
		CfReturnGiftQueryRepository: cfReturnGiftQueryRepositoryImpl,
	}
	cfReturnGiftQueryController := api.CfReturnGiftQueryController{
		Converters:               converters,
		CfReturnGiftQueryService: cfReturnGiftQueryServiceImpl,
	}
	cfInnReserveRequestCommandRepositoryImpl := &repository.CfInnReserveRequestCommandRepositoryImpl{
		DAO: dao,
	}
	cfInnReserveRequestQueryRepositoryImpl := &repository.CfInnReserveRequestQueryRepositoryImpl{
		DB: db,
	}
	paymentQueryRepositoryImpl := &repository.PaymentQueryRepositoryImpl{
		DAO: dao,
	}
	paymentCommandRepositoryImpl := &repository.PaymentCommandRepositoryImpl{
		DAO: dao,
	}
	cfInnReserveRequestCommandServiceImpl := &service.CfInnReserveRequestCommandServiceImpl{
		CfInnReserveRequestCommandRepository: cfInnReserveRequestCommandRepositoryImpl,
		CfInnReserveRequestQueryRepository:   cfInnReserveRequestQueryRepositoryImpl,
		PaymentQueryRepository:               paymentQueryRepositoryImpl,
		PaymentCommandRepository:             paymentCommandRepositoryImpl,
		MailCommandRepository:                mailCommandRepository,
		TransactionService:                   transactionServiceImpl,
	}
	cfInnReserveRequestCommandController := api.CfInnReserveRequestCommandController{
		Converters:                        converters,
		CfInnReserveRequestCommandService: cfInnReserveRequestCommandServiceImpl,
	}
	shippingQueryRepositoryImpl := &repository.ShippingQueryRepositoryImpl{
		DAO: dao,
	}
	shippingQueryServiceImpl := &service.ShippingQueryServiceImpl{
		ShippingQueryRepository: shippingQueryRepositoryImpl,
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
	cardQueryServiceImpl := &service.CardQueryServiceImpl{
		CardQueryRepository: cardQueryRepositoryImpl,
	}
	cardQueryController := api.CardQueryController{
		CardQueryService: cardQueryServiceImpl,
		Converters:       converters,
	}
	cardCommandRepositoryImpl := &payjp.CardCommandRepositoryImpl{
		PayjpClient: payjpService,
	}
	repositoryCardCommandRepositoryImpl := &repository.CardCommandRepositoryImpl{
		DAO:         dao,
		PayjpClient: payjpService,
	}
	chargeCommandRepositoryImpl := &payjp.ChargeCommandRepositoryImpl{
		PayjpClient: payjpService,
	}
	uploader := repository.ProvideS3Uploader(session)
	userCommandRepositoryImpl := &repository.UserCommandRepositoryImpl{
		DAO:           dao,
		MediaUploader: uploader,
		AWSConfig:     aws,
		AWSSession:    session,
	}
	customerQueryRepositoryImpl := &payjp.CustomerQueryRepositoryImpl{
		PayjpClient: payjpService,
	}
	customerCommandRepositoryImpl := &payjp.CustomerCommandRepositoryImpl{
		PayjpClient: payjpService,
	}
	cfReturnGiftCommandRepositoryImpl := &repository.CfReturnGiftCommandRepositoryImpl{
		DAO: dao,
	}
	userSalesHistoryCommandRepositoryImpl := &repository.UserSalesHistoryCommandRepositoryImpl{
		DAO: dao,
	}
	inquiryCodeGeneratorImpl := &helper.InquiryCodeGeneratorImpl{}
	cfProject := configConfig.CfProject
	chargeCommandServiceImpl := &service.ChargeCommandServiceImpl{
		PaymentCommandRepository:          paymentCommandRepositoryImpl,
		PaymentQueryRepository:            paymentQueryRepositoryImpl,
		PayjpCardCommandRepository:        cardCommandRepositoryImpl,
		CardQueryRepository:               cardQueryRepositoryImpl,
		CardCommandRepository:             repositoryCardCommandRepositoryImpl,
		CfProjectQueryRepository:          cfProjectQueryRepositoryImpl,
		ChargeCommandRepository:           chargeCommandRepositoryImpl,
		CfReturnGiftQueryRepository:       cfReturnGiftQueryRepositoryImpl,
		UserQueryRepository:               userQueryRepositoryImpl,
		UserCommandRepository:             userCommandRepositoryImpl,
		CustomerQueryRepository:           customerQueryRepositoryImpl,
		CustomerCommandRepository:         customerCommandRepositoryImpl,
		CfReturnGiftCommandRepository:     cfReturnGiftCommandRepositoryImpl,
		ShippingQueryRepository:           shippingQueryRepositoryImpl,
		ShippingCommandRepository:         shippingCommandRepositoryImpl,
		CfProjectCommandRepository:        cfProjectCommandRepositoryImpl,
		MailCommandRepository:             mailCommandRepository,
		UserSalesHistoryCommandRepository: userSalesHistoryCommandRepositoryImpl,
		InquiryCodeGenerator:              inquiryCodeGeneratorImpl,
		TransactionService:                transactionServiceImpl,
		CfProjectConfig:                   cfProject,
	}
	userValidatorDomainServiceImpl := &service2.UserValidatorDomainServiceImpl{
		UserQueryRepository: userQueryRepositoryImpl,
	}
	noticeCommandRepositoryImpl := &repository.NoticeCommandRepositoryImpl{
		DAO: dao,
	}
	noticeQueryRepositoryImpl := &repository.NoticeQueryRepositoryImpl{
		DAO: dao,
	}
	firebaseAppWrap, err := repository.ProvideFirebaseApp(session, configConfig)
	if err != nil {
		return nil, err
	}
	fcmClientWrap, err := repository.ProvideFcmClient(firebaseAppWrap)
	if err != nil {
		return nil, err
	}
	cloudMessageCommandRepository := repository.ProvideFcmRepo(fcmClientWrap)
	taggedUserDomainServiceImpl := service2.TaggedUserDomainServiceImpl{
		UserQueryRepository: userQueryRepositoryImpl,
	}
	noticeDomainServiceImpl := &service2.NoticeDomainServiceImpl{
		NoticeCommandRepository:       noticeCommandRepositoryImpl,
		UserQueryRepository:           userQueryRepositoryImpl,
		NoticeQueryRepository:         noticeQueryRepositoryImpl,
		CloudMessageCommandRepository: cloudMessageCommandRepository,
		TaggedUserDomainService:       taggedUserDomainServiceImpl,
	}
	userCommandServiceImpl := &service.UserCommandServiceImpl{
		UserCommandRepository:      userCommandRepositoryImpl,
		UserQueryRepository:        userQueryRepositoryImpl,
		WordpressQueryRepository:   wordpressQueryRepositoryImpl,
		UserValidatorDomainService: userValidatorDomainServiceImpl,
		CustomerCommandRepository:  customerCommandRepositoryImpl,
		CustomerQueryRepository:    customerQueryRepositoryImpl,
		AuthService:                authService,
		NoticeDomainService:        noticeDomainServiceImpl,
		TransactionService:         transactionServiceImpl,
	}
	chargeCommandScenarioImpl := &scenario.ChargeCommandScenarioImpl{
		ChargeCommandService:       chargeCommandServiceImpl,
		CardCommandRepository:      repositoryCardCommandRepositoryImpl,
		UserCommandRepository:      userCommandRepositoryImpl,
		ShippingCommandRepository:  shippingCommandRepositoryImpl,
		UserCommandService:         userCommandServiceImpl,
		PayjpCardCommandRepository: cardCommandRepositoryImpl,
		CustomerQueryRepository:    customerQueryRepositoryImpl,
		CustomerCommandRepository:  customerCommandRepositoryImpl,
		UserValidatorDomainService: userValidatorDomainServiceImpl,
		TransactionService:         transactionServiceImpl,
	}
	chargeCommandController := api.ChargeCommandController{
		ChargeCommandScenario: chargeCommandScenarioImpl,
		ChargeCommandService:  chargeCommandServiceImpl,
		Converters:            converters,
	}
	cardCommandServiceImpl := &service.CardCommandServiceImpl{
		CardCommandRepository:      repositoryCardCommandRepositoryImpl,
		CardQueryRepository:        cardQueryRepositoryImpl,
		PayjpCardCommandRepository: cardCommandRepositoryImpl,
		TransactionService:         transactionServiceImpl,
	}
	cardCommandController := api.CardCommandController{
		CardCommandService: cardCommandServiceImpl,
		Converters:         converters,
	}
	paymentQueryServiceImpl := &service.PaymentQueryServiceImpl{
		PaymentQueryRepository: paymentQueryRepositoryImpl,
	}
	paymentQueryController := api.PaymentQueryController{
		Converters:          converters,
		PaymentQueryService: paymentQueryServiceImpl,
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
		UserQueryRepository:  userQueryRepositoryImpl,
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
	comicQueryQueryScenarioImpl := &scenario.ComicQueryQueryScenarioImpl{
		ComicQueryService:   comicQueryServiceImpl,
		UserQueryRepository: userQueryRepositoryImpl,
	}
	comicQueryController := api.ComicQueryController{
		Converters:         converters,
		ComicQueryScenario: comicQueryQueryScenarioImpl,
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
	metasearchAreaQueryRepositoryImpl := &repository.MetasearchAreaQueryRepositoryImpl{
		DB: db,
	}
	reviewQueryServiceImpl := &service.ReviewQueryServiceImpl{
		ReviewQueryRepository:         reviewQueryRepositoryImpl,
		InnQueryRepository:            innQueryRepositoryImpl,
		MetasearchAreaQueryRepository: metasearchAreaQueryRepositoryImpl,
	}
	reviewQueryScenarioImpl := &scenario.ReviewQueryScenarioImpl{
		ReviewQueryService:  reviewQueryServiceImpl,
		UserQueryRepository: userQueryRepositoryImpl,
	}
	reviewQueryController := api.ReviewQueryController{
		Converters:          converters,
		ReviewQueryService:  reviewQueryServiceImpl,
		ReviewQueryScenario: reviewQueryScenarioImpl,
	}
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
		UserQueryRepository:   userQueryRepositoryImpl,
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
		UserQueryRepository:  userQueryRepositoryImpl,
	}
	featureQueryController := api.FeatureQueryController{
		Converters:           converters,
		FeatureQueryScenario: featureQueryScenarioImpl,
	}
	vlogQueryRepositoryImpl := &repository.VlogQueryRepositoryImpl{
		DB: db,
	}
	vlogQueryServiceImpl := &service.VlogQueryServiceImpl{
		VlogQueryRepository:        vlogQueryRepositoryImpl,
		TouristSpotQueryRepository: touristSpotQueryRepositoryImpl,
		CategoryIDMapFactory:       categoryIDMapFactory,
	}
	vlogQueryScenarioImpl := &scenario.VlogQueryScenarioImpl{
		VlogQueryService:     vlogQueryServiceImpl,
		CategoryIDMapFactory: categoryIDMapFactory,
		UserQueryRepository:  userQueryRepositoryImpl,
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
		AreaCategoryCommandService:  areaCategoryCommandServiceImpl,
		ThemeCategoryCommandService: themeCategoryCommandServiceImpl,
		WordpressQueryRepository:    wordpressQueryRepositoryImpl,
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
		PostCommandRepository:      postCommandRepositoryImpl,
		HashtagCommandRepository:   hashtagCommandRepositoryImpl,
		WordpressQueryRepository:   wordpressQueryRepositoryImpl,
		CfProjectCommandRepository: cfProjectCommandRepositoryImpl,
		WordpressService:           wordpressServiceImpl,
		TransactionService:         transactionServiceImpl,
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
	cfReturnGiftCommandServiceImpl := &service.CfReturnGiftCommandServiceImpl{
		CfReturnGiftCommandRepository: cfReturnGiftCommandRepositoryImpl,
		WordpressQueryRepository:      wordpressQueryRepositoryImpl,
		WordpressService:              wordpressServiceImpl,
		TransactionService:            transactionServiceImpl,
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
		CfProjectCommandService:    cfProjectCommandServiceImpl,
		CfReturnGiftCommandService: cfReturnGiftCommandServiceImpl,
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
		ReviewQueryRepository:      reviewQueryRepositoryImpl,
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
		InnQueryRepository:            innQueryRepositoryImpl,
		MetasearchAreaQueryRepository: metasearchAreaQueryRepositoryImpl,
		TouristSpotQueryRepository:    touristSpotQueryRepositoryImpl,
	}
	innQueryController := api.InnQueryController{
		Converters:      converters,
		InnQueryService: innQueryServiceImpl,
	}
	noticeQueryServiceImpl := &service.NoticeQueryServiceImpl{
		NoticeQueryRepository:   noticeQueryRepositoryImpl,
		NoticeCommandRepository: noticeCommandRepositoryImpl,
		ReviewQueryRepository:   reviewQueryRepositoryImpl,
		TransactionService:      transactionServiceImpl,
	}
	noticeQueryController := api.NoticeQueryController{
		Converters:         converters,
		NoticeQueryService: noticeQueryServiceImpl,
	}
	noticeCommandServiceImpl := &service.NoticeCommandServiceImpl{
		NoticeQueryRepository:   noticeQueryRepositoryImpl,
		NoticeCommandRepository: noticeCommandRepositoryImpl,
		TransactionService:      transactionServiceImpl,
	}
	noticeCommandController := api.NoticeCommandController{
		NoticeCommand: noticeCommandServiceImpl,
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
		Config:                               configConfig,
		DB:                                   db,
		Echo:                                 echoEcho,
		AuthorizeWrapper:                     authorize,
		CfProjectQueryController:             cfProjectQueryController,
		CfProjectCommandController:           cfProjectCommandController,
		CfReturnGiftQueryController:          cfReturnGiftQueryController,
		CfInnReserveRequestCommandController: cfInnReserveRequestCommandController,
		ShippingQueryController:              shippingQueryController,
		ShippingCommandController:            shippingCommandController,
		CardQueryController:                  cardQueryController,
		ChargeCommandController:              chargeCommandController,
		CardCommandController:                cardCommandController,
		PaymentQueryController:               paymentQueryController,
		PostQueryController:                  postQueryController,
		PostFavoriteCommandController:        postFavoriteCommandController,
		CategoryQueryController:              categoryQueryController,
		ComicQueryController:                 comicQueryController,
		ReviewQueryController:                reviewQueryController,
		ReviewCommandController:              reviewCommandController,
		ReviewFavoriteCommandController:      reviewFavoriteCommandController,
		RssController:                        rssController,
		HashtagQueryController:               hashtagQueryController,
		HashtagCommandController:             hashtagCommandController,
		SearchQueryController:                searchQueryController,
		FeatureQueryController:               featureQueryController,
		VlogQueryController:                  vlogQueryController,
		UserQueryController:                  userQueryController,
		UserCommandController:                userCommandController,
		HealthCheckController:                healthCheckController,
		WordpressCallbackController:          wordpressCallbackController,
		SitemapController:                    sitemapController,
		S3CommandController:                  s3CommandController,
		TouristSpotQueryController:           touristSpotQueryController,
		InteresetQueryController:             interestQueryController,
		ThemeQueryController:                 themeQueryController,
		AreaQueryController:                  areaQueryController,
		InnQueryController:                   innQueryController,
		NoticeQueryController:                noticeQueryController,
		NoticeCommandController:              noticeCommandController,
		ReportCommandController:              reportCommandController,
		ComicFavoriteCommandController:       comicFavoriteCommandController,
		VlogFavoriteCommandController:        vlogFavoriteCommandController,
	}
	return app, nil
}

var (
	_wireConfigValue = &client.Config{}
)

// wire.go:

var controllerSet = wire.NewSet(converter.ConvertersSet, api.ShippingQueryControllerSet, api.ShippingCommandControllerSet, api.ChargeCommandControllerSet, api.CardQueryControllerSet, api.CardCommandControllerSet, api.PostQueryControllerSet, api.PostFavoriteCommandControllerSet, api.CfProjectQueryControllerSet, api.CfReturnGiftQueryControllerSet, api.CategoryQueryControllerSet, api.CfProjectCommandControllerSet, api.ComicQueryControllerSet, api.ReviewQueryControllerSet, api.ReviewCommandControllerSet, api.ReviewFavoriteCommandControllerSet, api.RSSControllerSet, api.TouristSpotQeuryControllerSet, api.SearchQueryControllerSet, api.FeatureQueryControllerSet, api.VlogQueryControllerSet, api.HashtagQueryControllerSet, api.HashtagCommandControllerSet, api.UserQueryControllerSet, api.UserCommandControllerSet, api.HealthCheckControllerSet, api.ThemeQueryControllerSet, api.WordpressCallbackControllerSet, api.SitemapControllerSet, api.S3CommandControllerSet, api.InterestQueryControllerSet, api.AreaQueryControllerSet, api.InnQueryControllerSet, api.NoticeQueryControllerSet, api.NoticeCommandControllerSet, api.PaymentQueryControllerSet, api.CfReserveRequestCommandControllerSet, api.ReportCommandControllerSet, api.ComicFavoriteCommandControllerSet, api.VlogFavoriteCommandControllerSet)

var scenarioSet = wire.NewSet(scenario.ReviewCommandScenarioSet, scenario.ReviewQueryScenarioSet, scenario.PostQueryScenarioSet, scenario.FeatureQueryScenarioSet, scenario.VlogQueryScenarioSet, scenario.TouristSpotQueryScenarioSet, scenario.CfProjectQueryScenarioSet, scenario.ComicQueryScenarioSet, scenario.ChargeCommandScenarioSet)

var domainServiceSet = wire.NewSet(service2.NoticeDomainServiceSet, service2.TaggedUserDomainServiceSet, service2.UserValidatorDomainServiceSet)

var serviceSet = wire.NewSet(service.ShippingQueryServiceSet, service.ShippingCommandServiceSet, service.CfProjectCommandServiceSet, service.CfProjectQueryServiceSet, service.ChargeCommandServiceSet, service.CardCommandServiceSet, service.CardQueryServiceSet, service.PaymentQueryServiceSet, service.PostQueryServiceSet, service.PostCommandServiceSet, service.PostFavoriteCommandServiceSet, service.CategoryQueryServiceSet, service.CategoryCommandServiceSet, service.CfReturnGiftQueryServiceSet, service.CfReturnGiftCommandServiceSet, service.CfInnReserveRequestCommandServiceSet, service.AreaCategoryQueryServiceSet, service.AreaCategoryCommandServiceSet, service.ThemeCategoryQueryServiceSet, service.ThemeCategoryCommandServiceSet, service.ComicQueryServiceSet, service.ComicCommandServiceSet, service.ReviewQueryServiceSet, service.ReviewCommandServiceSet, service.ReviewFavoriteCommandServiceSet, service.RssServiceSet, service.WordpressServiceSet, service.TouristSpotQueryServiceSet, service.SearchQueryServiceSet, service.FeatureQueryServiceSet, service.FeatureCommandServiceSet, service.VlogQueryServiceSet, service.VlogCommandServiceSet, service.HashtagQueryServiceSet, service.HashtagCommandServiceSet, service.TouristSpotCommandServiceSet, service.SpotCategoryCommandServiceSet, service.SitemapServiceSet, service.WordpressCallbackServiceSet, service.UserQueryServiceSet, service.UserCommandServiceSet, service.S3CommandServiceSet, service.ProvideAuthService, service.InterestQueryServiceSet, service.InnQueryServiceSet, service.NoticeQueryServiceSet, service.NoticeCommandServiceSet, service.ReportCommandServiceSet, service.ComicFavoriteCommandServiceSet, service.VlogFavoriteCommandServiceSet, helper.InquiryCodeGeneratorSet)

var factorySet = wire.NewSet(factory.S3SignatureFactorySet, factory.CategoryIDMapFactorySet)
