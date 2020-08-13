package repository

import (
	"context"

	"github.com/google/wire"
	"github.com/pkg/errors"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/entity"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/repository"
)

type CfInnReserveRequestCommandRepositoryImpl struct {
	DAO
}

var CfInnReserveRequestCommandRepositorySet = wire.NewSet(
	wire.Struct(new(CfInnReserveRequestCommandRepositoryImpl), "*"),
	wire.Bind(new(repository.CfInnReserveRequestCommandRepository), new(*CfInnReserveRequestCommandRepositoryImpl)),
)

func (r *CfInnReserveRequestCommandRepositoryImpl) Store(ctx context.Context, reserveReq *entity.CfInnReserveRequest) error {
	if err := r.DB(ctx).Save(reserveReq).Error; err != nil {
		return errors.Wrap(err, "failed store cf_reserve_request")
	}
	return nil
}
