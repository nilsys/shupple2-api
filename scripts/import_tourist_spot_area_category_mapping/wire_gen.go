// Code generated by Wire. DO NOT EDIT.

//go:generate wire
//+build !wireinject

package main

import (
	"github.com/stayway-corp/stayway-media-api/pkg/adaptor/infrastructure/client"
	"github.com/stayway-corp/stayway-media-api/pkg/adaptor/infrastructure/repository"
	"github.com/stayway-corp/stayway-media-api/pkg/config"
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
	clientConfig := _wireConfigValue
	clientClient := client.NewClient(clientConfig)
	script := &Script{
		DB:     db,
		Client: clientClient,
		Config: configConfig,
	}
	return script, nil
}

var (
	_wireConfigValue = &client.Config{}
)