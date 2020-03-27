package repository

import (
	"context"

	"github.com/google/wire"
	"github.com/pkg/errors"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/entity"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/repository"
)

type CategoryCommandRepositoryImpl struct {
	DAO
}

var CategoryCommandRepositorySet = wire.NewSet(
	wire.Struct(new(CategoryCommandRepositoryImpl), "*"),
	wire.Bind(new(repository.CategoryCommandRepository), new(*CategoryCommandRepositoryImpl)),
)

func (r *CategoryCommandRepositoryImpl) Lock(c context.Context, id int) (*entity.Category, error) {
	var row entity.Category
	if err := r.LockDB(c).First(&row, id).Error; err != nil {
		return nil, ErrorToFindSingleRecord(err, "category(id=%d)", id)
	}
	return &row, nil
}

func (r *CategoryCommandRepositoryImpl) Store(c context.Context, category *entity.Category) error {
	return errors.Wrap(r.DB(c).Save(category).Error, "failed to save category")
}
