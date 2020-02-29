package service

import (
	"github.com/google/wire"
	"github.com/pkg/errors"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/entity"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/model"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/model/serror"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/repository"
)

type (
	CategoryQueryService interface {
		ShowAreaListByParams(parentCategoryID int, limit int, excludeID []int) ([]*entity.Category, error)
		ShowAreaByID(id int) (*entity.Category, error)
		IsTypeAreaGroupByID(id int) (bool, error)
		IsTypeAreaByID(id int) (bool, error)
	}

	CategoryQueryServiceImpl struct {
		Repository repository.CategoryQueryRepository
	}
)

var CategoryQueryServiceSet = wire.NewSet(
	wire.Struct(new(CategoryQueryServiceImpl), "*"),
	wire.Bind(new(CategoryQueryService), new(*CategoryQueryServiceImpl)),
)

func (r *CategoryQueryServiceImpl) ShowAreaListByParams(parentCategoryID int, limit int, excludeID []int) ([]*entity.Category, error) {
	typeMatch, err := r.IsTypeAreaGroupByID(parentCategoryID)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to IsTypeAreaGroupByID")
	}
	if !typeMatch {
		return nil, serror.New(nil, serror.CodeInvalidParam, "parentCategoryID's type is not area_group")
	}

	categories, err := r.Repository.FindListByParentCategoryID(parentCategoryID, limit, excludeID)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to get Categories by FindListByParentCategoryID")
	}

	return categories, nil
}

func (r *CategoryQueryServiceImpl) ShowAreaByID(id int) (*entity.Category, error) {
	typeMatch, err := r.IsTypeAreaByID(id)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to IsTypeAreaByID")
	}
	if !typeMatch {
		return nil, serror.New(nil, serror.CodeInvalidParam, "id:%d is not area", id)
	}

	category, err := r.Repository.FindByID(id)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to FindByID")
	}

	return category, nil
}

func (r *CategoryQueryServiceImpl) IsTypeAreaGroupByID(id int) (bool, error) {
	categoryType, err := r.Repository.FindTypeByID(id)
	if err != nil {
		return false, errors.Wrapf(err, "failed to FindTypeByID")
	}

	return *categoryType == model.CategoryTypeAreaGroup, nil
}

func (r *CategoryQueryServiceImpl) IsTypeAreaByID(id int) (bool, error) {
	categoryType, err := r.Repository.FindTypeByID(id)
	if err != nil {
		return false, errors.Wrapf(err, "failed to FindTypeByID")
	}

	return *categoryType == model.CategoryTypeArea, nil
}
