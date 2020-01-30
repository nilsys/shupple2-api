package repository

import (
	"github.com/google/wire"
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/entity"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/model/query"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/repository"
)

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

func (r *PostQueryRepositoryImpl) FindByParams(query *query.FindPostListQuery) ([]*entity.Post, error) {
	var posts []*entity.Post

	q := r.buildFindByParamsQuery(query)

	if err := q.Preload("Bodies").Preload("CategoryIDs").
		Limit(query.Limit).
		Offset(query.OffSet).
		Find(&posts).Error; err != nil {
		return nil, errors.Wrapf(err, "Failed get posts by params")
	}

	return posts, nil
}

func (r *PostQueryRepositoryImpl) buildFindByParamsQuery(query *query.FindPostListQuery) *gorm.DB {
	q := r.DB.Select("*")

	if query.AreaID != 0 {
		q = q.Where("id IN (SELECT post_id FROM post_category WHERE category_id = ? AND category_type = area", query.AreaID)
	}

	if query.SubAreaID != 0 {
		q = q.Where("id IN (SELECT post_id FROM post_category WHERE category_id = ? AND category_type = sub_area", query.SubAreaID)
	}

	if query.SubSubAreaID != 0 {
		q = q.Where("id IN (SELECT post_id FROM post_category WHERE category_id = ? AND category_type = sub_sub_area", query.SubSubAreaID)
	}

	if query.ThemeID != 0 {
		q = q.Where("id IN (SELECT post_id FROM post_category WHERE category_id = ? AND category_type = theme", query.ThemeID)
	}

	if query.HashTag != "" {
		q = q.Where("id IN (SELECT post_id FROM post_hashtag WHERE hashtag_id = (SELECT id FROM hashtag WHERE name = ?))", query.HashTag)
	}

	return q
}
