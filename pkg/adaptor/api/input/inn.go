package input

import "github.com/stayway-corp/stayway-media-api/pkg/domain/model/serror"

type (
	ListInn struct {
		AreaID        int `query:"areaId"`
		SubAreaID     int `query:"subAreaId"`
		SubSubAreaID  int `query:"subSubAreaId"`
		TouristSpotID int `query:"touristSpotId"`
	}
)

func (i ListInn) Validate() error {
	// 1つ以上のパラメータが指定されていた場合
	if (i.AreaID != 0 && i.SubAreaID != 0 && i.SubSubAreaID != 0) || (i.AreaID != 0 && i.SubSubAreaID != 0 && i.TouristSpotID != 0) || (i.SubAreaID != 0 && i.SubSubAreaID != 0 && i.TouristSpotID != 0) || (i.AreaID != 0 && i.SubAreaID != 0 && i.TouristSpotID != 0) {
		return serror.New(nil, serror.CodeInvalidParam, "invalid list inn parameters")
	}

	return nil
}
