package repository

import (
	"context"

	"github.com/uma-co82/shupple2-api/pkg/domain/entity"
)

type (
	ArrangeScheduleRequestQueryRepository interface {
		FindByMatchingUserID(userID int) ([]*entity.ArrangeScheduleRequest, error)
		FindByUserID(userID int) ([]*entity.ArrangeScheduleRequest, error)
	}

	ArrangeScheduleRequestCommandRepository interface {
		Store(ctx context.Context, request *entity.ArrangeScheduleRequest) error
	}
)
