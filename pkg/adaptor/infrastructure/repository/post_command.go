package repository

import (
	"context"

	"github.com/google/wire"
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/entity"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/repository"
)

type PostCommandRepositoryImpl struct {
	DB *gorm.DB
}

var PostCommandRepositorySet = wire.NewSet(
	wire.Struct(new(PostCommandRepositoryImpl), "*"),
	wire.Bind(new(repository.PostCommandRepository), new(*PostCommandRepositoryImpl)),
)

func (r *PostCommandRepositoryImpl) Store(c context.Context, post *entity.Post) error {
	return errors.Wrap(r.DB.Save(post).Error, "failed to save post")
}

func (r *PostCommandRepositoryImpl) DeleteByID(id int) error {
	e := &entity.Post{}
	e.ID = id
	return errors.Wrapf(r.DB.Delete(e).Error, "failed to delete post(id=%d)", id)
}
