package repository

import (
	"github.com/google/wire"
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/entity"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/model"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/repository"
)

type InterestQueryRepositoryImpl struct {
	DB *gorm.DB
}

var InterestQueryRepositorySet = wire.NewSet(
	wire.Struct(new(InterestQueryRepositoryImpl), "*"),
	wire.Bind(new(repository.InterestQueryRepository), new(*InterestQueryRepositoryImpl)),
)

func (r *InterestQueryRepositoryImpl) FindAllByGroup(group model.InterestGroup) ([]*entity.Interest, error) {
	var rows []*entity.Interest

	q := r.buildFindAllByGroupQuery(group)

	if err := q.Limit(defaultAcquisitionNumber).Order("interest_group").Find(&rows).Error; err != nil {
		return nil, errors.Wrap(err, "failed to find all interest")
	}

	return rows, nil
}

func (r *InterestQueryRepositoryImpl) buildFindAllByGroupQuery(group model.InterestGroup) *gorm.DB {
	q := r.DB

	if group != model.InterestGroupUndefined {
		q = q.Where("interest_group = ?", group)
	}

	return q
}
