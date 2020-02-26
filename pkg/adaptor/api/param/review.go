package param

import (
	"github.com/stayway-corp/stayway-media-api/pkg/domain/model"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/model/serror"
)

type (
	ListReviewParams struct {
		UserID        int               `query:"userId"`
		InnID         int               `query:"innId"`
		TouristSpotID int               `query:"touristSpotId"`
		HashTag       string            `query:"hashtag"`
		AreaID        int               `query:"areaId"`
		SubAreaID     int               `query:"subAreaId"`
		SubSubAreaID  int               `query:"subSubAreaID"`
		SortBy        model.MediaSortBy `query:"sortBy"`
		PerPage       int               `query:"perPage"`
		Page          int               `query:"page"`
	}

	ListFeedReviewParam struct {
		ID      int `param:"id" validate:"required"`
		Page    int `query:"page"`
		PerPage int `query:"perPage"`
	}
)

const getReviewsDefaultPerPage = 10

// いずれのクエリも飛んで来なかった場合エラーを返す
func (param *ListReviewParams) Validate() error {
	if param.UserID == 0 && param.InnID == 0 && param.TouristSpotID == 0 && param.HashTag == "" && param.AreaID == 0 && param.SubAreaID == 0 && param.SubSubAreaID == 0 {
		return serror.New(nil, serror.CodeInvalidParam, "Invalid show review param")
	}

	return nil
}

// PerPageがクエリで飛んで来なかった場合、デフォルト値である10を返す
func (param *ListReviewParams) GetLimit() int {
	if param.PerPage == 0 {
		return getReviewsDefaultPerPage
	}
	return param.PerPage
}

// offsetを返す(sqlで使う想定)
func (param *ListReviewParams) GetOffset() int {
	if param.Page == 1 || param.Page == 0 {
		return 0
	}
	return param.GetLimit()*(param.Page-1) + 1
}

// PerPageがクエリで飛んで来なかった場合、デフォルト値である10を返す
func (param *ListFeedReviewParam) GetLimit() int {
	if param.PerPage == 0 {
		return getReviewsDefaultPerPage
	}
	return param.PerPage
}

// offsetを返す(sqlで使う想定)
func (param *ListFeedReviewParam) GetOffset() int {
	if param.Page == 1 || param.Page == 0 {
		return 0
	}
	return param.GetLimit()*(param.Page-1) + 1
}
