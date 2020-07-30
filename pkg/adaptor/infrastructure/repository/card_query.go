package repository

import (
	"context"

	"github.com/google/wire"
	"github.com/payjp/payjp-go/v1"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/entity"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/repository"
)

type CardQueryRepositoryImpl struct {
	DAO
	PayjpClient *payjp.Service
}

var CardQueryRepositorySet = wire.NewSet(
	wire.Struct(new(CardQueryRepositoryImpl), "*"),
	wire.Bind(new(repository.CardQueryRepository), new(*CardQueryRepositoryImpl)),
)

func (r *CardQueryRepositoryImpl) FindLatestByUserID(c context.Context, userID int) (*entity.Card, error) {
	var row entity.Card
	if err := r.DB(c).Where("user_id = ? AND deleted_at IS NULL", userID).Order("created_at desc").First(&row).Error; err != nil {
		return nil, ErrorToFindSingleRecord(err, "card(user_id=%d)", userID)
	}
	return &row, nil
}

func (r *CardQueryRepositoryImpl) FindByID(id int) (*entity.Card, error) {
	var row entity.Card
	if err := r.DB(context.Background()).Where("deleted_at IS NULL AND id = ?", id).First(&row).Error; err != nil {
		return nil, ErrorToFindSingleRecord(err, "card(id=%d)", id)
	}
	return &row, nil
}
