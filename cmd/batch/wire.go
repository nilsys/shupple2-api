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
	service.LcategoryCommandServiceSet,
	service.WordpressCallbackServiceSet,
	service.UserQueryServiceSet,
	service.UserCommandServiceSet,
	service.AreaCategoryQueryServiceSet,
	service.AreaCategoryCommandServiceSet,
	service.ProvideAuthService,
	service.ThemeCategoryCommandServiceSet,
)

func InitializeBatch(configFilePath config.FilePath) (*Batch, error) {
	wire.Build(
		wire.Struct(new(Batch), "*"),
		config.GetConfig,
		wire.FieldsOf(new(*config.Config), "Wordpress", "Stayway", "AWS"),
		wire.FieldsOf(new(config.Stayway), "Media"),
		serviceSet,
		repository.RepositoriesSet,
		repository.ProvideS3Uploader,
	)

	return new(Batch), nil
}
