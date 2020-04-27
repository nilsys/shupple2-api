package input

type (
	// おすすめHashTag一覧
	ListRecommendHashTagParam struct {
		AreaID       int `json:"areaId"`
		SubAreaID    int `json:"subAreaId"`
		SubSubAreaID int `json:"subSubAreaId"`
	}
)
