package service

import (
	"context"

	payjp2 "github.com/stayway-corp/stayway-media-api/pkg/domain/repository/payjp"

	"github.com/google/wire"
	"github.com/payjp/payjp-go/v1"
	"github.com/pkg/errors"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/entity"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/repository"
)

type (
	CardQueryService interface {
		ShowCard(user *entity.User) (*payjp.CardResponse, error)
	}

	CardQueryServiceImpl struct {
		CardQueryRepository      repository.CardQueryRepository
		PayjpCardQueryRepository payjp2.CardQueryRepository
	}
)

var CardQueryServiceSet = wire.NewSet(
	wire.Struct(new(CardQueryServiceImpl), "*"),
	wire.Bind(new(CardQueryService), new(*CardQueryServiceImpl)),
)

func (s *CardQueryServiceImpl) ShowCard(user *entity.User) (*payjp.CardResponse, error) {
	creditCard, err := s.CardQueryRepository.FindLatestByUserID(context.Background(), user.ID)
	if err != nil {
		return nil, errors.Wrap(err, "failed find latest credit_card.user_id")
	}
	cardResponse, err := s.PayjpCardQueryRepository.Find(user.PayjpCustomerID(), creditCard.CardID)
	if err != nil {
		return nil, errors.Wrap(err, "failed find card from pay.jp")
	}

	return cardResponse, nil
}
