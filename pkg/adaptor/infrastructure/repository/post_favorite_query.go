package repository

import (
	"github.com/google/wire"
	"github.com/jinzhu/gorm"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/entity"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/repository"
)

type PostFavoriteQueryRepositoryImpl struct {
	DB *gorm.DB
}

var PostFavoriteQueryRepositorySet = wire.NewSet(
	wire.Struct(new(PostFavoriteQueryRepositoryImpl), "*"),
	wire.Bind(new(repository.PostFavoriteQueryRepository), new(*PostFavoriteQueryRepositoryImpl)),
)

func (r *PostFavoriteQueryRepositoryImpl) IsExist(userID, postID int) (bool, error) {
	var row entity.UserFavoritePost

	err := r.DB.Where("user_id = ? AND post_id = ?", userID, postID).First(&row).Error

	return ErrorToIsExist(err, "user_favorite_post(user_id=%d,post_id=%d)", userID, postID)
}
