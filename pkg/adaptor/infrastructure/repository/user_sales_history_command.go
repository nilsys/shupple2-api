package repository

import (
	"context"

	"github.com/pkg/errors"

	"github.com/google/wire"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/entity"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/repository"
)

type UserSalesHistoryCommandRepositoryImpl struct {
	DAO
}

var UserSalesHistoryCommandRepositorySet = wire.NewSet(
	wire.Struct(new(UserSalesHistoryCommandRepositoryImpl), "*"),
	wire.Bind(new(repository.UserSalesHistoryCommandRepository), new(*UserSalesHistoryCommandRepositoryImpl)),
)

func (r *UserSalesHistoryCommandRepositoryImpl) Store(ctx context.Context, history *entity.UserSalesHistoryTiny) error {
	if err := r.DB(ctx).Save(history).Error; err != nil {
		return errors.Wrap(err, "failed store user_sales_history")
	}
	return nil
}
