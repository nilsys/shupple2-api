package repository

import (
	"context"

	"github.com/google/wire"
	"github.com/pkg/errors"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/entity"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/repository"
)

type AreaCategoryCommandRepositoryImpl struct {
	DAO
}

var AreaCategoryCommandRepositorySet = wire.NewSet(
	wire.Struct(new(AreaCategoryCommandRepositoryImpl), "*"),
	wire.Bind(new(repository.AreaCategoryCommandRepository), new(*AreaCategoryCommandRepositoryImpl)),
)

func (r *AreaCategoryCommandRepositoryImpl) Lock(c context.Context, id int) (*entity.AreaCategory, error) {
	var row entity.AreaCategory
	if err := r.LockDB(c).First(&row, id).Error; err != nil {
		return nil, ErrorToFindSingleRecord(err, "areaCategory(id=%d)", id)
	}
	return &row, nil
}

func (r *AreaCategoryCommandRepositoryImpl) Store(c context.Context, areaCategory *entity.AreaCategory) error {
	return errors.Wrap(r.DB(c).Save(areaCategory).Error, "failed to save areaCategory")
}
