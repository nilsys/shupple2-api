package repository

import (
	"context"

	"github.com/google/wire"
	"github.com/pkg/errors"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/entity"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/repository"
)

type SpotCategoryCommandRepositoryImpl struct {
	DAO
}

var SpotCategoryCommandRepositorySet = wire.NewSet(
	wire.Struct(new(SpotCategoryCommandRepositoryImpl), "*"),
	wire.Bind(new(repository.SpotCategoryCommandRepository), new(*SpotCategoryCommandRepositoryImpl)),
)

func (r *SpotCategoryCommandRepositoryImpl) Lock(c context.Context, id int) (*entity.SpotCategory, error) {
	var row entity.SpotCategory
	if err := r.LockDB(c).First(&row, id).Error; err != nil {
		return nil, ErrorToFindSingleRecord(err, "spotCategory(id=%d)", id)
	}
	return &row, nil
}

func (r *SpotCategoryCommandRepositoryImpl) Store(c context.Context, spotCategory *entity.SpotCategory) error {
	return errors.Wrap(r.DB(c).Save(spotCategory).Error, "failed to save spotCategory")
}
