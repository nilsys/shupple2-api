package repository

import (
	"github.com/google/wire"
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
	"github.com/uma-co82/shupple2-api/pkg/domain/entity"
	"github.com/uma-co82/shupple2-api/pkg/domain/repository"
)

type UserQueryRepositoryImpl struct {
	DB *gorm.DB
}

var UserQueryRepositorySet = wire.NewSet(
	wire.Struct(new(UserQueryRepositoryImpl), "*"),
	wire.Bind(new(repository.UserQueryRepository), new(*UserQueryRepositoryImpl)),
)

func (r *UserQueryRepositoryImpl) FindByFirebaseID(id string) (*entity.UserTiny, error) {
	var row entity.UserTiny
	if err := r.DB.Where("firebase_id = ?", id).First(&row).Error; err != nil {
		return nil, errors.Wrap(err, "failed find user")
	}
	return &row, nil
}

func (r *UserQueryRepositoryImpl) FindByID(id int) (*entity.UserTiny, error) {
	var row entity.UserTiny
	if err := r.DB.First(&row, id).Error; err != nil {
		return nil, errors.Wrap(err, "failed find user")
	}
	return &row, nil
}
