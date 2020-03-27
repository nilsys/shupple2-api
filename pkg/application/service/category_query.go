package service

import (
	"github.com/google/wire"
	"github.com/pkg/errors"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/entity"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/model/serror"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/repository"
)

type (
	CategoryQueryService interface {
		ShowBySlug(slug string) (entity.Category, error)
	}

	CategoryQueryServiceImpl struct {
		AreaCategoryRepository  repository.AreaCategoryQueryRepository
		ThemeCategoryRepository repository.ThemeCategoryQueryRepository
		LcategoryRepository     repository.LcategoryQueryRepository
	}
)

var CategoryQueryServiceSet = wire.NewSet(
	wire.Struct(new(CategoryQueryServiceImpl), "*"),
	wire.Bind(new(CategoryQueryService), new(*CategoryQueryServiceImpl)),
)

func (r *CategoryQueryServiceImpl) ShowBySlug(slug string) (entity.Category, error) {
	areaCategory, err := r.AreaCategoryRepository.FindBySlug(slug)
	if err != nil && serror.IsErrorCode(err, serror.CodeNotFound) {
		return nil, errors.Wrap(err, "failed to find area category by slug")
	}
	if areaCategory != nil {
		return areaCategory, nil
	}

	themeCategory, err := r.ThemeCategoryRepository.FindBySlug(slug)
	if err != nil && serror.IsErrorCode(err, serror.CodeNotFound) {
		return nil, errors.Wrap(err, "failed to find theme category by slug")
	}
	if themeCategory != nil {
		return themeCategory, nil
	}

	lcategory, err := r.AreaCategoryRepository.FindBySlug(slug)
	if err != nil && serror.IsErrorCode(err, serror.CodeNotFound) {
		return nil, errors.Wrap(err, "failed to find lcategory by slug")
	}

	return lcategory, nil
}
