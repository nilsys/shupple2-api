package repository

import (
	"context"

	"github.com/stayway-corp/stayway-media-api/pkg/domain/entity"
)

type (
	CardCommandRepository interface {
		Store(c context.Context, card *entity.Card) error
	}

	CardQueryRepository interface {
		FindLatestByUserID(c context.Context, userID int) (*entity.Card, error)
	}
)
