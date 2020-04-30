package input

type (
	// おすすめHashTag一覧
	ListRecommendHashtagParam struct {
		AreaID       int `json:"areaId"`
		SubAreaID    int `json:"subAreaId"`
		SubSubAreaID int `json:"subSubAreaId"`
		PerPage      int `json:"perPage"`
	}
)

const listHashtagDefaultPerPage = 10

func (p *ListRecommendHashtagParam) GetLimit() int {
	if p.PerPage == 0 {
		return listHashtagDefaultPerPage
	}
	return p.PerPage
}
