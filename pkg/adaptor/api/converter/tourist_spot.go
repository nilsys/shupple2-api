package converter

import (
	"github.com/stayway-corp/stayway-media-api/pkg/adaptor/api/param"
	"github.com/stayway-corp/stayway-media-api/pkg/adaptor/api/response"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/entity"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/model"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/model/query"
)

// i/oの構造体からレポジトリで使用するクエリ発行用構造体へコンバート
func ConvertTouristSpotListParamToQuery(param *param.ListTouristSpotParams) *query.FindTouristSpotListQuery {
	return &query.FindTouristSpotListQuery{
		AreaID:         param.AreaID,
		SubAreaID:      param.SubAreaID,
		SubSubAreaID:   param.SubSubAreaID,
		SpotCategoryID: param.TouristSpotCategoryID,
		ExcludeSpotIDs: param.ExcludeSpotIDs,
		Limit:          param.GetLimit(),
		OffSet:         param.GetOffset(),
	}
}

func ConvertTouristSpotToOutput(touristSpots []*entity.TouristSpot) []*response.TouristSpot {
	responseTouristSpots := make([]*response.TouristSpot, len(touristSpots))

	for i, touristSpot := range touristSpots {
		responseTouristSpots[i] = convertTouristSpotToOutput(touristSpot)
	}

	return responseTouristSpots
}

// outputの構造体へconvert
func convertTouristSpotToOutput(touristSpot *entity.TouristSpot) *response.TouristSpot {
	return &response.TouristSpot{
		ID:        touristSpot.ID,
		Name:      touristSpot.Name,
		Thumbnail: touristSpot.Thumbnail,
		URL:       touristSpot.WebsiteURL,
	}
}

func ConvertQueryTouristSpotToOutput(queryTouristSpot *entity.QueryTouristSpot) *response.ShowTouristSpot {
	lcategories := make([]*response.Lcategory, len(queryTouristSpot.Lcategories))

	for i, lcategory := range queryTouristSpot.Lcategories {
		lcategories[i] = response.NewLcategory(lcategory.ID, lcategory.Name)
	}

	return &response.ShowTouristSpot{
		ID:                      queryTouristSpot.ID,
		Slug:                    queryTouristSpot.Slug,
		Name:                    queryTouristSpot.Name,
		Thumbnail:               queryTouristSpot.Thumbnail,
		WebsiteURL:              queryTouristSpot.WebsiteURL,
		City:                    queryTouristSpot.City,
		Address:                 queryTouristSpot.Address,
		Latitude:                queryTouristSpot.Lat,
		Longitude:               queryTouristSpot.Lng,
		AccessCar:               queryTouristSpot.AccessCar,
		AccessTrain:             queryTouristSpot.AccessTrain,
		AccessBus:               queryTouristSpot.AccessBus,
		OpeningHours:            queryTouristSpot.OpeningHours,
		Tel:                     queryTouristSpot.TEL,
		Price:                   queryTouristSpot.Price,
		InstagramURL:            queryTouristSpot.InstagramURL,
		SearchInnURL:            queryTouristSpot.SearchInnURL,
		Rate:                    queryTouristSpot.Rate,
		VendorRate:              queryTouristSpot.VendorRate,
		AreaCategories:          ConvertAreaCategoriesToOutput(queryTouristSpot.AreaCategories),
		ThemeCategoryCategories: ConvertThemeCategoriesToOutput(queryTouristSpot.ThemeCategories),
		Lcategories:             lcategories,
		CreatedAt:               model.TimeResponse(queryTouristSpot.CreatedAt),
		UpdatedAt:               model.TimeResponse(queryTouristSpot.UpdatedAt),
	}
}

func ConvertRecommendTouristSpotListParamToQuery(param *param.ListRecommendTouristSpotParam) *query.FindRecommendTouristSpotListQuery {
	return &query.FindRecommendTouristSpotListQuery{
		ID:                    param.ID,
		TouristSpotCategoryID: param.TouristSpotCategoryID,
		Limit:                 param.GetLimit(),
		OffSet:                param.GetOffset(),
	}
}
