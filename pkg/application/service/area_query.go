package service

import (
	"github.com/google/wire"
	"github.com/pkg/errors"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/entity"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/model"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/repository"
)

type (
	AreaQueryService interface {
		ListAreaByParams(areaGroupID model.AreaGroup, limit int, excludeID []int) ([]*entity.AreaCategoryWithPostCount, error)
		ListSubAreaByParams(areaID, themeID int, limit int, excludeID []int) ([]*entity.AreaCategoryWithPostCount, error)
		ListSubSubAreaByParams(subAreaID, themeID int, limit int, excludeID []int) ([]*entity.AreaCategoryWithPostCount, error)

		ShowAreaByID(id int) (*entity.AreaCategoryDetail, error)
		ShowSubAreaByID(id int) (*entity.AreaCategoryDetail, error)
		ShowSubSubAreaByID(id int) (*entity.AreaCategoryDetail, error)
	}

	AreaQueryServiceImpl struct {
		Repository repository.AreaCategoryQueryRepository
	}
)

var AreaCategoryQueryServiceSet = wire.NewSet(
	wire.Struct(new(AreaQueryServiceImpl), "*"),
	wire.Bind(new(AreaQueryService), new(*AreaQueryServiceImpl)),
)

func (r *AreaQueryServiceImpl) ListAreaByParams(areaGroup model.AreaGroup, limit int, excludeID []int) ([]*entity.AreaCategoryWithPostCount, error) {
	areaCategories, err := r.Repository.FindAreaListHavingPostByAreaGroup(areaGroup, limit, excludeID)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to list area")
	}
	return areaCategories, nil
}

func (r *AreaQueryServiceImpl) ListSubAreaByParams(areaID, themeID int, limit int, excludeID []int) ([]*entity.AreaCategoryWithPostCount, error) {
	areaCategories, err := r.Repository.FindSubAreaListHavingPostByAreaIDAndThemeID(areaID, themeID, limit, excludeID)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to list sub area")
	}
	return areaCategories, nil
}

func (r *AreaQueryServiceImpl) ListSubSubAreaByParams(subAreaID, themeID int, limit int, excludeID []int) ([]*entity.AreaCategoryWithPostCount, error) {
	areaCategories, err := r.Repository.FindSubSubAreaListHavingPostBySubAreaIDAndThemeID(subAreaID, themeID, limit, excludeID)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to list sub sub area")
	}
	return areaCategories, nil
}

func (r *AreaQueryServiceImpl) ShowAreaByID(id int) (*entity.AreaCategoryDetail, error) {
	areaCategory, err := r.Repository.FindDetailByIDAndType(id, model.AreaCategoryTypeArea)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to show area")
	}
	return areaCategory, nil
}

func (r *AreaQueryServiceImpl) ShowSubAreaByID(id int) (*entity.AreaCategoryDetail, error) {
	areaCategory, err := r.Repository.FindDetailByIDAndType(id, model.AreaCategoryTypeSubArea)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to show sub area")
	}
	return areaCategory, nil
}

func (r *AreaQueryServiceImpl) ShowSubSubAreaByID(id int) (*entity.AreaCategoryDetail, error) {
	areaCategory, err := r.Repository.FindDetailByIDAndType(id, model.AreaCategoryTypeSubSubArea)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to show sub sub area")
	}
	return areaCategory, nil
}
