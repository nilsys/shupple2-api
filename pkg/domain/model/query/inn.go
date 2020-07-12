package query

import (
	"strconv"

	"github.com/stayway-corp/stayway-media-api/pkg/domain/entity"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/model"
)

type (
	FindInn struct {
		MetasearchAreaID       int
		MetasearchSubAreaID    int
		MetasearchSubSubAreaID int
		Longitude              float64
		Latitude               float64
	}
)

func (q *FindInn) SetMetaserachID(metasearchAreas []*entity.MetasearchArea) {
	for _, metasearchArea := range metasearchAreas {
		switch metasearchArea.MetasearchAreaType {
		case model.AreaCategoryTypeArea:
			q.MetasearchAreaID = metasearchArea.MetasearchAreaID
		case model.AreaCategoryTypeSubArea:
			q.MetasearchSubAreaID = metasearchArea.MetasearchAreaID
		case model.AreaCategoryTypeSubSubArea:
			q.MetasearchSubSubAreaID = metasearchArea.MetasearchAreaID
		}
	}

	// 広めに検索したいので、一番上のレイヤーのみセットする
	// FYI: 宿検索APIは複数のレイヤーのエリアを指定すると一番下のレイヤーのみを採用する
	if q.MetasearchAreaID != 0 {
		q.MetasearchSubAreaID = 0
		q.MetasearchSubSubAreaID = 0
	} else if q.MetasearchSubAreaID != 0 {
		q.MetasearchSubSubAreaID = 0
	}
}

func (q *FindInn) GetGeoCode() string {
	return strconv.FormatFloat(q.Longitude, 'f', -1, 64) + "," + strconv.FormatFloat(q.Latitude, 'f', -1, 64)
}
