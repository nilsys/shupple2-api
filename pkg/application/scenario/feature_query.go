package scenario

import (
	"github.com/google/wire"
	"github.com/pkg/errors"
	"github.com/stayway-corp/stayway-media-api/pkg/application/service"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/entity"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/factory"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/model/query"
)

type (
	FeatureQueryScenario interface {
		Show(id int, ouser *entity.OptionalUser) (*entity.FeatureDetailWithPosts, map[int]*entity.AreaCategory, map[int]*entity.ThemeCategory, *entity.UserRelationFlgMap, error)
		List(query *query.FindListPaginationQuery) (*entity.FeatureList, error)
	}

	FeatureQueryScenarioImpl struct {
		factory.CategoryIDMapFactory
		service.FeatureQueryService
		service.UserQueryService
	}
)

var FeatureQueryScenarioSet = wire.NewSet(
	wire.Struct(new(FeatureQueryScenarioImpl), "*"),
	wire.Bind(new(FeatureQueryScenario), new(*FeatureQueryScenarioImpl)),
)

func (s *FeatureQueryScenarioImpl) Show(id int, ouser *entity.OptionalUser) (*entity.FeatureDetailWithPosts, map[int]*entity.AreaCategory, map[int]*entity.ThemeCategory, *entity.UserRelationFlgMap, error) {
	idRelationFlgMap := &entity.UserRelationFlgMap{}

	feature, err := s.FeatureQueryService.Show(id)
	if err != nil {
		return nil, nil, nil, nil, errors.Wrap(err, "failed feature_query_service.Show")
	}

	areaCategoriesMap, themeCategoriesMap, err := s.CategoryIDMapFactory.GenerateCategoryIDMap(feature.AreaCategoryIDs(), feature.ThemeCategoryIDs())
	if err != nil {
		return nil, nil, nil, nil, errors.Wrap(err, "failed gen category map")
	}

	if ouser.IsAuthorized() {
		idRelationFlgMap, err = s.UserQueryService.RelationFlgMaps(ouser.ID, []int{feature.UserID})
		if err != nil {
			return nil, nil, nil, nil, errors.Wrap(err, "failed find is doing flg")
		}
	}

	return feature, areaCategoriesMap, themeCategoriesMap, idRelationFlgMap, nil
}

func (s *FeatureQueryScenarioImpl) List(query *query.FindListPaginationQuery) (*entity.FeatureList, error) {
	return s.FeatureQueryService.List(query)
}
