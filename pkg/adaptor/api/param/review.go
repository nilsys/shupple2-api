package param

import (
	"unicode/utf8"

	"github.com/stayway-corp/stayway-media-api/pkg/domain/model"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/model/serror"
)

type (
	ListReviewParams struct {
		UserID                 int                `query:"userId"`
		InnID                  int                `query:"innId"`
		TouristSpotID          int                `query:"touristSpotId"`
		HashTag                string             `query:"hashtag"`
		AreaID                 int                `query:"areaId"`
		SubAreaID              int                `query:"subAreaId"`
		SubSubAreaID           int                `query:"subSubAreaID"`
		MetasearchAreaID       int                `query:"metasearchAreaId"`
		MetasearchSubAreaID    int                `query:"metasearchSubAreaId"`
		MetasearchSubSubAreaID int                `query:"metasearchSubSubAreaId"`
		ChildID                int                `query:"childId"`
		SortBy                 model.ReviewSortBy `query:"sortBy"`
		Keyward                string             `query:"q"`
		PerPage                int                `query:"perPage"`
		Page                   int                `query:"page"`
	}

	ListFeedReviewParam struct {
		UserID  int `param:"id" validate:"required"`
		Page    int `query:"page"`
		PerPage int `query:"perPage"`
	}

	StoreFavoriteReviewParam struct {
		ReviewID int `param:"id" validate:"required"`
	}

	StoreReviewParam struct {
		TravelDate    model.YearMonth        `json:"travelDate" validate:"required"`
		Accompanying  model.AccompanyingType `json:"accompanying" validate:"required"`
		TouristSpotID int                    `json:"touristSpotId"`
		InnID         int                    `json:"innId"`
		Score         int                    `json:"score" validate:"required"`
		Body          string                 `json:"body" validate:"required"`
		MediaUUIDs    []MediasUUID           `json:"mediaUuids"`
	}

	MediasUUID struct {
		UUID     string `json:"uuid"`
		MimeType string `json:"mimeType"`
	}

	ShowReview struct {
		ID int `param:"id" validate:"required"`
	}

	ListReviewCommentParam struct {
		ID      int `param:"id" validate:"required"`
		PerPage int `query:"perPage"`
	}
)

const getReviewsDefaultPerPage = 10

// 投稿内容の最低文字数
const storeBodyMinimumLimit = 50

// いずれのクエリも飛んで来なかった場合エラーを返す
func (param *ListReviewParams) Validate() error {
	if param.UserID == 0 && param.InnID == 0 && param.TouristSpotID == 0 && param.HashTag == "" && param.AreaID == 0 && param.SubAreaID == 0 && param.SubSubAreaID == 0 && param.MetasearchAreaID == 0 && param.MetasearchSubAreaID == 0 && param.MetasearchSubSubAreaID == 0 && param.ChildID == 0 && param.Keyward == "" {
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

func (param *StoreReviewParam) Validate() error {
	if utf8.RuneCountInString(param.Body) < storeBodyMinimumLimit {
		return serror.New(nil, serror.CodeInvalidParam, "Invalid store review body")
	}

	if (param.TouristSpotID != 0 && param.InnID != 0) || (param.TouristSpotID == 0 && param.InnID == 0) {
		return serror.New(nil, serror.CodeInvalidParam, "Invalid store review tourist_spot & inn_id")
	}

	return nil
}

// PerPageがクエリで飛んで来なかった場合、デフォルト値である10を返す
func (param *ListReviewCommentParam) GetLimit() int {
	if param.PerPage == 0 {
		return getReviewsDefaultPerPage
	}
	return param.PerPage
}
