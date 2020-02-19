package converter

import (
	"github.com/stayway-corp/stayway-media-api/pkg/adaptor/api/param"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/model"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/model/query"
)

// i/oの構造体からレポジトリで使用するクエリ発行用構造体へコンバート
func ConvertFindReviewListParamToQuery(param *param.ListReviewParams) *query.ShowReviewListQuery {
	// MEMO: validate後実行される想定
	sortBy, _ := model.ParseSortBy(param.SortBy)
	return &query.ShowReviewListQuery{
		UserID:        param.UserID,
		InnID:         param.InnID,
		TouristSpotID: param.TouristSpotID,
		HashTag:       param.HashTag,
		AreaID:        param.AreaID,
		SubAreaID:     param.SubAreaID,
		SubSubAreaID:  param.SubSubAreaID,
		SortBy:        sortBy,
		Limit:         param.GetLimit(),
		OffSet:        param.GetOffset(),
	}
}

func ConvertListFeedReviewParamToQuery(param *param.ListFeedReviewParam) *query.FindListPaginationQuery {
	return &query.FindListPaginationQuery{
		Limit:  param.GetLimit(),
		Offset: param.GetOffset(),
	}
}
