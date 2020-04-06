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

func ConvertVlogListToOutput(queryVlogs *entity.VlogList) response.VlogList {
	responseVlogs := make([]*response.Vlog, len(queryVlogs.Vlogs))

	for i, queryVlog := range queryVlogs.Vlogs {
		responseVlogs[i] = convertVlogToOutput(queryVlog)
	}

	return response.VlogList{
		TotalNumber: queryVlogs.TotalNumber,
		Vlogs:       responseVlogs,
	}
}

func ConvertVlogDetail(vlog *entity.VlogDetail) *response.VlogDetail {
	touristSpots := make([]*response.TouristSpot, len(vlog.TouristSpots))
	for i, touristSpot := range vlog.TouristSpots {
		touristSpots[i] = response.NewTouristSpots(touristSpot.ID, touristSpot.Name, touristSpot.Thumbnail)
	}

	editors := make([]*response.Creator, len(vlog.Editors))
	for i, editor := range vlog.Editors {
		e := response.NewCreatorFromUser(editor)
		editors[i] = &e
	}

	return &response.VlogDetail{
		ID:              vlog.ID,
		Thumbnail:       vlog.Thumbnail,
		Title:           vlog.Title,
		Body:            vlog.Body,
		Series:          vlog.Series,
		YoutubeURL:      vlog.YoutubeURL,
		Views:           vlog.Views,
		ShootingDate:    vlog.YearMonth,
		PlayTime:        vlog.PlayTime,
		Timeline:        vlog.Timeline,
		FacebookCount:   vlog.FacebookCount,
		TwitterCount:    vlog.TwitterCount,
		Editors:         editors,
		AreaCategories:  ConvertAreaCategoriesToOutput(vlog.AreaCategories),
		ThemeCategories: ConvertThemeCategoriesToOutput(vlog.ThemeCategories),
		TouristSpot:     touristSpots,
		CreatedAt:       model.TimeResponse(vlog.CreatedAt),
		UpdatedAt:       model.TimeResponse(vlog.UpdatedAt),
	}
}

func convertVlogToOutput(queryVlog *entity.VlogForList) *response.Vlog {
	return &response.Vlog{
		ID:              queryVlog.ID,
		Thumbnail:       queryVlog.Thumbnail,
		AreaCategories:  ConvertAreaCategoriesToOutput(queryVlog.AreaCategories),
		ThemeCategories: ConvertThemeCategoriesToOutput(queryVlog.ThemeCategories),
		Title:           queryVlog.VlogTiny.Title,
	}
}
