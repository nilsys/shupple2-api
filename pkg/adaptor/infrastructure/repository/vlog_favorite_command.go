package repository

import (
	"context"

	"github.com/pkg/errors"

	"github.com/google/wire"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/entity"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/repository"
)

type VlogFavoriteCommandRepositoryImpl struct {
	DAO
}

var VlogFavoriteCommandRepositorySet = wire.NewSet(
	wire.Struct(new(VlogFavoriteCommandRepositoryImpl), "*"),
	wire.Bind(new(repository.VlogFavoriteCommandRepository), new(*VlogFavoriteCommandRepositoryImpl)),
)

func (r *VlogFavoriteCommandRepositoryImpl) Store(c context.Context, favorite *entity.UserFavoriteVlog) error {
	return errors.Wrap(r.DB(c).Save(favorite).Error, "failed to save favorite")
}

func (r *VlogFavoriteCommandRepositoryImpl) Delete(c context.Context, unfavorite *entity.UserFavoriteVlog) error {
	return errors.Wrap(r.DB(c).Where("user_id = ? AND vlog_id = ?", unfavorite.UserID, unfavorite.VlogID).Delete(&unfavorite).Error, "failed to delete")
}
