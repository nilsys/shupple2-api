package repository

import (
	"github.com/google/wire"
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
	"github.com/uma-co82/shupple2-api/pkg/domain/entity"
	"github.com/uma-co82/shupple2-api/pkg/domain/model"
	"github.com/uma-co82/shupple2-api/pkg/domain/repository"
)

type UserQueryRepositoryImpl struct {
	DB *gorm.DB
}

var UserQueryRepositorySet = wire.NewSet(
	wire.Struct(new(UserQueryRepositoryImpl), "*"),
	wire.Bind(new(repository.UserQueryRepository), new(*UserQueryRepositoryImpl)),
)

func (r *UserQueryRepositoryImpl) FindByFirebaseID(id string) (*entity.UserTiny, error) {
	var row entity.UserTiny
	if err := r.DB.Where("firebase_id = ?", id).First(&row).Error; err != nil {
		return nil, errors.Wrap(err, "failed find user")
	}
	return &row, nil
}

func (r *UserQueryRepositoryImpl) FindTinyByID(id int) (*entity.UserTiny, error) {
	var row entity.UserTiny
	if err := r.DB.First(&row, id).Error; err != nil {
		return nil, errors.Wrap(err, "failed find user")
	}
	return &row, nil
}

func (r *UserQueryRepositoryImpl) FindByID(id int) (*entity.User, error) {
	var row entity.User
	if err := r.DB.First(&row, id).Error; err != nil {
		return nil, errors.Wrap(err, "failed find user")
	}
	return &row, nil
}

func (r *UserQueryRepositoryImpl) FindMatchingUserByID(id int) (*entity.User, error) {
	var row entity.User
	if err := r.DB.Where("id = (SELECT matching_user_id FROM user_matching_history WHERE user_id = ? ORDER BY created_at DESC LIMIT 1)", id).First(&row).Error; err != nil {
		return nil, errors.Wrap(err, "failed find user")
	}
	return &row, nil
}

func (r *UserQueryRepositoryImpl) FindAvailableMatchingUser(gender model.Gender, reason model.MatchingReason, id int) (*entity.UserTiny, error) {
	var row entity.UserTiny
	if err := r.DB.
		Where("is_matching = false AND gender = ? AND matching_reason = ? AND id NOT IN (?) AND id NOT IN (SELECT matching_user_id FROM user_matching_history WHERE user_id = ?)", gender.Reverse(), reason, id, id).
		First(&row).Error; err != nil {
		return nil, errors.Wrap(err, "failed find user")
	}
	return &row, nil
}

func (r *UserQueryRepositoryImpl) FindMatchingHistoryByUserIDAndMatchingUserID(userID, matchingUserID int) (*entity.UserMatchingHistory, error) {
	var row entity.UserMatchingHistory
	if err := r.DB.Where("user_id = ? AND matching_user_id = ?", userID, matchingUserID).First(&row).Error; err != nil {
		return nil, ErrorToFindSingleRecord(err, "(user_matching_history=%d, %d)", userID, matchingUserID)
	}
	return &row, nil
}

/*
	マッチング後、評価していないUser一覧
*/
func (r *UserQueryRepositoryImpl) FindNotMainMatchingReviewMatchingUsersByID(id int) ([]*entity.User, error) {
	var rows []*entity.User
	if err := r.DB.
		Joins("INNER JOIN user_matching_history ON user.id = user_matching_history.matching_user_id").
		Where("user_matching_history.user_id = ? AND user_matching_history.user_main_matching_approve IS NULL", id).Find(&rows).Error; err != nil {
		return nil, errors.Wrap(err, "failed find user")
	}
	return rows, nil
}

func (r *UserQueryRepositoryImpl) FindMainMatchingUserByID(id int) ([]*entity.User, error) {
	var rows []*entity.User
	if err := r.DB.
		Joins("INNER JOIN user_matching_history ON user.id = user_matching_history.matching_user_id").
		Where("user_matching_history.user_id = ? AND user_matching_history.user_main_matching_approve = true AND user_matching_history.matching_user_main_matching_approve = true", id).Find(&rows).Error; err != nil {
	}
	return rows, nil
}

func (r *UserQueryRepositoryImpl) FindImageByUUID(uuid string) (*entity.UserImage, error) {
	var row entity.UserImage
	if err := r.DB.First(&row, uuid).Error; err != nil {
		return nil, ErrorToFindSingleRecord(err, "user_image(id=%s)", uuid)
	}
	return &row, nil
}
