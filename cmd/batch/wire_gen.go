// Code generated by Wire. DO NOT EDIT.

//go:generate wire
//+build !wireinject

package main

import (
	"github.com/google/wire"
	"github.com/stayway-corp/stayway-media-api/pkg/adaptor/infrastructure/repository"
	"github.com/stayway-corp/stayway-media-api/pkg/adaptor/infrastructure/repository/payjp"
	"github.com/stayway-corp/stayway-media-api/pkg/application/facade"
	service2 "github.com/stayway-corp/stayway-media-api/pkg/application/service"
	"github.com/stayway-corp/stayway-media-api/pkg/config"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/service"
)

// Injectors from wire.go:

func InitializeBatch(configFilePath config.FilePath) (*Batch, error) {
	configConfig, err := config.GetConfig(configFilePath)
	if err != nil {
		return nil, err
	}
	db, err := repository.ProvideDB(configConfig)
	if err != nil {
		return nil, err
	}
	dao := repository.DAO{
		UnderlyingDB: db,
	}
	session, err := repository.ProvideAWSSession(configConfig)
	if err != nil {
		return nil, err
	}
	uploader := repository.ProvideS3Uploader(session)
	aws := configConfig.AWS
	userCommandRepositoryImpl := &repository.UserCommandRepositoryImpl{
		DAO:           dao,
		MediaUploader: uploader,
		AWSConfig:     aws,
		AWSSession:    session,
	}
	userQueryRepositoryImpl := &repository.UserQueryRepositoryImpl{
		DB:         db,
		AWSConfig:  aws,
		AWSSession: session,
	}
	wordpressQueryRepositoryImpl := repository.NewWordpressQueryRepositoryImpl(configConfig)
	userValidatorDomainServiceImpl := &service.UserValidatorDomainServiceImpl{
		UserQueryRepository: userQueryRepositoryImpl,
	}
	payjpService := repository.ProvidePayjp(configConfig)
	customerCommandRepositoryImpl := &payjp.CustomerCommandRepositoryImpl{
		PayjpClient: payjpService,
	}
	customerQueryRepositoryImpl := &payjp.CustomerQueryRepositoryImpl{
		PayjpClient: payjpService,
	}
	authService, err := service2.ProvideAuthService(configConfig)
	if err != nil {
		return nil, err
	}
	noticeCommandRepositoryImpl := &repository.NoticeCommandRepositoryImpl{
		DAO: dao,
	}
	taggedUserDomainServiceImpl := service.TaggedUserDomainServiceImpl{
		UserQueryRepository: userQueryRepositoryImpl,
	}
	noticeDomainServiceImpl := &service.NoticeDomainServiceImpl{
		NoticeCommandRepository: noticeCommandRepositoryImpl,
		TaggedUserDomainService: taggedUserDomainServiceImpl,
	}
	transactionServiceImpl := &repository.TransactionServiceImpl{
		DB: db,
	}
	userCommandServiceImpl := &service2.UserCommandServiceImpl{
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
	hashtagCommandServiceImpl := &service2.HashtagCommandServiceImpl{
		HashtagQueryRepository:   hashtagQueryRepositoryImpl,
		HashtagCommandRepository: hashtagCommandRepositoryImpl,
	}
	wordpressServiceImpl := &service2.WordpressServiceImpl{
		WordpressQueryRepository:     wordpressQueryRepositoryImpl,
		UserQueryRepository:          userQueryRepositoryImpl,
		AreaCategoryQueryRepository:  areaCategoryQueryRepositoryImpl,
		ThemeCategoryQueryRepository: themeCategoryQueryRepositoryImpl,
		SpotCategoryQueryRepository:  spotCategoryQueryRepositoryImpl,
		HashtagCommandService:        hashtagCommandServiceImpl,
	}
	areaCategoryCommandServiceImpl := &service2.AreaCategoryCommandServiceImpl{
		AreaCategoryCommandRepository: areaCategoryCommandRepositoryImpl,
		AreaCategoryQueryRepository:   areaCategoryQueryRepositoryImpl,
		WordpressQueryRepository:      wordpressQueryRepositoryImpl,
		WordpressService:              wordpressServiceImpl,
		TransactionService:            transactionServiceImpl,
	}
	themeCategoryCommandRepositoryImpl := &repository.ThemeCategoryCommandRepositoryImpl{
		DAO: dao,
	}
	themeCategoryCommandServiceImpl := &service2.ThemeCategoryCommandServiceImpl{
		ThemeCategoryCommandRepository: themeCategoryCommandRepositoryImpl,
		ThemeCategoryQueryRepository:   themeCategoryQueryRepositoryImpl,
		WordpressQueryRepository:       wordpressQueryRepositoryImpl,
		WordpressService:               wordpressServiceImpl,
		TransactionService:             transactionServiceImpl,
	}
	categoryCommandServiceImpl := &service2.CategoryCommandServiceImpl{
		AreaCategoryCommandService:  areaCategoryCommandServiceImpl,
		ThemeCategoryCommandService: themeCategoryCommandServiceImpl,
		WordpressQueryRepository:    wordpressQueryRepositoryImpl,
	}
	comicCommandRepositoryImpl := &repository.ComicCommandRepositoryImpl{
		DAO: dao,
	}
	comicCommandServiceImpl := &service2.ComicCommandServiceImpl{
		ComicCommandRepository:   comicCommandRepositoryImpl,
		WordpressQueryRepository: wordpressQueryRepositoryImpl,
		WordpressService:         wordpressServiceImpl,
		TransactionService:       transactionServiceImpl,
	}
	featureCommandRepositoryImpl := &repository.FeatureCommandRepositoryImpl{
		DAO: dao,
	}
	featureCommandServiceImpl := &service2.FeatureCommandServiceImpl{
		FeatureCommandRepository: featureCommandRepositoryImpl,
		WordpressQueryRepository: wordpressQueryRepositoryImpl,
		WordpressService:         wordpressServiceImpl,
		TransactionService:       transactionServiceImpl,
	}
	spotCategoryCommandRepositoryImpl := &repository.SpotCategoryCommandRepositoryImpl{
		DAO: dao,
	}
	spotCategoryCommandServiceImpl := &service2.SpotCategoryCommandServiceImpl{
		SpotCategoryCommandRepository: spotCategoryCommandRepositoryImpl,
		WordpressQueryRepository:      wordpressQueryRepositoryImpl,
		WordpressService:              wordpressServiceImpl,
		TransactionService:            transactionServiceImpl,
	}
	postCommandRepositoryImpl := &repository.PostCommandRepositoryImpl{
		DAO: dao,
	}
	cfProjectCommandRepositoryImpl := &repository.CfProjectCommandRepositoryImpl{
		DAO: dao,
	}
	postCommandServiceImpl := &service2.PostCommandServiceImpl{
		PostCommandRepository:      postCommandRepositoryImpl,
		HashtagCommandRepository:   hashtagCommandRepositoryImpl,
		WordpressQueryRepository:   wordpressQueryRepositoryImpl,
		CfProjectCommandRepository: cfProjectCommandRepositoryImpl,
		WordpressService:           wordpressServiceImpl,
		TransactionService:         transactionServiceImpl,
	}
	touristSpotCommandRepositoryImpl := &repository.TouristSpotCommandRepositoryImpl{
		DAO: dao,
	}
	touristSpotCommandServiceImpl := &service2.TouristSpotCommandServiceImpl{
		TouristSpotCommandRepository: touristSpotCommandRepositoryImpl,
		WordpressQueryRepository:     wordpressQueryRepositoryImpl,
		WordpressService:             wordpressServiceImpl,
		TransactionService:           transactionServiceImpl,
	}
	vlogCommandRepositoryImpl := &repository.VlogCommandRepositoryImpl{
		DAO: dao,
	}
	vlogCommandServiceImpl := &service2.VlogCommandServiceImpl{
		VlogCommandRepository:    vlogCommandRepositoryImpl,
		WordpressQueryRepository: wordpressQueryRepositoryImpl,
		WordpressService:         wordpressServiceImpl,
		TransactionService:       transactionServiceImpl,
	}
	mailCommandRepository := repository.ProvideMailer(configConfig, session)
	cfProjectCommandServiceImpl := &service2.CfProjectCommandServiceImpl{
		CfProjectCommandRepository: cfProjectCommandRepositoryImpl,
		UserQueryRepository:        userQueryRepositoryImpl,
		MailCommandRepository:      mailCommandRepository,
		WordpressQueryRepository:   wordpressQueryRepositoryImpl,
		WordpressService:           wordpressServiceImpl,
		TransactionService:         transactionServiceImpl,
	}
	cfReturnGiftCommandRepositoryImpl := &repository.CfReturnGiftCommandRepositoryImpl{
		DAO: dao,
	}
	cfReturnGiftCommandServiceImpl := &service2.CfReturnGiftCommandServiceImpl{
		CfReturnGiftCommandRepository: cfReturnGiftCommandRepositoryImpl,
		WordpressQueryRepository:      wordpressQueryRepositoryImpl,
		WordpressService:              wordpressServiceImpl,
		TransactionService:            transactionServiceImpl,
	}
	wordpressCallbackServiceImpl := &service2.WordpressCallbackServiceImpl{
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
	postQueryRepositoryImpl := &repository.PostQueryRepositoryImpl{
		DB: db,
	}
	reviewQueryRepositoryImpl := &repository.ReviewQueryRepositoryImpl{
		DB: db,
	}
	reviewCommandRepositoryImpl := &repository.ReviewCommandRepositoryImpl{
		DAO:        dao,
		AWSSession: session,
		AWSConfig:  aws,
	}
	vlogQueryRepositoryImpl := &repository.VlogQueryRepositoryImpl{
		DB: db,
	}
	featureQueryRepositoryImpl := &repository.FeatureQueryRepositoryImpl{
		DB: db,
	}
	cfProjectQueryRepositoryImpl := &repository.CfProjectQueryRepositoryImpl{
		DAO: dao,
	}
	cfProjectFacadeImpl := &facade.CfProjectFacadeImpl{
		CfProjectCommandService:    cfProjectCommandServiceImpl,
		CfReturnGiftCommandService: cfReturnGiftCommandServiceImpl,
		CfProjectQueryRepository:   cfProjectQueryRepositoryImpl,
		PostQueryRepository:        postQueryRepositoryImpl,
		WordpressQueryRepository:   wordpressQueryRepositoryImpl,
	}
	batch := &Batch{
		Config:                     configConfig,
		WordpressCallbackService:   wordpressCallbackServiceImpl,
		PostQueryRepository:        postQueryRepositoryImpl,
		PostCommandRepository:      postCommandRepositoryImpl,
		ReviewQueryRepository:      reviewQueryRepositoryImpl,
		ReviewCommandRepository:    reviewCommandRepositoryImpl,
		VlogQueryRepository:        vlogQueryRepositoryImpl,
		VlogCommandRepository:      vlogCommandRepositoryImpl,
		FeatureQueryRepository:     featureQueryRepositoryImpl,
		FeatureCommandRepository:   featureCommandRepositoryImpl,
		CfProjectQueryRepository:   cfProjectQueryRepositoryImpl,
		CfProjectCommandRepository: cfProjectCommandRepositoryImpl,
		UserQueryRepository:        userQueryRepositoryImpl,
		MailCommandRepository:      mailCommandRepository,
		CfProjectCommandService:    cfProjectCommandServiceImpl,
		CfProjectFacade:            cfProjectFacadeImpl,
	}
	return batch, nil
}

// wire.go:

var serviceSet = wire.NewSet(service2.PostQueryServiceSet, service2.PostCommandServiceSet, service2.AreaCategoryQueryServiceSet, service2.AreaCategoryCommandServiceSet, service2.ThemeCategoryCommandServiceSet, service2.CategoryQueryServiceSet, service2.ComicQueryServiceSet, service2.ComicCommandServiceSet, service2.ReviewQueryServiceSet, service2.WordpressServiceSet, service2.SearchQueryServiceSet, service2.FeatureQueryServiceSet, service2.FeatureCommandServiceSet, service2.VlogQueryServiceSet, service2.VlogCommandServiceSet, service2.HashtagQueryServiceSet, service2.HashtagCommandServiceSet, service2.TouristSpotCommandServiceSet, service2.CategoryCommandServiceSet, service2.SpotCategoryCommandServiceSet, service2.WordpressCallbackServiceSet, service2.UserQueryServiceSet, service2.UserCommandServiceSet, service2.CfProjectCommandServiceSet, service2.CfReturnGiftCommandServiceSet, service2.ProvideAuthService)

var facadeSet = wire.NewSet(facade.CfProjectFacadeSet)

var domainServiceSet = wire.NewSet(service.NoticeDomainServiceSet, service.TaggedUserDomainServiceSet, service.UserValidatorDomainServiceSet)
