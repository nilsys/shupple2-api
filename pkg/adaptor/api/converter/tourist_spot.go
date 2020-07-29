package converter

import (
	"github.com/stayway-corp/stayway-media-api/pkg/adaptor/api/input"
	"github.com/stayway-corp/stayway-media-api/pkg/adaptor/api/output"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/entity"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/model"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/model/query"
)

// i/oの構造体からレポジトリで使用するクエリ発行用構造体へコンバート
func (c Converters) ConvertTouristSpotListParamToQuery(param *input.ListTouristSpotParams) *query.FindTouristSpotListQuery {
	return &query.FindTouristSpotListQuery{
		AreaID:         param.AreaID,
		SubAreaID:      param.SubAreaID,
		SubSubAreaID:   param.SubSubAreaID,
		SpotCategoryID: param.SpotCategoryID,
		ExcludeSpotIDs: param.ExcludeSpotIDs,
		Limit:          param.GetLimit(),
		OffSet:         param.GetOffset(),
	}
}

func (c Converters) ConvertTouristSpotListToOutput(touristSpots *entity.TouristSpotList, areaCategories map[int]*entity.AreaCategory, themeCategories map[int]*entity.ThemeCategory) *output.TouristSpotList {
	responseTouristSpots := make([]*output.TouristSpot, len(touristSpots.TouristSpots))

	for i, touristSpot := range touristSpots.TouristSpots {
		responseTouristSpots[i] = c.ConvertTouristSpotToOutput(touristSpot, areaCategories, themeCategories)
	}

	return &output.TouristSpotList{
		TotalNumber:  touristSpots.TotalNumber,
		TouristSpots: responseTouristSpots,
	}
}

func (c Converters) ConvertTouristSpotToOutput(touristSpot *entity.TouristSpotDetail, areaCategories map[int]*entity.AreaCategory, themeCategories map[int]*entity.ThemeCategory) *output.TouristSpot {
	thumbnail := touristSpot.Thumbnail
	if touristSpot.HasAlternativeThumbnail() {
		thumbnail = touristSpot.AlternativeImageURL(c.filesURL())
	}

	areaCategoriesRes := make([]*output.AreaCategoryDetail, len(touristSpot.AreaCategories))
	for i, areaCate := range touristSpot.AreaCategories {
		areaCategoriesRes[i] = c.ConvertAreaCategoryDetailFromAreaCategory(areaCate, areaCategories)
	}

	themeCategoriesRes := make([]*output.ThemeCategoryDetail, len(touristSpot.ThemeCategories))
	for i, themeCate := range touristSpot.ThemeCategories {
		themeCategoriesRes[i] = c.ConvertThemeCategoryDetailFromThemeCategory(themeCate, themeCategories)
	}

	spotCategories := make([]*output.SpotCategory, len(touristSpot.SpotCategories))
	for i, spotCategory := range touristSpot.SpotCategories {
		spotCategories[i] = output.NewSpotCategory(spotCategory.ID, spotCategory.Name, spotCategory.Slug)
	}

	return &output.TouristSpot{
		ID:              touristSpot.ID,
		Slug:            touristSpot.Slug,
		Name:            touristSpot.Name,
		Thumbnail:       thumbnail,
		URL:             touristSpot.WebsiteURL,
		City:            touristSpot.City,
		Address:         touristSpot.Address,
		Latitude:        touristSpot.Lat,
		Longitude:       touristSpot.Lng,
		AccessCar:       touristSpot.AccessCar,
		AccessTrain:     touristSpot.AccessTrain,
		AccessBus:       touristSpot.AccessBus,
		OpeningHours:    touristSpot.OpeningHours,
		Tel:             touristSpot.TEL,
		Price:           touristSpot.Price,
		InstagramURL:    touristSpot.InstagramURL,
		SearchInnURL:    touristSpot.SearchInnURL,
		Rate:            touristSpot.Rate,
		VendorRate:      touristSpot.VendorRate,
		ReviewCount:     touristSpot.ReviewCount,
		AreaCategories:  areaCategoriesRes,
		ThemeCategories: themeCategoriesRes,
		SpotCategories:  spotCategories,
		EditedAt:        model.TimeResponse(touristSpot.EditedAt),
		CreatedAt:       model.TimeResponse(touristSpot.CreatedAt),
		UpdatedAt:       model.TimeResponse(touristSpot.UpdatedAt),
	}
}

func (c Converters) ConvertRecommendTouristSpotListParamToQuery(param *input.ListRecommendTouristSpotParam) *query.FindRecommendTouristSpotListQuery {
	return &query.FindRecommendTouristSpotListQuery{
		ID:                    param.ID,
		TouristSpotCategoryID: param.TouristSpotCategoryID,
		Limit:                 param.GetLimit(),
		OffSet:                param.GetOffset(),
	}
}
