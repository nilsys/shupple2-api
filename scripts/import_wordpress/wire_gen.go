// Code generated by Wire. DO NOT EDIT.

//go:generate wire
//+build !wireinject

package main

import (
	"github.com/google/wire"
	"github.com/stayway-corp/stayway-media-api/pkg/adaptor/infrastructure/repository"
	"github.com/stayway-corp/stayway-media-api/pkg/application/service"
	"github.com/stayway-corp/stayway-media-api/pkg/config"
)

// Injectors from wire.go:

func InitializeScript(configFilePath config.ConfigFilePath) (*Script, error) {
	configConfig, err := config.GetConfig(configFilePath)
	if err != nil {
		return nil, err
	}
	db, err := repository.ProvideDB(configConfig)
	if err != nil {
		return nil, err
	}
	wordpress := configConfig.Wordpress
	stayway := configConfig.Stayway
	staywayMedia := stayway.Media
	wordpressQueryRepositoryImpl := repository.NewWordpressQueryRepositoryImpl(wordpress, staywayMedia)
	session, err := repository.ProvideAWSSession(configConfig)
	if err != nil {
		return nil, err
	}
	uploader := repository.ProvideS3Uploader(session)
	aws := configConfig.AWS
	userCommandRepositoryImpl := &repository.UserCommandRepositoryImpl{
		DB:            db,
		MediaUploader: uploader,
		AWSConfig:     aws,
	}
	categoryCommandRepositoryImpl := &repository.CategoryCommandRepositoryImpl{
		DB: db,
	}
	userQueryRepositoryImpl := &repository.UserQueryRepositoryImpl{
		DB: db,
	}
	categoryQueryRepositoryImpl := &repository.CategoryQueryRepositoryImpl{
		DB: db,
	}
	hashtagQueryRepositoryImpl := &repository.HashtagQueryRepositoryImpl{
		DB: db,
	}
	dao := repository.DAO{
		DB_: db,
	}
	hashtagCommandRepositoryImpl := &repository.HashtagCommandRepositoryImpl{
		DAO: dao,
	}
	hashtagCommandServiceImpl := &service.HashtagCommandServiceImpl{
		HashtagQueryRepository:   hashtagQueryRepositoryImpl,
		HashtagCommandRepository: hashtagCommandRepositoryImpl,
	}
	wordpressServiceImpl := &service.WordpressServiceImpl{
		WordpressQueryRepository: wordpressQueryRepositoryImpl,
		UserQueryRepository:      userQueryRepositoryImpl,
		CategoryQueryRepository:  categoryQueryRepositoryImpl,
		HashtagCommandService:    hashtagCommandServiceImpl,
	}
	categoryCommandServiceImpl := &service.CategoryCommandServiceImpl{
		CategoryCommandRepository: categoryCommandRepositoryImpl,
		WordpressQueryRepository:  wordpressQueryRepositoryImpl,
		WordpressService:          wordpressServiceImpl,
	}
	authService, err := service.ProvideAuthService(configConfig)
	if err != nil {
		return nil, err
	}
	userCommandServiceImpl := &service.UserCommandServiceImpl{
		UserCommandRepository:    userCommandRepositoryImpl,
		UserQueryRepository:      userQueryRepositoryImpl,
		WordpressQueryRepository: wordpressQueryRepositoryImpl,
		AuthService:              authService,
	}
	comicCommandRepositoryImpl := &repository.ComicCommandRepositoryImpl{
		DB: db,
	}
	comicCommandServiceImpl := &service.ComicCommandServiceImpl{
		ComicCommandRepository:   comicCommandRepositoryImpl,
		WordpressQueryRepository: wordpressQueryRepositoryImpl,
		WordpressService:         wordpressServiceImpl,
	}
	featureCommandRepositoryImpl := &repository.FeatureCommandRepositoryImpl{
		DB: db,
	}
	featureCommandServiceImpl := &service.FeatureCommandServiceImpl{
		FeatureCommandRepository: featureCommandRepositoryImpl,
		WordpressQueryRepository: wordpressQueryRepositoryImpl,
		WordpressService:         wordpressServiceImpl,
	}
	lcategoryCommandRepositoryImpl := &repository.LcategoryCommandRepositoryImpl{
		DB: db,
	}
	lcategoryCommandServiceImpl := &service.LcategoryCommandServiceImpl{
		LcategoryCommandRepository: lcategoryCommandRepositoryImpl,
		WordpressQueryRepository:   wordpressQueryRepositoryImpl,
		WordpressService:           wordpressServiceImpl,
	}
	postCommandRepositoryImpl := &repository.PostCommandRepositoryImpl{
		DAO: dao,
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
	touristSpotCommandRepositoryImpl := &repository.TouristSpotCommandRepositoryImpl{
		DAO: dao,
	}
	touristSpotCommandServiceImpl := &service.TouristSpotCommandServiceImpl{
		TouristSpotCommandRepository: touristSpotCommandRepositoryImpl,
		WordpressQueryRepository:     wordpressQueryRepositoryImpl,
		WordpressService:             wordpressServiceImpl,
	}
	vlogCommandRepositoryImpl := &repository.VlogCommandRepositoryImpl{
		DB: db,
	}
	vlogCommandServiceImpl := &service.VlogCommandServiceImpl{
		VlogCommandRepository:    vlogCommandRepositoryImpl,
		WordpressQueryRepository: wordpressQueryRepositoryImpl,
		WordpressService:         wordpressServiceImpl,
	}
	script := &Script{
		DB:                  db,
		Config:              configConfig,
		WordpressRepo:       wordpressQueryRepositoryImpl,
		UserRepo:            userCommandRepositoryImpl,
		CategoryCommandRepo: categoryCommandRepositoryImpl,
		CategoryService:     categoryCommandServiceImpl,
		UserService:         userCommandServiceImpl,
		ComicService:        comicCommandServiceImpl,
		FeatureService:      featureCommandServiceImpl,
		LcategoryService:    lcategoryCommandServiceImpl,
		PostService:         postCommandServiceImpl,
		TouristSpotService:  touristSpotCommandServiceImpl,
		VlogService:         vlogCommandServiceImpl,
	}
	return script, nil
}

// wire.go:

var serviceSet = wire.NewSet(service.ProvideAuthService, service.PostQueryServiceSet, service.PostCommandServiceSet, service.WordpressServiceSet, service.UserCommandServiceSet, service.CategoryCommandServiceSet, service.ComicCommandServiceSet, service.FeatureCommandServiceSet, service.LcategoryCommandServiceSet, service.TouristSpotCommandServiceSet, service.VlogCommandServiceSet, service.HashtagCommandServiceSet)
