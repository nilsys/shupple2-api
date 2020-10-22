package repository

import (
	"context"

	"github.com/google/wire"
	"github.com/pkg/errors"
	"github.com/uma-co82/shupple2-api/pkg/domain/entity"
	"github.com/uma-co82/shupple2-api/pkg/domain/repository"
)

type UserCommandRepositoryImpl struct {
	DAO
}

var UserCommandRepositorySet = wire.NewSet(
	wire.Struct(new(UserCommandRepositoryImpl), "*"),
	wire.Bind(new(repository.UserCommandRepository), new(*UserCommandRepositoryImpl)),
)

func (r *UserCommandRepositoryImpl) Store(ctx context.Context, user *entity.UserTiny) error {
	if err := r.DB(ctx).Save(user).Error; err != nil {
		return errors.Wrap(err, "failed store user")
	}
	return nil
}

func (r *UserCommandRepositoryImpl) StoreUserImages(ctx context.Context, images []*entity.UserImage) error {
	db := r.DB(ctx)
	for _, image := range images {
		if err := db.Save(image).Error; err != nil {
			return errors.Wrap(err, "failed store user_image")
		}
	}
	return nil
}

func (r *UserCommandRepositoryImpl) StoreUserMatchingHistory(ctx context.Context, history *entity.UserMatchingHistory) error {
	if err := r.DB(ctx).Save(history).Error; err != nil {
		return errors.Wrap(err, "failed store user_matching_history")
	}
	return nil
}

func (r *UserCommandRepositoryImpl) UpdateForMatchingByIDs(ctx context.Context, ids []int) error {
	if err := r.DB(ctx).Exec("UPDATE user SET is_matching = true WHERE id IN (?)", ids).Error; err != nil {
		return errors.Wrap(err, "failed update user.is_matching = true")
	}
	return nil
}

func (r *UserCommandRepositoryImpl) UpdateForNotMatchingByID(ctx context.Context, id int) error {
	if err := r.DB(ctx).Exec("UPDATE user SET is_matching = false WHERE id = ?", id).Error; err != nil {
		return errors.Wrap(err, "failed update user.is_matching = false")
	}
	return nil
}
