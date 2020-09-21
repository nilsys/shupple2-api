//+build wireinject

package main

import (
	"github.com/google/wire"
	"github.com/stayway-corp/stayway-media-api/pkg/adaptor/infrastructure/repository"
	"github.com/stayway-corp/stayway-media-api/pkg/application/service"
	"github.com/stayway-corp/stayway-media-api/pkg/config"
)

var serviceSet = wire.NewSet(
	service.MediaCommandServiceSet,
)

type EnvString string

func InitializeLambda(env EnvString) (*Lambda, error) {
	wire.Build(
		wire.Struct(new(Lambda), "*"),
		getConfig,
		wire.FieldsOf(new(*config.Config), "AWS"),
		serviceSet,
		repository.RepositoriesSet,
	)

	return new(Lambda), nil
}

func getConfig(envString EnvString) (*config.Config, error) {
	if envString == "" {
		return config.GetConfig(config.DefaultConfigFilePath)
	}

	env, err := config.ParseEnv(string(envString))
	if err != nil {
		return nil, err
	}

	return config.GetConfigFromSSM(env)
}
