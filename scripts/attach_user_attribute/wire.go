//+build wireinject

package main

import (
	"github.com/google/wire"
	"github.com/stayway-corp/stayway-media-api/pkg/adaptor/infrastructure/repository"
	"github.com/stayway-corp/stayway-media-api/pkg/config"
)

func InitializeScript(path config.FilePath) (*Script, error) {
	wire.Build(
		wire.Struct(new(Script), "*"),
		config.GetConfig,
		wire.FieldsOf(new(*config.Config), "AWS"),
		repository.ProvideAWSSession,
		repository.ProvideDB,
		repository.UserQueryRepositorySet,
	)

	return new(Script), nil
}
