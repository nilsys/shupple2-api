package repository

import (
	"github.com/google/wire"
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/entity"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/model"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/model/query"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/repository"
)

// User参照系レポジトリ実装
type UserQueryRepositoryImpl struct {
	DB *gorm.DB
}

var UserQueryRepositorySet = wire.NewSet(
	wire.Struct(new(UserQueryRepositoryImpl), "*"),
	wire.Bind(new(repository.UserQueryRepository), new(*UserQueryRepositoryImpl)),
)

func (r *UserQueryRepositoryImpl) FindByID(id int) (*entity.User, error) {
	var row entity.User
	if err := r.DB.First(&row, id).Error; err != nil {
		return nil, ErrorToFindSingleRecord(err, "user(id=%d)", id)
	}
	return &row, nil
}

func (r *UserQueryRepositoryImpl) FindByCognitoID(cognitoID string) (*entity.User, error) {
	var row entity.User
	if err := r.DB.Where("cognito_id = ?", cognitoID).First(&row).Error; err != nil {
		return nil, ErrorToFindSingleRecord(err, "user(cognito_id=%s)", cognitoID)
	}
	return &row, nil
}

func (r *UserQueryRepositoryImpl) FindByWordpressID(wordpressUserID int) (*entity.User, error) {
	var row entity.User
	if err := r.DB.Where("wordpress_id = ?", wordpressUserID).First(&row).Error; err != nil {
		return nil, ErrorToFindSingleRecord(err, "user(wordpress_id=%d)", wordpressUserID)
	}
	return &row, nil
}

func (r *UserQueryRepositoryImpl) FindByMigrationCode(code string) (*entity.User, error) {
	var row entity.User
	if err := r.DB.Where("migration_code = ?", code).First(&row).Error; err != nil {
		return nil, ErrorToFindSingleRecord(err, "user(migration_code=%s)", code)
	}
	return &row, nil
}

// TODO: テスト
func (r *UserQueryRepositoryImpl) FindUserRankingListByParams(query *query.FindUserRankingListQuery) ([]*entity.QueryRankingUser, error) {
	var rows []*entity.QueryRankingUser

	q := r.buildFindUserRankingListQuery(query)

	if err := q.
		Table("user").
		Limit(query.Limit).
		Offset(query.Offset).
		Find(&rows).Error; err != nil {
		return nil, errors.Wrap(err, "failed find user ranking list")
	}

	return rows, nil
}

func (r *UserQueryRepositoryImpl) IsExistByUID(uid string) (bool, error) {
	var row entity.User

	err := r.DB.Where("uid = ?", uid).First(&row).Error

	return ErrorToIsExist(err, "user(uid=%s)", uid)
}

// name部分一致検索
func (r *UserQueryRepositoryImpl) SearchByName(name string) ([]*entity.User, error) {
	var rows []*entity.User

	if err := r.DB.Where("MATCH(name) AGAINST(?)", name).Limit(defaultSearchSuggestionsNumber).Find(&rows).Error; err != nil {
		return nil, errors.Wrap(err, "failed to find user list by like name")
	}

	return rows, nil
}

// idで指定されたユーザーがフォローしているユーザー
func (r *UserQueryRepositoryImpl) FindFollowingByID(query *query.FindFollowUser) ([]*entity.User, error) {
	var rows []*entity.User

	if err := r.DB.Where("id IN (SELECT target_id FROM user_following WHERE user_id = ?)", query.ID).
		Limit(query.Limit).
		Offset(query.Offset).
		Find(&rows).Error; err != nil {
		return nil, errors.Wrapf(err, "failed find follow user list user_id=%d", query.ID)
	}

	return rows, nil
}

// idで指定されたユーザーがフォローされているユーザー
func (r *UserQueryRepositoryImpl) FindFollowedByID(query *query.FindFollowUser) ([]*entity.User, error) {
	var rows []*entity.User

	if err := r.DB.Where("id IN (SELECT user_id FROM user_followed WHERE target_id = ?)", query.ID).
		Limit(query.Limit).
		Offset(query.Offset).
		Find(&rows).Error; err != nil {
		return nil, errors.Wrapf(err, "failed find follower user list user_id=%d", query.ID)
	}

	return rows, nil
}

func (r *UserQueryRepositoryImpl) FindFavoritePostUser(postID int, query *query.FindListPaginationQuery) ([]*entity.User, error) {
	var rows []*entity.User

	if err := r.DB.Joins("INNER JOIN (SELECT user_id, created_at FROM user_favorite_post WHERE post_id = ?) uf ON user.id = uf.user_id", postID).
		Order("uf.created_at DESC").
		Limit(query.Limit).
		Offset(query.Offset).
		Find(&rows).Error; err != nil {
		return nil, errors.Wrap(err, "failed to find post favorite user list")
	}

	return rows, nil
}

func (r *UserQueryRepositoryImpl) FindFavoriteReviewUser(reviewID int, query *query.FindListPaginationQuery) ([]*entity.User, error) {
	var rows []*entity.User

	if err := r.DB.Joins("INNER JOIN (SELECT user_id, created_at FROM user_favorite_review WHERE review_id = ?) uf ON user.id = uf.user_id", reviewID).
		Order("uf.created_at DESC").
		Limit(query.Limit).
		Offset(query.Offset).
		Find(&rows).Error; err != nil {
		return nil, errors.Wrap(err, "failed to find review favorite user list")
	}

	return rows, nil
}

func (r *UserQueryRepositoryImpl) FindFavoritePostUserByUserID(postID, userID int, query *query.FindListPaginationQuery) ([]*entity.User, error) {
	var rows []*entity.User

	// user_favorite_post AS f
	// user AS u
	// user_following AS uf
	if err := r.DB.Unscoped().Table("user_favorite_post f").Select("f.*, u.*").
		Joins("JOIN user u ON f.user_id = u.id AND u.deleted_at IS NULL").
		Joins("LEFT JOIN user_following uf ON u.id = uf.target_id and uf.user_id=? WHERE f.post_id=?", userID, postID).
		Order("uf.created_at DESC").Order("f.created_at DESC").Find(&rows).Error; err != nil {
		return nil, errors.Wrap(err, "failed to find favorite post user")
	}

	return rows, nil
}

func (r *UserQueryRepositoryImpl) FindFavoriteReviewUserByUserID(reviewID, userID int, query *query.FindListPaginationQuery) ([]*entity.User, error) {
	var rows []*entity.User

	// user_favorite_review AS f
	// user AS u
	// user_following AS uf
	if err := r.DB.Unscoped().Table("user_favorite_review f").Select("f.*, u.*").
		Joins("JOIN user u ON f.user_id = u.id AND u.deleted_at IS NULL").
		Joins("LEFT JOIN user_following uf ON u.id = uf.target_id and uf.user_id=? WHERE f.review_id=?", userID, reviewID).
		Order("uf.created_at DESC, f.created_at DESC").Find(&rows).Error; err != nil {
		return nil, errors.Wrap(err, "failed to find favorite review user")
	}

	return rows, nil
}

// TODO: クエリ見直し
func (r *UserQueryRepositoryImpl) buildFindUserRankingListQuery(query *query.FindUserRankingListQuery) *gorm.DB {
	q := r.DB

	if query.SortBy == model.UserSortByRANKING {
		return q.Joins("INNER JOIN (SELECT MAX(user_id) AS user_id, SUM(favorite_count) AS favorite_count FROM (SELECT user_id, favorite_count FROM review WHERE tourist_spot_id IN (SELECT tourist_spot_id FROM tourist_spot_category WHERE category_id = ?) AND updated_at BETWEEN ? AND ? UNION ALL SELECT user_id, favorite_count FROM post WHERE id IN (SELECT post_id FROM post_category WHERE category_id = ?) AND updated_at BETWEEN ? AND ?) AS user_favorite_count GROUP BY user_id) user_favorite_count ON user.id = user_favorite_count.user_id", query.CategoryID, query.FromDate, query.ToDate, query.CategoryID, query.FromDate, query.ToDate).
			Order("user_favorite_count.favorite_count DESC")
	}

	// デフォルトはおすすめ順
	return q.Joins("INNER JOIN (SELECT MAX(user_id) AS user_id, SUM(views) AS views_count FROM (SELECT user_id, views FROM review WHERE tourist_spot_id IN (SELECT tourist_spot_id FROM tourist_spot_category WHERE category_id = ?) AND updated_at BETWEEN ? AND ? UNION ALL SELECT user_id, views FROM post WHERE id IN (SELECT post_id FROM post_category WHERE category_id = ?) AND updated_at BETWEEN ? AND ?) AS user_views_count GROUP BY user_id) user_views_count ON user.id = user_views_count.user_id", query.CategoryID, query.FromDate, query.ToDate, query.CategoryID, query.FromDate, query.ToDate).
		Order("user_views_count.views_count DESC")
}
