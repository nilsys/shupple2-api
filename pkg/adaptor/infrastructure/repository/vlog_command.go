package repository

import (
	"github.com/google/wire"
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/entity"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/repository"
)

type VlogCommandRepositoryImpl struct {
	DB *gorm.DB
}

var VlogCommandRepositorySet = wire.NewSet(
	wire.Struct(new(VlogCommandRepositoryImpl), "*"),
	wire.Bind(new(repository.VlogCommandRepository), new(*VlogCommandRepositoryImpl)),
)

func (r *VlogCommandRepositoryImpl) Store(vlog *entity.Vlog) error {
	return errors.Wrap(r.DB.Save(vlog).Error, "failed to save vlog")
}

func (r *VlogCommandRepositoryImpl) DeleteByID(id int) error {
	e := &entity.Vlog{}
	e.ID = id
	return errors.Wrapf(r.DB.Delete(e).Error, "failed to delete vlog(id=%d)", id)
}
