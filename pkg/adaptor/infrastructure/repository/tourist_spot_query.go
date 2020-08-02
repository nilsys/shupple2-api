package repository

import (
	"github.com/google/wire"
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/entity"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/model/query"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/repository"
)

type TouristSpotQueryRepositoryImpl struct {
	DB *gorm.DB
}

var TouristSpotQueryRepositorySet = wire.NewSet(
	wire.Struct(new(TouristSpotQueryRepositoryImpl), "*"),
	wire.Bind(new(repository.TouristSpotQueryRepository), new(*TouristSpotQueryRepositoryImpl)),
)

func (r *TouristSpotQueryRepositoryImpl) FindAll() ([]*entity.TouristSpot, error) {
	var rows []*entity.TouristSpot

	if err := r.DB.Find(&rows).Error; err != nil {
		return nil, errors.Wrap(err, "failed to find all tourist_spot")
	}

	return rows, nil
}

func (r *TouristSpotQueryRepositoryImpl) FindByID(id int) (*entity.TouristSpot, error) {
	var row entity.TouristSpot

	if err := r.DB.First(&row, id).Error; err != nil {
		return nil, ErrorToFindSingleRecord(err, "touristSpot(id=%d)", id)
	}
	return &row, nil
}

func (r *TouristSpotQueryRepositoryImpl) FindDetailByID(id int) (*entity.TouristSpotDetail, error) {
	var row entity.TouristSpotDetail
	if err := r.DB.Select("*").Joins("INNER JOIN (SELECT count(id) AS review_count FROM review WHERE tourist_spot_id = ? AND deleted_at IS NULL) AS r", id).First(&row, id).Error; err != nil {
		return nil, ErrorToFindSingleRecord(err, "touristSpot detail(id=%d)", id)
	}
	return &row, nil
}

// https://github.com/stayway-corp/stayway-media-api/issues/56
func (r *TouristSpotQueryRepositoryImpl) FindListByParams(query *query.FindTouristSpotListQuery) (*entity.TouristSpotList, error) {
	var rows entity.TouristSpotList

	q := r.buildFindListByParamsQuery(query)

	if err := q.
		Select("*").
		Joins("LEFT JOIN (SELECT tourist_spot_id, count(id) AS review_count FROM review WHERE deleted_at IS NULL GROUP BY tourist_spot_id) rc ON tourist_spot.id = rc.tourist_spot_id").
		Limit(query.Limit).
		Offset(query.OffSet).
		Order("vendor_rate desc").
		Find(&rows.TouristSpots).Offset(0).Count(&rows.TotalNumber).Error; err != nil {
		return nil, errors.Wrapf(err, "Failed get tourist_spots by params")
	}

	return &rows, nil
}

// https://github.com/stayway-corp/stayway-media-api/issues/56
func (r *TouristSpotQueryRepositoryImpl) FindRecommendListByParams(query *query.FindRecommendTouristSpotListQuery) (*entity.TouristSpotList, error) {
	var rows entity.TouristSpotList

	q := r.buildFindRecommendListQuery(query)
	cq := r.buildCountRecommendListQuery(query)

	if err := q.
		Joins("LEFT JOIN (SELECT tourist_spot_id, count(id) AS review_count FROM review WHERE deleted_at IS NULL GROUP BY tourist_spot_id) rc ON tourist_spot.id = rc.tourist_spot_id").
		Limit(query.Limit).
		Offset(query.OffSet).
		Order("vendor_rate desc").
		Find(&rows.TouristSpots).Error; err != nil {
		return nil, errors.Wrap(err, "failed get recommend tourist_spots by params")
	}

	if err := cq.
		Count(&rows.TotalNumber).Error; err != nil {
		return nil, errors.Wrap(err, "failed get recommend tourist_spots.count by params")
	}

	return &rows, nil
}

// name部分一致検索
func (r *TouristSpotQueryRepositoryImpl) SearchByName(name string) ([]*entity.TouristSpotTiny, error) {
	var rows []*entity.TouristSpotTiny

	if err := r.DB.Where("MATCH(name) AGAINST(?)", name).Limit(defaultSearchSuggestionsNumber).Find(&rows).Error; err != nil {
		return nil, errors.Wrap(err, "failed to find tourist_spot list by like name")
	}

	return rows, nil
}

func (r *TouristSpotQueryRepositoryImpl) FindReviewCountByIDs(ids []int) (*entity.TouristSpotReviewCountList, error) {
	var rows entity.TouristSpotReviewCountList

	if err := r.DB.Table("tourist_spot").
		Select("tourist_spot.id AS tourist_spot_id, rc.review_count AS review_count").
		Joins("LEFT JOIN (SELECT tourist_spot_id, count(id) AS review_count FROM review WHERE deleted_at IS NULL GROUP BY tourist_spot_id) rc ON tourist_spot.id = rc.tourist_spot_id").
		Where("id IN (?)", ids).Find(&rows.List).Error; err != nil {
		return nil, errors.Wrap(err, "failed find tourist_spot review_count")
	}

	return &rows, nil
}

// TODO: 速度ヤバそうなんでstg環境で確認後チューニング
// golang側で緯度経度
func (r *TouristSpotQueryRepositoryImpl) buildFindRecommendListQuery(query *query.FindRecommendTouristSpotListQuery) *gorm.DB {
	q := r.DB

	if query.ID != 0 {
		q = q.Select("*, (6371 * acos(cos(radians((SELECT lat FROM tourist_spot WHERE id = ?)))* cos(radians(lat))* cos(radians(lng) - radians((SELECT lng FROM tourist_spot WHERE id = ?)))+ sin(radians((SELECT lat FROM tourist_spot WHERE id = ?)))* sin(radians(lat)))) AS distance", query.ID, query.ID, query.ID).Not("id = ?", query.ID).Having("distance <= ?", defaultRangeSearchKm).Order("distance")
	}
	if query.TouristSpotCategoryID != 0 {
		q = q.Where("id IN (SELECT tourist_spot_id FROM tourist_spot_spot_category WHERE spot_category_id = ?)", query.TouristSpotCategoryID)
	}

	return q
}

func (r *TouristSpotQueryRepositoryImpl) buildCountRecommendListQuery(query *query.FindRecommendTouristSpotListQuery) *gorm.DB {
	q := r.DB

	q = q.Table("tourist_spot").
		Joins("LEFT JOIN (SELECT tourist_spot_id, count(id) AS review_count FROM review GROUP BY tourist_spot_id) rc ON tourist_spot.id = rc.tourist_spot_id")

	if query.ID != 0 {
		q = q.Not("id = ?", query.ID).Where("6371 * acos(cos(radians((SELECT lat FROM tourist_spot WHERE id = ?)))* cos(radians(lat))* cos(radians(lng) - radians((SELECT lng FROM tourist_spot WHERE id = ?)))+ sin(radians((SELECT lat FROM tourist_spot WHERE id = ?)))* sin(radians(lat))) <= ?", query.ID, query.ID, query.ID, defaultRangeSearchKm)
	}
	if query.TouristSpotCategoryID != 0 {
		q = q.Where("id IN (SELECT tourist_spot_id FROM tourist_spot_spot_category WHERE spot_category_id = ?)", query.TouristSpotCategoryID)
	}

	return q
}

func (r *TouristSpotQueryRepositoryImpl) buildFindListByParamsQuery(query *query.FindTouristSpotListQuery) *gorm.DB {
	q := r.DB

	if query.AreaID != 0 {
		q = q.Where("id IN (SELECT tourist_spot_id FROM tourist_spot_area_category WHERE area_category_id IN (SELECT id FROM area_category WHERE area_id = ?))", query.AreaID)
	}
	if query.SubAreaID != 0 {
		q = q.Where("id IN (SELECT tourist_spot_id FROM tourist_spot_area_category WHERE area_category_id IN (SELECT id FROM area_category WHERE sub_area_id = ?))", query.SubAreaID)
	}
	if query.SubSubAreaID != 0 {
		q = q.Where("id IN (SELECT tourist_spot_id FROM tourist_spot_area_category WHERE area_category_id IN (SELECT id FROM area_category WHERE sub_sub_area_id = ?))", query.SubSubAreaID)
	}
	if query.SpotCategoryID != 0 {
		q = q.Where("id IN (SELECT tourist_spot_id FROM tourist_spot_spot_category WHERE spot_category_id = ?)", query.SpotCategoryID)
	}
	if len(query.ExcludeSpotIDs) > 0 {
		q = q.Not("id IN (SELECT tourist_spot_id FROM tourist_spot_area_category WHERE area_category_id IN (?))", query.ExcludeSpotIDs)
	}

	return q
}
