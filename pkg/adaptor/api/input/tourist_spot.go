package input

import (
	"github.com/stayway-corp/stayway-media-api/pkg/domain/model/serror"
)

type (
	ListTouristSpotParams struct {
		AreaID         int   `query:"areaId"`
		SubAreaID      int   `query:"subAreaId"`
		SubSubAreaID   int   `query:"subSubAreaId"`
		SpotCategoryID int   `query:"spotCategoryId"`
		ExcludeSpotIDs []int `query:"excludeSpotId"`
		PerPage        int   `query:"perPage"`
		Page           int   `query:"page"`
	}

	ShowTouristSpotParam struct {
		ID int `param:"id" validate:"required"`
	}

	ListRecommendTouristSpotParam struct {
		ID                    int     `query:"touristSpotId"`
		TouristSpotCategoryID int     `query:"spotCategoryId"`
		Lat                   float64 `query:"lat"`
		Lng                   float64 `query:"lng"`
		PerPage               int     `query:"perPage"`
		Page                  int     `query:"page"`
	}
)

const findTouristSpotListDefaultPerPage = 30

// PerPageがクエリで飛んで来なかった場合、デフォルト値である30を返す
func (param *ListTouristSpotParams) GetLimit() int {
	if param.PerPage == 0 {
		return findTouristSpotListDefaultPerPage
	}
	return param.PerPage
}

// いずれのクエリも飛んで来なかった場合エラーを返す
func (param *ListTouristSpotParams) Validate() error {
	if param.AreaID == 0 && param.SubAreaID == 0 && param.SubSubAreaID == 0 && param.SpotCategoryID == 0 {
		return serror.New(nil, serror.CodeInvalidParam, "Invalid show review input")
	}

	return nil
}

// offsetを返す(sqlで使う想定)
func (param *ListTouristSpotParams) GetOffset() int {
	if param.Page == 1 || param.Page == 0 {
		return 0
	}
	return param.GetLimit() * (param.Page - 1)
}

func (param *ListRecommendTouristSpotParam) GetLimit() int {
	if param.PerPage == 0 {
		return findTouristSpotListDefaultPerPage
	}
	return param.PerPage
}

func (param *ListRecommendTouristSpotParam) GetOffset() int {
	if param.Page == 1 || param.Page == 0 {
		return 0
	}
	return param.GetLimit() * (param.Page - 1)
}
