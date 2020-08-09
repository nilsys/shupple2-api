package service

import (
	"github.com/google/wire"
	"github.com/pkg/errors"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/entity"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/model/query"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/repository"
)

type (
	PaymentQueryService interface {
		ListByUser(user *entity.User, projectID int, query *query.FindListPaginationQuery) (*entity.PaymentList, error)
	}

	PaymentQueryServiceImpl struct {
		repository.PaymentQueryRepository
	}
)

var PaymentQueryServiceSet = wire.NewSet(
	wire.Struct(new(PaymentQueryServiceImpl), "*"),
	wire.Bind(new(PaymentQueryService), new(*PaymentQueryServiceImpl)),
)

func (s *PaymentQueryServiceImpl) ListByUser(user *entity.User, projectID int, query *query.FindListPaginationQuery) (*entity.PaymentList, error) {
	payments, err := s.PaymentQueryRepository.FindByUserID(user.ID, projectID, query)
	if err != nil {
		return nil, errors.Wrap(err, "failed payment_query_repository.FindByUserID")
	}
	return payments, nil
}
