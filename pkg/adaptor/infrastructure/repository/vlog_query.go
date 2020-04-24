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

func (r *VlogQueryRepositoryImpl) FindListByParams(query *query.FindVlogListQuery) (*entity.VlogList, error) {
	var rows entity.VlogList

	q := r.buildFindByParamsQuery(query)

	// フリーワード検索の場合
	if query.Keyward != "" {
		if err := q.
			Select("*, CASE WHEN MATCH(title) AGAINST(?) THEN 'TRUE' ELSE 'FALSE' END is_matched_title", query.Keyward).
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

// クエリ構造体を用い、FindListByParams()で使用するsqlクエリを作成
func (r *VlogQueryRepositoryImpl) buildFindByParamsQuery(query *query.FindVlogListQuery) *gorm.DB {
	q := r.DB

	if query.AreaID != 0 {
		q = q.Where("id IN (SELECT vlog_id FROM vlog_area_category WHERE area_category_id IN (SELECT id FROM area_category WHERE area_id = ?))", query.AreaID)
	}

	if query.SubAreaID != 0 {
		q = q.Where("id IN (SELECT vlog_id FROM vlog_area_category WHERE area_category_id IN (SELECT id FROM area_category WHERE sub_area_id = ?))", query.SubAreaID)
	}

	if query.SubSubAreaID != 0 {
		q = q.Where("id IN (SELECT vlog_id FROM vlog_area_category WHERE area_category_id IN (SELECT id FROM area_category WHERE sub_sub_area_id = ?))", query.SubSubAreaID)
	}

	if query.TouristSpotID != 0 {
		q = q.Where("id IN (SELECT vlog_id FROM vlog_tourist_spot WHERE tourist_spot_id = ?)", query.TouristSpotID)
	}

	// TODO: titleに引っかかる物が優先順位が高い、その後body
	if query.Keyward != "" {
		q = q.Where("MATCH(title) AGAINST(?)", query.Keyward).Or("MATCH(body) AGAINST(?)", query.Keyward)
	}

	return q
}
