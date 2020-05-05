package converter

import (
	"github.com/stayway-corp/stayway-media-api/pkg/adaptor/api/input"
	"github.com/stayway-corp/stayway-media-api/pkg/adaptor/api/output"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/entity"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/model"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/model/query"
)

// i/oの構造体からレポジトリで使用するクエリ構造体へconvert
func (c Converters) ConvertListVlogParamToQuery(param *input.ListVlogParam) *query.FindVlogListQuery {
	return &query.FindVlogListQuery{
		AreaID:        param.AreaID,
		SubAreaID:     param.SubAreaID,
		SubSubAreaID:  param.SubSubAreaID,
		TouristSpotID: param.TouristSpotID,
		UserID:        param.UserID,
		SortBy:        param.SortBy,
		Keyword:       param.Keyward,
		Limit:         param.GetLimit(),
		OffSet:        param.GetOffSet(),
	}
}

func (c Converters) ConvertVlogListToOutput(queryVlogs *entity.VlogList) output.VlogList {
	responseVlogs := make([]*output.Vlog, len(queryVlogs.Vlogs))

	for i, queryVlog := range queryVlogs.Vlogs {
		responseVlogs[i] = c.convertVlogToOutput(queryVlog)
	}

	return output.VlogList{
		TotalNumber: queryVlogs.TotalNumber,
		Vlogs:       responseVlogs,
	}
}

func (c Converters) ConvertVlogDetail(vlog *entity.VlogDetail) *output.VlogDetail {
	touristSpots := make([]*output.TouristSpot, len(vlog.TouristSpots))
	for i, touristSpot := range vlog.TouristSpots {
		touristSpots[i] = output.NewTouristSpots(touristSpot.ID, touristSpot.Name, touristSpot.Thumbnail)
	}

	editors := make([]*output.Creator, len(vlog.Editors))
	for i, editor := range vlog.Editors {
		e := c.NewCreatorFromUser(editor)
		editors[i] = &e
	}

	return &output.VlogDetail{
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
		Creator:         c.NewCreatorFromUser(vlog.User),
		Editors:         editors,
		AreaCategories:  c.ConvertAreaCategoriesToOutput(vlog.AreaCategories),
		ThemeCategories: c.ConvertThemeCategoriesToOutput(vlog.ThemeCategories),
		TouristSpot:     touristSpots,
		CreatedAt:       model.TimeResponse(vlog.CreatedAt),
		UpdatedAt:       model.TimeResponse(vlog.UpdatedAt),
	}
}

func (c Converters) convertVlogToOutput(queryVlog *entity.VlogForList) *output.Vlog {
	return &output.Vlog{
		ID:              queryVlog.ID,
		Thumbnail:       queryVlog.Thumbnail,
		AreaCategories:  c.ConvertAreaCategoriesToOutput(queryVlog.AreaCategories),
		ThemeCategories: c.ConvertThemeCategoriesToOutput(queryVlog.ThemeCategories),
		Title:           queryVlog.VlogTiny.Title,
	}
}
