package param

type (
	GetPost struct {
		ID int `param:"id" validate:"required"`
	}

	StorePostParam struct {
		Title string
		Body  string
	}

	// 記事取得API
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

func (findPostListParam FindPostListParam) Validate() error {
	if findPostListParam.AreaID == 0 && findPostListParam.SubAreaID == 0 && findPostListParam.SubSubAreaID == 0 && findPostListParam.ThemeID == 0 && findPostListParam.HashTag == "" {
		return
	}
}
