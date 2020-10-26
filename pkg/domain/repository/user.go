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
		FindNotConfirmMatchingUsersByID(id int) ([]*entity.User, error)
		FindConfirmMatchingUserByID(id int) ([]*entity.User, error)
	}

	UserCommandRepository interface {
		Store(ctx context.Context, user *entity.UserTiny) error
		StoreUserImages(ctx context.Context, images []*entity.UserImage) error
		StoreUserMatchingHistory(ctx context.Context, history *entity.UserMatchingHistory) error
		UpdateIsMatchingToTrueByIDs(ctx context.Context, ids []int) error
		UpdateIsMatchingToFalseByID(ctx context.Context, id int) error
		UpdateUserMatchingHistoryUserConfirmed(ctx context.Context, userID, matchingUserID int, isConfirm bool) error
		UpdateUserMatchingHistoryMatchingUserConfirmed(ctx context.Context, userID, matchingUserID int, isConfirm bool) error
	}
)
