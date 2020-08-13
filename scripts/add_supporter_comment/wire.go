//+build wireinject

package main

import (
	"github.com/google/wire"
	"github.com/stayway-corp/stayway-media-api/pkg/adaptor/infrastructure/repository"
	"github.com/stayway-corp/stayway-media-api/pkg/application/scenario"
	"github.com/stayway-corp/stayway-media-api/pkg/application/service"
	"github.com/stayway-corp/stayway-media-api/pkg/application/service/helper"
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
	service.AreaCategoryCommandServiceSet,
	service.ThemeCategoryCommandServiceSet,
	service.ComicCommandServiceSet,
	service.FeatureCommandServiceSet,
	service.SpotCategoryCommandServiceSet,
	service.TouristSpotCommandServiceSet,
	service.VlogCommandServiceSet,
	service.HashtagCommandServiceSet,
	service.ReviewCommandServiceSet,
	service.ReviewQueryServiceSet,
	service.CardCommandServiceSet,
	service.ChargeCommandServiceSet,
	service.ShippingCommandServiceSet,
	scenario.ReviewCommandScenarioSet,
	domain_service.NoticeDomainServiceSet,
	domain_service.TaggedUserDomainServiceSet,
	helper.InquiryCodeGeneratorSet,
)

func InitializeScript(configFilePath config.FilePath) (*Script, error) {
	wire.Build(
		wire.Struct(new(Script), "*"),
		config.GetConfig,
		wire.FieldsOf(new(*config.Config), "AWS", "CfProject"),
		serviceSet,
		repository.RepositoriesSet,
	)

	return new(Script), nil
}
