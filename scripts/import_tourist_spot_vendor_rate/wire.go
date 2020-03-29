//+build wireinject

package main

import (
	"github.com/google/wire"
	"github.com/stayway-corp/stayway-media-api/pkg/adaptor/infrastructure/repository"
	"github.com/stayway-corp/stayway-media-api/pkg/config"
)

func InitializeScript(configFilePath config.FilePath) (*Script, error) {
	wire.Build(
		wire.Struct(new(Script), "*"),
		config.GetConfig,
		repository.RepositoriesSet,
	)

	return new(Script), nil
}
