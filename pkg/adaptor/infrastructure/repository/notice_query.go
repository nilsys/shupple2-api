package repository

import (
	"github.com/google/wire"
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/entity"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/repository"
)

type NoticeQueryRepositoryImpl struct {
	DB *gorm.DB
}

var NoticeQueryRepositorySet = wire.NewSet(
	wire.Struct(new(NoticeQueryRepositoryImpl), "*"),
	wire.Bind(new(repository.NoticeQueryRepository), new(*NoticeQueryRepositoryImpl)),
)

func (r NoticeQueryRepositoryImpl) ListNotice(userID int, limit int) ([]*entity.Notice, error) {
	var results []*entity.Notice
	err := r.DB.
		Where("user_id = ?", userID).
		Order("created_at DESC").
		Limit(limit).
		Find(&results).
		Error

	if err != nil {
		return nil, errors.Wrap(err, "failed find notices")
	}

	return results, nil
}
