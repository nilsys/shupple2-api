package repository

import (
	"context"

	"github.com/google/wire"
	"github.com/pkg/errors"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/entity"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/repository"
)

// Notice更新系レポジトリ実装
type NoticeCommandRepositoryImpl struct {
	DAO
}

var NoticeCommandRepositorySet = wire.NewSet(
	wire.Struct(new(NoticeCommandRepositoryImpl), "*"),
	wire.Bind(new(repository.NoticeCommandRepository), new(*NoticeCommandRepositoryImpl)),
)

func (r *NoticeCommandRepositoryImpl) StoreNotice(c context.Context, notice *entity.Notice) error {
	if err := r.DB(c).Save(notice).Error; err != nil {
		return errors.Wrap(err, "failed store notice")
	}
	return nil
}

func (r *NoticeCommandRepositoryImpl) MarkAsRead(c context.Context, noticeID, userID int) error {
	if err := r.DB(c).Exec("UPDATE notice SET is_read = true WHERE id = ? AND user_id = ?", noticeID, userID).Error; err != nil {
		return errors.Wrap(err, "Failed update to be marked as read")
	}

	return nil
}
