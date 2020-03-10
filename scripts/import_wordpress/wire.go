//+build wireinject

package main

import (
	"github.com/google/wire"
	"github.com/stayway-corp/stayway-media-api/pkg/adaptor/infrastructure/repository"
	"github.com/stayway-corp/stayway-media-api/pkg/application/service"
	"github.com/stayway-corp/stayway-media-api/pkg/config"
)

var serviceSet = wire.NewSet(
	service.PostQueryServiceSet,
	service.PostCommandServiceSet,
	service.WordpressServiceSet,
	service.CategoryCommandServiceSet,
	service.ComicCommandServiceSet,
	service.FeatureCommandServiceSet,
	service.LcategoryCommandServiceSet,
	service.TouristSpotCommandServiceSet,
	service.VlogCommandServiceSet,
	service.HashtagCommandServiceSet,
)

func InitializeScript(configFilePath config.ConfigFilePath) (*Script, error) {
	wire.Build(
		wire.Struct(new(Script), "*"),
		config.GetConfig,
		wire.FieldsOf(new(*config.Config), "Wordpress", "AWS"),
		serviceSet,
		repository.RepositoriesSet,
		repository.ProvideAWSSession,
		repository.ProvideS3Uploader,
	)

	return new(Script), nil
}
