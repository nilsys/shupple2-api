package service

import (
	"github.com/google/wire"
	"github.com/pkg/errors"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/entity"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/repository"
)

type (
	NoticeQueryService interface {
		ListNotice(user *entity.User) (*entity.NoticeList, error)
	}

	NoticeQueryServiceImpl struct {
		repository.NoticeQueryRepository
		repository.NoticeCommandRepository
		repository.ReviewQueryRepository
		TransactionService
	}
)

var NoticeQueryServiceSet = wire.NewSet(
	wire.Struct(new(NoticeQueryServiceImpl), "*"),
	wire.Bind(new(NoticeQueryService), new(*NoticeQueryServiceImpl)),
)

var noticeLimit = 100

func (s *NoticeQueryServiceImpl) ListNotice(user *entity.User) (*entity.NoticeList, error) {
	notices, err := s.NoticeQueryRepository.ListNotice(user.ID, noticeLimit)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to get notices")
	}

	return notices, nil
}
