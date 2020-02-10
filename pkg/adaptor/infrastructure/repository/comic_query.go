package repository

import (
	"github.com/google/wire"
	"github.com/jinzhu/gorm"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/entity"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/repository"
)

type ComicQueryRepositoryImpl struct {
	DB *gorm.DB
}

var ComicQueryRepositorySet = wire.NewSet(
	wire.Struct(new(ComicQueryRepositoryImpl), "*"),
	wire.Bind(new(repository.ComicQueryRepository), new(*ComicQueryRepositoryImpl)),
)

func (r *ComicQueryRepositoryImpl) FindByID(id int) (*entity.Comic, error) {
	var row entity.Comic
	if err := r.DB.First(&row, id).Error; err != nil {
		return nil, ErrorToFindSingleRecord(err, "comic(id=%d)", id)
	}
	return &row, nil
}
