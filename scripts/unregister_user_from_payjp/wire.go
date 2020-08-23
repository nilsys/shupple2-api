//+build wireinject

package main

import (
	"github.com/google/wire"
	"github.com/stayway-corp/stayway-media-api/pkg/adaptor/infrastructure/repository"
	"github.com/stayway-corp/stayway-media-api/pkg/adaptor/infrastructure/repository/payjp"
	"github.com/stayway-corp/stayway-media-api/pkg/config"
)

func InitializeScript(path config.FilePath) (*Script, error) {
	wire.Build(
		wire.Struct(new(Script), "*"),
		config.GetConfig,
		repository.ProvideDB,
		payjp.CustomerCommandRepositorySet,
		repository.ProvidePayjp,
		wire.Struct(new(repository.DAO), "*"),
		repository.TransactionServiceSet,
	)

	return new(Script), nil
}
