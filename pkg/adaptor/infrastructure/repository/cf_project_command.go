package repository

import (
	"context"

	"github.com/google/wire"
	"github.com/pkg/errors"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/entity"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/repository"
)

type (
	CfProjectCommandRepositoryImpl struct {
		DAO
	}
)

var CfProjectCommandRepositorySet = wire.NewSet(
	wire.Struct(new(CfProjectCommandRepositoryImpl), "*"),
	wire.Bind(new(repository.CfProjectCommandRepository), new(*CfProjectCommandRepositoryImpl)),
)

func (r *CfProjectCommandRepositoryImpl) StoreUserFavoriteCfProject(c context.Context, fav *entity.UserFavoriteCfProject) error {
	if err := r.DB(c).Save(fav).Error; err != nil {
		return errors.Wrap(err, "failed store user_favorite_cf_project")
	}
	return nil
}

func (r *CfProjectCommandRepositoryImpl) StoreSupportComment(c context.Context, comment *entity.CfProjectSupportCommentTable) error {
	if err := r.DB(c).Save(comment).Error; err != nil {
		return errors.Wrap(err, "failed store cf_project_support_comment")
	}
	return nil
}

func (r *CfProjectCommandRepositoryImpl) DeleteUserFavoriteCfProject(c context.Context, fav *entity.UserFavoriteCfProject) error {
	if err := r.DB(c).Delete(fav).Error; err != nil {
		return errors.Wrap(err, "failed delete user_favorite_cf_project")
	}
	return nil
}

func (r *CfProjectCommandRepositoryImpl) IncrementFavoriteCountByID(c context.Context, projectID int) error {
	if err := r.DB(c).Exec("UPDATE cf_project SET favorite_count=favorite_count+1 WHERE id = ?", projectID).Error; err != nil {
		return errors.Wrap(err, "failed increment cf_project.favorite_count")
	}
	return nil
}

func (r *CfProjectCommandRepositoryImpl) DecrementFavoriteCountByID(c context.Context, projectID int) error {
	if err := r.DB(c).Exec("UPDATE cf_project SET favorite_count=favorite_count-1 WHERE id = ?", projectID).Error; err != nil {
		return errors.Wrap(err, "failed decrement cf_project.favorite_count")
	}
	return nil
}

func (r *CfProjectCommandRepositoryImpl) IncrementSupportCommentCount(c context.Context, id int) error {
	if err := r.DB(c).Exec("UPDATE cf_project SET support_comment_count=support_comment_count+1 WHERE id = ?", id).Error; err != nil {
		return errors.Wrap(err, "failed to increment cf_project.support_comment_count")
	}
	return nil
}
