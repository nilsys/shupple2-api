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

// TODO: クエリ用のentityに変える
func (r *PostQueryRepositoryImpl) FindByID(id int) (*entity.Post, error) {
	var row entity.Post
	if err := r.DB.First(&row, id).Error; err != nil {
		return nil, ErrorToFindSingleRecord(err, "post(id=%d)", id)
	}
	return &row, nil
}

// 検索条件に指定されたクエリ構造体を用い、postを複数参照
func (r *PostQueryRepositoryImpl) FindListByParams(query *query.FindPostListQuery) ([]*entity.QueryPost, error) {
	var posts []*entity.QueryPost

	q := r.buildFindByParamsQuery(query)

	if err := q.
		Table("post").
		Order(query.SortBy.GetPostOrderQuery()).
		Limit(query.Limit).
		Offset(query.OffSet).
		Find(&posts).Error; err != nil {
		return nil, errors.Wrapf(err, "Failed get posts by params")
	}

	return posts, nil
}

// クエリ構造体を用い、検索クエリを作成
// TODO: category取得にtype付け足す？
func (r *PostQueryRepositoryImpl) buildFindByParamsQuery(query *query.FindPostListQuery) *gorm.DB {
	q := r.DB

	if query.AreaID != 0 {
		q = q.Where("id IN (SELECT post_id FROM post_category WHERE category_id = ?)", query.AreaID)
	}

	if query.SubAreaID != 0 {
		q = q.Where("id IN (SELECT post_id FROM post_category WHERE category_id = ?)", query.SubAreaID)
	}

	if query.SubSubAreaID != 0 {
		q = q.Where("id IN (SELECT post_id FROM post_category WHERE category_id = ?)", query.SubSubAreaID)
	}

	if query.ThemeID != 0 {
		q = q.Where("id IN (SELECT post_id FROM post_category WHERE category_id = ?)", query.ThemeID)
	}

	if query.HashTag != "" {
		q = q.Where("id IN (SELECT post_id FROM post_hashtag WHERE hashtag_id = (SELECT id FROM hashtag WHERE name = ?))", query.HashTag)
	}

	return q
}
