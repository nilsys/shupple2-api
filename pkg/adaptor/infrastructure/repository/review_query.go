package repository

import (
	"github.com/google/wire"
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/entity"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/model/query"
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

	q := r.buildShowReviewListQuery(query)

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

// ユーザーIDからフォローしているハッシュタグ or ユーザーのreview一覧を参照
func (r *ReviewQueryRepositoryImpl) FindFeedReviewListByUserID(userID int, query *query.FindListPaginationQuery) ([]*entity.Review, error) {
	var reviews []*entity.Review

	q := r.buildFindFeedListQuery(userID)

	if err := q.
		Order("updated_at desc").
		Limit(query.Limit).
		Offset(query.Offset).
		Find(&reviews).Error; err != nil {
		return nil, errors.Wrap(err, "failed find feed reviews")
	}

	return reviews, nil
}

// パスパラメータで飛んで来た値によって検索クエリを切り替える
// MEMO: presentation、application層などでバリデーションがparamsにバリデーションが掛かっている事が前提
func (r *ReviewQueryRepositoryImpl) buildShowReviewListQuery(query *query.ShowReviewListQuery) *gorm.DB {
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

	if query.MetasearchAreaID != 0 {
		q = q.Where("tourist_spot_id IN (SELECT tourist_spot_id FROM tourist_spot_category WHERE category_id IN (SELECT id FROM category WHERE metasearch_area_id = ?))", query.MetasearchAreaID)
	}

	if query.MetasearchSubAreaID != 0 {
		q = q.Where("tourist_spot_id IN (SELECT tourist_spot_id FROM tourist_spot_category WHERE category_id IN (SELECT id FROM category WHERE metasearch_sub_area_id = ?))", query.MetasearchSubAreaID)
	}

	if query.MetasearchSubSubAreaID != 0 {
		q = q.Where("tourist_spot_id IN (SELECT tourist_spot_id FROM tourist_spot_category WHERE category_id IN (SELECT id FROM category WHERE metasearch_sub_sub_area_id = ?))", query.MetasearchSubSubAreaID)
	}

	if len(query.InnIDs) > 0 {
		q = q.Where("inn_id IN (?)", query.InnIDs)
	}

	if query.Keyward != "" {
		q = q.Where("MATCH(body) AGAINST(?)", query.Keyward)
	}

	return q
}

func (r *ReviewQueryRepositoryImpl) buildFindFeedListQuery(userID int) *gorm.DB {
	q := r.DB

	if userID != 0 {
		q = q.Where("user_id IN (SELECT target_id FROM user_follow WHERE user_id = ?)", userID).Or("id IN (SELECT review_id FROM review_hashtag WHERE hashtag_id IN (SELECT hashtag_id FROM user_follow_hashtag WHERE user_id = ?))", userID)
	}

	return q
}
