package repository

import (
	"context"

	"github.com/stayway-corp/stayway-media-api/pkg/domain/entity"
)

type (
	NoticeCommandRepository interface {
		StoreNotice(c context.Context, notice *entity.Notice) error
		MarkAsRead(noticeIDs []int) error
	}

	NoticeQueryRepository interface {
		ListNotice(userID int, limit int) ([]*entity.Notice, error)
	}
)
