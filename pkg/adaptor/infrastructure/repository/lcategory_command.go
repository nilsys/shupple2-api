package repository

import (
	"context"

	"github.com/google/wire"
	"github.com/pkg/errors"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/entity"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/repository"
)

type LcategoryCommandRepositoryImpl struct {
	DAO
}

var LcategoryCommandRepositorySet = wire.NewSet(
	wire.Struct(new(LcategoryCommandRepositoryImpl), "*"),
	wire.Bind(new(repository.LcategoryCommandRepository), new(*LcategoryCommandRepositoryImpl)),
)

func (r *LcategoryCommandRepositoryImpl) Lock(c context.Context, id int) (*entity.Lcategory, error) {
	var row entity.Lcategory
	if err := r.LockDB(c).First(&row, id).Error; err != nil {
		return nil, ErrorToFindSingleRecord(err, "lcategory(id=%d)", id)
	}
	return &row, nil
}

func (r *LcategoryCommandRepositoryImpl) Store(c context.Context, lcategory *entity.Lcategory) error {
	return errors.Wrap(r.DB(c).Save(lcategory).Error, "failed to save lcategory")
}
