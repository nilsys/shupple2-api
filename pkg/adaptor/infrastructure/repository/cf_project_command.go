package repository

import (
	"context"

	"github.com/google/wire"
	"github.com/jinzhu/gorm"
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

func (r *CfProjectCommandRepositoryImpl) Store(cfProject *entity.CfProject) error {
	return Transaction(r.DB(context.Background()), func(db *gorm.DB) error {
		if err := db.Set("gorm:insert_modifier", "IGNORE").Create(&cfProject.CfProjectTiny).Error; err != nil {
			return errors.Wrap(err, "failed to insert cf_project")
		}

		cfProject.Snapshot.SnapshotID = 0
		if err := db.Create(&cfProject.Snapshot).Error; err != nil {
			return errors.Wrap(err, "failed to insert cf_project_snapshot")
		}

		if err := db.Exec("UPDATE cf_project SET latest_snapshot_id = ? WHERE id = ?", cfProject.Snapshot.SnapshotID, cfProject.ID).Error; err != nil {
			return errors.Wrap(err, "failed to update latest_snapshot_id")
		}

		return nil
	})
}

func (r *CfProjectCommandRepositoryImpl) Lock(c context.Context, id int) (*entity.CfProject, error) {
	var rows entity.CfProject
	if err := r.LockDB(c).Find(&rows, id).Error; err != nil {
		return nil, ErrorToFindSingleRecord(err, "cf_project(id=%d)", id)
	}
	return &rows, nil
}

func (r *CfProjectCommandRepositoryImpl) StoreUserFavoriteCfProject(c context.Context, fav *entity.UserFavoriteCfProject) error {
	if err := r.DB(c).Save(fav).Error; err != nil {
		return errors.Wrap(err, "failed store user_favorite_cf_project")
	}
	return nil
}

func (r *CfProjectCommandRepositoryImpl) StoreSupportComment(c context.Context, comment *entity.CfProjectSupportCommentTiny) error {
	if err := r.DB(c).Save(comment).Error; err != nil {
		return errors.Wrap(err, "failed store cf_project_support_comment")
	}
	return nil
}

func (r *CfProjectCommandRepositoryImpl) UndeleteByID(c context.Context, id int) error {
	e := &entity.CfProjectTiny{}
	e.ID = id
	return errors.Wrapf(
		r.DB(c).Unscoped().Model(e).Update("deleted_at", gorm.Expr("NULL")).Error,
		"failed to cfproject post(id=%d)", id)
}

func (r *CfProjectCommandRepositoryImpl) DeleteByID(id int) error {
	return errors.Wrapf(r.DB(context.Background()).Delete(&entity.CfProjectTiny{ID: id}).Error, "failed to delete cfproject(id=%d)", id)
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

func (r *CfProjectCommandRepositoryImpl) IncrementAchievedPrice(c context.Context, id, price int) error {
	if err := r.DB(c).Exec("UPDATE cf_project SET achieved_price=achieved_price+? WHERE id = ?", price, id).Error; err != nil {
		return errors.Wrap(err, "failed to increment cf_project.achieved_price")
	}
	return nil
}

func (r *CfProjectCommandRepositoryImpl) MarkAsIsSentAchievementNoticeMail(id int) error {
	if err := r.DB(context.Background()).Exec("UPDATE cf_project SET is_send_achievement_mail = true WHERE id = ?", id).Error; err != nil {
		return errors.Wrap(err, "failed update is_send")
	}
	return nil
}
