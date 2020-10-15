package repository

import (
	"context"

	"github.com/google/wire"
	"github.com/pkg/errors"
	"github.com/uma-co82/shupple2-api/pkg/domain/entity"
	"github.com/uma-co82/shupple2-api/pkg/domain/repository"
)

type UserCommandRepositoryImpl struct {
	DAO
}

var UserCommandRepositorySet = wire.NewSet(
	wire.Struct(new(UserCommandRepositoryImpl), "*"),
	wire.Bind(new(repository.UserCommandRepository), new(*UserCommandRepositoryImpl)),
)

func (r *UserCommandRepositoryImpl) Store(ctx context.Context, user *entity.UserTiny) error {
	if err := r.DB(ctx).Save(user).Error; err != nil {
		return errors.Wrap(err, "failed store user")
	}
	return nil
}
