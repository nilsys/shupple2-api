package repository

import (
	"time"

	"github.com/google/wire"
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/entity"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/model"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/model/query"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/repository"
)

// Post参照系レポジトリ実装
type PostQueryRepositoryImpl struct {
	DB *gorm.DB
}

var PostQueryRepositorySet = wire.NewSet(
	wire.Struct(new(PostQueryRepositoryImpl), "*"),
	wire.Bind(new(repository.PostQueryRepository), new(*PostQueryRepositoryImpl)),
)

func (r *PostQueryRepositoryImpl) FindByLastID(lastID, limit int) ([]*entity.Post, error) {
	var rows []*entity.Post

	if err := r.DB.Where("id > ?", lastID).Order("id").Limit(limit).Find(&rows).Error; err != nil {
		return nil, errors.Wrap(err, "failed to find post")
	}

	return rows, nil
}

func (r *PostQueryRepositoryImpl) FindByID(id int) (*entity.Post, error) {
	var row entity.Post

	if err := r.DB.First(&row, id).Error; err != nil {
		return nil, ErrorToFindSingleRecord(err, "post(id=%d)", id)
	}
	return &row, nil
}

func (r *PostQueryRepositoryImpl) FindPostDetailWithHashtagByID(id int) (*entity.PostDetailWithHashtagAndIsFavorite, error) {
	var row entity.PostDetailWithHashtagAndIsFavorite

	if err := r.DB.First(&row, id).Error; err != nil {
		return nil, ErrorToFindSingleRecord(err, "post(id=%d)", id)
	}

	return &row, nil
}

func (r *PostQueryRepositoryImpl) FindPostDetailWithHashtagAndIsFavoriteByID(id, userID int) (*entity.PostDetailWithHashtagAndIsFavorite, error) {
	var row entity.PostDetailWithHashtagAndIsFavorite

	if err := r.DB.
		Select("post.*, CASE WHEN user_favorite_post.post_id IS NULL THEN 'FALSE' ELSE 'TRUE' END is_favorite").
		Joins("LEFT JOIN user_favorite_post ON post.id = user_favorite_post.post_id AND user_favorite_post.user_id = ?", userID).
		First(&row, id).
		Error; err != nil {
		return nil, ErrorToFindSingleRecord(err, "post(id=%d)", id)
	}

	return &row, nil
}

func (r *PostQueryRepositoryImpl) FindPostDetailWithHashtagBySlug(slug string) (*entity.PostDetailWithHashtag, error) {
	var row entity.PostDetailWithHashtag

	if err := r.DB.Where("slug = ?", slug).First(&row).Error; err != nil {
		return nil, ErrorToFindSingleRecord(err, "post(slug=%s)", slug)
	}

	return &row, nil
}

// 検索条件に指定されたクエリ構造体を用い、postを複数参照
func (r *PostQueryRepositoryImpl) FindListByParams(query *query.FindPostListQuery) (*entity.PostList, error) {
	var postList entity.PostList

	q := r.buildFindListByParamsQuery(query)

	// フリーワード検索の場合
	if query.Keyward != "" {
		if err := q.
			Select("*, CASE WHEN MATCH(title) AGAINST(?) THEN 'TRUE' ELSE 'FALSE' END is_matched_title", query.Keyward).
			Order("is_matched_title desc").
			Order(query.SortBy.GetPostOrderQuery()).
			Limit(query.Limit).
			Offset(query.OffSet).
			Find(&postList.Posts).
			Offset(0).
			Count(&postList.TotalNumber).Error; err != nil {
			return nil, errors.Wrapf(err, "Failed get posts by params")
		}

		return &postList, nil
	}

	if err := q.
		Order(query.SortBy.GetPostOrderQuery()).
		Limit(query.Limit).
		Offset(query.OffSet).
		Find(&postList.Posts).
		Offset(0).
		Count(&postList.TotalNumber).Error; err != nil {
		return nil, errors.Wrapf(err, "Failed get posts by params")
	}

	return &postList, nil
}

func (r *PostQueryRepositoryImpl) FindListWithIsFavoriteByParams(query *query.FindPostListQuery, userID int) (*entity.PostList, error) {
	var postList entity.PostList

	q := r.buildFindListByParamsQuery(query)

	if err := q.
		Select("post.*, CASE WHEN user_favorite_post.post_id IS NULL THEN 'FALSE' ELSE 'TRUE' END is_favorite").
		Joins("LEFT JOIN user_favorite_post ON post.id = user_favorite_post.post_id AND user_favorite_post.user_id = ?", userID).
		Order(query.SortBy.GetPostOrderQuery()).
		Limit(query.Limit).
		Offset(query.OffSet).
		Find(&postList.Posts).Count(&postList.TotalNumber).Error; err != nil {
		return nil, errors.Wrapf(err, "Failed get posts by params")
	}

	return &postList, nil
}

// ユーザーIDからフォローしているハッシュタグ or ユーザーのpost一覧を参照
func (r *PostQueryRepositoryImpl) FindFeedListByUserID(targetUserID int, query *query.FindListPaginationQuery) (*entity.PostList, error) {
	var rows entity.PostList

	q := r.buildFindFeedListQuery(targetUserID)

	if err := q.
		Order("created_at DESC").
		Limit(query.Limit).
		Offset(query.Offset).
		Find(&rows.Posts).Offset(0).Count(&rows.TotalNumber).Error; err != nil {
		return nil, errors.Wrapf(err, "failed find feed posts")
	}

	return &rows, nil
}

func (r *PostQueryRepositoryImpl) FindFeedListWithIsFavoriteByUserID(userID, targetUserID int, query *query.FindListPaginationQuery) (*entity.PostList, error) {
	var rows entity.PostList

	q := r.buildFindFeedListQuery(targetUserID)

	if err := q.
		Select("post.*, CASE WHEN user_favorite_post.post_id IS NULL THEN 'FALSE' ELSE 'TRUE' END is_favorite").
		Joins("LEFT JOIN user_favorite_post ON post.id = user_favorite_post.post_id AND user_favorite_post.user_id = ?", userID).
		Order("created_at DESC").
		Limit(query.Limit).
		Offset(query.Offset).
		Find(&rows.Posts).Offset(0).Count(&rows.TotalNumber).Error; err != nil {
		return nil, errors.Wrapf(err, "failed find feed posts")
	}

	return &rows, nil
}

func (r *PostQueryRepositoryImpl) FindFavoriteListByUserID(targetUserID int, query *query.FindListPaginationQuery) (*entity.PostList, error) {
	var rows entity.PostList

	if err := r.DB.
		Joins("INNER JOIN (SELECT post_id, created_at FROM user_favorite_post WHERE user_id = ?) uf ON post.id = uf.post_id", targetUserID).
		Order("uf.created_at DESC").
		Limit(query.Limit).
		Offset(query.Offset).
		Find(&rows.Posts).Offset(0).Count(&rows.TotalNumber).Error; err != nil {
		return nil, errors.Wrapf(err, "failed find favorite posts by userID=%d", targetUserID)
	}
	return &rows, nil
}

func (r *PostQueryRepositoryImpl) FindFavoriteListWithIsFavoriteByUserID(userID, targetUserID int, query *query.FindListPaginationQuery) (*entity.PostList, error) {
	var rows entity.PostList

	if err := r.DB.
		Select("post.*, CASE WHEN user_favorite_post.post_id IS NULL THEN 'FALSE' ELSE 'TRUE' END is_favorite").
		Joins("LEFT JOIN user_favorite_post ON post.id = user_favorite_post.post_id AND user_favorite_post.user_id = ?", userID).
		Joins("INNER JOIN (SELECT post_id, created_at FROM user_favorite_post WHERE user_id = ?) uf ON post.id = uf.post_id", targetUserID).
		Order("uf.created_at DESC").
		Limit(query.Limit).
		Offset(query.Offset).
		Find(&rows.Posts).Offset(0).Count(&rows.TotalNumber).Error; err != nil {
		return nil, errors.Wrapf(err, "failed find favorite posts by userID=%d", targetUserID)
	}
	return &rows, nil
}

// クエリ構造体を用い、検索クエリを作成
func (r *PostQueryRepositoryImpl) buildFindListByParamsQuery(query *query.FindPostListQuery) *gorm.DB {
	q := r.DB

	if query.UserID != 0 {
		q = q.Where("post.user_id = ?", query.UserID)
	}

	if query.AreaID != 0 {
		q = q.Where("id IN (SELECT post_id FROM post_area_category WHERE area_category_id IN (SELECT id FROM area_category WHERE area_id = ?))", query.AreaID)
	}

	if query.SubAreaID != 0 {
		q = q.Where("id IN (SELECT post_id FROM post_area_category WHERE area_category_id IN (SELECT id FROM area_category WHERE sub_area_id = ?))", query.SubAreaID)
	}

	if query.SubSubAreaID != 0 {
		q = q.Where("id IN (SELECT post_id FROM post_area_category WHERE area_category_id IN (SELECT id FROM area_category WHERE sub_sub_area_id = ?))", query.SubSubAreaID)
	}

	// 一つ上のエリアに紐づいた記事を返す
	if query.ChildAreaID != 0 {
		q = q.Where("id IN (SELECT post_id FROM post_area_category WHERE area_category_id IN (SELECT area_group FROM area_category WHERE area_id = ?))", query.ChildAreaID)
	}
	if query.ChildSubAreaID != 0 {
		q = q.Where("id IN (SELECT post_id FROM post_area_category WHERE area_category_id IN (SELECT area_id FROM area_category WHERE sub_area_id = ?))", query.ChildSubAreaID)
	}
	if query.ChildSubSubAreaID != 0 {
		q = q.Where("id IN (SELECT post_id FROM post_area_category WHERE area_category_id IN (SELECT sub_area_id FROM area_category WHERE sub_sub_area_id = ?))", query.ChildSubSubAreaID)
	}
	//

	if query.MetasearchAreaID != 0 {
		q = q.Where("id IN (SELECT post_id FROM post_area_category WHERE area_category_id IN (SELECT id FROM area_category WHERE `metasearch_area_id` = ?))", query.MetasearchAreaID)
	}

	if query.MetasearchSubAreaID != 0 {
		q = q.Where("id IN (SELECT post_id FROM post_area_category WHERE area_category_id IN (SELECT id FROM area_category WHERE `metasearch_sub_area_id` = ?))", query.MetasearchSubAreaID)
	}

	if query.MetasearchSubSubAreaID != 0 {
		q = q.Where("id IN (SELECT post_id FROM post_area_category WHERE area_category_id IN (SELECT id FROM area_category WHERE `metasearch_sub_sub_area_id` = ?))", query.MetasearchSubSubAreaID)
	}

	if query.InnTypeID != 0 {
		q = q.Where("id IN (SELECT post_id FROM post_area_category WHERE area_category_id IN (SELECT id FROM area_category WHERE `metasearch_inn_type_id` = ?))", query.InnTypeID)
	}

	if query.InnDiscerningType != 0 {
		q = q.Where("id IN (SELECT post_id FROM post_area_category WHERE area_category_id IN (SELECT id FROM area_category WHERE `metasearch_discerning_condition_id` = ?))", query.InnDiscerningType)
	}

	if query.ThemeID != 0 {
		q = q.Where("id IN (SELECT post_id FROM post_theme_category WHERE theme_category_id IN (SELECT id FROM theme_category WHERE theme_id = ?))", query.ThemeID)
	}

	if query.HashTag != "" {
		q = q.Where("id IN (SELECT post_id FROM post_hashtag WHERE hashtag_id = (SELECT id FROM hashtag WHERE name = ?))", query.HashTag)
	}

	// TODO: titleに引っかかる物が優先順位が高い、その後body
	if query.Keyward != "" {
		q = q.Where("MATCH(title) AGAINST(?)", query.Keyward).Or("id IN (SELECT post_id FROM post_body WHERE MATCH(body) AGAINST(?))", query.Keyward)
	}

	if query.SortBy == model.MediaSortByRANKING {
		q = q.Where("updated_at BETWEEN ? AND ?", time.Date(time.Now().Year(), time.Now().Month()-recommendMonthPeriod, time.Now().Day(), 0, 0, 0, 0, time.Local), time.Now())
	}

	return q
}

func (r *PostQueryRepositoryImpl) buildFindFeedListQuery(userID int) *gorm.DB {
	q := r.DB

	if userID != 0 {
		q = q.Where("post.user_id IN (SELECT target_id FROM user_following WHERE user_id = ?)", userID).Or("id IN (SELECT post_id FROM post_hashtag WHERE hashtag_id IN (SELECT hashtag_id FROM user_follow_hashtag WHERE user_id = ?))", userID)
	}

	return q
}

func (r *PostQueryRepositoryImpl) IsExist(id int) (bool, error) {
	var row entity.Post

	err := r.DB.First(&row, id).Error

	return ErrorToIsExist(err, "post(id=%d)", id)
}
