package repository

import (
	"context"

	"github.com/stayway-corp/stayway-media-api/pkg/domain/entity"
)

type (
	CfReturnGiftQueryRepository interface {
		LockCfReturnGiftList(c context.Context, ids []int) (*entity.CfReturnGiftList, error)
		FindSoldCountByReturnGiftIDs(c context.Context, ids []int) (*entity.CfReturnGiftSoldCountList, error)
		FindByCfProjectID(projectID int) (*entity.CfReturnGiftList, error)
	}
)
