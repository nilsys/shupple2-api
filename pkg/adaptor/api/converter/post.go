package converter

import (
	"github.com/stayway-corp/stayway-media-api/pkg/adaptor/api/param"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/model/query"
)

func ConvertFindPostListParamToQuery(param param.FindPostListParam) *query.FindPostListQuery {
	return &query.FindPostListQuery{
		AreaID:       param.AreaID,
		SubAreaID:    param.SubAreaID,
		SubSubAreaID: param.SubSubAreaID,
		ThemeID:      param.ThemeID,
		HashTag:      param.HashTag,
		SortBy:       param.SortBy,
		Limit:        param.GetLimit(),
		OffSet:       param.GetOffSet(),
	}
}
