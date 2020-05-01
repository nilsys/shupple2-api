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
		ShowBySlug(slug string) (entity.Category, model.AreaGroup, error)
	}

	CategoryQueryServiceImpl struct {
		AreaCategoryRepository  repository.AreaCategoryQueryRepository
		ThemeCategoryRepository repository.ThemeCategoryQueryRepository
		SpotCategoryRepository  repository.SpotCategoryQueryRepository
	}
)

var CategoryQueryServiceSet = wire.NewSet(
	wire.Struct(new(CategoryQueryServiceImpl), "*"), wire.Bind(new(CategoryQueryService), new(*CategoryQueryServiceImpl)),
)

func (r *CategoryQueryServiceImpl) ShowBySlug(slug string) (entity.Category, model.AreaGroup, error) {
	areaCategory, err := r.AreaCategoryRepository.FindBySlug(slug)
	// area_categoryで見つからなければtheme_categoryから探す
	if err != nil && !serror.IsErrorCode(err, serror.CodeNotFound) {
		return nil, model.AreaGroupUndefined, errors.Wrap(err, "failed to find area category by slug")
	}
	if areaCategory != nil {
		return areaCategory, areaCategory.AreaGroup, nil
	}

	themeCategory, err := r.ThemeCategoryRepository.FindBySlug(slug)
	// theme_categoryでも見つからなければ404
	if err != nil {
		return nil, model.AreaGroupUndefined, errors.Wrap(err, "failed to find theme category by slug")
	}

	return themeCategory, model.AreaGroupUndefined, nil
}
