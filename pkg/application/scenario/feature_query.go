package scenario

import (
	"github.com/google/wire"
	"github.com/pkg/errors"
	"github.com/stayway-corp/stayway-media-api/pkg/application/service"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/entity"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/factory"
)

type (
	FeatureQueryScenario interface {
		Show(id int) (*entity.FeatureDetailWithPosts, map[int]*entity.AreaCategory, map[int]*entity.ThemeCategory, error)
	}

	FeatureQueryScenarioImpl struct {
		factory.CategoryIDMapFactory
		service.FeatureQueryService
	}
)

var FeatureQueryScenarioSet = wire.NewSet(
	wire.Struct(new(FeatureQueryScenarioImpl), "*"),
	wire.Bind(new(FeatureQueryScenario), new(*FeatureQueryScenarioImpl)),
)

func (s *FeatureQueryScenarioImpl) Show(id int) (*entity.FeatureDetailWithPosts, map[int]*entity.AreaCategory, map[int]*entity.ThemeCategory, error) {
	feature, err := s.FeatureQueryService.ShowQuery(id)
	if err != nil {
		return nil, nil, nil, errors.Wrap(err, "failed feature_query_service.ShowQuery")
	}

	areaCategoriesMap, themeCategoriesMap, err := s.CategoryIDMapFactory.GenerateCategoryIDMap(feature.AreaCategoryIDs(), feature.ThemeCategoryIDs())
	if err != nil {
		return nil, nil, nil, errors.Wrap(err, "failed gen category map")
	}

	return feature, areaCategoriesMap, themeCategoriesMap, nil
}
