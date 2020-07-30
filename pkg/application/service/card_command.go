package service

import (
	"context"

	"github.com/stayway-corp/stayway-media-api/pkg/domain/model/serror"

	"github.com/stayway-corp/stayway-media-api/pkg/domain/model"

	payjp2 "github.com/stayway-corp/stayway-media-api/pkg/domain/repository/payjp"

	"github.com/google/wire"
	"github.com/pkg/errors"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/entity"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/repository"
)

type (
	CardCommandService interface {
		Register(user *entity.User, cardToken string) error
		Delete(user *entity.User, cardID int) error
	}

	CardCommandServiceImpl struct {
		CardCommandRepository      repository.CardCommandRepository
		CardQueryRepository        repository.CardQueryRepository
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

		if err := s.CardCommandRepository.Store(c, entity.NewCard(user.ID, card.ID, card.Last4, model.CardExpiredFromMonthAndYear(card.ExpMonth, card.ExpYear))); err != nil {
			return errors.Wrap(err, "failed store card")
		}
		return nil
	})
}

func (s *CardCommandServiceImpl) Delete(user *entity.User, cardID int) error {
	card, err := s.CardQueryRepository.FindByID(cardID)
	if err != nil {
		return errors.Wrap(err, "failed find card")
	}

	if !user.IsSelfID(card.UserID) {
		return serror.New(nil, serror.CodeForbidden, "forbidden")
	}

	return s.TransactionService.Do(func(ctx context.Context) error {
		if err := s.CardCommandRepository.Delete(ctx, card); err != nil {
			return errors.Wrap(err, "failed delete card")
		}

		if err := s.PayjpCardCommandRepository.Delete(user.PayjpCustomerID(), card.CardID); err != nil {
			return errors.Wrap(err, "failed delete card from pay.jp")
		}

		return nil
	})
}
