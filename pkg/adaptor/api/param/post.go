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
	ShowPostListParam struct {
		AreaID       int    `query:"areaId"`
		SubAreaID    int    `query:"subAreaId"`
		SubSubAreaID int    `query:"subSubAreaId"`
		ThemeID      int    `query:"themeId"`
		HashTag      string `query:"hashTag"`
		SortBy       string `query:"sortBy"`
		Page         int    `query:"page"`
		PerPage      int    `query:"perPage"`
	}
)

const findPostListDefaultPerPage = 10

// いずれのクエリも飛んでこない場合 or sortの値が期待値以外の場合エラーを返す
func (showListParam ShowPostListParam) Validate() error {
	if showListParam.AreaID == 0 && showListParam.SubAreaID == 0 && showListParam.SubSubAreaID == 0 && showListParam.ThemeID == 0 && showListParam.HashTag == "" {
		return serror.New(nil, serror.CodeInvalidParam, "Invalid find post list param")
	}

	if showListParam.SortBy != "" {
		if showListParam.SortBy != model.RANKING.Value() || showListParam.SortBy != model.NEW.Value() {
			return serror.New(nil, serror.CodeInvalidParam, "Invalid find post list sortBy")
		}
	}

	return nil
}

// PerPageがクエリで飛んで来なかった場合、デフォルト値である10を返す
func (showListParam ShowPostListParam) GetLimit() int {
	if showListParam.PerPage == 0 {
		return findPostListDefaultPerPage
	}
	return showListParam.PerPage
}

// offSetを返す(sqlで使う想定)
func (showListParam ShowPostListParam) GetOffSet() int {
	if showListParam.Page == 1 || showListParam.Page == 0 {
		return 0
	}
	return showListParam.GetLimit()*(showListParam.Page-1) + 1
}
