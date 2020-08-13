package repository

import (
	"context"

	"github.com/stayway-corp/stayway-media-api/pkg/domain/entity"
)

type (
	UserSalesHistoryCommandRepository interface {
		Store(ctx context.Context, history *entity.UserSalesHistoryTiny) error
	}
)
