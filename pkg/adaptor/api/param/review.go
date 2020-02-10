package param

import (
	"github.com/stayway-corp/stayway-media-api/pkg/domain/model"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/model/serror"
)

type ListReviewParams struct {
	UserID        int    `query:"userId"`
	InnID         int    `query:"innId"`
	TouristSpotID int    `query:"touristSpotId"`
	HashTag       string `query:"hashtag"`
	AreaID        int    `query:"areaId"`
	SubAreaID     int    `query:"subAreaId"`
	SubSubAreaID  int    `query:"subSubAreaID"`
	SortBy        string `query:"sortBy"`
	PerPage       int    `query:"perPage"`
	Page          int    `query:"page"`
}

const getReviewsDefaultPerPage = 10

// PerPageがクエリで飛んで来なかった場合、デフォルト値である10を返す
func (param *ListReviewParams) GetLimit() int {
	if param.PerPage == 0 {
		return getReviewsDefaultPerPage
	}
	return param.PerPage
}

// いずれのクエリも飛んで来なかった場合エラーを返す
func (param *ListReviewParams) Validate() error {
	if param.UserID == 0 && param.InnID == 0 && param.TouristSpotID == 0 && param.HashTag == "" && param.AreaID == 0 && param.SubAreaID == 0 && param.SubSubAreaID == 0 {
		return serror.New(nil, serror.CodeInvalidParam, "Invalid show review param")
	}

	if param.SortBy != "" {
		if _, err := model.ParseSortBy(param.SortBy); err != nil {
			return serror.New(err, serror.CodeInvalidParam, "Invalid show review list sortBy")
		}
	}

	return nil
}

// offsetを返す(sqlで使う想定)
func (param *ListReviewParams) GetOffset() int {
	if param.Page == 1 || param.Page == 0 {
		return 0
	}
	return param.GetLimit()*(param.Page-1) + 1
}
