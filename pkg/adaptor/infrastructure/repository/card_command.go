package repository

import (
	"context"

	"github.com/google/wire"
	"github.com/payjp/payjp-go/v1"
	"github.com/pkg/errors"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/entity"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/repository"
)

type CardCommandRepositoryImpl struct {
	DAO
	PayjpClient *payjp.Service
}

var CardCommandRepositorySet = wire.NewSet(
	wire.Struct(new(CardCommandRepositoryImpl), "*"),
	wire.Bind(new(repository.CardCommandRepository), new(*CardCommandRepositoryImpl)),
)

func (r *CardCommandRepositoryImpl) Store(c context.Context, card *entity.Card) error {
	if err := r.DB(c).Save(card).Error; err != nil {
		return errors.Wrap(err, "failed store credit_card")
	}
	return nil
}
