package service

import (
	"github.com/google/wire"
	"github.com/payjp/payjp-go/v1"
	"github.com/pkg/errors"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/entity"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/model/query"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/repository"
	payjp2 "github.com/stayway-corp/stayway-media-api/pkg/domain/repository/payjp"
)

type (
	PaymentQueryService interface {
		ListByUser(user *entity.User, query *query.FindListPaginationQuery) (*entity.PaymentList, map[string]*payjp.CardResponse, error)
	}

	PaymentQueryServiceImpl struct {
		repository.PaymentQueryRepository
		PayjpCardQueryRepository payjp2.CardQueryRepository
	}
)

var PaymentQueryServiceSet = wire.NewSet(
	wire.Struct(new(PaymentQueryServiceImpl), "*"),
	wire.Bind(new(PaymentQueryService), new(*PaymentQueryServiceImpl)),
)

func (s *PaymentQueryServiceImpl) ListByUser(user *entity.User, query *query.FindListPaginationQuery) (*entity.PaymentList, map[string]*payjp.CardResponse, error) {
	payments, err := s.PaymentQueryRepository.FindByUserID(user.ID, query)
	if err != nil {
		return nil, nil, errors.Wrap(err, "failed payment_query_repository.FindByUserID")
	}
	cards, err := s.PayjpCardQueryRepository.FindList(user.PayjpCustomerID(), payments.CardIDs())
	if err != nil {
		return nil, nil, errors.Wrap(err, "failed find card from payjp")
	}
	return payments, s.cardResToIDMap(cards), nil
}

func (s *PaymentQueryServiceImpl) cardResToIDMap(cards []*payjp.CardResponse) map[string]*payjp.CardResponse {
	list := make(map[string]*payjp.CardResponse, len(cards))
	for _, card := range cards {
		list[card.ID] = card
	}
	return list
}
