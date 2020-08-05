package query

import "github.com/stayway-corp/stayway-media-api/pkg/domain/model"

type (
	FindCfProjectQuery struct {
		AreaID       int
		SubAreaID    int
		SubSubAreaID int
		UserID       int
		SortBy       model.CfProjectSortBy
	}
)
