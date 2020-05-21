package repository

import (
	"github.com/google/wire"
	"github.com/jinzhu/gorm"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/entity"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/repository"
)

type VlogFavoriteQueryRepositoryImpl struct {
	DB *gorm.DB
}

var VlogFavoriteQueryRepositorySet = wire.NewSet(
	wire.Struct(new(VlogFavoriteQueryRepositoryImpl), "*"),
	wire.Bind(new(repository.VlogFavoriteQueryRepository), new(*VlogFavoriteQueryRepositoryImpl)),
)

func (r *VlogFavoriteQueryRepositoryImpl) IsExist(userID, vlogID int) (bool, error) {
	var row entity.UserFavoriteVlog

	err := r.DB.Where("user_id = ? AND vlog_id = ?", userID, vlogID).First(&row).Error

	return ErrorToIsExist(err, "user_favorite_vlog(user_id=%d,vlog_id=%d)", userID, vlogID)
}
