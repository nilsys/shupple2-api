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
		MetasearchInnTypeID    int               `json:"metasearchInnTypeId"`
		MetasearchTagID        int               `json:"metasearchTagId"`
		ThemeID                int               `query:"themeId"`
		ThemeSlug              string            `query:"themeSlug"`
		HashTag                string            `query:"hashTag"`
		SortBy                 model.MediaSortBy `query:"sortBy"`
		Keyward                string            `query:"q"`
		CfProjectID            int               `query:"cfProjectId"`
		// 特別なフラグが増えてきたらenumを検討
		NotHaveAreaID bool `query:"notHaveAreaId"`
		Page          int  `query:"page"`
		PerPage       int  `query:"perPage"`
	}

	// ユーザーフィード記事取得APIパラメータ
	ListFavoritePostParam struct {
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
	if param.UserID == 0 && param.AreaID == 0 && param.SubAreaID == 0 && param.SubSubAreaID == 0 && param.ThemeID == 0 && param.MetasearchAreaID == 0 && param.MetasearchSubAreaID == 0 && param.MetasearchSubSubAreaID == 0 && param.MetasearchInnTypeID == 0 && param.MetasearchTagID == 0 && param.HashTag == "" && param.Keyward == "" && param.CfProjectID == 0 && param.SortBy == 0 && !param.NotHaveAreaID && param.ThemeSlug == "" {
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
func (param ListFavoritePostParam) GetLimit() int {
	if param.PerPage == 0 {
		return findPostListDefaultPerPage
	}
	return param.PerPage
}

// offSetを返す(sqlで使う想定)
func (param ListFavoritePostParam) GetOffSet() int {
	if param.Page == 1 || param.Page == 0 {
		return 0
	}
	return param.GetLimit() * (param.Page - 1)
}

func (i *PaginationQuery) GetPostLimit() int {
	if i.PerPage == 0 {
		return findPostListDefaultPerPage
	}
	return i.PerPage
}

func (i *PaginationQuery) GetPostOffset() int {
	if i.Page == 1 || i.Page == 0 {
		return 0
	}
	return i.GetPostLimit() * (i.Page - 1)
}
