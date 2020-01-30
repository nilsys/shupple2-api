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
	FindPostListParam struct {
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
func (findPostListParam FindPostListParam) Validate() error {
	if findPostListParam.AreaID == 0 && findPostListParam.SubAreaID == 0 && findPostListParam.SubSubAreaID == 0 && findPostListParam.ThemeID == 0 && findPostListParam.HashTag == "" {
		return serror.New(nil, serror.CodeInvalidParam, "Invalid find post list param")
	}

	if findPostListParam.SortBy != model.RANKING.Value() || findPostListParam.SortBy != model.NEW.Value() {
		return serror.New(nil, serror.CodeInvalidParam, "Invalid find post list sortBy")
	}

	return nil
}

// PerPageがクエリで飛んで来なかった場合、デフォルト値である10を返す
func (findPostListParam FindPostListParam) GetPerPage() int {
	if findPostListParam.PerPage == 0 {
		return findPostListDefaultPerPage
	}
	return findPostListParam.PerPage
}

// offSetを返す(sqlで使う想定)
func (findPostListParam FindPostListParam) GetOffSet() int {
	if findPostListParam.Page == 1 || findPostListParam.Page == 0 {
		return 0
	}
	return findPostListParam.GetPerPage()*(findPostListParam.Page-1) + 1
}
