package repository

import (
	"github.com/google/wire"
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/entity"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/repository"
)

type HashtagQueryRepositoryImpl struct {
	DB *gorm.DB
}

var HashtagQueryRepositorySet = wire.NewSet(
	wire.Struct(new(HashtagQueryRepositoryImpl), "*"),
	wire.Bind(new(repository.HashtagQueryRepository), new(*HashtagQueryRepositoryImpl)),
)

func (r *HashtagQueryRepositoryImpl) FindByNames(names []string) (map[string]*entity.Hashtag, error) {
	var rows []*entity.Hashtag

	if err := r.DB.Where("name in (?)", names).Find(&rows).Error; err != nil {
		return nil, errors.Wrap(err, "failed to find hashtag by names")
	}

	result := make(map[string]*entity.Hashtag, len(rows))
	for _, row := range rows {
		result[row.Name] = row
	}

	return result, nil
}

func (r *HashtagQueryRepositoryImpl) FindRecommendList(areaID, subAreaID, subSubAreaID int) ([]*entity.Hashtag, error) {
	var rows []*entity.Hashtag

	q := r.buildFindRecommendListQuery(areaID, subAreaID, subSubAreaID)

	if err := q.
		Order("score DESC").
		Find(&rows).
		Limit(defaultAcquisitionNumber).
		Error; err != nil {
		return nil, errors.Wrapf(err, "failed to find get recommend reviews")
	}

	return rows, nil
}

func (r *HashtagQueryRepositoryImpl) SearchByName(name string) ([]*entity.Hashtag, error) {
	var rows []*entity.Hashtag

	if err := r.DB.Where("MATCH(name) AGAINST(?)", name).Limit(defaultSearchSuggestionsNumber).Find(&rows).Error; err != nil {
		return nil, errors.Wrap(err, "failed to find hashtag list by like name")
	}

	return rows, nil
}

func (r *HashtagQueryRepositoryImpl) buildFindRecommendListQuery(areaID, subAreaID, subSubAreaID int) *gorm.DB {
	q := r.DB

	if areaID != 0 {
		q = q.Where("id IN (SELECT id FROM hashtag_category WHERE category_id = ?)", areaID)
	}
	if subAreaID != 0 {
		q = q.Where("id IN (SELECT id FROM hashtag_category WHERE category_id = ?)", subAreaID)
	}
	if subSubAreaID != 0 {
		q = q.Where("id IN (SELECT id FROM hashtag_category WHERE category_id = ?)", subSubAreaID)
	}

	return q
}
