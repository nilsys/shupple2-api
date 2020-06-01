package repository

import (
	"context"

	"github.com/google/wire"
	"github.com/pkg/errors"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/entity"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/repository"
)

type ComicFavoriteCommandRepositoryImpl struct {
	DAO
}

var ComicFavoriteCommandRepositorySet = wire.NewSet(
	wire.Struct(new(ComicFavoriteCommandRepositoryImpl), "*"),
	wire.Bind(new(repository.ComicFavoriteCommandRepository), new(*ComicFavoriteCommandRepositoryImpl)),
)

func (r *ComicFavoriteCommandRepositoryImpl) Store(c context.Context, favorite *entity.UserFavoriteComic) error {
	return errors.Wrap(r.DB(c).Save(favorite).Error, "failed to save favorite")
}

func (r *ComicFavoriteCommandRepositoryImpl) Delete(c context.Context, unfavorite *entity.UserFavoriteComic) error {
	return errors.Wrapf(r.DB(c).Where("user_id = ? AND comic_id = ?", unfavorite.UserID, unfavorite.ComicID).Delete(unfavorite).Error, "failed to delete comic(comic_id=%d,user_id=%d)", unfavorite.ComicID, unfavorite.UserID)
}
