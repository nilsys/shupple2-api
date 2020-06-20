package repository

import (
	"context"

	"github.com/google/wire"
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/entity"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/repository"
)

type PostCommandRepositoryImpl struct {
	DAO
}

var PostCommandRepositorySet = wire.NewSet(
	wire.Struct(new(PostCommandRepositoryImpl), "*"),
	wire.Bind(new(repository.PostCommandRepository), new(*PostCommandRepositoryImpl)),
)

func (r *PostCommandRepositoryImpl) Lock(c context.Context, id int) (*entity.Post, error) {
	var row entity.Post
	if err := r.LockDB(c).Unscoped().First(&row, id).Error; err != nil {
		return nil, ErrorToFindSingleRecord(err, "post(id=%d)", id)
	}
	return &row, nil
}

func (r *PostCommandRepositoryImpl) Store(c context.Context, post *entity.Post) error {
	return errors.Wrap(r.DB(c).Save(post).Error, "failed to save post")
}

func (r *PostCommandRepositoryImpl) UndeleteByID(c context.Context, id int) error {
	e := &entity.Post{}
	e.ID = id
	return errors.Wrapf(
		r.DB(c).Unscoped().Model(e).Update("deleted_at", gorm.Expr("NULL")).Error,
		"failed to delete post(id=%d)", id)
}

func (r *PostCommandRepositoryImpl) DeleteByID(c context.Context, id int) error {
	e := &entity.Post{}
	e.ID = id
	return errors.Wrapf(r.DB(c).Delete(e).Error, "failed to delete post(id=%d)", id)
}

func (r *PostCommandRepositoryImpl) IncrementFavoriteCount(c context.Context, postID int) error {
	return errors.Wrapf(r.DB(c).Exec("UPDATE post SET favorite_count = favorite_count + 1 WHERE id = ?", postID).Error, "failed to update")
}

func (r *PostCommandRepositoryImpl) DecrementFavoriteCount(c context.Context, postID int) error {
	return errors.Wrapf(r.DB(c).Exec("UPDATE post SET favorite_count = favorite_count - 1 WHERE id = ?", postID).Error, "failed to update")
}

func (r *PostCommandRepositoryImpl) UpdateViewsByID(id, views int) error {
	if err := r.DB(context.Background()).Exec("UPDATE post SET views = ?, updated_at = updated_at WHERE id = ?", views, id).Error; err != nil {
		return errors.Wrap(err, "failed to update post.views")
	}
	return nil
}

func (r *PostCommandRepositoryImpl) UpdateMonthlyViewsByID(id, views int) error {
	if err := r.DB(context.Background()).Exec("UPDATE post SET monthly_views = ?, updated_at = updated_at WHERE id = ?", views, id).Error; err != nil {
		return errors.Wrap(err, "failed to update post.monthly_views")
	}
	return nil
}

func (r *PostCommandRepositoryImpl) UpdateWeeklyViewsByID(id, views int) error {
	if err := r.DB(context.Background()).Exec("UPDATE post SET weekly_views = ?, updated_at = updated_at WHERE id = ?", views, id).Error; err != nil {
		return errors.Wrap(err, "failed to update post.weekly_views")
	}
	return nil
}
