package query

import "strconv"

type (
	FindInn struct {
		MetasearchAreaId       int
		MetasearchSubAreaID    int
		MetasearchSubSubAreaID int
		Longitude              float64
		Latitude               float64
	}
)

func (q *FindInn) SetMetaserachID(areaID, subAreaID, subSubAreaID int) {
	q.MetasearchAreaId = areaID
	q.MetasearchSubAreaID = subAreaID
	q.MetasearchSubSubAreaID = subSubAreaID
}

func (q *FindInn) GetGeoCode() string {
	return strconv.FormatFloat(q.Longitude, 'f', -1, 64) + strconv.FormatFloat(q.Latitude, 'f', -1, 64)
}
