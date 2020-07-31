package repository

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cognitoidentityprovider"
	"github.com/google/wire"
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
	"github.com/stayway-corp/stayway-media-api/pkg/config"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/entity"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/model"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/model/query"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/model/serror"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/repository"
)

// User参照系レポジトリ実装
type UserQueryRepositoryImpl struct {
	DB         *gorm.DB
	AWSConfig  config.AWS
	AWSSession *session.Session
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

func (r *UserQueryRepositoryImpl) FindByUIDs(uIDs []string) ([]*entity.User, error) {
	var rows []*entity.User
	if err := r.DB.Where("uid IN (?)", uIDs).Find(&rows).Error; err != nil {
		return nil, errors.Wrap(err, "failed to find user by UIDs")
	}
	return rows, nil
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

func (r *UserQueryRepositoryImpl) FindByMigrationCode(code string) (*entity.UserTiny, error) {
	var row entity.UserTiny
	if err := r.DB.Where("migration_code = ?", code).First(&row).Error; err != nil {
		return nil, ErrorToFindSingleRecord(err, "user(migration_code=%s)", code)
	}
	return &row, nil
}

func (r *UserQueryRepositoryImpl) FindUserRankingListByParams(query *query.FindUserRankingListQuery) ([]*entity.UserDetail, error) {
	var rows []*entity.UserDetail

	q := r.buildFindUserRankingListQuery(query)
	// MEMO: validationを掛けているのであり得ないが
	if q == nil {
		return nil, serror.New(nil, serror.CodeInvalidParam, "Invalid list user ranking search input")
	}

	if err := q.
		Limit(query.Limit).
		Offset(query.Offset).
		Find(&rows).Error; err != nil {
		return nil, errors.Wrap(err, "failed find user ranking list")
	}

	return rows, nil
}

func (r *UserQueryRepositoryImpl) FindByUID(uid string) (*entity.UserTiny, error) {
	var row entity.UserTiny

	if err := r.DB.Where("uid = ?", uid).First(&row).Error; err != nil {
		return nil, ErrorToFindSingleRecord(err, "user(uid=%s)", uid)
	}

	return &row, nil
}

func (r *UserQueryRepositoryImpl) FindRecommendFollowUserList(interestIDs []int) ([]*entity.UserTiny, error) {
	var rows []*entity.UserTiny

	if err := r.DB.
		Joins("INNER JOIN (SELECT user_id, SUM(weekly_views) AS views_count FROM (SELECT user_id, weekly_views FROM review UNION ALL SELECT user_id, weekly_views FROM post) AS user_views_count GROUP BY user_id) user_views_count ON user.id = user_views_count.user_id").
		Order("user_views_count.views_count DESC").
		Where("id IN (SELECT user_id FROM user_interest WHERE interest_id IN (?))", interestIDs).
		Limit(defaultFollowRecommendUserLimit).
		Find(&rows).Error; err != nil {
		return nil, errors.Wrap(err, "failed find follow recommend user list")
	}

	return rows, nil
}

func (r *UserQueryRepositoryImpl) IsFollow(targetID int, userID int) (bool, error) {
	var row entity.UserFollowing

	err := r.DB.Where("user_id = ? AND target_id = ?", userID, targetID).First(&row).Error

	return ErrorToIsExist(err, "user_following(user_id=%s, target_id=%s)", userID, targetID)
}

func (r *UserQueryRepositoryImpl) FindUserDetailWithCountByID(id int) (*entity.UserDetailWithMediaCount, error) {
	var row entity.UserDetailWithMediaCount

	if err := r.DB.Select("*").Where("user.id = ?", id).
		Joins("LEFT JOIN (SELECT COUNT(id) as review_count, MAX(user_id) as user_id FROM review WHERE user_id = ? AND deleted_at IS NULL) AS r ON user.id = r.user_id", id).
		Joins("LEFT JOIN (SELECT COUNT(id) as post_count, MAX(user_id) as user_id FROM post WHERE user_id = ? AND deleted_at IS NULL) AS p ON user.id = p.user_id", id).
		Joins("LEFT JOIN (SELECT COUNT(id) as vlog_count, MAX(user_id) as user_id FROM vlog WHERE user_id = ? AND deleted_at IS NULL) AS v ON user.id = v.user_id", id).
		Joins("LEFT JOIN (SELECT COUNT(target_id) as followed_count, MAX(user_id) as user_id FROM user_followed WHERE user_id = ?) AS ufi ON user.id = ufi.user_id", id).
		Joins("LEFT JOIN (SELECT COUNT(target_id) as following_count, MAX(user_id) as user_id FROM user_following WHERE user_id = ?) AS ufe ON user.id = ufe.user_id", id).
		First(&row).Error; err != nil {
		return nil, ErrorToFindSingleRecord(err, "user(id=%d)", id)
	}

	return &row, nil
}

func (r *UserQueryRepositoryImpl) IsExistByUID(uid string) (bool, error) {
	var row entity.User

	err := r.DB.Unscoped().Where("uid = ?", uid).First(&row).Error

	return ErrorToIsExist(err, "user(uid=%s)", uid)
}

func (r *UserQueryRepositoryImpl) FindByCognitoUserName(cognitoUserName []string) ([]*entity.UserTiny, error) {
	var row []*entity.UserTiny
	if err := r.DB.Where("cognito_user_name IN (?)", cognitoUserName).Find(&row).Error; err != nil {
		return nil, errors.Wrap(err, "failed find by cognito_user_name")
	}
	return row, nil
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
func (r *UserQueryRepositoryImpl) FindFollowingByID(query *query.FindFollowUser) ([]*entity.UserTinyWithIsFollow, error) {
	var rows []*entity.UserTinyWithIsFollow

	if err := r.DB.
		Joins("LEFT JOIN user_following ON user.id = user_following.user_id").
		Where("id IN (SELECT target_id FROM user_following WHERE user_id = ?)", query.ID).
		Limit(query.Limit).
		Offset(query.Offset).
		Order("user_following.created_at DESC").
		Find(&rows).Error; err != nil {
		return nil, errors.Wrapf(err, "failed find follow user list user_id=%d", query.ID)
	}

	return rows, nil
}

// idで指定されたユーザーがフォローされているユーザー
func (r *UserQueryRepositoryImpl) FindFollowedByID(query *query.FindFollowUser) ([]*entity.UserTinyWithIsFollow, error) {
	var rows []*entity.UserTinyWithIsFollow

	if err := r.DB.
		Joins("LEFT JOIN user_following ON user.id = user_following.target_id").
		Where("id IN (SELECT user_id FROM user_following WHERE target_id = ?)", query.ID).
		Limit(query.Limit).
		Offset(query.Offset).
		Order("user_following.created_at DESC").
		Find(&rows).Error; err != nil {
		return nil, errors.Wrapf(err, "failed find follower user list user_id=%d", query.ID)
	}

	return rows, nil
}

func (r *UserQueryRepositoryImpl) FindFollowingWithIsFollowByID(userID int, query *query.FindFollowUser) ([]*entity.UserTinyWithIsFollow, error) {
	var rows []*entity.UserTinyWithIsFollow

	if err := r.DB.
		Select("*, CASE WHEN user_following.target_id IS NULL THEN 'FALSE' ELSE 'TRUE' END is_follow").
		Joins("LEFT JOIN user_following ON user.id = user_following.target_id AND user_following.user_id = ?", userID).
		Where("id IN (SELECT target_id FROM user_following WHERE user_id = ?)", query.ID).
		Limit(query.Limit).
		Offset(query.Offset).
		Order(fmt.Sprintf("id = %d DESC, is_follow DESC, user_following.created_at DESC", userID)).
		Find(&rows).Error; err != nil {
		return nil, errors.Wrapf(err, "failed find follow user list user_id=%d", query.ID)
	}

	return rows, nil
}

func (r *UserQueryRepositoryImpl) FindFollowedWithIsFollowByID(userID int, query *query.FindFollowUser) ([]*entity.UserTinyWithIsFollow, error) {
	var rows []*entity.UserTinyWithIsFollow

	if err := r.DB.
		Select("*, CASE WHEN user_following.target_id IS NULL THEN 'FALSE' ELSE 'TRUE' END is_follow").
		Joins("LEFT JOIN user_following ON user.id = user_following.target_id AND user_following.user_id = ?", userID).
		Where("id IN (SELECT user_id FROM user_following WHERE target_id = ?)", query.ID).
		Limit(query.Limit).
		Offset(query.Offset).
		Order(fmt.Sprintf("id = %d DESC, is_follow DESC, user_following.created_at DESC", userID)).
		Find(&rows).Error; err != nil {
		return nil, errors.Wrapf(err, "failed find follow user list user_id=%d", query.ID)
	}

	return rows, nil
}

func (r *UserQueryRepositoryImpl) FindFavoritePostUser(postID int, query *query.FindListPaginationQuery) ([]*entity.UserTiny, error) {
	var rows []*entity.UserTiny

	if err := r.DB.Joins("INNER JOIN (SELECT user_id, created_at FROM user_favorite_post WHERE post_id = ?) uf ON user.id = uf.user_id", postID).
		Order("uf.created_at DESC").
		Limit(query.Limit).
		Offset(query.Offset).
		Find(&rows).Error; err != nil {
		return nil, errors.Wrap(err, "failed to find post favorite user list")
	}

	return rows, nil
}

func (r *UserQueryRepositoryImpl) FindFavoriteReviewUser(reviewID int, query *query.FindListPaginationQuery) ([]*entity.UserTiny, error) {
	var rows []*entity.UserTiny

	if err := r.DB.Joins("INNER JOIN (SELECT user_id, created_at FROM user_favorite_review WHERE review_id = ?) uf ON user.id = uf.user_id", reviewID).
		Order("uf.created_at DESC").
		Limit(query.Limit).
		Offset(query.Offset).
		Find(&rows).Error; err != nil {
		return nil, errors.Wrap(err, "failed to find review favorite user list")
	}

	return rows, nil
}

func (r *UserQueryRepositoryImpl) FindFavoriteVlogUser(vlogID int, query *query.FindListPaginationQuery) ([]*entity.UserTiny, error) {
	var rows []*entity.UserTiny

	if err := r.DB.Joins("INNER JOIN (SELECT user_id, created_at FROM user_favorite_vlog WHERE vlog_id = ?) uf ON user.id = uf.user_id", vlogID).
		Order("uf.created_at DESC").
		Limit(query.Limit).
		Offset(query.Offset).
		Find(&rows).Error; err != nil {
		return nil, errors.Wrap(err, "failed to find vlog favorite user list")
	}

	return rows, nil
}

func (r *UserQueryRepositoryImpl) FindFavoriteComicUser(comicID int, query *query.FindListPaginationQuery) ([]*entity.UserTiny, error) {
	var rows []*entity.UserTiny

	if err := r.DB.Joins("INNER JOIN (SELECT user_id, created_at FROM user_favorite_comic WHERE comic_id = ?) uf ON user.id = uf.user_id", comicID).
		Order("uf.created_at DESC").
		Limit(query.Limit).
		Offset(query.Offset).
		Find(&rows).Error; err != nil {
		return nil, errors.Wrap(err, "failed to find comic favorite user list")
	}

	return rows, nil
}

// フォローしているユーザーがorder上位
func (r *UserQueryRepositoryImpl) FindFavoritePostUserByUserID(postID, userID int, query *query.FindListPaginationQuery) ([]*entity.UserTiny, error) {
	var rows []*entity.UserTiny

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

// フォローしているユーザーがorder上位
func (r *UserQueryRepositoryImpl) FindFavoriteReviewUserByUserID(reviewID, userID int, query *query.FindListPaginationQuery) ([]*entity.UserTiny, error) {
	var rows []*entity.UserTiny

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

// フォローしているユーザーがorder上位
func (r *UserQueryRepositoryImpl) FindFavoriteVlogUserByUserID(vlogID, userID int, query *query.FindListPaginationQuery) ([]*entity.UserTiny, error) {
	var rows []*entity.UserTiny

	// user_favorite_vlog AS f
	// user AS u
	// user_following AS uf
	if err := r.DB.Unscoped().Table("user_favorite_vlog f").Select("f.*, u.*").
		Joins("JOIN user u ON f.user_id = u.id AND u.deleted_at IS NULL").
		Joins("LEFT JOIN user_following uf ON u.id = uf.target_id and uf.user_id=? WHERE f.vlog_id=?", userID, vlogID).
		Order("uf.created_at DESC, f.created_at DESC").Find(&rows).Error; err != nil {
		return nil, errors.Wrap(err, "failed to find favorite vlog user")
	}

	return rows, nil
}

// フォローしているユーザーがorder上位
func (r *UserQueryRepositoryImpl) FindFavoriteComicUserByUserID(comicID, userID int, query *query.FindListPaginationQuery) ([]*entity.UserTiny, error) {
	var rows []*entity.UserTiny

	// user_favorite_comic AS f
	// user AS u
	// user_following AS uf
	if err := r.DB.Unscoped().Table("user_favorite_comic f").Select("f.*, u.*").
		Joins("JOIN user u ON f.user_id = u.id AND u.deleted_at IS NULL").
		Joins("LEFT JOIN user_following uf ON u.id = uf.target_id and uf.user_id=? WHERE f.comic_id=?", userID, comicID).
		Order("uf.created_at DESC, f.created_at DESC").Find(&rows).Error; err != nil {
		return nil, errors.Wrap(err, "failed to find favorite comic user")
	}

	return rows, nil
}

func (r *UserQueryRepositoryImpl) FindCfProjectSupporterByCfProjectID(cfProjectID int) (*entity.UserTinyList, error) {
	var rows entity.UserTinyList
	if err := r.DB.Where("id IN (SELECT user_id FROM payment LEFT JOIN payment_cf_return_gift ON payment.id = payment_cf_return_gift.payment_id WHERE payment_cf_return_gift.cf_project_id = ?)", cfProjectID).Find(&rows.List).Error; err != nil {
		return nil, errors.Wrap(err, "failed find support cf_project user")
	}
	return &rows, nil
}

// TODO: クエリ見直し
func (r *UserQueryRepositoryImpl) buildFindUserRankingListQuery(query *query.FindUserRankingListQuery) *gorm.DB {
	q := r.DB

	if query.SortBy == model.UserSortByRANKING {
		if query.AreaID != 0 {
			return q.Joins("INNER JOIN (SELECT MAX(user_id) AS user_id, SUM(favorite_count) AS favorite_count FROM (SELECT user_id, fp.favorite_count FROM post INNER JOIN (SELECT post_id, COUNT(post_id) AS favorite_count FROM user_favorite_post WHERE created_at BETWEEN ? AND ? GROUP BY post_id ORDER BY favorite_count DESC) AS fp ON post.id = fp.post_id WHERE id IN (SELECT post_id FROM post_area_category WHERE area_category_id IN (SELECT id FROM area_category WHERE area_id = ?)) UNION ALL SELECT user_id, fr.favorite_count FROM review INNER JOIN (SELECT review_id, COUNT(review_id) AS favorite_count FROM user_favorite_review WHERE created_at BETWEEN ? AND ? GROUP BY review_id ORDER BY favorite_count DESC) AS fr ON review.id = fr.review_id WHERE tourist_spot_id IN(SELECT tourist_spot_id FROM tourist_spot_area_category WHERE area_category_id IN (SELECT id FROM area_category WHERE area_id = ?))) AS user_favorite_count GROUP BY user_id) user_favorite_count ON user.id = user_favorite_count.user_id", query.FromDate, query.ToDate, query.AreaID, query.FromDate, query.ToDate, query.AreaID).
				Order("user_favorite_count.favorite_count DESC")
		}
		if query.SubAreaID != 0 {
			return q.Joins("INNER JOIN (SELECT MAX(user_id) AS user_id, SUM(favorite_count) AS favorite_count FROM (SELECT user_id, fp.favorite_count FROM post INNER JOIN (SELECT post_id, COUNT(post_id) AS favorite_count FROM user_favorite_post WHERE created_at BETWEEN ? AND ? GROUP BY post_id ORDER BY favorite_count DESC) AS fp ON post.id = fp.post_id WHERE id IN (SELECT post_id FROM post_area_category WHERE area_category_id IN (SELECT id FROM area_category WHERE sub_area_id = ?)) UNION ALL SELECT user_id, fr.favorite_count FROM review INNER JOIN (SELECT review_id, COUNT(review_id) AS favorite_count FROM user_favorite_review WHERE created_at BETWEEN ? AND ? GROUP BY review_id ORDER BY favorite_count DESC) AS fr ON review.id = fr.review_id WHERE tourist_spot_id IN(SELECT tourist_spot_id FROM tourist_spot_area_category WHERE area_category_id IN (SELECT id FROM area_category WHERE sub_area_id = ?))) AS user_favorite_count GROUP BY user_id) user_favorite_count ON user.id = user_favorite_count.user_id", query.FromDate, query.ToDate, query.SubAreaID, query.FromDate, query.ToDate, query.SubAreaID).
				Order("user_favorite_count.favorite_count DESC")
		}
		if query.SubSubAreaID != 0 {
			return q.Joins("INNER JOIN (SELECT MAX(user_id) AS user_id, SUM(favorite_count) AS favorite_count FROM (SELECT user_id, fp.favorite_count FROM post INNER JOIN (SELECT post_id, COUNT(post_id) AS favorite_count FROM user_favorite_post WHERE created_at BETWEEN ? AND ? GROUP BY post_id ORDER BY favorite_count DESC) AS fp ON post.id = fp.post_id WHERE id IN (SELECT post_id FROM post_area_category WHERE area_category_id IN (SELECT id FROM area_category WHERE sub_sub_area_id = ?)) UNION ALL SELECT user_id, fr.favorite_count FROM review INNER JOIN (SELECT review_id, COUNT(review_id) AS favorite_count FROM user_favorite_review WHERE created_at BETWEEN ? AND ? GROUP BY review_id ORDER BY favorite_count DESC) AS fr ON review.id = fr.review_id WHERE tourist_spot_id IN(SELECT tourist_spot_id FROM tourist_spot_area_category WHERE area_category_id IN (SELECT id FROM area_category WHERE sub_sub_area_id = ?))) AS user_favorite_count GROUP BY user_id) user_favorite_count ON user.id = user_favorite_count.user_id", query.FromDate, query.ToDate, query.SubSubAreaID, query.FromDate, query.ToDate, query.SubSubAreaID).
				Order("user_favorite_count.favorite_count DESC")
		}
		// どれも指定されていないとき
		return q.Joins("INNER JOIN (SELECT MAX(user_id) AS user_id, SUM(favorite_count) AS favorite_count FROM (SELECT user_id, fp.favorite_count FROM post INNER JOIN (SELECT post_id, COUNT(post_id) AS favorite_count FROM user_favorite_post WHERE created_at BETWEEN ? AND ? GROUP BY post_id ORDER BY favorite_count DESC) AS fp ON post.id = fp.post_id UNION ALL SELECT user_id, fr.favorite_count FROM review INNER JOIN (SELECT review_id, COUNT(review_id) AS favorite_count FROM user_favorite_review WHERE created_at BETWEEN ? AND ? GROUP BY review_id ORDER BY favorite_count DESC) AS fr ON review.id = fr.review_id) AS user_favorite_count GROUP BY user_id) user_favorite_count ON user.id = user_favorite_count.user_id", query.FromDate, query.ToDate, query.FromDate, query.ToDate).
			Order("user_favorite_count.favorite_count DESC")
	}

	// 以下RECOMMENDの時
	if query.AreaID != 0 {
		return q.Joins("INNER JOIN (SELECT MAX(user_id) AS user_id, SUM(weekly_views) AS views_count FROM (SELECT user_id, weekly_views FROM review WHERE tourist_spot_id IN (SELECT tourist_spot_id FROM tourist_spot_area_category WHERE area_category_id IN (SELECT id FROM area_category WHERE area_id = ?)) UNION ALL SELECT user_id, weekly_views FROM post WHERE id IN (SELECT post_id FROM post_area_category WHERE area_category_id IN (SELECT id FROM area_category WHERE area_id = ?))) AS user_views_count GROUP BY user_id) user_views_count ON user.id = user_views_count.user_id", query.AreaID, query.AreaID).
			Order("user_views_count.views_count DESC")
	}
	if query.SubAreaID != 0 {
		return q.Joins("INNER JOIN (SELECT MAX(user_id) AS user_id, SUM(weekly_views) AS views_count FROM (SELECT user_id, weekly_views FROM review WHERE tourist_spot_id IN (SELECT tourist_spot_id FROM tourist_spot_area_category WHERE area_category_id IN (SELECT id FROM area_category WHERE sub_area_id = ?)) UNION ALL SELECT user_id, weekly_views FROM post WHERE id IN (SELECT post_id FROM post_area_category WHERE area_category_id IN (SELECT id FROM area_category WHERE sub_area_id = ?))) AS user_views_count GROUP BY user_id) user_views_count ON user.id = user_views_count.user_id", query.SubAreaID, query.SubAreaID).
			Order("user_views_count.views_count DESC")
	}
	if query.SubSubAreaID != 0 {
		return q.Joins("INNER JOIN (SELECT MAX(user_id) AS user_id, SUM(weekly_views) AS views_count FROM (SELECT user_id, weekly_views FROM review WHERE tourist_spot_id IN (SELECT tourist_spot_id FROM tourist_spot_area_category WHERE area_category_id IN (SELECT id FROM area_category WHERE sub_sub_area_id = ?)) UNION ALL SELECT user_id, weekly_views FROM post WHERE id IN (SELECT post_id FROM post_area_category WHERE area_category_id IN (SELECT id FROM area_category WHERE sub_sub_area_id = ?))) AS user_views_count GROUP BY user_id) user_views_count ON user.id = user_views_count.user_id", query.SubSubAreaID, query.SubSubAreaID).
			Order("user_views_count.views_count DESC")
	}

	return q.Joins("INNER JOIN (SELECT MAX(user_id) AS user_id, SUM(weekly_views) AS views_count FROM (SELECT user_id, weekly_views FROM review UNION ALL SELECT user_id, weekly_views FROM post) AS user_views_count GROUP BY user_id) user_views_count ON user.id = user_views_count.user_id").
		Order("user_views_count.views_count DESC")
}

// TODO: ここに置いて良いのか
// cognitoのpoolから電話番号が使用されているか検索
func (r *UserQueryRepositoryImpl) FindConfirmedUserTypeByPhoneNumberFromCognito(number string) ([]*cognitoidentityprovider.UserType, error) {
	svc := cognitoidentityprovider.New(r.AWSSession)

	// TODO: 無限対応
	input := &cognitoidentityprovider.ListUsersInput{
		UserPoolId: aws.String(r.AWSConfig.UserPoolID),
		Filter:     aws.String(fmt.Sprintf("phone_number = \"%s\"", number)),
	}
	output, err := svc.ListUsers(input)
	if err != nil {
		return nil, errors.Wrap(err, "failed list users")
	}

	users := make([]*cognitoidentityprovider.UserType, 0, len(output.Users))
	// 複数検索条件指定できないのでここで確認する
	for _, user := range output.Users {
		if *user.UserStatus == "CONFIRMED" {
			users = append(users, user)
		}
	}

	return users, nil
}
