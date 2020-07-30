// Code generated by Wire. DO NOT EDIT.

//go:generate wire
//+build !wireinject

package main

import (
	"github.com/google/wire"
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
	dao := repository.DAO{
		UnderlyingDB: db,
	}
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
	payjpService := repository.ProvidePayjp(configConfig)
	customerCommandRepositoryImpl := &payjp.CustomerCommandRepositoryImpl{
		PayjpClient: payjpService,
	}
	customerQueryRepositoryImpl := &payjp.CustomerQueryRepositoryImpl{
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
		CustomerQueryRepository:   customerQueryRepositoryImpl,
		AuthService:               authService,
		NoticeDomainService:       noticeDomainServiceImpl,
		TransactionService:        transactionServiceImpl,
	}
	script := &Script{
		DB:            db,
		Config:        configConfig,
		MediaUploader: uploader,
		AWSConfig:     aws,
		WordpressRepo: wordpressQueryRepositoryImpl,
		UserRepo:      userCommandRepositoryImpl,
		UserService:   userCommandServiceImpl,
	}
	return script, nil
}

// wire.go:

var serviceSet = wire.NewSet(service.ProvideAuthService, service.PostQueryServiceSet, service.PostCommandServiceSet, service.WordpressServiceSet, service.UserCommandServiceSet, service.CategoryCommandServiceSet, service.AreaCategoryCommandServiceSet, service.ThemeCategoryCommandServiceSet, service.ComicCommandServiceSet, service.FeatureCommandServiceSet, service.SpotCategoryCommandServiceSet, service.TouristSpotCommandServiceSet, service.VlogCommandServiceSet, service.HashtagCommandServiceSet, service.ReviewCommandServiceSet, service.ReviewQueryServiceSet, scenario.ReviewCommandScenarioSet, service2.NoticeDomainServiceSet, service2.TaggedUserDomainServiceSet)
