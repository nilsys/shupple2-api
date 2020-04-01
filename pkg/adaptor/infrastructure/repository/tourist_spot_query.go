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

func (r *TouristSpotQueryRepositoryImpl) FindDetailByID(id int) (*entity.QueryTouristSpot, error) {
	var row entity.QueryTouristSpot
	if err := r.DB.First(&row, id).Error; err != nil {
		return nil, ErrorToFindSingleRecord(err, "touristSpot detail(id=%d)", id)
	}
	return &row, nil
}

// TODO: rateに関して
// https://github.com/stayway-corp/stayway-media-api/issues/56
func (r *TouristSpotQueryRepositoryImpl) FindListByParams(query *query.FindTouristSpotListQuery) ([]*entity.TouristSpot, error) {
	var rows []*entity.TouristSpot

	q := r.buildFindListByParamsQuery(query)

	if err := q.
		Limit(query.Limit).
		Offset(query.OffSet).
		Order("vendor_rate desc").
		Find(&rows).Error; err != nil {
		return nil, errors.Wrapf(err, "Failed get tourist_spots by params")
	}

	return rows, nil
}

// TODO: rateに関して
// https://github.com/stayway-corp/stayway-media-api/issues/56
func (r *TouristSpotQueryRepositoryImpl) FindRecommendListByParams(query *query.FindRecommendTouristSpotListQuery) ([]*entity.TouristSpot, error) {
	var rows []*entity.TouristSpot

	q := r.buildFindRecommendListQuery(query)

	if err := q.
		Limit(query.Limit).
		Offset(query.OffSet).
		Order("vendor_rate desc").
		Find(&rows).Error; err != nil {
		return nil, errors.Wrap(err, "failed get recommend tourist_spots by params")
	}
	return rows, nil
}

// name部分一致検索
func (r *TouristSpotQueryRepositoryImpl) SearchByName(name string) ([]*entity.TouristSpot, error) {
	var rows []*entity.TouristSpot

	if err := r.DB.Where("MATCH(name) AGAINST(?)", name).Limit(defaultSearchSuggestionsNumber).Find(&rows).Error; err != nil {
		return nil, errors.Wrap(err, "failed to find tourist_spot list by like name")
	}

	return rows, nil
}

// TODO: 速度ヤバそうなんでstg環境で確認後チューニング
func (r *TouristSpotQueryRepositoryImpl) buildFindRecommendListQuery(query *query.FindRecommendTouristSpotListQuery) *gorm.DB {
	q := r.DB

	if query.ID != 0 {
		q = q.Select("*, (6371 * acos(cos(radians((SELECT lat FROM tourist_spot WHERE id = ?)))* cos(radians(lat))* cos(radians(lng) - radians((SELECT lng FROM tourist_spot WHERE id = ?)))+ sin(radians((SELECT lat FROM tourist_spot WHERE id = ?)))* sin(radians(lat)))) AS distance", query.ID, query.ID, query.ID).Having("distance <= ?", defaultRangeSearchKm).Order("distance")
	}
	if query.TouristSpotCategoryID != 0 {
		q = q.Where("id IN (SELECT tourist_spot_id FROM tourist_spot_lcategory WHERE lcategory_id = ?)", query.TouristSpotCategoryID)
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
	if query.LcategoryID != 0 {
		q = q.Where("id IN (SELECT tourist_spot_id FROM tourist_spot_lcategory WHERE lcategory_id = ?)", query.LcategoryID)
	}
	if len(query.ExcludeSpotIDs) > 0 {
		q = q.Not("id IN (SELECT tourist_spot_id FROM tourist_spot_area_category WHERE area_category_id IN (?))", query.ExcludeSpotIDs)
	}

	return q
}
