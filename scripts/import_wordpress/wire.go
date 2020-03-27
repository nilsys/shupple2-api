//+build wireinject

package main

import (
	"github.com/google/wire"
	"github.com/stayway-corp/stayway-media-api/pkg/adaptor/infrastructure/client"
	"github.com/stayway-corp/stayway-media-api/pkg/adaptor/infrastructure/repository"
	"github.com/stayway-corp/stayway-media-api/pkg/application/scenario"
	"github.com/stayway-corp/stayway-media-api/pkg/application/service"
	"github.com/stayway-corp/stayway-media-api/pkg/config"
	domain_service "github.com/stayway-corp/stayway-media-api/pkg/domain/service"
)

var serviceSet = wire.NewSet(
	service.ProvideAuthService,
	service.PostQueryServiceSet,
	service.PostCommandServiceSet,
	service.WordpressServiceSet,
	service.UserCommandServiceSet,
	service.CategoryCommandServiceSet,
	service.ComicCommandServiceSet,
	service.FeatureCommandServiceSet,
	service.LcategoryCommandServiceSet,
	service.TouristSpotCommandServiceSet,
	service.VlogCommandServiceSet,
	service.HashtagCommandServiceSet,
	service.ReviewCommandServiceSet,
	service.ReviewQueryServiceSet,
	scenario.ReviewCommandScenarioSet,
	domain_service.NoticeDomainServiceSet,
	domain_service.TaggedUserDomainServiceSet,
)

func InitializeScript(configFilePath config.FilePath) (*Script, error) {
	wire.Build(
		wire.Struct(new(Script), "*"),
		config.GetConfig,
		wire.FieldsOf(new(*config.Config), "Wordpress", "AWS", "Stayway"),
		client.NewClient,
		wire.Value(&client.Config{}),
		wire.FieldsOf(new(config.Stayway), "Media", "Metasearch"),
		serviceSet,
		repository.RepositoriesSet,
		repository.ProvideS3Uploader,
	)

	return new(Script), nil
}
