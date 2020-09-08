package service

import (
	"context"

	"github.com/google/wire"

	"github.com/pkg/errors"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/entity"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/repository"
)

type (
	NoticeCommandService interface {
		MarkAsRead(user *entity.User, noticeID int) (int, error)
	}

	NoticeCommandServiceImpl struct {
		repository.NoticeQueryRepository
		repository.NoticeCommandRepository
		TransactionService
	}
)

var NoticeCommandServiceSet = wire.NewSet(
	wire.Struct(new(NoticeCommandServiceImpl), "*"),
	wire.Bind(new(NoticeCommandService), new(*NoticeCommandServiceImpl)),
)

// 指定されたpush_noticeを既読へ変更、その上で未読数を返す
// MEMO: アプリ側で分ける必要が発生すれば分ける
func (s *NoticeCommandServiceImpl) MarkAsRead(user *entity.User, noticeID int) (int, error) {
	var (
		unreadCount int
		err         error
	)

	err = s.TransactionService.Do(func(ctx context.Context) error {
		if err = s.NoticeCommandRepository.MarkAsRead(ctx, noticeID, user.ID); err != nil {
			return errors.Wrap(err, "failed mark as read push_notice")
		}

		unreadCount, err = s.NoticeQueryRepository.UnreadPushNoticeCount(ctx, user.ID)
		if err != nil {
			return errors.Wrap(err, "failed count unread push_notice")
		}

		return nil
	})

	if err != nil {
		return 0, errors.Wrap(err, "failed mark as read push notice transaction")
	}

	return unreadCount, nil
}
