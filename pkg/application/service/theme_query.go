package service

import (
	"github.com/google/wire"
	"github.com/pkg/errors"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/entity"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/repository"
)

type (
	ThemeCategoryQueryService interface {
		ListThemeByParams(areaID, subAreaID, subSubAreaID int, excludeIDs []int) ([]*entity.ThemeCategoryWithPostCount, error)
		ListSubThemeByParams(themeID, areaID, subAreaID, subSubAreaID int, excludeIDs []int) ([]*entity.ThemeCategoryWithPostCount, error)
	}

	ThemeCategoryQueryServiceImpl struct {
		repository.ThemeCategoryQueryRepository
	}
)

var ThemeCategoryQueryServiceSet = wire.NewSet(
	wire.Struct(new(ThemeCategoryQueryServiceImpl), "*"),
	wire.Bind(new(ThemeCategoryQueryService), new(*ThemeCategoryQueryServiceImpl)),
)

func (r *ThemeCategoryQueryServiceImpl) ListThemeByParams(areaID, subAreaID, subSubAreaID int, excludeIDs []int) ([]*entity.ThemeCategoryWithPostCount, error) {
	categories, err := r.ThemeCategoryQueryRepository.FindThemesByAreaCategoryID(areaID, subAreaID, subSubAreaID, excludeIDs)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to find theme list")
	}

	return categories, nil
}

func (r *ThemeCategoryQueryServiceImpl) ListSubThemeByParams(themeID, areaID, subAreaID, subSubAreaID int, excludeIDs []int) ([]*entity.ThemeCategoryWithPostCount, error) {
	categories, err := r.ThemeCategoryQueryRepository.FindSubThemesByAreaCategoryIDAndParentThemeID(themeID, areaID, subAreaID, subSubAreaID, excludeIDs)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to find subTheme list")
	}

	return categories, nil
}
