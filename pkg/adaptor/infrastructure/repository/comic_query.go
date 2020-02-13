package repository

import (
	"github.com/google/wire"
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/entity"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/model/query"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/repository"
)

type ComicQueryRepositoryImpl struct {
	DB *gorm.DB
}

var ComicQueryRepositorySet = wire.NewSet(
	wire.Struct(new(ComicQueryRepositoryImpl), "*"),
	wire.Bind(new(repository.ComicQueryRepository), new(*ComicQueryRepositoryImpl)),
)

func (r *ComicQueryRepositoryImpl) FindByID(id int) (*entity.QueryComic, error) {
	var row entity.QueryComic
	if err := r.DB.Table("comic").First(&row, id).Error; err != nil {
		return nil, ErrorToFindSingleRecord(err, "comic(id=%d)", id)
	}
	return &row, nil
}

// 作成日時降順に指定されたlimit, offsetで一覧取得
func (r *ComicQueryRepositoryImpl) FindListOrderByCreatedAt(query *query.FindListPaginationQuery) ([]*entity.Comic, error) {
	var rows []*entity.Comic

	if err := r.DB.Order("created_at desc").Limit(query.Limit).Offset(query.Offset).Find(&rows).Error; err != nil {
		return nil, errors.Wrapf(err, "Failed find comics")
	}

	return rows, nil
}
