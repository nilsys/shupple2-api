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

func (c Converters) ConvertTouristSpotToOutput(touristSpots []*entity.TouristSpot) []*output.TouristSpot {
	responseTouristSpots := make([]*output.TouristSpot, len(touristSpots))

	for i, touristSpot := range touristSpots {
		responseTouristSpots[i] = c.convertTouristSpotToOutput(touristSpot)
	}

	return responseTouristSpots
}

// outputの構造体へconvert
func (c Converters) convertTouristSpotToOutput(touristSpot *entity.TouristSpot) *output.TouristSpot {
	return &output.TouristSpot{
		ID:        touristSpot.ID,
		Name:      touristSpot.Name,
		Thumbnail: touristSpot.Thumbnail,
		URL:       touristSpot.WebsiteURL,
	}
}

func (c Converters) ConvertQueryTouristSpotToOutput(queryTouristSpot *entity.TouristSpotDetail) *output.ShowTouristSpot {
	spotCategories := make([]*output.SpotCategory, len(queryTouristSpot.SpotCategories))

	for i, spotCategory := range queryTouristSpot.SpotCategories {
		spotCategories[i] = output.NewSpotCategory(spotCategory.ID, spotCategory.Name, spotCategory.Slug)
	}

	return &output.ShowTouristSpot{
		ID:              queryTouristSpot.ID,
		Slug:            queryTouristSpot.Slug,
		Name:            queryTouristSpot.Name,
		Thumbnail:       queryTouristSpot.Thumbnail,
		WebsiteURL:      queryTouristSpot.WebsiteURL,
		City:            queryTouristSpot.City,
		Address:         queryTouristSpot.Address,
		Latitude:        queryTouristSpot.Lat,
		Longitude:       queryTouristSpot.Lng,
		AccessCar:       queryTouristSpot.AccessCar,
		AccessTrain:     queryTouristSpot.AccessTrain,
		AccessBus:       queryTouristSpot.AccessBus,
		OpeningHours:    queryTouristSpot.OpeningHours,
		Tel:             queryTouristSpot.TEL,
		Price:           queryTouristSpot.Price,
		InstagramURL:    queryTouristSpot.InstagramURL,
		SearchInnURL:    queryTouristSpot.SearchInnURL,
		Rate:            queryTouristSpot.Rate,
		VendorRate:      queryTouristSpot.VendorRate,
		ReviewCount:     queryTouristSpot.ReviewCount,
		AreaCategories:  c.ConvertAreaCategoriesToOutput(queryTouristSpot.AreaCategories),
		ThemeCategories: c.ConvertThemeCategoriesToOutput(queryTouristSpot.ThemeCategories),
		SpotCategories:  spotCategories,
		CreatedAt:       model.TimeResponse(queryTouristSpot.CreatedAt),
		UpdatedAt:       model.TimeResponse(queryTouristSpot.UpdatedAt),
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
