package input

type (
	ShowFeatureListParam struct {
		Page    int `query:"page"`
		PerPage int `query:"perPage"`
	}

	ShowFeatureParam struct {
		ID int `param:"id" validate:"required"`
	}
)

const findFeatureListDefaultLimit = 10

// PerPageがクエリで飛んで来なかった場合、デフォルト値である10を返す
func (param ShowFeatureListParam) GetLimit() int {
	if param.PerPage == 0 {
		return findFeatureListDefaultLimit
	}
	return param.PerPage
}

// offSetを返す(sqlで使う想定)
func (param ShowFeatureListParam) GetOffSet() int {
	if param.Page == 1 || param.Page == 0 {
		return 0
	}
	return param.GetLimit()*(param.Page-1) + 1
}
