package repository

import (
	"context"

	"github.com/stayway-corp/stayway-media-api/pkg/domain/entity"
)

type (
	NoticeCommandRepository interface {
		StoreNotice(c context.Context, notice *entity.Notice) error
		MarkAsRead(c context.Context, noticeID, userID int) error
	}

	NoticeQueryRepository interface {
		ListNotice(userID int, limit int) (*entity.NoticeList, error)
		UnreadPushNoticeCount(c context.Context, userID int) (int, error)
	}
)
