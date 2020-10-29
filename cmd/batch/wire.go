//+build wireinject

package main

import (
	"github.com/google/wire"
	"github.com/uma-co82/shupple2-api/pkg/adaptor/infrastructure/repository"
	"github.com/uma-co82/shupple2-api/pkg/config"
)

func InitializeBatch(configFilePath config.FilePath) (*Batch, error) {
	wire.Build(
		wire.Struct(new(Batch), "*"),
		config.GetConfig,
		repository.RepositoriesSet,
	)

	return new(Batch), nil
}
