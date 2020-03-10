package repository

import (
	"github.com/google/wire"
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/entity"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/repository"
)

type InterestQueryRepositoryImpl struct {
	DB *gorm.DB
}

var InterestQueryRepositorySet = wire.NewSet(
	wire.Struct(new(InterestQueryRepositoryImpl), "*"),
	wire.Bind(new(repository.InterestQueryRepository), new(*InterestQueryRepositoryImpl)),
)

func (r *InterestQueryRepositoryImpl) FindAll() ([]*entity.Interest, error) {
	var rows []*entity.Interest

	if err := r.DB.Limit(defaultAcquisitionNumber).Order("created_at desc").Find(&rows).Error; err != nil {
		return nil, errors.Wrap(err, "failed to find all interest")
	}

	return rows, nil
}
