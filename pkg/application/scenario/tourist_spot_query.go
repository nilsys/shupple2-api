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
	TouristSpotQueryScenario interface {
		Show(id int) (*entity.TouristSpotDetail, map[int]*entity.AreaCategory, map[int]*entity.ThemeCategory, error)
		List(query *query.FindTouristSpotListQuery) (*entity.TouristSpotList, map[int]*entity.AreaCategory, map[int]*entity.ThemeCategory, error)
		ListRecommend(query *query.FindRecommendTouristSpotListQuery) (*entity.TouristSpotList, map[int]*entity.AreaCategory, map[int]*entity.ThemeCategory, error)
	}

	TouristSpotQueryScenarioImpl struct {
		factory.CategoryIDMapFactory
		service.TouristSpotQueryService
	}
)

var TouristSpotQueryScenarioSet = wire.NewSet(
	wire.Struct(new(TouristSpotQueryScenarioImpl), "*"),
	wire.Bind(new(TouristSpotQueryScenario), new(*TouristSpotQueryScenarioImpl)),
)

func (s *TouristSpotQueryScenarioImpl) Show(id int) (*entity.TouristSpotDetail, map[int]*entity.AreaCategory, map[int]*entity.ThemeCategory, error) {
	touristSpot, err := s.TouristSpotQueryService.Show(id)
	if err != nil {
		return nil, nil, nil, errors.Wrap(err, "failed find tourist_spot")
	}

	areaCategoriesMap, themeCategoriesMap, err := s.CategoryIDMapFactory.GenerateCategoryIDMap(touristSpot.AreaCategoryIDs(), touristSpot.ThemeCategoryIDs())
	if err != nil {
		return nil, nil, nil, errors.Wrap(err, "failed gen category map")
	}

	return touristSpot, areaCategoriesMap, themeCategoriesMap, nil
}

func (s *TouristSpotQueryScenarioImpl) List(query *query.FindTouristSpotListQuery) (*entity.TouristSpotList, map[int]*entity.AreaCategory, map[int]*entity.ThemeCategory, error) {
	touristSpots, err := s.TouristSpotQueryService.List(query)
	if err != nil {
		return nil, nil, nil, errors.Wrap(err, "failed list tourist_spot")
	}

	areaCategoriesMap, themeCategoriesMap, err := s.CategoryIDMapFactory.GenerateCategoryIDMap(touristSpots.AreaCategoryIDs(), touristSpots.ThemeCategoryIDs())
	if err != nil {
		return nil, nil, nil, errors.Wrap(err, "failed gen category map")
	}

	return touristSpots, areaCategoriesMap, themeCategoriesMap, nil
}

func (s *TouristSpotQueryScenarioImpl) ListRecommend(query *query.FindRecommendTouristSpotListQuery) (*entity.TouristSpotList, map[int]*entity.AreaCategory, map[int]*entity.ThemeCategory, error) {
	touristSpots, err := s.TouristSpotQueryService.ListRecommend(query)
	if err != nil {
		return nil, nil, nil, errors.Wrap(err, "failed list recommend tourist_spot")
	}

	areaCategoriesMap, themeCategoriesMap, err := s.CategoryIDMapFactory.GenerateCategoryIDMap(touristSpots.AreaCategoryIDs(), touristSpots.ThemeCategoryIDs())
	if err != nil {
		return nil, nil, nil, errors.Wrap(err, "failed gen category map")
	}

	return touristSpots, areaCategoriesMap, themeCategoriesMap, nil
}
