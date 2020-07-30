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
	dao := repository.DAO{
		UnderlyingDB: db,
	}
	cfReturnGiftQueryRepositoryImpl := &repository.CfReturnGiftQueryRepositoryImpl{
		DAO: dao,
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
	payjpService := repository.ProvidePayjp(configConfig)
	customerQueryRepositoryImpl := &payjp.CustomerQueryRepositoryImpl{
		PayjpClient: payjpService,
	}
	customerCommandRepositoryImpl := &payjp.CustomerCommandRepositoryImpl{
		PayjpClient: payjpService,
	}
	cardCommandRepositoryImpl := &repository.CardCommandRepositoryImpl{
		DAO:         dao,
		PayjpClient: payjpService,
	}
	payjpCardCommandRepositoryImpl := &payjp.CardCommandRepositoryImpl{
		PayjpClient: payjpService,
	}
	transactionServiceImpl := &repository.TransactionServiceImpl{
		DB: db,
	}
	cardCommandServiceImpl := &service.CardCommandServiceImpl{
		CardCommandRepository:      cardCommandRepositoryImpl,
		PayjpCardCommandRepository: payjpCardCommandRepositoryImpl,
		TransactionService:         transactionServiceImpl,
	}
	paymentCommandRepositoryImpl := &repository.PaymentCommandRepositoryImpl{
		DAO: dao,
	}
	paymentQueryRepositoryImpl := &repository.PaymentQueryRepositoryImpl{
		DAO: dao,
	}
	cardQueryRepositoryImpl := &repository.CardQueryRepositoryImpl{
		DAO:         dao,
		PayjpClient: payjpService,
	}
	cfProjectQueryRepositoryImpl := &repository.CfProjectQueryRepositoryImpl{
		DAO: dao,
	}
	chargeCommandRepositoryImpl := &payjp.ChargeCommandRepositoryImpl{
		PayjpClient: payjpService,
	}
	cfReturnGiftCommandRepositoryImpl := &repository.CfReturnGiftCommandRepositoryImpl{
		DAO: dao,
	}
	shippingQueryRepositoryImpl := &repository.ShippingQueryRepositoryImpl{
		DAO: dao,
	}
	cfProjectCommandRepositoryImpl := &repository.CfProjectCommandRepositoryImpl{
		DAO: dao,
	}
	mailCommandRepository := repository.ProvideMailer(configConfig, session)
	chargeCommandServiceImpl := &service.ChargeCommandServiceImpl{
		PaymentCommandRepository:      paymentCommandRepositoryImpl,
		PaymentQueryRepository:        paymentQueryRepositoryImpl,
		CardQueryRepository:           cardQueryRepositoryImpl,
		CfProjectQueryRepository:      cfProjectQueryRepositoryImpl,
		ChargeCommandRepository:       chargeCommandRepositoryImpl,
		CfReturnGiftQueryRepository:   cfReturnGiftQueryRepositoryImpl,
		UserQueryRepository:           userQueryRepositoryImpl,
		CfReturnGiftCommandRepository: cfReturnGiftCommandRepositoryImpl,
		ShippingQueryRepository:       shippingQueryRepositoryImpl,
		CfProjectCommandRepository:    cfProjectCommandRepositoryImpl,
		MailCommandRepository:         mailCommandRepository,
		TransactionService:            transactionServiceImpl,
	}
	shippingCommandRepositoryImpl := &repository.ShippingCommandRepositoryImpl{
		DAO: dao,
	}
	shippingCommandServiceImpl := &service.ShippingCommandServiceImpl{
		ShippingCommandRepository: shippingCommandRepositoryImpl,
		ShippingQueryRepository:   shippingQueryRepositoryImpl,
	}
	script := &Script{
		Config:                      configConfig,
		CfReturnGiftQueryRepository: cfReturnGiftQueryRepositoryImpl,
		UserQueryRepository:         userQueryRepositoryImpl,
		CustomerQueryRepository:     customerQueryRepositoryImpl,
		CustomerCommandRepository:   customerCommandRepositoryImpl,
		CardCommandService:          cardCommandServiceImpl,
		ChargeCommandService:        chargeCommandServiceImpl,
		ShippingCommandService:      shippingCommandServiceImpl,
	}
	return script, nil
}

// wire.go:

var serviceSet = wire.NewSet(service.ProvideAuthService, service.PostQueryServiceSet, service.PostCommandServiceSet, service.WordpressServiceSet, service.UserCommandServiceSet, service.CategoryCommandServiceSet, service.AreaCategoryCommandServiceSet, service.ThemeCategoryCommandServiceSet, service.ComicCommandServiceSet, service.FeatureCommandServiceSet, service.SpotCategoryCommandServiceSet, service.TouristSpotCommandServiceSet, service.VlogCommandServiceSet, service.HashtagCommandServiceSet, service.ReviewCommandServiceSet, service.ReviewQueryServiceSet, service.CardCommandServiceSet, service.ChargeCommandServiceSet, service.ShippingCommandServiceSet, scenario.ReviewCommandScenarioSet, service2.NoticeDomainServiceSet, service2.TaggedUserDomainServiceSet)