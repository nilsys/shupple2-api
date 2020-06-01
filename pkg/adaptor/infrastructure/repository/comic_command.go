package repository

import (
	"context"

	"github.com/google/wire"
	"github.com/pkg/errors"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/entity"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/repository"
)

type ComicCommandRepositoryImpl struct {
	DAO
}

var ComicCommandRepositorySet = wire.NewSet(
	wire.Struct(new(ComicCommandRepositoryImpl), "*"),
	wire.Bind(new(repository.ComicCommandRepository), new(*ComicCommandRepositoryImpl)),
)

func (r *ComicCommandRepositoryImpl) Lock(c context.Context, id int) (*entity.Comic, error) {
	var row entity.Comic
	if err := r.LockDB(c).First(&row, id).Error; err != nil {
		return nil, ErrorToFindSingleRecord(err, "comic(id=%d)", id)
	}
	return &row, nil
}

func (r *ComicCommandRepositoryImpl) Store(c context.Context, comic *entity.Comic) error {
	return errors.Wrap(r.DB(c).Save(comic).Error, "failed to save comic")
}

func (r *ComicCommandRepositoryImpl) DeleteByID(id int) error {
	return errors.Wrapf(r.DB(context.Background()).Delete(&entity.Comic{ID: id}).Error, "failed to delete comic(id=%d)", id)
}

func (r *ComicCommandRepositoryImpl) IncrementFavoriteCount(c context.Context, id int) error {
	return errors.Wrapf(r.DB(c).Exec("UPDATE comic SET favorite_count = favorite_count + 1 WHERE id = ?", id).Error, "failed to update comic(id=%d)", id)
}

func (r *ComicCommandRepositoryImpl) DecrementFavoriteCount(c context.Context, id int) error {
	return errors.Wrapf(r.DB(c).Exec("UPDATE comic SET favorite_count = favorite_count - 1 WHERE id = ?", id).Error, "failed to update comic(id=%d)", id)
}
