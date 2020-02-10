package repository

import (
	"github.com/google/wire"
	"github.com/jinzhu/gorm"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/entity"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/repository"
)

type VlogQueryRepositoryImpl struct {
	DB *gorm.DB
}

var VlogQueryRepositorySet = wire.NewSet(
	wire.Struct(new(VlogQueryRepositoryImpl), "*"),
	wire.Bind(new(repository.VlogQueryRepository), new(*VlogQueryRepositoryImpl)),
)

func (r *VlogQueryRepositoryImpl) FindByID(id int) (*entity.Vlog, error) {
	var row entity.Vlog
	if err := r.DB.First(&row, id).Error; err != nil {
		return nil, ErrorToFindSingleRecord(err, "vlog(id=%d)", id)
	}
	return &row, nil
}
