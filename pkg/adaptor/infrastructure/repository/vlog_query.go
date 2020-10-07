package repository

import (
	"github.com/google/wire"
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/entity"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/model/query"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/repository"
)

type VlogQueryRepositoryImpl struct {
	DB *gorm.DB
}

var VlogQueryRepositorySet = wire.NewSet(
	wire.Struct(new(VlogQueryRepositoryImpl), "*"),
	wire.Bind(new(repository.VlogQueryRepository), new(*VlogQueryRepositoryImpl)),
)

func (r *VlogQueryRepositoryImpl) FindByID(id int) (*entity.Vlog, error) {
	var row entity.Vlog
	if err := r.DB.First(&row, id).Error; err != nil {
		return nil, ErrorToFindSingleRecord(err, "vlog(id=%d)", id)
	}
	return &row, nil
}

func (r *VlogQueryRepositoryImpl) FindDetailByID(id int) (*entity.VlogDetail, error) {
	var row entity.VlogDetail

	if err := r.DB.First(&row, id).Error; err != nil {
		return nil, ErrorToFindSingleRecord(err, "vlog(id=%d)", id)
	}
	return &row, nil
}

func (r *VlogQueryRepositoryImpl) FindByLastID(lastID, limit int) (entity.VlogTinyList, error) {
	var rows []*entity.VlogTiny

	if err := r.DB.Where("id > ?", lastID).Limit(limit).Find(&rows).Error; err != nil {
		return nil, errors.Wrap(err, "failed find vlog")
	}

	return rows, nil
}

func (r *VlogQueryRepositoryImpl) FindDetailWithIsFavoriteByID(id, userID int) (*entity.VlogDetail, error) {
	var row entity.VlogDetail

	if err := r.DB.
		Select("*, CASE WHEN user_favorite_vlog.vlog_id IS NULL THEN 'FALSE' ELSE 'TRUE' END is_favorite").
		Joins("LEFT JOIN user_favorite_vlog ON vlog.id = user_favorite_vlog.vlog_id AND user_favorite_vlog.user_id = ?", userID).
		First(&row, id).Error; err != nil {
		return nil, ErrorToFindSingleRecord(err, "vlog(id=%d)", id)
	}
	return &row, nil
}

func (r *VlogQueryRepositoryImpl) FindListByParams(query *query.FindVlogListQuery) (*entity.VlogList, error) {
	var rows entity.VlogList

	q := r.buildFindByParamsQuery(query)

	// フリーワード検索の場合
	if query.Keyword != "" {
		if err := q.
			Select("*, CASE WHEN vlog.title LIKE ? THEN 'TRUE' ELSE 'FALSE' END is_matched_title", query.SQLLikeKeyword()).
			Order("is_matched_title desc").
			Order(query.SortBy.GetVlogOrderQuery()).
			Limit(query.Limit).
			Offset(query.OffSet).
			Find(&rows.Vlogs).Offset(0).Count(&rows.TotalNumber).Error; err != nil {
			return nil, errors.Wrap(err, "failed find vlogs by params")
		}

		return &rows, nil
	}

	if err := q.
		Order(query.SortBy.GetVlogOrderQuery()).
		Limit(query.Limit).
		Offset(query.OffSet).
		Find(&rows.Vlogs).Offset(0).Count(&rows.TotalNumber).Error; err != nil {
		return nil, errors.Wrap(err, "failed find vlogs by params")
	}

	return &rows, nil
}

func (r *VlogQueryRepositoryImpl) FindWithIsFavoriteListByParams(query *query.FindVlogListQuery, userID int) (*entity.VlogList, error) {
	var rows entity.VlogList

	q := r.buildFindByParamsQuery(query)

	// フリーワード検索の場合
	if query.Keyword != "" {
		if err := q.
			Select("*, CASE WHEN vlog.title LIKE ? THEN 'TRUE' ELSE 'FALSE' END is_matched_title, CASE WHEN user_favorite_vlog.vlog_id IS NULL THEN 'FALSE' ELSE 'TRUE' END is_favorite", query.SQLLikeKeyword()).
			Joins("LEFT JOIN user_favorite_vlog ON vlog.id = user_favorite_vlog.vlog_id AND user_favorite_vlog.user_id = ?", userID).
			Order("is_matched_title desc").
			Order(query.SortBy.GetVlogOrderQuery()).
			Limit(query.Limit).
			Offset(query.OffSet).
			Find(&rows.Vlogs).Offset(0).Count(&rows.TotalNumber).Error; err != nil {
			return nil, errors.Wrap(err, "failed find vlogs by params")
		}

		return &rows, nil
	}

	if err := q.
		Select("*, CASE WHEN title LIKE ? THEN 'TRUE' ELSE 'FALSE' END is_matched_title, CASE WHEN user_favorite_vlog.vlog_id IS NULL THEN 'FALSE' ELSE 'TRUE' END is_favorite", query.SQLLikeKeyword()).
		Joins("LEFT JOIN user_favorite_vlog ON vlog.id = user_favorite_vlog.vlog_id AND user_favorite_vlog.user_id = ?", userID).
		Order(query.SortBy.GetVlogOrderQuery()).
		Limit(query.Limit).
		Offset(query.OffSet).
		Find(&rows.Vlogs).Offset(0).Count(&rows.TotalNumber).Error; err != nil {
		return nil, errors.Wrap(err, "failed find vlogs by params")
	}

	return &rows, nil
}

func (r *VlogQueryRepositoryImpl) FindAll() ([]*entity.Vlog, error) {
	var rows []*entity.Vlog

	if err := r.DB.Find(&rows).Error; err != nil {
		return nil, errors.Wrap(err, "failed to find all vlog")
	}

	return rows, nil
}

// クエリ構造体を用い、FindListByParams()で使用するsqlクエリを作成
func (r *VlogQueryRepositoryImpl) buildFindByParamsQuery(query *query.FindVlogListQuery) *gorm.DB {
	q := r.DB

	if query.AreaID != 0 {
		q = q.Where("vlog.id IN (SELECT vlog_id FROM vlog_area_category WHERE area_category_id IN (SELECT id FROM area_category WHERE area_id = ?))", query.AreaID)
	}

	if query.SubAreaID != 0 {
		q = q.Where("vlog.id IN (SELECT vlog_id FROM vlog_area_category WHERE area_category_id IN (SELECT id FROM area_category WHERE sub_area_id = ?))", query.SubAreaID)
	}

	if query.SubSubAreaID != 0 {
		q = q.Where("vlog.id IN (SELECT vlog_id FROM vlog_area_category WHERE area_category_id IN (SELECT id FROM area_category WHERE sub_sub_area_id = ?))", query.SubSubAreaID)
	}

	if query.TouristSpotID != 0 {
		q = q.Where("vlog.id IN (SELECT vlog_id FROM vlog_tourist_spot WHERE tourist_spot_id = ?)", query.TouristSpotID)
	}

	if query.UserID != 0 {
		q = q.Where("vlog.user_id = ?", query.UserID)
	}

	if query.Keyword != "" {
		q = q.Where("vlog.title LIKE ?", query.SQLLikeKeyword()).Or("vlog.body LIKE ?", query.SQLLikeKeyword())
	}

	return q
}

func (r *VlogQueryRepositoryImpl) IsExist(id int) (bool, error) {
	var row entity.Vlog

	err := r.DB.First(&row, id).Error

	return ErrorToIsExist(err, "vlog(id=%d)", id)
}
