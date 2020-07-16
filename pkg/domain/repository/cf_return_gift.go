package repository

import (
	"context"

	"github.com/stayway-corp/stayway-media-api/pkg/domain/entity"
)

type (
	CfReturnGiftCommandRepository interface {
		Store(*entity.CfReturnGift) error
		LockByIDs(c context.Context, ids []int) (*entity.CfReturnGiftList, error)
		UndeleteByID(c context.Context, id int) error
		DeleteByID(id int) error
	}

	CfReturnGiftQueryRepository interface {
		FindSoldCountByReturnGiftIDs(c context.Context, ids []int) (*entity.CfReturnGiftSoldCountList, error)
		FindByCfProjectID(projectID int) (*entity.CfReturnGiftList, error)
	}
)
