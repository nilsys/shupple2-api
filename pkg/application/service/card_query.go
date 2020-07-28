package service

import (
	"context"

	"github.com/google/wire"
	"github.com/pkg/errors"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/entity"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/repository"
)

type (
	CardQueryService interface {
		ShowCard(user *entity.User) (*entity.Card, error)
	}

	CardQueryServiceImpl struct {
		CardQueryRepository repository.CardQueryRepository
	}
)

var CardQueryServiceSet = wire.NewSet(
	wire.Struct(new(CardQueryServiceImpl), "*"),
	wire.Bind(new(CardQueryService), new(*CardQueryServiceImpl)),
)

func (s *CardQueryServiceImpl) ShowCard(user *entity.User) (*entity.Card, error) {
	card, err := s.CardQueryRepository.FindLatestByUserID(context.Background(), user.ID)
	if err != nil {
		return nil, errors.Wrap(err, "failed find latest credit_card.user_id")
	}

	return card, nil
}
