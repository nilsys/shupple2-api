package repository

import (
	"context"

	"github.com/google/wire"
	"github.com/pkg/errors"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/entity"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/repository"
)

type NoticeQueryRepositoryImpl struct {
	DAO
}

var NoticeQueryRepositorySet = wire.NewSet(
	wire.Struct(new(NoticeQueryRepositoryImpl), "*"),
	wire.Bind(new(repository.NoticeQueryRepository), new(*NoticeQueryRepositoryImpl)),
)

func (r NoticeQueryRepositoryImpl) ListNotice(userID int, limit int) (*entity.NoticeList, error) {
	var results entity.NoticeList
	err := r.DB(context.Background()).
		Where("user_id = ?", userID).
		Order("created_at DESC").
		Limit(limit).
		Find(&results.List).
		Error

	if err != nil {
		return nil, errors.Wrap(err, "failed find notices")
	}

	return &results, nil
}

func (r NoticeQueryRepositoryImpl) UnreadPushNoticeCount(c context.Context, userID int) (int, error) {
	var count int

	if err := r.DB(c).
		Table("notice").
		Where("user_id = ? AND is_read = false", userID).
		Count(&count).
		Error; err != nil {
		return 0, errors.Wrap(err, "failed count unread push_notice")
	}

	return count, nil
}
