package repository

import (
	"context"

	"github.com/google/wire"
	"github.com/pkg/errors"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/entity"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/repository"
)

type ThemeCategoryCommandRepositoryImpl struct {
	DAO
}

var ThemeCategoryCommandRepositorySet = wire.NewSet(
	wire.Struct(new(ThemeCategoryCommandRepositoryImpl), "*"),
	wire.Bind(new(repository.ThemeCategoryCommandRepository), new(*ThemeCategoryCommandRepositoryImpl)),
)

func (r *ThemeCategoryCommandRepositoryImpl) Lock(c context.Context, id int) (*entity.ThemeCategory, error) {
	var row entity.ThemeCategory
	if err := r.LockDB(c).First(&row, id).Error; err != nil {
		return nil, ErrorToFindSingleRecord(err, "themeCategory(id=%d)", id)
	}
	return &row, nil
}

func (r *ThemeCategoryCommandRepositoryImpl) Store(c context.Context, themeCategory *entity.ThemeCategory) error {
	return errors.Wrap(r.DB(c).Save(themeCategory).Error, "failed to save themeCategory")
}
