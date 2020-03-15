package dto

import (
	"github.com/stayway-corp/stayway-media-api/pkg/domain/entity"
)

// stayway側のAPIレスポンス型
// /api/innsのレスポンス型
type (
	Inns struct {
		Count     int `json:"Count"`
		Rectangle struct {
			LatNE float64 `json:"LatNE"`
			LngNE float64 `json:"LngNE"`
			LatSW float64 `json:"LatSW"`
			LngSW float64 `json:"LngSW"`
		} `json:"Rectangle"`
		Inns []struct {
			ID          int    `json:"ID"`
			Name        string `json:"Name"`
			Description string `json:"Description"`
			URL         string `json:"URL"`
			Brand       struct {
				ID   int    `json:"ID"`
				Name string `json:"Name"`
			} `json:"Brand"`
			Thumbnail    string      `json:"Thumbnail"`
			Latitude     float64     `json:"Latitude"`
			Longitude    float64     `json:"Longitude"`
			CountryCode  string      `json:"CountryCode"`
			Address      string      `json:"Address"`
			PostCode     string      `json:"PostCode"`
			Phone        string      `json:"Phone"`
			Access       string      `json:"Access"`
			StayDaysMin  int         `json:"StayDaysMin"`
			StayDaysMax  int         `json:"StayDaysMax"`
			CheckInTime  string      `json:"CheckInTime"`
			CheckOutTime string      `json:"CheckOutTime"`
			RoomCount    int         `json:"RoomCount"`
			BedCount     int         `json:"BedCount"`
			Score        float64     `json:"Score"`
			ReviewCount  int         `json:"ReviewCount"`
			MinPrice     interface{} `json:"MinPrice"`
			MaxPrice     interface{} `json:"MaxPrice"`
			InnTypes     []struct {
				ID   int    `json:"ID"`
				Name string `json:"Name"`
				URL  string `json:"URL"`
			} `json:"InnTypes"`
			AvailableCards []struct {
				ID   int    `json:"ID"`
				Name string `json:"Name"`
			} `json:"AvailableCards"`
			RoomTypes []struct {
				ID   int    `json:"ID"`
				Name string `json:"Name"`
			} `json:"RoomTypes"`
			TrueTags []struct {
				ID   int    `json:"ID"`
				Name string `json:"Name"`
				URL  string `json:"URL"`
			} `json:"TrueTags"`
			FalseTags []struct {
				ID   int    `json:"ID"`
				Name string `json:"Name"`
				URL  string `json:"URL"`
			} `json:"FalseTags"`
		} `json:"Inns"`
		HasNext bool `json:"HasNext"`
	}

	// /api/inns/:id/areaのレスポンス型
	InnArea struct {
		ID         int
		Name       string
		URL        string
		Area       Area
		SubArea    Area
		SubSubArea Area
	}

	Area struct {
		ID   int
		Name string
	}
)

// InnsからIDだけを抽出
func (inns Inns) InnsToIDs() []int {
	var ids []int
	for _, inn := range inns.Inns {
		ids = append(ids, inn.ID)
	}

	return ids
}

// (エリアID, サブエリアID, サブサブエリアID)順に返る
func (innArea *InnArea) ToInnAreaTypeIDs() *entity.InnAreaTypeIDs {
	return &entity.InnAreaTypeIDs{
		AreaID:       innArea.Area.ID,
		SubAreaID:    innArea.SubArea.ID,
		SubSubAreaID: innArea.SubSubArea.ID,
	}
}
