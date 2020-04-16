package input

import "github.com/stayway-corp/stayway-media-api/pkg/domain/model/serror"

type (
	// おすすめHashTag一覧
	ListRecommendHashTagParam struct {
		AreaID       int `json:"areaId"`
		SubAreaID    int `json:"subAreaId"`
		SubSubAreaID int `json:"subSubAreaId"`
	}
)

func (param ListRecommendHashTagParam) Validate() error {
	if param.AreaID == 0 && param.SubAreaID == 0 && param.SubSubAreaID == 0 {
		return serror.New(nil, serror.CodeInvalidParam, "Invalid find post list input")
	}
	return nil
}
