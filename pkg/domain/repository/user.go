package repository

import (
	"context"

	"github.com/uma-co82/shupple2-api/pkg/domain/model"

	"github.com/uma-co82/shupple2-api/pkg/domain/entity"
)

type (
	UserQueryRepository interface {
		FindByFirebaseID(id string) (*entity.UserTiny, error)
		FindTinyByID(id int) (*entity.UserTiny, error)
		FindByID(id int) (*entity.User, error)
		FindMatchingUserByID(id int) (*entity.User, error)
		FindAvailableMatchingUser(gender model.Gender, reason model.MatchingReason, id int) (*entity.UserTiny, error)
		FindMatchingHistoryByUserIDAndMatchingUserID(userID, matchingUserID int) (*entity.UserMatchingHistory, error)
		FindPendingMainMatchingMatchingUsersByID(id int) ([]*entity.User, error)
		FindMainMatchingUserByID(id int) ([]*entity.User, error)
		FindImageByUUID(uuid string) (*entity.UserImage, error)
		IsExistMainMatchingUserMatchingHistory(id, matchingUserID int) (bool, error)
	}

	UserCommandRepository interface {
		Store(ctx context.Context, user *entity.UserTiny) error
		StoreUserImages(ctx context.Context, image *entity.UserImage) error
		StoreUserMatchingHistory(ctx context.Context, history *entity.UserMatchingHistory) error
		UpdateLatestMatchingUserID(ctx context.Context, id, matchingUserID int) error
		UpdateUserMatchingHistoryUserMainMatchingApprove(ctx context.Context, userID, matchingUserID int, isApprove bool) error
		UpdateUserMatchingHistoryMatchingUserMainMatchingApprove(ctx context.Context, userID, matchingUserID int, isApprove bool) error
		UpdateMatchingExpiredUserLatestMatchingUserID() error
	}
)
