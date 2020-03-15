package repository

import (
	"fmt"
	"time"

	"github.com/google/wire"
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/entity"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/model"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/model/query"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/repository"
)

// Review参照系レポジトリ実装
type ReviewQueryRepositoryImpl struct {
	DB *gorm.DB
}

const recommendMonthPeriod = 1

var ReviewQueryRepositorySet = wire.NewSet(
	wire.Struct(new(ReviewQueryRepositoryImpl), "*"),
	wire.Bind(new(repository.ReviewQueryRepository), new(*ReviewQueryRepositoryImpl)),
)

// パスパラメータで飛んで来た検索条件を用いreviewを検索
// TODO: comment_countは新しくreviewテーブルにカラム 増やすのでそこから取れる
func (r *ReviewQueryRepositoryImpl) ShowReviewListByParams(query *query.ShowReviewListQuery) ([]*entity.QueryReview, error) {
	var reviews []*entity.QueryReview

	q := r.buildShowReviewListQuery(query)

	if err := q.
		Select("review.*, count(review_comment.id) AS comment_count").
		Joins("LEFT JOIN review_comment ON review.id = review_comment.review_id").
		Preload("Medias").
		Limit(query.Limit).
		Offset(query.OffSet).
		Order(query.SortBy.GetReviewOrderQuery()).
		Group("review.id").
		Find(&reviews).Error; err != nil {
		return nil, errors.Wrapf(err, "Failed get reviews by params")
	}

	return reviews, nil
}

// ユーザーIDからフォローしているハッシュタグ or ユーザーのreview一覧を参照
// TODO: comment_countは新しくreviewテーブルにカラム 増やすのでそこから取れる
func (r *ReviewQueryRepositoryImpl) FindFeedReviewListByUserID(userID int, query *query.FindListPaginationQuery) ([]*entity.QueryReview, error) {
	var reviews []*entity.QueryReview

	q := r.buildFindFeedListQuery(userID)

	if err := q.
		Select("review.*, count(review_comment.id) AS comment_count").
		Joins("LEFT JOIN review_comment ON review.id = review_comment.review_id").
		Limit(query.Limit).
		Offset(query.Offset).
		Order("updated_at DESC").
		Group("review.id").
		Find(&reviews).Error; err != nil {
		return nil, errors.Wrap(err, "failed find feed reviews")
	}

	return reviews, nil
}

// TODO: comment_countは新しくreviewテーブルにカラム 増やすのでそこから取れる
func (r *ReviewQueryRepositoryImpl) FindQueryReviewByID(id int) (*entity.QueryReview, error) {
	var row entity.QueryReview

	q := r.DB

	// review AS review
	// review_comment AS rc
	if err := q.
		Select("review.*, count(rc.id) AS comment_count").
		Joins("INNER JOIN review_comment AS rc ON review.id = rc.review_id").
		First(&row, id).
		Error; err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("failed find review by id=%d", id))
	}

	return &row, nil
}

// パスパラメータで飛んで来た値によって検索クエリを切り替える
func (r *ReviewQueryRepositoryImpl) buildShowReviewListQuery(query *query.ShowReviewListQuery) *gorm.DB {
	q := r.DB

	if query.UserID != 0 {
		q = q.Where("review.user_id = ?", query.UserID)
	}

	if query.InnID != 0 {
		q = q.Where("review.inn_id = ?", query.InnID)
	}

	if query.TouristSpotID != 0 {
		q = q.Where("review.tourist_spot_id = ?", query.TouristSpotID)
	}

	if query.HashTag != "" {
		q = q.Where("review.id IN (SELECT review_id FROM review_hashtag WHERE hashtag_id = (SELECT id FROM hashtag WHERE name = ?))", query.HashTag)
	}

	if query.AreaID != 0 {
		q = q.Where("review.tourist_spot_id IN (SELECT tourist_spot_id FROM tourist_spot_category WHERE category_id = ?)", query.AreaID)
	}

	if query.SubAreaID != 0 {
		q = q.Where("review.tourist_spot_id IN (SELECT tourist_spot_id FROM tourist_spot_category WHERE category_id = ?)", query.SubAreaID)
	}

	if query.SubSubAreaID != 0 {
		q = q.Where("review.tourist_spot_id IN (SELECT tourist_spot_id FROM tourist_spot_category WHERE category_id = ?)", query.SubSubAreaID)
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

	if query.ChildID != 0 {
		q = q.Where("tourist_spot_id IN (SELECT tourist_spot_id FROM tourist_spot_category WHERE category_id IN (SELECT parent_id FROM category WHERE id = ?))", query.ChildID)
	}

	if len(query.InnIDs) > 0 {
		q = q.Where("review.inn_id IN (?)", query.InnIDs)
	}

	if query.SortBy == model.ReviewSortByRECOMMEND {
		q = q.Where("review.updated_at BETWEEN ? AND ?", time.Date(time.Now().Year(), time.Now().Month()-recommendMonthPeriod, time.Now().Day(), 0, 0, 0, 0, time.Local), time.Now())
	}

	if query.Keyward != "" {
		q = q.Where("MATCH(body) AGAINST(?)", query.Keyward)
	}

	return q
}

func (r *ReviewQueryRepositoryImpl) buildFindFeedListQuery(userID int) *gorm.DB {
	q := r.DB

	if userID != 0 {
		q = q.Where("review.user_id IN (SELECT target_id FROM user_follow WHERE user_id = ?)", userID).Or("review.id IN (SELECT review_id FROM review_hashtag WHERE hashtag_id IN (SELECT hashtag_id FROM user_follow_hashtag WHERE user_id = ?))", userID)
	}

	return q
}
