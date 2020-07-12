// Code generated by Wire. DO NOT EDIT.

//go:generate wire
//+build !wireinject

package main

import (
	"github.com/google/wire"
	"github.com/stayway-corp/stayway-media-api/pkg/adaptor/infrastructure/client"
	"github.com/stayway-corp/stayway-media-api/pkg/adaptor/infrastructure/repository"
	"github.com/stayway-corp/stayway-media-api/pkg/adaptor/infrastructure/repository/payjp"
	"github.com/stayway-corp/stayway-media-api/pkg/application/scenario"
	"github.com/stayway-corp/stayway-media-api/pkg/application/service"
	"github.com/stayway-corp/stayway-media-api/pkg/config"
	service2 "github.com/stayway-corp/stayway-media-api/pkg/domain/service"
)

// Injectors from wire.go:

func InitializeScript(configFilePath config.FilePath) (*Script, error) {
	configConfig, err := config.GetConfig(configFilePath)
	if err != nil {
		return nil, err
	}
	db, err := repository.ProvideDB(configConfig)
	if err != nil {
		return nil, err
	}
	session, err := repository.ProvideAWSSession(configConfig)
	if err != nil {
		return nil, err
	}
	uploader := repository.ProvideS3Uploader(session)
	aws := configConfig.AWS
	wordpressQueryRepositoryImpl := repository.NewWordpressQueryRepositoryImpl(configConfig)
	userQueryRepositoryImpl := &repository.UserQueryRepositoryImpl{
		DB:         db,
		AWSConfig:  aws,
		AWSSession: session,
	}
	dao := repository.DAO{
		UnderlyingDB: db,
	}
	userCommandRepositoryImpl := &repository.UserCommandRepositoryImpl{
		DAO:           dao,
		MediaUploader: uploader,
		AWSConfig:     aws,
		AWSSession:    session,
	}
	payjpService := repository.ProvidePayjp(configConfig)
	customerCommandRepositoryImpl := &payjp.CustomerCommandRepositoryImpl{
		PayjpClient: payjpService,
	}
	authService, err := service.ProvideAuthService(configConfig)
	if err != nil {
		return nil, err
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
	transactionServiceImpl := &repository.TransactionServiceImpl{
		DB: db,
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
	areaCategoryCommandRepositoryImpl := &repository.AreaCategoryCommandRepositoryImpl{
		DAO: dao,
	}
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
	postCommandRepositoryImpl := &repository.PostCommandRepositoryImpl{
		DAO: dao,
	}
	postCommandServiceImpl := &service.PostCommandServiceImpl{
		PostCommandRepository:    postCommandRepositoryImpl,
		HashtagCommandRepository: hashtagCommandRepositoryImpl,
		WordpressQueryRepository: wordpressQueryRepositoryImpl,
		WordpressService:         wordpressServiceImpl,
		TransactionService:       transactionServiceImpl,
	}
	touristSpotCommandRepositoryImpl := &repository.TouristSpotCommandRepositoryImpl{
		DAO: dao,
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
	reviewCommandRepositoryImpl := &repository.ReviewCommandRepositoryImpl{
		DAO:        dao,
		AWSSession: session,
		AWSConfig:  aws,
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
	script := &Script{
		DB:                    db,
		Config:                configConfig,
		MediaUploader:         uploader,
		AWSConfig:             aws,
		WordpressRepo:         wordpressQueryRepositoryImpl,
		UserQueryRepository:   userQueryRepositoryImpl,
		UserRepo:              userCommandRepositoryImpl,
		UserService:           userCommandServiceImpl,
		CategoryService:       categoryCommandServiceImpl,
		ComicService:          comicCommandServiceImpl,
		FeatureService:        featureCommandServiceImpl,
		SpotCategoryService:   spotCategoryCommandServiceImpl,
		PostService:           postCommandServiceImpl,
		TouristSpotService:    touristSpotCommandServiceImpl,
		VlogService:           vlogCommandServiceImpl,
		ReviewCommandScenario: reviewCommandScenarioImpl,
	}
	return script, nil
}

var (
	_wireConfigValue = &client.Config{}
)

// wire.go:

var serviceSet = wire.NewSet(service.ProvideAuthService, service.PostQueryServiceSet, service.PostCommandServiceSet, service.WordpressServiceSet, service.UserCommandServiceSet, service.CategoryCommandServiceSet, service.AreaCategoryCommandServiceSet, service.ThemeCategoryCommandServiceSet, service.ComicCommandServiceSet, service.FeatureCommandServiceSet, service.SpotCategoryCommandServiceSet, service.TouristSpotCommandServiceSet, service.VlogCommandServiceSet, service.HashtagCommandServiceSet, service.ReviewCommandServiceSet, service.ReviewQueryServiceSet, scenario.ReviewCommandScenarioSet, service2.NoticeDomainServiceSet, service2.TaggedUserDomainServiceSet)
