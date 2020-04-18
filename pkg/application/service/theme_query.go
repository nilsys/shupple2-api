package service

import (
	"github.com/google/wire"
	"github.com/pkg/errors"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/entity"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/repository"
)

type (
	ThemeCategoryQueryService interface {
		ListThemeByParams(excludeIDs []int, categoryID int) ([]*entity.ThemeCategoryWithPostCount, error)
	}

	ThemeCategoryQueryServiceImpl struct {
		repository.ThemeCategoryQueryRepository
	}
)

var ThemeCategoryQueryServiceSet = wire.NewSet(
	wire.Struct(new(ThemeCategoryQueryServiceImpl), "*"),
	wire.Bind(new(ThemeCategoryQueryService), new(*ThemeCategoryQueryServiceImpl)),
)

func (r *ThemeCategoryQueryServiceImpl) ListThemeByParams(excludeIDs []int, categoryID int) ([]*entity.ThemeCategoryWithPostCount, error) {
	if categoryID == 0 {
		categories, err := r.ThemeCategoryQueryRepository.FindAll(excludeIDs)
		if err != nil {
			return nil, errors.Wrapf(err, "failed to find all themes")
		}
		return categories, nil
	}

	categories, err := r.ThemeCategoryQueryRepository.FindThemesByAreaCategoryID(excludeIDs, categoryID)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to find themes by areaCategoryID")
	}

	return categories, nil
}
