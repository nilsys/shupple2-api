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
	AreaQueryService interface {
		ListAreaByParams(areaGroupID int, limit int, excludeID []int) ([]*entity.Category, error)
		ListSubAreaByParams(areaID int, limit int, excludeID []int) ([]*entity.Category, error)
		ListSubSubAreaByParams(subAreaID int, limit int, excludeID []int) ([]*entity.Category, error)

		ShowAreaByID(id int) (*entity.Category, error)
		ShowSubAreaByID(id int) (*entity.Category, error)
		ShowSubSubAreaByID(id int) (*entity.Category, error)
	}

	AreaQueryServiceImpl struct {
		Repository repository.CategoryQueryRepository
	}
)

var AreaQueryServiceSet = wire.NewSet(
	wire.Struct(new(AreaQueryServiceImpl), "*"),
	wire.Bind(new(AreaQueryService), new(*AreaQueryServiceImpl)),
)

func (r *AreaQueryServiceImpl) ListAreaByParams(areaGroupID int, limit int, excludeID []int) ([]*entity.Category, error) {
	categories, err := r.showList(areaGroupID, model.CategoryTypeAreaGroup, model.CategoryTypeArea, limit, excludeID)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to ShowList")
	}
	return categories, nil
}

func (r *AreaQueryServiceImpl) ListSubAreaByParams(areaID int, limit int, excludeID []int) ([]*entity.Category, error) {
	categories, err := r.showList(areaID, model.CategoryTypeArea, model.CategoryTypeSubArea, limit, excludeID)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to ShowList")
	}
	return categories, nil
}

func (r *AreaQueryServiceImpl) ListSubSubAreaByParams(subAreaID int, limit int, excludeID []int) ([]*entity.Category, error) {
	categories, err := r.showList(subAreaID, model.CategoryTypeSubArea, model.CategoryTypeSubSubArea, limit, excludeID)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to ShowList")
	}
	return categories, nil
}

func (r *AreaQueryServiceImpl) ShowAreaByID(id int) (*entity.Category, error) {
	category, err := r.showByID(id, model.CategoryTypeArea)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to showByID")
	}
	return category, nil
}

func (r *AreaQueryServiceImpl) ShowSubAreaByID(id int) (*entity.Category, error) {
	category, err := r.showByID(id, model.CategoryTypeSubArea)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to showByID")
	}
	return category, nil
}

func (r *AreaQueryServiceImpl) ShowSubSubAreaByID(id int) (*entity.Category, error) {
	category, err := r.showByID(id, model.CategoryTypeSubSubArea)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to showByID")
	}
	return category, nil
}

// parentIDをもとにCategoryテーブル内の該当レコードのTypeがtargetTypeと一致するか確認したのち、categoryを取得する
func (r *AreaQueryServiceImpl) showList(parentID int, parentCategoryType, categoryType model.CategoryType, limit int, excludeID []int) ([]*entity.Category, error) {
	typeMatch, err := r.isTypeByID(parentID, parentCategoryType)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to IsTypeByID")
	}
	if !typeMatch {
		return nil, serror.New(nil, serror.CodeInvalidParam, "id:%d is not %s", parentID, parentCategoryType)
	}

	categories, err := r.Repository.FindListByParentID(parentID, limit, excludeID)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to get Categories by FindListByID")
	}

	return r.filterByCategoryType(categories, categoryType), nil
}

// idをもとにCategoryテーブル内の該当レコードのTypeがcategoryTypeと一致するか確認したのち、categoryを取得する
func (r *AreaQueryServiceImpl) showByID(id int, categoryType model.CategoryType) (*entity.Category, error) {
	typeMatch, err := r.isTypeByID(id, categoryType)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to IsTypeByID")
	}
	if !typeMatch {
		return nil, serror.New(nil, serror.CodeInvalidParam, "id:%d is not %s", id, categoryType)
	}

	category, err := r.Repository.FindByID(id)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to FindByID")
	}

	return category, nil
}

func (r *AreaQueryServiceImpl) isTypeByID(id int, targetCategoryType model.CategoryType) (bool, error) {
	categoryType, err := r.Repository.FindTypeByID(id)
	if err != nil {
		return false, errors.Wrapf(err, "failed to FindTypeByID")
	}
	return *categoryType == targetCategoryType, nil
}

func (r *AreaQueryServiceImpl) filterByCategoryType(categories []*entity.Category, categoryType model.CategoryType) []*entity.Category {
	resp := []*entity.Category{}
	for _, category := range categories {
		if category.Type == categoryType {
			resp = append(resp, category)
		}
	}
	return resp
}
