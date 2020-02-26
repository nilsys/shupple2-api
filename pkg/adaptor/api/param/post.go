package param

import (
	"github.com/stayway-corp/stayway-media-api/pkg/domain/model"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/model/serror"
)

type (
	GetPost struct {
		ID int `param:"id" validate:"required"`
	}

	StorePostParam struct {
		Title string
		Body  string
	}

	// 記事一覧取得APIパラメータ
	ListPostParam struct {
		UserID       int               `query:"userId"`
		AreaID       int               `query:"areaId"`
		SubAreaID    int               `query:"subAreaId"`
		SubSubAreaID int               `query:"subSubAreaId"`
		ThemeID      int               `query:"themeId"`
		HashTag      string            `query:"hashTag"`
		SortBy       model.MediaSortBy `query:"sortBy"`
		Page         int               `query:"page"`
		PerPage      int               `query:"perPage"`
	}

	// ユーザーフィード記事取得APIパラメータ
	ListFeedPostParam struct {
		UserID  int `param:"id" validate:"required"`
		Page    int `query:"page"`
		PerPage int `query:"perPage"`
	}
)

const findPostListDefaultPerPage = 10

// いずれのクエリも飛んでこない場合 or sortの値が期待値以外の場合エラーを返す
func (param ListPostParam) Validate() error {
	if param.UserID == 0 && param.AreaID == 0 && param.SubAreaID == 0 && param.SubSubAreaID == 0 && param.ThemeID == 0 && param.HashTag == "" {
		return serror.New(nil, serror.CodeInvalidParam, "Invalid find post list param")
	}

	return nil
}

// PerPageがクエリで飛んで来なかった場合、デフォルト値である10を返す
func (param ListPostParam) GetLimit() int {
	if param.PerPage == 0 {
		return findPostListDefaultPerPage
	}
	return param.PerPage
}

// offSetを返す(sqlで使う想定)
func (showListParam ListPostParam) GetOffSet() int {
	if showListParam.Page == 1 || showListParam.Page == 0 {
		return 0
	}
	return showListParam.GetLimit()*(showListParam.Page-1) + 1
}

// PerPageがクエリで飛んで来なかった場合、デフォルト値である10を返す
func (param ListFeedPostParam) GetLimit() int {
	if param.PerPage == 0 {
		return findPostListDefaultPerPage
	}
	return param.PerPage
}

// offSetを返す(sqlで使う想定)
func (param ListFeedPostParam) GetOffSet() int {
	if param.Page == 1 || param.Page == 0 {
		return 0
	}
	return param.GetLimit()*(param.Page-1) + 1
}
