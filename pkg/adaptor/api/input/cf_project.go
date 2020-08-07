package input

import "github.com/stayway-corp/stayway-media-api/pkg/domain/model"

const (
	defaultCfProjectPerPage = 10
)

type (
	ListCfProject struct {
		AreaID       int                   `query:"areaId"`
		SubAreaID    int                   `query:"subAreaId"`
		SubSubAreaID int                   `query:"subSubAreaId"`
		UserID       int                   `query:"userId"`
		SortBy       model.CfProjectSortBy `query:"sortBy"`
	}

	ListCfProjectSupportComment struct {
		IDParam
		PerPage int `query:"perPage"`
	}
)

// PerPageがクエリで飛んで来なかった場合、デフォルト値である10を返す
func (i *ListCfProjectSupportComment) GetLimit() int {
	if i.PerPage == 0 {
		return defaultCfProjectPerPage
	}
	return i.PerPage
}

func (i *PaginationQuery) GetCfProjectLimit() int {
	if i.PerPage == 0 {
		return defaultCfProjectPerPage
	}
	return i.PerPage
}

func (i *PaginationQuery) GetCfProjectOffset() int {
	if i.Page == 1 || i.Page == 0 {
		return 0
	}
	return i.GetCfProjectLimit() * (i.Page - 1)
}
