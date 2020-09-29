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
func (r *ReviewQueryRepositoryImpl) ShowReviewListByParams(query *query.ShowReviewListQuery) (*entity.ReviewDetailWithIsFavoriteList, error) {
	var rows entity.ReviewDetailWithIsFavoriteList

	q := r.buildShowReviewListQuery(query)

	if err := q.
		Limit(query.Limit).
		Offset(query.OffSet).
		Order(query.SortBy.GetReviewOrderQuery()).
		Find(&rows.List).Offset(0).Count(&rows.TotalNumber).Error; err != nil {
		return nil, errors.Wrapf(err, "Failed get reviews by params")
	}

	return &rows, nil
}

// パスパラメータで飛んで来た検索条件を用いreviewを検索
// UserIDで指定されたUserがお気に入りしているかどうかのフラグ
func (r *ReviewQueryRepositoryImpl) ShowReviewWithIsFavoriteListByParams(query *query.ShowReviewListQuery, userID int) (*entity.ReviewDetailWithIsFavoriteList, error) {
	var rows entity.ReviewDetailWithIsFavoriteList

	q := r.buildShowReviewListQuery(query)

	if err := q.
		Select("review.*, CASE WHEN user_favorite_review.review_id IS NULL THEN 'FALSE' ELSE 'TRUE' END is_favorite").
		Limit(query.Limit).
		Offset(query.OffSet).
		Order(query.SortBy.GetReviewOrderQueryForJoin()).
		Joins("LEFT JOIN user_favorite_review ON review.id = user_favorite_review.review_id AND user_favorite_review.user_id = ?", userID).
		Find(&rows.List).Offset(0).Count(&rows.TotalNumber).Error; err != nil {
		return nil, errors.Wrapf(err, "Failed get reviews by params")
	}

	return &rows, nil
}

func (r *ReviewQueryRepositoryImpl) FindFeedReviewWithIsFavoriteListByUserID(userID int, query *query.FindListPaginationQuery) (*entity.ReviewDetailWithIsFavoriteList, error) {
	var rows entity.ReviewDetailWithIsFavoriteList

	q := r.buildFindFeedListQuery(userID)

	if err := q.
		Select("review.*, CASE WHEN user_favorite_review.review_id IS NULL THEN 'FALSE' ELSE 'TRUE' END is_favorite").
		Limit(query.Limit).
		Offset(query.Offset).
		Order("created_at DESC").
		Joins("LEFT JOIN user_favorite_review ON review.id = user_favorite_review.review_id AND user_favorite_review.user_id = ?", userID).
		Find(&rows.List).Offset(0).Count(&rows.TotalNumber).Error; err != nil {
		return nil, errors.Wrap(err, "failed find feed reviews")
	}

	return &rows, nil
}

func (r *ReviewQueryRepositoryImpl) FindQueryReviewByID(id int) (*entity.ReviewDetailWithIsFavorite, error) {
	var row entity.ReviewDetailWithIsFavorite

	q := r.DB

	if err := q.
		First(&row, id).
		Error; err != nil {
		return nil, ErrorToFindSingleRecord(err, "review(id=%d)", id)
	}

	return &row, nil
}

func (r *ReviewQueryRepositoryImpl) FindQueryReviewWithIsFavoriteByID(id, userID int) (*entity.ReviewDetailWithIsFavorite, error) {
	var row entity.ReviewDetailWithIsFavorite

	q := r.DB

	if err := q.
		Select("review.*, CASE WHEN user_favorite_review.review_id IS NULL THEN 'FALSE' ELSE 'TRUE' END is_favorite").
		Joins("LEFT JOIN user_favorite_review ON review.id = user_favorite_review.review_id AND user_favorite_review.user_id = ?", userID).
		First(&row, id).
		Error; err != nil {
		return nil, ErrorToFindSingleRecord(err, "review(id=%d)", id)
	}

	return &row, nil
}

func (r *ReviewQueryRepositoryImpl) FindFavoriteReviewListByUserID(userID int, query *query.FindListPaginationQuery) (*entity.ReviewDetailWithIsFavoriteList, error) {
	var rows entity.ReviewDetailWithIsFavoriteList

	if err := r.DB.
		Joins("INNER JOIN (SELECT review_id, created_at FROM user_favorite_review WHERE user_id = ?) uf ON review.id = uf.review_id", userID).
		Order("uf.created_at DESC").
		Limit(query.Limit).
		Offset(query.Offset).
		Find(&rows.List).Offset(0).Count(&rows.TotalNumber).Error; err != nil {
		return nil, errors.Wrapf(err, "failed find favorite reviews by userID=%d", userID)
	}

	return &rows, nil
}

func (r *ReviewQueryRepositoryImpl) FindFavoriteReviewWithIsFavoriteListByUserID(userID, targetUserID int, query *query.FindListPaginationQuery) (*entity.ReviewDetailWithIsFavoriteList, error) {
	var rows entity.ReviewDetailWithIsFavoriteList

	if err := r.DB.
		Select("review.*, CASE WHEN user_favorite_review.review_id IS NULL THEN 'FALSE' ELSE 'TRUE' END is_favorite").
		Joins("LEFT JOIN user_favorite_review ON review.id = user_favorite_review.review_id AND user_favorite_review.user_id = ?", userID).
		Joins("INNER JOIN (SELECT review_id, created_at FROM user_favorite_review WHERE user_id = ?) uf ON review.id = uf.review_id", targetUserID).
		Order("uf.created_at DESC").
		Limit(query.Limit).
		Offset(query.Offset).
		Find(&rows.List).Offset(0).Count(&rows.TotalNumber).Error; err != nil {
		return nil, errors.Wrapf(err, "failed find favorite reviews by userID=%d", targetUserID)
	}

	return &rows, nil
}

func (r *ReviewQueryRepositoryImpl) FindAll() ([]*entity.Review, error) {
	var rows []*entity.Review

	if err := r.DB.Find(&rows).Error; err != nil {
		return nil, errors.Wrap(err, "failed to find all review")
	}

	return rows, nil
}

// パスパラメータで飛んで来た値によって検索クエリを切り替える
func (r *ReviewQueryRepositoryImpl) buildShowReviewListQuery(query *query.ShowReviewListQuery) *gorm.DB {
	q := r.DB

	if query.UserID != 0 {
		q = q.Where("review.user_id = ?", query.UserID)
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
		q = q.Where("tourist_spot_id IN (SELECT tourist_spot_id FROM tourist_spot_area_category WHERE area_category_id IN (SELECT id FROM area_category WHERE area_id = ?))", query.AreaID)
	}

	if query.SubAreaID != 0 {
		q = q.Where("tourist_spot_id IN (SELECT tourist_spot_id FROM tourist_spot_area_category WHERE area_category_id IN (SELECT id FROM area_category WHERE sub_area_id = ?))", query.SubAreaID)
	}

	if query.SubSubAreaID != 0 {
		q = q.Where("tourist_spot_id IN (SELECT tourist_spot_id FROM tourist_spot_area_category WHERE area_category_id IN (SELECT id FROM area_category WHERE sub_sub_area_id = ?))", query.SubSubAreaID)
	}

	if query.MetasearchAreaID != 0 {
		q = q.Where("tourist_spot_id IN (SELECT tourist_spot_id FROM tourist_spot_area_category WHERE area_category_id IN (SELECT id FROM area_category WHERE metasearch_area_id = ?))", query.MetasearchAreaID)
	}

	if query.MetasearchSubAreaID != 0 {
		q = q.Where("tourist_spot_id IN (SELECT tourist_spot_id FROM tourist_spot_area_category WHERE area_category_id IN (SELECT id FROM area_category WHERE metasearch_sub_area_id = ?))", query.MetasearchSubAreaID)
	}

	if query.MetasearchSubSubAreaID != 0 {
		q = q.Where("tourist_spot_id IN (SELECT tourist_spot_id FROM tourist_spot_area_category WHERE area_category_id IN (SELECT id FROM area_category WHERE metasearch_sub_sub_area_id = ?))", query.MetasearchSubSubAreaID)
	}

	if len(query.InnIDs) > 0 {
		q = q.Or("inn_id IN (?)", query.InnIDs)
	}

	if query.Keyword != "" {
		q = q.Where("body LIKE ?", query.SQLLikeKeyword())
	}

	if query.ExcludeID != 0 {
		q = q.Not("id = ?", query.ExcludeID)
	}

	return q
}

func (r *ReviewQueryRepositoryImpl) buildFindFeedListQuery(userID int) *gorm.DB {
	q := r.DB

	if userID != 0 {
		q = q.Where("review.user_id IN (SELECT target_id FROM user_following WHERE user_id = ?)", userID).Or("review.id IN (SELECT review_id FROM review_hashtag WHERE hashtag_id IN (SELECT hashtag_id FROM user_follow_hashtag WHERE user_id = ?))", userID)
	}

	return q
}

func (r *ReviewQueryRepositoryImpl) FindByID(id int) (*entity.Review, error) {
	var review entity.Review

	if err := r.DB.Find(&review, id).Error; err != nil {
		return nil, errors.Wrapf(err, "Failed to get review")
	}

	return &review, nil
}

func (r *ReviewQueryRepositoryImpl) IsExist(id int) (bool, error) {
	var review entity.Review

	err := r.DB.Find(&review, id).Error

	return ErrorToIsExist(err, "review(id=%d)", id)
}

func (r *ReviewQueryRepositoryImpl) FindReviewCommentListByReviewID(reviewID int, limit int) ([]*entity.ReviewCommentWithIsFavorite, error) {
	var comments []*entity.ReviewCommentWithIsFavorite

	err := r.DB.
		Where("review_id = ?", reviewID).
		Order("created_at DESC").
		Limit(limit).
		Find(&comments).
		Error

	if err != nil {
		return nil, errors.Wrap(err, "failed find review comments.")
	}

	return comments, nil
}

func (r *ReviewQueryRepositoryImpl) FindReviewCommentWithIsFavoriteListByReviewID(reviewID int, limit int, userID int) ([]*entity.ReviewCommentWithIsFavorite, error) {
	var comments []*entity.ReviewCommentWithIsFavorite

	err := r.DB.
		Select("review_comment.*, CASE WHEN user_favorite_review_comment.review_comment_id IS NULL THEN 'FALSE' ELSE 'TRUE' END is_favorite").
		Joins("LEFT JOIN user_favorite_review_comment ON review_comment.id = user_favorite_review_comment.review_comment_id AND user_favorite_review_comment.user_id = ?", userID).
		Where("review_id = ?", reviewID).
		Order("created_at DESC").
		Limit(limit).
		Find(&comments).
		Error

	if err != nil {
		return nil, errors.Wrap(err, "failed find review comments.")
	}

	return comments, nil
}

func (r *ReviewQueryRepositoryImpl) FindReviewCommentReplyListByReviewCommentID(reviewCommentID int) ([]*entity.ReviewCommentReplyWithIsFavorite, error) {
	var rows []*entity.ReviewCommentReplyWithIsFavorite

	if err := r.DB.Where("review_comment_id = ?", reviewCommentID).
		Order("created_at DESC").
		Limit(defaultAcquisitionNumber).
		Find(&rows).Error; err != nil {
		return nil, errors.Wrap(err, "failed to find review comment replies")
	}

	return rows, nil
}

func (r *ReviewQueryRepositoryImpl) FindReviewCommentReplyWithIsFavoriteListByReviewCommentID(reviewCommentID int, userID int) ([]*entity.ReviewCommentReplyWithIsFavorite, error) {
	var rows []*entity.ReviewCommentReplyWithIsFavorite

	if err := r.DB.
		Select("review_comment_reply.*, CASE WHEN user_favorite_review_comment_reply.review_comment_reply_id IS NULL THEN 'FALSE' ELSE 'TRUE' END is_favorite").
		Joins("LEFT JOIN user_favorite_review_comment_reply ON review_comment_reply.id = user_favorite_review_comment_reply.review_comment_reply_id AND user_favorite_review_comment_reply.user_id = ?", userID).
		Where("review_comment_id = ?", reviewCommentID).
		Order("created_at DESC").
		Limit(defaultAcquisitionNumber).
		Find(&rows).Error; err != nil {
		return nil, errors.Wrap(err, "failed to find review comment replies")
	}

	return rows, nil
}

func (r *ReviewQueryRepositoryImpl) IsExistReviewComment(id int) (bool, error) {
	var row entity.ReviewCommentTiny

	err := r.DB.Find(&row, id).Error

	return ErrorToIsExist(err, "review_comment(id=%d)", id)
}

func (r *ReviewQueryRepositoryImpl) IsExistReviewCommentFavorite(userID, reviewCommentID int) (bool, error) {
	var row entity.UserFavoriteReviewComment

	err := r.DB.Where("user_id = ? AND review_comment_id = ?", userID, reviewCommentID).First(&row).Error

	return ErrorToIsExist(err, "review_comment_favorite(user_id=%d,review_comment_id=%d)", userID, reviewCommentID)
}

func (r *ReviewQueryRepositoryImpl) FindReviewCommentByID(id int) (*entity.ReviewCommentTiny, error) {
	var row entity.ReviewCommentTiny

	if err := r.DB.Find(&row, id).Error; err != nil {
		return nil, ErrorToFindSingleRecord(err, "review_comment(id=%d)", id)
	}

	return &row, nil
}

func (r *ReviewQueryRepositoryImpl) FindReviewCommentDetailByID(id int) (*entity.ReviewCommentDetail, error) {
	var row entity.ReviewCommentDetail

	if err := r.DB.Find(&row, id).Error; err != nil {
		return nil, ErrorToFindSingleRecord(err, "review_comment(id=%d)", id)
	}

	return &row, nil
}

func (r *ReviewQueryRepositoryImpl) FindReviewCommentReplyByID(id int) (*entity.ReviewCommentReplyTiny, error) {
	var row entity.ReviewCommentReplyTiny

	if err := r.DB.Find(&row, id).Error; err != nil {
		return nil, ErrorToFindSingleRecord(err, "review_comment(id=%d)", id)
	}

	return &row, nil
}

func (r *ReviewQueryRepositoryImpl) FindReviewCommentReplyDetailByID(id int) (*entity.ReviewCommentReplyDetail, error) {
	var row entity.ReviewCommentReplyDetail

	if err := r.DB.Find(&row, id).Error; err != nil {
		return nil, ErrorToFindSingleRecord(err, "review_comment(id=%d)", id)
	}

	return &row, nil
}

func (r *ReviewQueryRepositoryImpl) IsExistReviewCommentReply(id int) (bool, error) {
	var row entity.ReviewCommentReplyTiny

	err := r.DB.Find(&row, id).Error

	return ErrorToIsExist(err, "review_comment_reply(id=%d)", id)
}

// tourist_spotに紐づくreviewの中で、mediaを持ち、最新のreviewを取得
func (r *ReviewQueryRepositoryImpl) FindLatestHasMediaReviewByTouristSpotIDs(touristSpotIDs []int) (*entity.ReviewList, error) {
	var rows entity.ReviewList

	if err := r.DB.Joins("INNER JOIN (SELECT tourist_spot_id, MAX(created_at) as latestDate FROM review GROUP BY tourist_spot_id) tm ON review.tourist_spot_id = tm.tourist_spot_id AND review.created_at = tm.latestDate").
		Where("review.id IN (SELECT review_id FROM review_media) AND review.tourist_spot_id IN (?)", touristSpotIDs).
		Find(&rows.List).Error; err != nil {
		return nil, errors.Wrap(err, "failed find review")
	}

	return &rows, nil
}

// Userに紐づくTouristSpot or Inn のReviewを返す
func (r *ReviewQueryRepositoryImpl) FindRelationLocationReview(query *query.FindListPaginationQuery, userID int) (*entity.ReviewDetailWithIsFavoriteList, error) {
	var rows entity.ReviewDetailWithIsFavoriteList

	if err := r.DB.Where("tourist_spot_id IN (SELECT tourist_spot_id FROM user_tourist_spot WHERE user_id = ?) OR inn_id IN (SELECT inn_id FROM user_inn WHERE user_id = ?)", userID, userID).
		Offset(query.Offset).
		Limit(query.Limit).
		Order("created_at DESC").
		Find(&rows.List).Offset(0).Count(&rows.TotalNumber).Error; err != nil {
		return nil, errors.Wrap(err, "failed find review")
	}

	return &rows, nil
}
