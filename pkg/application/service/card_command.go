package service

import (
	"context"

	payjp2 "github.com/stayway-corp/stayway-media-api/pkg/domain/repository/payjp"

	"github.com/google/wire"
	"github.com/pkg/errors"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/entity"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/repository"
)

type (
	CardCommandService interface {
		Register(user *entity.User, cardToken string) error
	}

	CardCommandServiceImpl struct {
		CardCommandRepository      repository.CardCommandRepository
		PayjpCardCommandRepository payjp2.CardCommandRepository
		TransactionService
	}
)

var CardCommandServiceSet = wire.NewSet(
	wire.Struct(new(CardCommandServiceImpl), "*"),
	wire.Bind(new(CardCommandService), new(*CardCommandServiceImpl)),
)

// カード情報登録
func (s *CardCommandServiceImpl) Register(user *entity.User, cardToken string) error {
	return s.TransactionService.Do(func(c context.Context) error {
		// pay.jp側へカード登録
		card, err := s.PayjpCardCommandRepository.Register(user.PayjpCustomerID(), cardToken)
		if err != nil {
			return errors.Wrap(err, "failed register card")
		}

		if err := s.CardCommandRepository.Store(c, entity.NewCard(user.ID, card.ID)); err != nil {
			return errors.Wrap(err, "failed store credit_card")
		}
		return nil
	})
}
