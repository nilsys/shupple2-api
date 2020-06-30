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
	VlogQueryScenario interface {
		Show(id int, ouser *entity.OptionalUser) (*entity.VlogDetail, map[int]*entity.TouristSpotReviewCount, map[int]*entity.AreaCategory, map[int]*entity.ThemeCategory, error)
		ListByParams(query *query.FindVlogListQuery, ouser *entity.OptionalUser) (*entity.VlogList, map[int]*entity.AreaCategory, map[int]*entity.ThemeCategory, error)
	}

	VlogQueryScenarioImpl struct {
		service.VlogQueryService
		factory.CategoryIDMapFactory
	}
)

var VlogQueryScenarioSet = wire.NewSet(
	wire.Struct(new(VlogQueryScenarioImpl), "*"),
	wire.Bind(new(VlogQueryScenario), new(*VlogQueryScenarioImpl)),
)

func (s *VlogQueryScenarioImpl) Show(id int, ouser *entity.OptionalUser) (*entity.VlogDetail, map[int]*entity.TouristSpotReviewCount, map[int]*entity.AreaCategory, map[int]*entity.ThemeCategory, error) {
	vlog, touristSpotReviewCountList, err := s.VlogQueryService.Show(id, ouser)
	if err != nil {
		return nil, nil, nil, nil, errors.Wrap(err, "failed find vlog by id")
	}

	areaCategoriesMap, themeCategoriesMap, err := s.CategoryIDMapFactory.GenerateCategoryIDMap(vlog.AreaCategoryIDs(), vlog.ThemeCategoryIDs())
	if err != nil {
		return nil, nil, nil, nil, errors.Wrap(err, "failed gen category map")
	}

	return vlog, touristSpotReviewCountList.IDMap(), areaCategoriesMap, themeCategoriesMap, nil
}

func (s *VlogQueryScenarioImpl) ListByParams(query *query.FindVlogListQuery, ouser *entity.OptionalUser) (*entity.VlogList, map[int]*entity.AreaCategory, map[int]*entity.ThemeCategory, error) {
	vlogs, err := s.VlogQueryService.ListByParams(query, ouser)
	if err != nil {
		return nil, nil, nil, errors.Wrap(err, "failed to find vlogs")
	}

	areaCategoriesMap, themeCategoriesMap, err := s.CategoryIDMapFactory.GenerateCategoryIDMap(vlogs.AreaCategoryIDs(), vlogs.ThemeCategoryIDs())
	if err != nil {
		return nil, nil, nil, errors.Wrap(err, "failed gen category map")
	}

	return vlogs, areaCategoriesMap, themeCategoriesMap, nil
}
