package repository

import (
	"context"

	"github.com/stayway-corp/stayway-media-api/pkg/domain/entity"
)

type (
	ReturnGiftQueryRepository interface {
		LockReturnGiftsWitLatestSummary(c context.Context, ids []int) (*entity.CfReturnGiftList, error)
		FindSoldCountByReturnGiftIDs(c context.Context, ids []int) (*entity.CfReturnGiftSoldCountList, error)
	}
)
