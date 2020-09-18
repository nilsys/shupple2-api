//+build wireinject

package main

import (
	"github.com/google/wire"
	"github.com/stayway-corp/stayway-media-api/pkg/adaptor/infrastructure/client"
	"github.com/stayway-corp/stayway-media-api/pkg/adaptor/infrastructure/repository"
	"github.com/stayway-corp/stayway-media-api/pkg/application/facade"
	"github.com/stayway-corp/stayway-media-api/pkg/application/service"
	"github.com/stayway-corp/stayway-media-api/pkg/config"
	domain_service "github.com/stayway-corp/stayway-media-api/pkg/domain/service"
)

var serviceSet = wire.NewSet(
	service.PostQueryServiceSet,
	service.PostCommandServiceSet,
	service.AreaCategoryQueryServiceSet,
	service.AreaCategoryCommandServiceSet,
	service.ThemeCategoryCommandServiceSet,
	service.CategoryQueryServiceSet,
	service.ComicQueryServiceSet,
	service.ComicCommandServiceSet,
	service.ReviewQueryServiceSet,
	service.WordpressServiceSet,
	service.SearchQueryServiceSet,
	service.FeatureQueryServiceSet,
	service.FeatureCommandServiceSet,
	service.VlogQueryServiceSet,
	service.VlogCommandServiceSet,
	service.HashtagQueryServiceSet,
	service.HashtagCommandServiceSet,
	service.TouristSpotCommandServiceSet,
	service.CategoryCommandServiceSet,
	service.SpotCategoryCommandServiceSet,
	service.WordpressCallbackServiceSet,
	service.UserQueryServiceSet,
	service.UserCommandServiceSet,
	service.CfProjectCommandServiceSet,
	service.CfReturnGiftCommandServiceSet,
	service.ProvideAuthService,
)

var facadeSet = wire.NewSet(
	facade.CfProjectFacadeSet,
	facade.ImportSnsShareCountFacadeSet,
)

var domainServiceSet = wire.NewSet(
	domain_service.NoticeDomainServiceSet,
	domain_service.TaggedUserDomainServiceSet,
	domain_service.UserValidatorDomainServiceSet,
)

func InitializeBatch(configFilePath config.FilePath) (*Batch, error) {
	wire.Build(
		wire.Struct(new(Batch), "*"),
		config.GetConfig,
		wire.FieldsOf(new(*config.Config), "AWS"),
		client.NewClient,
		wire.Value(&client.Config{}),
		serviceSet,
		facadeSet,
		domainServiceSet,
		repository.RepositoriesSet,
		repository.ProvideS3Uploader,
		repository.ProvideFacebookSession,
	)

	return new(Batch), nil
}
