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

func (c Converters) ConvertVlogListToOutput(queryVlogs *entity.VlogList, areaCategories map[int]*entity.AreaCategory, themeCategories map[int]*entity.ThemeCategory) output.VlogList {
	responseVlogs := make([]*output.Vlog, len(queryVlogs.Vlogs))

	for i, queryVlog := range queryVlogs.Vlogs {
		responseVlogs[i] = c.convertVlogToOutput(queryVlog, areaCategories, themeCategories)
	}

	return output.VlogList{
		TotalNumber: queryVlogs.TotalNumber,
		Vlogs:       responseVlogs,
	}
}

func (c Converters) ConvertVlogDetail(vlog *entity.VlogDetail, touristSpotReviewCount map[int]*entity.TouristSpotReviewCount, areaCategories map[int]*entity.AreaCategory, themeCategories map[int]*entity.ThemeCategory, idIsFollowMap map[int]bool) *output.VlogDetail {
	touristSpots := make([]*output.TouristSpotTiny, len(vlog.TouristSpots))
	for i, touristSpot := range vlog.TouristSpots {
		touristSpots[i] = output.NewTouristSpotTinyFromEntity(touristSpot, touristSpotReviewCount[touristSpot.ID].ReviewCount)
	}

	editors := make([]*output.Creator, len(vlog.Editors))
	for i, editor := range vlog.Editors {
		e := c.NewCreatorFromUser(editor, idIsFollowMap[editor.ID])
		editors[i] = &e
	}

	areaCategoriesRes := make([]*output.AreaCategoryDetail, len(vlog.AreaCategories))
	for i, areaCate := range vlog.AreaCategories {
		areaCategoriesRes[i] = c.ConvertAreaCategoryDetailFromAreaCategory(areaCate, areaCategories)
	}

	themeCategoriesRes := make([]*output.ThemeCategoryDetail, len(vlog.ThemeCategories))
	for i, themeCate := range vlog.ThemeCategories {
		themeCategoriesRes[i] = c.ConvertThemeCategoryDetailFromThemeCategory(themeCate, themeCategories)
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
		FavoriteCount:   vlog.FavoriteCount,
		IsFavorite:      vlog.IsFavorite,
		Creator:         c.NewCreatorFromUser(vlog.User, idIsFollowMap[vlog.UserID]),
		Editors:         editors,
		AreaCategories:  areaCategoriesRes,
		ThemeCategories: themeCategoriesRes,
		TouristSpot:     touristSpots,
		EditedAt:        model.TimeResponse(vlog.EditedAt),
		CreatedAt:       model.TimeResponse(vlog.CreatedAt),
		UpdatedAt:       model.TimeResponse(vlog.UpdatedAt),
	}
}

func (c Converters) convertVlogToOutput(vlog *entity.VlogForList, areaCategories map[int]*entity.AreaCategory, themeCategories map[int]*entity.ThemeCategory) *output.Vlog {
	areaCategoriesRes := make([]*output.AreaCategoryDetail, len(vlog.AreaCategories))
	for i, areaCate := range vlog.AreaCategories {
		areaCategoriesRes[i] = c.ConvertAreaCategoryDetailFromAreaCategory(areaCate, areaCategories)
	}

	themeCategoriesRes := make([]*output.ThemeCategoryDetail, len(vlog.ThemeCategories))
	for i, themeCate := range vlog.ThemeCategories {
		themeCategoriesRes[i] = c.ConvertThemeCategoryDetailFromThemeCategory(themeCate, themeCategories)
	}

	return &output.Vlog{
		ID:              vlog.ID,
		Thumbnail:       vlog.Thumbnail,
		AreaCategories:  areaCategoriesRes,
		ThemeCategories: themeCategoriesRes,
		Title:           vlog.VlogTiny.Title,
		FavoriteCount:   vlog.FavoriteCount,
		IsFavorite:      vlog.IsFavorite,
	}
}
