package repository

import (
	"context"

	"github.com/uma-co82/shupple2-api/pkg/domain/entity"
)

type (
	UserQueryRepository interface {
		FindByFirebaseID(id string) (*entity.UserTiny, error)
		FindByID(id int) (*entity.UserTiny, error)
	}

	UserCommandRepository interface {
		Store(ctx context.Context, user *entity.UserTiny) error
	}
)
