package repository

import (
	"github.com/google/wire"
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
	"github.com/uma-co82/shupple2-api/pkg/domain/entity"
	"github.com/uma-co82/shupple2-api/pkg/domain/repository"
)

type ArrangeScheduleRequestQueryRepositoryImpl struct {
	DB *gorm.DB
}

var ArrangeScheduleRequestQueryRepositorySet = wire.NewSet(
	wire.Struct(new(ArrangeScheduleRequestQueryRepositoryImpl), "*"),
	wire.Bind(new(repository.ArrangeScheduleRequestQueryRepository), new(*ArrangeScheduleRequestQueryRepositoryImpl)),
)

func (r *ArrangeScheduleRequestQueryRepositoryImpl) FindByMatchingUserID(userID int) ([]*entity.ArrangeScheduleRequest, error) {
	var rows []*entity.ArrangeScheduleRequest
	if err := r.DB.Where("matching_user_id = ?", userID).Find(&rows).Error; err != nil {
		return nil, errors.Wrap(err, "failed find arrange_schedule_request")
	}
	return rows, nil
}

func (r *ArrangeScheduleRequestQueryRepositoryImpl) FindByUserID(userID int) ([]*entity.ArrangeScheduleRequest, error) {
	var rows []*entity.ArrangeScheduleRequest
	if err := r.DB.Where("user_id = ?", userID).Find(&rows).Error; err != nil {
		return nil, errors.Wrap(err, "failed find arrange_schedule_request")
	}
	return rows, nil
}
