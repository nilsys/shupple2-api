package repository

import (
	"github.com/google/wire"
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/entity"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/repository"
)

type HashTagQueryRepositoryImpl struct {
	DB *gorm.DB
}

var HashTagQueryRepositorySet = wire.NewSet(
	wire.Struct(new(HashTagQueryRepositoryImpl), "*"),
	wire.Bind(new(repository.HashTagQueryRepository), new(*HashTagQueryRepositoryImpl)),
)

func (r *HashTagQueryRepositoryImpl) FindRecommendList(areaID, subAreaID, subSubAreaID int) ([]*entity.HashTag, error) {
	var rows []*entity.HashTag

	q := r.buildFindRecommendListQuery(areaID, subAreaID, subSubAreaID)

	if err := q.
		Table("hashtag").
		Select("*, (SELECT COUNT(post_id) FROM post_hashtag WHERE hashtag_id = id) AS post_rank, (SELECT COUNT(review_id) FROM review_hashtag WHERE hashtag_id = id) AS review_rank").
		Order("post_rank + review_rank DESC").Find(&rows).
		Limit(defaultAcquisitionNumber).
		Error; err != nil {
		return nil, errors.Wrapf(err, "failed to find get recommend reviews")
	}

	return rows, nil
}

func (r *HashTagQueryRepositoryImpl) buildFindRecommendListQuery(areaID, subAreaID, subSubAreaID int) *gorm.DB {
	q := r.DB

	if areaID != 0 {
		q = q.Where("id IN (SELECT id FROM category_hashtag WHERE category_id = ?)", areaID)
	}
	if subAreaID != 0 {
		q = q.Where("id IN (SELECT id FROM category_hashtag WHERE category_id = ?)", subAreaID)
	}
	if subSubAreaID != 0 {
		q = q.Where("id IN (SELECT id FROM category_hashtag WHERE category_id = ?)", subSubAreaID)
	}

	return q
}
