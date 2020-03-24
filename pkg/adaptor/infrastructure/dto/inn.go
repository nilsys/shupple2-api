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

// TODO: 構造体毎、配列含めて各々のコンバータ用意する
func (inns Inns) ConvertToEntity() *entity.Inns {
	innsEntity := make([]entity.Inn, len(inns.Inns))
	for i, InnsDto := range inns.Inns {
		brand := entity.InnBrand{
			ID:   InnsDto.Brand.ID,
			Name: InnsDto.Brand.Name,
		}
		innTypes := make([]entity.InnType, len(inns.Inns[i].InnTypes))
		for n, innType := range inns.Inns[i].InnTypes {
			innTypes[n] = entity.InnType{
				ID:   innType.ID,
				Name: innType.Name,
				URL:  innType.URL,
			}
		}
		availableCards := make([]entity.AvailableCard, len(inns.Inns[i].AvailableCards))
		for t, availableCard := range inns.Inns[i].AvailableCards {
			availableCards[t] = entity.AvailableCard{
				ID:   availableCard.ID,
				Name: availableCard.Name,
			}
		}
		roomTypes := make([]entity.RoomType, len(inns.Inns[i].RoomTypes))
		for s, roomType := range inns.Inns[i].RoomTypes {
			roomTypes[s] = entity.RoomType{
				ID:   roomType.ID,
				Name: roomType.Name,
			}
		}
		trueTags := make([]entity.InnTag, len(inns.Inns[i].TrueTags))
		for m, trueTag := range inns.Inns[i].TrueTags {
			trueTags[m] = entity.InnTag{
				ID:   trueTag.ID,
				Name: trueTag.Name,
				URL:  trueTag.URL,
			}
		}
		falseTags := make([]entity.InnTag, len(inns.Inns[i].FalseTags))
		for c, falseTag := range inns.Inns[i].FalseTags {
			falseTags[c] = entity.InnTag{
				ID:   falseTag.ID,
				Name: falseTag.Name,
				URL:  falseTag.URL,
			}
		}
		innsEntity[i] = entity.Inn{
			ID:             InnsDto.ID,
			Name:           InnsDto.Name,
			Description:    InnsDto.Description,
			URL:            InnsDto.URL,
			Brand:          brand,
			Thumbnail:      InnsDto.Thumbnail,
			Latitude:       InnsDto.Latitude,
			Longitude:      InnsDto.Longitude,
			CountryCode:    InnsDto.CountryCode,
			Address:        InnsDto.Address,
			PostCode:       InnsDto.PostCode,
			Phone:          InnsDto.Phone,
			Access:         InnsDto.Access,
			StayDaysMin:    InnsDto.StayDaysMin,
			StayDaysMax:    InnsDto.StayDaysMax,
			CheckInTime:    InnsDto.CheckInTime,
			CheckOutTime:   InnsDto.CheckOutTime,
			RoomCount:      InnsDto.RoomCount,
			BedCount:       InnsDto.BedCount,
			Score:          InnsDto.Score,
			ReviewCount:    InnsDto.ReviewCount,
			MinPrice:       InnsDto.MinPrice,
			MaxPrice:       InnsDto.MaxPrice,
			InnTypes:       innTypes,
			AvailableCards: availableCards,
			RoomTypes:      roomTypes,
			TrueTags:       trueTags,
			FalseTags:      falseTags,
		}
	}
	return &entity.Inns{
		Count: inns.Count,
		Rectangle: struct {
			LatNE float64 `json:"latNE"`
			LngNE float64 `json:"lngNE"`
			LatSW float64 `json:"latSW"`
			LngSW float64 `json:"lngSW"`
		}{LatNE: inns.Rectangle.LatNE, LngNE: inns.Rectangle.LngNE, LatSW: inns.Rectangle.LatSW, LngSW: inns.Rectangle.LngSW},
		Inns:    innsEntity,
		HasNext: inns.HasNext,
	}
}
