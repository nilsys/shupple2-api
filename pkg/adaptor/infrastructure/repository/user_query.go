package repository

import (
	"github.com/google/wire"
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
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
	if err := r.DB.First(&row, id).Error; err != nil {
		return nil, ErrorToFindSingleRecord(err, "user(id=%d)", id)
	}
	return &row, nil
}

func (r *UserQueryRepositoryImpl) FindByWordpressID(wordpressUserID int) (*entity.User, error) {
	var row entity.User
	if err := r.DB.Where("wordpress_id = ?", wordpressUserID).First(&row).Error; err != nil {
		return nil, ErrorToFindSingleRecord(err, "user(wordpress_id=%d)", wordpressUserID)
	}
	return &row, nil
}

// name部分一致検索
func (r *UserQueryRepositoryImpl) SearchByName(name string) ([]*entity.User, error) {
	var rows []*entity.User

	if err := r.DB.Where("MATCH(name) AGAINST(?)", name).Limit(10).Find(&rows).Error; err != nil {
		return nil, errors.Wrap(err, "failed to find user list by like name")
	}

	return rows, nil
}
