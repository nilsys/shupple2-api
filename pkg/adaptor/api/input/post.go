package input

import (
	"github.com/stayway-corp/stayway-media-api/pkg/domain/model"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/model/serror"
)

type (
	GetPost struct {
		ID int `param:"id" validate:"required"`
	}

	ShowPostBySlug struct {
		Slug PathString `param:"slug" validate:"required"`
	}

	// 記事一覧取得APIパラメータ
	ListPostParam struct {
		UserID                 int               `query:"userId"`
		AreaID                 int               `query:"areaId"`
		SubAreaID              int               `query:"subAreaId"`
		SubSubAreaID           int               `query:"subSubAreaId"`
		ChildAreaID            int               `query:"childAreaID"`
		ChildSubAreaID         int               `query:"childSubAreaID"`
		ChildSubSubAreaID      int               `query:"childSubSubAreaID"`
		MetasearchAreaID       int               `query:"metasearchAreaId"`
		MetasearchSubAreaID    int               `query:"metasearchSubAreaId"`
		MetasearchSubSubAreaID int               `query:"metasearchSubSubAreaId"`
		InnTypeID              int               `json:"innTypeId"`
		InnDiscerningType      int               `json:"innDiscerningType"`
		ThemeID                int               `query:"themeId"`
		HashTag                string            `query:"hashTag"`
		SortBy                 model.MediaSortBy `query:"sortBy"`
		Keyward                string            `query:"q"`
		Page                   int               `query:"page"`
		PerPage                int               `query:"perPage"`
	}

	// ユーザーフィード記事取得APIパラメータ
	ListFeedPostParam struct {
		UserID  int `param:"id" validate:"required"`
		Page    int `query:"page"`
		PerPage int `query:"perPage"`
	}

	StoreFavoritePostParam struct {
		PostID int `param:"id" validate:"required"`
	}
)

const findPostListDefaultPerPage = 10

// いずれのクエリも飛んでこない場合 or sortの値が期待値以外の場合エラーを返す
func (param ListPostParam) Validate() error {
	if param.UserID == 0 && param.AreaID == 0 && param.SubAreaID == 0 && param.SubSubAreaID == 0 && param.ThemeID == 0 && param.MetasearchAreaID == 0 && param.MetasearchSubAreaID == 0 && param.MetasearchSubSubAreaID == 0 && param.InnTypeID == 0 && param.InnDiscerningType == 0 && param.HashTag == "" && param.Keyward == "" && param.SortBy == 0 {
		return serror.New(nil, serror.CodeInvalidParam, "Invalid find post list input")
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
func (param ListPostParam) GetOffSet() int {
	if param.Page == 1 || param.Page == 0 {
		return 0
	}
	return param.GetLimit() * (param.Page - 1)
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
	return param.GetLimit() * (param.Page - 1)
}
