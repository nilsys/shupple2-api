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

func (r *ComicQueryRepositoryImpl) FindByID(id int) (*entity.ComicDetail, error) {
	var row entity.ComicDetail
	if err := r.DB.First(&row, id).Error; err != nil {
		return nil, ErrorToFindSingleRecord(err, "comic(id=%d)", id)
	}
	return &row, nil
}

func (r *ComicQueryRepositoryImpl) FindWithIsFavoriteByID(id, userID int) (*entity.ComicDetail, error) {
	var row entity.ComicDetail
	if err := r.DB.
		Select("comic.*, CASE WHEN user_favorite_comic.comic_id IS NULL THEN 'FALSE' ELSE 'TRUE' END is_favorite").
		Joins("LEFT JOIN user_favorite_comic ON comic.id = user_favorite_comic.comic_id AND user_favorite_comic.user_id = ?", userID).
		First(&row, id).Error; err != nil {
		return nil, ErrorToFindSingleRecord(err, "comic(id=%d)", id)
	}
	return &row, nil
}

// 作成日時降順に指定されたlimit, offsetで一覧取得
func (r *ComicQueryRepositoryImpl) FindListOrderByCreatedAt(query *query.FindListPaginationQuery) (*entity.ComicList, error) {
	var rows entity.ComicList

	// query.ExcludeIDのdefaultは0
	if err := r.DB.
		Order("created_at desc").Offset(query.Offset).Limit(query.Limit).Not("id", query.ExcludeID).Find(&rows.List).Offset(0).Count(&rows.TotalNumber).Error; err != nil {
		return nil, errors.Wrapf(err, "Failed find comics")
	}

	return &rows, nil
}

// 作成日時降順に指定されたlimit, offsetで一覧取得
func (r *ComicQueryRepositoryImpl) FindWithIsFavoriteListOrderByCreatedAt(query *query.FindListPaginationQuery, userID int) (*entity.ComicList, error) {
	var rows entity.ComicList

	// query.ExcludeIDのdefaultは0
	if err := r.DB.
		Select("comic.*, CASE WHEN user_favorite_comic.comic_id IS NULL THEN 'FALSE' ELSE 'TRUE' END is_favorite").
		Joins("LEFT JOIN user_favorite_comic ON comic.id = user_favorite_comic.comic_id AND user_favorite_comic.user_id = ?", userID).
		Order("created_at desc").Offset(query.Offset).Limit(query.Limit).Not("id", query.ExcludeID).Find(&rows.List).Offset(0).Count(&rows.TotalNumber).Error; err != nil {
		return nil, errors.Wrapf(err, "Failed find comics")
	}

	return &rows, nil
}

func (r *ComicQueryRepositoryImpl) IsExist(id int) (bool, error) {
	var row entity.Comic

	err := r.DB.First(&row, id).Error

	return ErrorToIsExist(err, "comic(id=%d)", id)
}
