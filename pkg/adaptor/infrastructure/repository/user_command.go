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

func (r *UserCommandRepositoryImpl) StoreUserImages(ctx context.Context, image *entity.UserImage) error {
	if err := r.DB(ctx).Save(image).Error; err != nil {
		return errors.Wrap(err, "failed store user_image")
	}
	return nil
}

func (r *UserCommandRepositoryImpl) StoreUserMatchingHistory(ctx context.Context, history *entity.UserMatchingHistory) error {
	if err := r.DB(ctx).Save(history).Error; err != nil {
		return errors.Wrap(err, "failed store user_matching_history")
	}
	return nil
}

func (r *UserCommandRepositoryImpl) UpdateLatestMatchingUserID(ctx context.Context, id, matchingUserID int) error {
	if err := r.DB(ctx).Exec("UPDATE user SET latest_matching_user_id = ? WHERE id = ?", matchingUserID, id).Error; err != nil {
		return errors.Wrap(err, "failed update user.latest_matching_user_id")
	}
	return nil
}

func (r *UserCommandRepositoryImpl) UpdateUserMatchingHistoryUserMainMatchingApprove(ctx context.Context, userID, matchingUserID int, isApprove bool) error {
	if err := r.DB(ctx).Exec("UPDATE user_matching_history SET user_main_matching_approve = ? WHERE user_id = ? AND matching_user_id = ?", isApprove, userID, matchingUserID).Error; err != nil {
		return errors.Wrap(err, "failed update user_matching_history")
	}
	return nil
}

func (r *UserCommandRepositoryImpl) UpdateUserMatchingHistoryMatchingUserMainMatchingApprove(ctx context.Context, userID, matchingUserID int, isApprove bool) error {
	if err := r.DB(ctx).Exec("UPDATE user_matching_history SET matching_user_main_matching_approve = ? WHERE user_id = ? AND matching_user_id = ?", isApprove, userID, matchingUserID).Error; err != nil {
		return errors.Wrap(err, "failed update user_matching_history")
	}
	return nil
}

func (r *UserCommandRepositoryImpl) UpdateMatchingExpiredUserLatestMatchingUserID() error {
	if err := r.DB(context.Background()).Exec("UPDATE user u JOIN user_matching_history um ON u.id = um.user_id AND u.latest_matching_user_id = um.matching_user_id SET u.latest_matching_user_id = NULL WHERE um.matching_expired_at < NOW()").Error; err != nil {
		return errors.Wrap(err, "failed update user latest_matching_user_id")
	}
	return nil
}
