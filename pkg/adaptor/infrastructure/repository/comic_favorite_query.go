package repository

import (
	"github.com/google/wire"
	"github.com/jinzhu/gorm"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/entity"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/repository"
)

type ComicFavoriteQueryRepositoryImpl struct {
	DB *gorm.DB
}

var ComicFavoriteQueryRepositorySet = wire.NewSet(
	wire.Struct(new(ComicFavoriteQueryRepositoryImpl), "*"),
	wire.Bind(new(repository.ComicFavoriteQueryRepository), new(*ComicFavoriteQueryRepositoryImpl)),
)

func (r *ComicFavoriteQueryRepositoryImpl) IsExist(userID, comicID int) (bool, error) {
	var row entity.UserFavoriteComic

	err := r.DB.Where("user_id = ? AND comic_id = ?", userID, comicID).First(&row).Error

	return ErrorToIsExist(err, "user_favorite_comic(user_id=%d,comic_id=%d)", userID, comicID)
}
