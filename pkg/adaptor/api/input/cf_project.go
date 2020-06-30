package input

import "github.com/stayway-corp/stayway-media-api/pkg/domain/model"

const (
	defaultCfProjectSupportCommentPerPage = 10
)

type (
	ListCfProject struct {
		AreaID       int                   `query:"areaId"`
		SubAreaID    int                   `query:"subAreaId"`
		SubSubAreaID int                   `query:"subSubAreaId"`
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
		return defaultCfProjectSupportCommentPerPage
	}
	return i.PerPage
}
