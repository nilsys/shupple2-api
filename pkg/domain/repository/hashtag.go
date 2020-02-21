package repository

import "github.com/stayway-corp/stayway-media-api/pkg/domain/entity"

type (
	HashtagQueryRepository interface {
		FindRecommendList(areaID, subAreaID, subSubAreaID int) ([]*entity.Hashtag, error)
		// name部分一致検索
		SearchByName(name string) ([]*entity.Hashtag, error)
	}
)
