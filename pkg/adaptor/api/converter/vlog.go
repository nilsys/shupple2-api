package converter

import (
	"github.com/stayway-corp/stayway-media-api/pkg/adaptor/api/param"
	"github.com/stayway-corp/stayway-media-api/pkg/adaptor/api/response"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/entity"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/model"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/model/query"
)

// i/oの構造体からレポジトリで使用するクエリ構造体へconvert
func ConvertListVlogParamToQuery(param *param.ListVlogParam) *query.FindVlogListQuery {
	return &query.FindVlogListQuery{
		AreaID:        param.AreaID,
		SubAreaID:     param.SubAreaID,
		SubSubAreaID:  param.SubSubAreaID,
		TouristSpotID: param.TouristSpotID,
		SortBy:        param.SortBy,
		Keyward:       param.Keyward,
		Limit:         param.GetLimit(),
		OffSet:        param.GetOffSet(),
	}
}

func ConvertVlogListToOutput(queryVlogs []*entity.QueryVlog) []*response.Vlog {
	responseVlogs := make([]*response.Vlog, len(queryVlogs))

	for i, queryVlog := range queryVlogs {
		responseVlogs[i] = convertVlogToOutput(queryVlog)
	}

	return responseVlogs
}

func convertVlogToOutput(queryVlog *entity.QueryVlog) *response.Vlog {
	var areaCategories []response.Category
	var themeCategories []response.Category

	for _, category := range queryVlog.WordpressCategories {
		if category.Type == model.CategoryTypeArea || category.Type == model.CategoryTypeSubArea || category.Type == model.CategoryTypeSubSubArea {
			areaCategories = append(areaCategories, response.NewCategory(category.ID, category.Name))
		}
		if category.Type == model.CategoryTypeTheme {
			themeCategories = append(themeCategories, response.NewCategory(category.ID, category.Name))
		}
	}

	return &response.Vlog{
		Thumbnail:       queryVlog.GenerateThumbnailURL(),
		AreaCategories:  areaCategories,
		ThemeCategories: themeCategories,
		Title:           queryVlog.VlogTiny.Title,
	}
}
