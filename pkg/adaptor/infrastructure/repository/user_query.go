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

func (r *UserQueryRepositoryImpl) FindByWordpressID(wordpressUserID int) (*entity.User, error) {
	var row entity.User
	if err := r.DB.Where("wordpress_id = ?", wordpressUserID).First(&row).Error; err != nil {
		return nil, ErrorToFindSingleRecord(err, "user(wordpress_id=%d)", wordpressUserID)
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

// name部分一致検索
func (r *UserQueryRepositoryImpl) SearchByName(name string) ([]*entity.User, error) {
	var rows []*entity.User

	if err := r.DB.Where("MATCH(name) AGAINST(?)", name).Limit(10).Find(&rows).Error; err != nil {
		return nil, errors.Wrap(err, "failed to find user list by like name")
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
