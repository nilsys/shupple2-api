package query

import "github.com/stayway-corp/stayway-media-api/pkg/domain/model"

type (
	// repositoryで使用する検索用構造体
	FindVlogListQuery struct {
		AreaID        int
		SubAreaID     int
		SubSubAreaID  int
		TouristSpotID int
		SortBy        model.MediaSortBy
		Limit         int
		OffSet        int
	}
)
