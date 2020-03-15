package repository

import (
	"github.com/google/wire"
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/entity"
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

func (r *PostQueryRepositoryImpl) FindByID(id int) (*entity.Post, error) {
	var row entity.Post

	if err := r.DB.First(&row, id).Error; err != nil {
		return nil, ErrorToFindSingleRecord(err, "post(id=%d)", id)
	}
	return &row, nil
}

func (r *PostQueryRepositoryImpl) FindQueryShowByID(id int) (*entity.PostDetailWithHashtag, error) {
	var row entity.PostDetailWithHashtag

	if err := r.DB.First(&row, id).Error; err != nil {
		return nil, ErrorToFindSingleRecord(err, "post(id=%d)", id)
	}

	return &row, nil
}

// 検索条件に指定されたクエリ構造体を用い、postを複数参照
func (r *PostQueryRepositoryImpl) FindListByParams(query *query.FindPostListQuery) ([]*entity.PostDetail, error) {
	var posts []*entity.PostDetail

	q := r.buildFindListByParamsQuery(query)

	if err := q.
		Order(query.SortBy.GetPostOrderQuery()).
		Limit(query.Limit).
		Offset(query.OffSet).
		Find(&posts).Error; err != nil {
		return nil, errors.Wrapf(err, "Failed get posts by params")
	}

	return posts, nil
}

// クエリ構造体を用い、FindListByParams()で使用するsqlクエリを作成
// ユーザーIDからフォローしているハッシュタグ or ユーザーのpost一覧を参照
func (r *PostQueryRepositoryImpl) FindFeedListByUserID(userID int, query *query.FindListPaginationQuery) ([]*entity.PostDetail, error) {
	var posts []*entity.PostDetail

	q := r.buildFindFeedListQuery(userID)

	if err := q.
		Order("updated_at DESC").
		Limit(query.Limit).
		Offset(query.Offset).
		Find(&posts).Error; err != nil {
		return nil, errors.Wrapf(err, "failed find feed posts")
	}

	return posts, nil
}

// クエリ構造体を用い、検索クエリを作成
func (r *PostQueryRepositoryImpl) buildFindListByParamsQuery(query *query.FindPostListQuery) *gorm.DB {
	q := r.DB

	if query.UserID != 0 {
		q = q.Where("user_id = ?", query.UserID)
	}

	if query.AreaID != 0 {
		q = q.Where("id IN (SELECT post_id FROM post_category WHERE category_id = ?)", query.AreaID)
	}

	if query.SubAreaID != 0 {
		q = q.Where("id IN (SELECT post_id FROM post_category WHERE category_id = ?)", query.SubAreaID)
	}

	if query.SubSubAreaID != 0 {
		q = q.Where("id IN (SELECT post_id FROM post_category WHERE category_id = ?)", query.SubSubAreaID)
	}

	if query.MetasearchAreaID != 0 {
		q = q.Where("id IN (SELECT post_id FROM post_category WHERE `category_id` IN (SELECT id FROM category WHERE `metasearch_area_id` = ?))", query.MetasearchAreaID)
	}

	if query.MetasearchSubAreaID != 0 {
		q = q.Where("id IN (SELECT post_id FROM post_category WHERE `category_id` IN (SELECT id FROM category WHERE `metasearch_sub_area_id` = ?))", query.MetasearchSubAreaID)
	}

	if query.MetasearchSubSubAreaID != 0 {
		q = q.Where("id IN (SELECT post_id FROM post_category WHERE `category_id` IN (SELECT id FROM category WHERE `metasearch_sub_sub_area_id` = ?))", query.MetasearchSubSubAreaID)
	}

	if query.InnTypeID != 0 {
		q = q.Where("id IN (SELECT post_id FROM post_category WHERE `category_id` IN (SELECT id FROM category WHERE `metasearch_inn_type_id` = ?))", query.InnTypeID)
	}

	if query.InnDiscerningType != 0 {
		q = q.Where("id IN (SELECT post_id FROM post_category WHERE `category_id` IN (SELECT id FROM category WHERE `metasearch_discerning_condition_id` = ?))", query.InnDiscerningType)
	}

	if query.ThemeID != 0 {
		q = q.Where("id IN (SELECT post_id FROM post_category WHERE category_id = ?)", query.ThemeID)
	}

	if query.HashTag != "" {
		q = q.Where("id IN (SELECT post_id FROM post_hashtag WHERE hashtag_id = (SELECT id FROM hashtag WHERE name = ?))", query.HashTag)
	}

	// TODO: titleに引っかかる物が優先順位が高い、その後body
	if query.Keyward != "" {
		q = q.Where("MATCH(title) AGAINST(?)", query.Keyward).Or("id IN (SELECT post_id FROM post_body WHERE MATCH(body) AGAINST(?))", query.Keyward)
	}

	return q
}

func (r *PostQueryRepositoryImpl) buildFindFeedListQuery(userID int) *gorm.DB {
	q := r.DB

	if userID != 0 {
		q = q.Where("user_id IN (SELECT target_id FROM user_follow WHERE user_id = ?)", userID).Or("id IN (SELECT post_id FROM post_hashtag WHERE hashtag_id IN (SELECT hashtag_id FROM user_follow_hashtag WHERE user_id = ?))", userID)
	}

	return q
}
