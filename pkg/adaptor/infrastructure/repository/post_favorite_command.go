package repository

import (
	"context"

	"github.com/google/wire"
	"github.com/pkg/errors"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/entity"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/repository"
)

type PostFavoriteCommandRepositoryImpl struct {
	DAO
}

var PostFavoriteCommandRepositorySet = wire.NewSet(
	wire.Struct(new(PostFavoriteCommandRepositoryImpl), "*"),
	wire.Bind(new(repository.PostFavoriteCommandRepository), new(*PostFavoriteCommandRepositoryImpl)),
)

func (r *PostFavoriteCommandRepositoryImpl) Store(c context.Context, favorite *entity.UserFavoritePost) error {
	return errors.Wrap(r.DB(c).Save(favorite).Error, "failed to save favorite")
}

func (r *PostFavoriteCommandRepositoryImpl) Delete(c context.Context, unfavorite *entity.UserFavoritePost) error {
	return errors.Wrapf(r.DB(c).Where("user_id = ? AND post_id = ?", unfavorite.UserID, unfavorite.PostID).Delete(&unfavorite).Error, "failed to delete")
}
