package service

import (
	"github.com/google/wire"
	"github.com/pkg/errors"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/entity"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/repository"
)

type (
	NoticeQueryService interface {
		ListNotice(user *entity.User) ([]*entity.Notice, error)
	}

	NoticeQueryServiceImpl struct {
		repository.NoticeQueryRepository
		repository.NoticeCommandRepository
		TransactionService
	}
)

var NoticeQueryServiceSet = wire.NewSet(
	wire.Struct(new(NoticeQueryServiceImpl), "*"),
	wire.Bind(new(NoticeQueryService), new(*NoticeQueryServiceImpl)),
)

var noticeLimit = 100

func (s *NoticeQueryServiceImpl) ListNotice(user *entity.User) ([]*entity.Notice, error) {
	notices, err := s.NoticeQueryRepository.ListNotice(user.ID, noticeLimit)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to get notices")
	}

	// 未読のお知らせを取得する
	unreadNoticeIds := make([]int, 0, len(notices))
	for _, notice := range notices {
		if !notice.IsRead {
			unreadNoticeIds = append(unreadNoticeIds, notice.ID)
		}
	}

	// 未読のお知らせを既読にする
	if err := s.NoticeCommandRepository.MarkAsRead(unreadNoticeIds); err != nil {
		return nil, errors.Wrap(err, "Failed to change finished reading notices")
	}

	return notices, nil
}
