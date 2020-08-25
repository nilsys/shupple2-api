// Code generated by Wire. DO NOT EDIT.

//go:generate wire
//+build !wireinject

package main

import (
	"github.com/stayway-corp/stayway-media-api/pkg/adaptor/infrastructure/repository"
	"github.com/stayway-corp/stayway-media-api/pkg/adaptor/infrastructure/repository/payjp"
	"github.com/stayway-corp/stayway-media-api/pkg/config"
)

// Injectors from wire.go:

func InitializeScript(path config.FilePath) (*Script, error) {
	configConfig, err := config.GetConfig(path)
	if err != nil {
		return nil, err
	}
	service := repository.ProvidePayjp(configConfig)
	customerCommandRepositoryImpl := &payjp.CustomerCommandRepositoryImpl{
		PayjpClient: service,
	}
	db, err := repository.ProvideDB(configConfig)
	if err != nil {
		return nil, err
	}
	dao := repository.DAO{
		UnderlyingDB: db,
	}
	transactionServiceImpl := &repository.TransactionServiceImpl{
		DB: db,
	}
	script := &Script{
		PayjpClient:               service,
		CustomerCommandRepository: customerCommandRepositoryImpl,
		DAO:                       dao,
		TransactionService:        transactionServiceImpl,
	}
	return script, nil
}