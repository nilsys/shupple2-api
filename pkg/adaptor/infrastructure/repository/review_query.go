package repository

import (
	"github.com/google/wire"
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/entity"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/query"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/repository"
)

// Review参照系レポジトリ実装
type ReviewQueryRepositoryImpl struct {
	DB *gorm.DB
}

var ReviewQueryRepositorySet = wire.NewSet(
	wire.Struct(new(ReviewQueryRepositoryImpl), "*"),
	wire.Bind(new(repository.ReviewQueryRepository), new(*ReviewQueryRepositoryImpl)),
)

// パスパラメータで飛んで来た検索条件を用いreviewを検索
func (r *ReviewQueryRepositoryImpl) ShowReviewListByParams(query *query.ShowReviewListQuery) ([]*entity.Review, error) {
	var reviews []*entity.Review

	q := r.buildQuery(query)

	if err := q.
		Preload("Medias").
		Limit(query.Limit).
		Offset(query.OffSet).
		Order(query.SortBy.GetReviewOrderQuery()).
		Find(&reviews).Error; err != nil {
		return nil, errors.Wrapf(err, "Failed get reviews by params")
	}

	return reviews, nil
}

// パスパラメータで飛んで来た値によって検索クエリを切り替える
// MEMO: presentation、application層などでバリデーションがparamsにバリデーションが掛かっている事が前提
func (r *ReviewQueryRepositoryImpl) buildQuery(query *query.ShowReviewListQuery) *gorm.DB {
	q := r.DB

	if query.UserID != 0 {
		q = q.Where("user_id = ?", query.UserID)
	}

	if query.InnID != 0 {
		q = q.Where("inn_id = ?", query.InnID)
	}

	if query.TouristSpotID != 0 {
		q = q.Where("tourist_spot_id = ?", query.TouristSpotID)
	}

	if query.HashTag != "" {
		q = q.Where("id IN (SELECT review_id FROM review_hashtag WHERE hashtag_id = (SELECT id FROM hashtag WHERE name = ?))", query.HashTag)
	}

	if query.AreaID != 0 {
		q = q.Where("tourist_spot_id IN (SELECT tourist_spot_id FROM tourist_spot_category WHERE category_id = ?)", query.AreaID)
	}

	if query.SubAreaID != 0 {
		q = q.Where("tourist_spot_id IN (SELECT tourist_spot_id FROM tourist_spot_category WHERE category_id = ?)", query.SubAreaID)
	}

	if query.SubSubAreaID != 0 {
		q = q.Where("tourist_spot_id IN (SELECT tourist_spot_id FROM tourist_spot_category WHERE category_id = ?)", query.SubSubAreaID)
	}

	if len(query.InnIDs) > 0 {
		q = q.Where("inn_id IN (?)", query.InnIDs)
	}

	return q
}
