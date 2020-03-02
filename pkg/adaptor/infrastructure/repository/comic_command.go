package repository

import (
	"github.com/google/wire"
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/entity"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/repository"
)

type ComicCommandRepositoryImpl struct {
	DB *gorm.DB
}

var ComicCommandRepositorySet = wire.NewSet(
	wire.Struct(new(ComicCommandRepositoryImpl), "*"),
	wire.Bind(new(repository.ComicCommandRepository), new(*ComicCommandRepositoryImpl)),
)

func (r *ComicCommandRepositoryImpl) Store(comic *entity.Comic) error {
	return errors.Wrap(r.DB.Save(comic).Error, "failed to save comic")
}

func (r *ComicCommandRepositoryImpl) DeleteByID(id int) error {
	return errors.Wrapf(r.DB.Delete(&entity.Comic{ID: id}).Error, "failed to delete comic(id=%d)", id)
}
