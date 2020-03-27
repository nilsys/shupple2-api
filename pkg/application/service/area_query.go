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
		ListAreaByParams(areaGroupID model.AreaGroup, limit int, excludeID []int) ([]*entity.AreaCategory, error)
		ListSubAreaByParams(areaID int, limit int, excludeID []int) ([]*entity.AreaCategory, error)
		ListSubSubAreaByParams(subAreaID int, limit int, excludeID []int) ([]*entity.AreaCategory, error)

		ShowAreaByID(id int) (*entity.AreaCategory, error)
		ShowSubAreaByID(id int) (*entity.AreaCategory, error)
		ShowSubSubAreaByID(id int) (*entity.AreaCategory, error)
	}

	AreaQueryServiceImpl struct {
		Repository repository.AreaCategoryQueryRepository
	}
)

var AreaCategoryQueryServiceSet = wire.NewSet(
	wire.Struct(new(AreaQueryServiceImpl), "*"),
	wire.Bind(new(AreaQueryService), new(*AreaQueryServiceImpl)),
)

func (r *AreaQueryServiceImpl) ListAreaByParams(areaGroup model.AreaGroup, limit int, excludeID []int) ([]*entity.AreaCategory, error) {
	areaCategories, err := r.Repository.FindAreaListByAreaGroup(areaGroup, limit, excludeID)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to list area")
	}
	return areaCategories, nil
}

func (r *AreaQueryServiceImpl) ListSubAreaByParams(areaID int, limit int, excludeID []int) ([]*entity.AreaCategory, error) {
	areaCategories, err := r.Repository.FindSubAreaListByAreaID(areaID, limit, excludeID)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to list sub area")
	}
	return areaCategories, nil
}

func (r *AreaQueryServiceImpl) ListSubSubAreaByParams(subAreaID int, limit int, excludeID []int) ([]*entity.AreaCategory, error) {
	areaCategories, err := r.Repository.FindSubSubAreaListBySubAreaID(subAreaID, limit, excludeID)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to list sub sub area")
	}
	return areaCategories, nil
}

func (r *AreaQueryServiceImpl) ShowAreaByID(id int) (*entity.AreaCategory, error) {
	areaCategory, err := r.Repository.FindByIDAndType(id, model.AreaCategoryTypeArea)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to show area")
	}
	return areaCategory, nil
}

func (r *AreaQueryServiceImpl) ShowSubAreaByID(id int) (*entity.AreaCategory, error) {
	areaCategory, err := r.Repository.FindByIDAndType(id, model.AreaCategoryTypeSubArea)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to show sub area")
	}
	return areaCategory, nil
}

func (r *AreaQueryServiceImpl) ShowSubSubAreaByID(id int) (*entity.AreaCategory, error) {
	areaCategory, err := r.Repository.FindByIDAndType(id, model.AreaCategoryTypeSubSubArea)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to show sub sub area")
	}
	return areaCategory, nil
}
