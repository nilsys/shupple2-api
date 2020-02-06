package repository

import (
	"github.com/google/wire"
	"github.com/jinzhu/gorm"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/entity"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/repository"
)

// User参照系レポジトリ実装
type UserQueryRepositoryImpl struct {
	DB *gorm.DB
}

var UserQueryRepositorySet = wire.NewSet(
	wire.Struct(new(UserQueryRepositoryImpl), "*"),
	wire.Bind(new(repository.UserQueryRepository), new(*UserQueryRepositoryImpl)),
)

func (r *UserQueryRepositoryImpl) FindByID(id int) (*entity.User, error) {
	var row entity.User
	if err := r.DB.First(&row).Error; err != nil {
		return nil, ErrorToFindSingleRecord(err, "user(id=%d)", id)
	}
	return &row, nil
}
