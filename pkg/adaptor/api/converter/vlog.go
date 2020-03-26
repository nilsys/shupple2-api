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

func ConvertVlogDetailListToOutput(queryVlogs *entity.VlogDetailList) response.VlogList {
	responseVlogs := make([]*response.Vlog, len(queryVlogs.Vlogs))

	for i, queryVlog := range queryVlogs.Vlogs {
		responseVlogs[i] = convertVlogToOutput(queryVlog)
	}

	return response.VlogList{
		TotalNumber: queryVlogs.TotalNumber,
		Vlogs:       responseVlogs,
	}
}

func ConvertVlogDetailWithTouristSpots(vlog *entity.VlogDetailWithTouristSpots) *response.VlogDetail {
	areaCategories := make([]response.Category, 0, len(vlog.Categories))
	themeCategories := make([]response.Category, 0, len(vlog.Categories))
	touristSpots := make([]response.TouristSpot, len(vlog.TouristSpots))

	for _, category := range vlog.Categories {
		if category.Type.IsAreaKind() {
			areaCategories = append(areaCategories, response.NewCategory(category.ID, category.Name, category.Type))
		} else {
			themeCategories = append(themeCategories, response.NewCategory(category.ID, category.Name, category.Type))
		}
	}

	for i, touristSpot := range vlog.TouristSpots {
		touristSpots[i] = response.NewTouristSpots(touristSpot.ID, touristSpot.Name, touristSpot.Thumbnail)
	}

	return &response.VlogDetail{
		ID:              vlog.ID,
		Thumbnail:       vlog.Thumbnail,
		Title:           vlog.Title,
		Body:            vlog.Body,
		Series:          vlog.Series,
		Creator:         response.NewCreator(vlog.User.ID, vlog.User.UID, vlog.User.GenerateThumbnailURL(), vlog.User.Name, vlog.User.Profile),
		CreatorSNS:      vlog.UserSNS,
		EditorName:      vlog.EditorName,
		YoutubeURL:      vlog.YoutubeURL,
		Views:           vlog.Views,
		PhotoYearMonth:  vlog.YearMonth,
		PlayTime:        vlog.PlayTime,
		Timeline:        vlog.Timeline,
		FacebookCount:   vlog.FacebookCount,
		TwitterCount:    vlog.TwitterCount,
		AreaCategories:  areaCategories,
		ThemeCategories: themeCategories,
		CreatedAt:       model.TimeResponse(vlog.CreatedAt),
		UpdatedAt:       model.TimeResponse(vlog.UpdatedAt),
		TouristSpot:     touristSpots,
	}
}

func convertVlogToOutput(queryVlog *entity.VlogDetail) *response.Vlog {
	var areaCategories []response.Category
	var themeCategories []response.Category

	for _, category := range queryVlog.Categories {
		if category.Type == model.CategoryTypeArea || category.Type == model.CategoryTypeSubArea || category.Type == model.CategoryTypeSubSubArea {
			areaCategories = append(areaCategories, response.NewCategory(category.ID, category.Name, category.Type))
		}
		if category.Type == model.CategoryTypeTheme {
			themeCategories = append(themeCategories, response.NewCategory(category.ID, category.Name, category.Type))
		}
	}

	return &response.Vlog{
		ID:              queryVlog.ID,
		Thumbnail:       queryVlog.Thumbnail,
		AreaCategories:  areaCategories,
		ThemeCategories: themeCategories,
		Title:           queryVlog.VlogTiny.Title,
	}
}
