package repository

import (
	"context"

	"github.com/pkg/errors"

	"github.com/google/wire"
	"github.com/uma-co82/shupple2-api/pkg/domain/entity"
	"github.com/uma-co82/shupple2-api/pkg/domain/repository"
)

type ArrangeScheduleRequestCommandRepositoryImpl struct {
	DAO
}

var ArrangeScheduleRequestCommandRepositorySet = wire.NewSet(
	wire.Struct(new(ArrangeScheduleRequestCommandRepositoryImpl), "*"),
	wire.Bind(new(repository.ArrangeScheduleRequestCommandRepository), new(*ArrangeScheduleRequestCommandRepositoryImpl)),
)

func (r *ArrangeScheduleRequestCommandRepositoryImpl) Store(ctx context.Context, request *entity.ArrangeScheduleRequestTiny) error {
	if err := r.DB(ctx).Save(request).Error; err != nil {
		return errors.Wrap(err, "failed store arrange_schedule_request")
	}
	return nil
}
