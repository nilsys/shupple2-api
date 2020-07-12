package entity

type (
	InnAreaTypeIDs struct {
		AreaID       int
		SubAreaID    int
		SubSubAreaID int
	}

	InnType struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
		URL  string `json:"url"`
	}
	AvailableCard struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	}
	RoomType struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	}
	InnTag struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
		URL  string `json:"url"`
	}
	InnBrand struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	}

	Inn struct {
		ID             int             `json:"id"`
		Name           string          `json:"name"`
		Description    string          `json:"description"`
		URL            string          `json:"url"`
		Brand          InnBrand        `json:"brand"`
		Thumbnail      string          `json:"thumbnail"`
		Latitude       float64         `json:"latitude"`
		Longitude      float64         `json:"longitude"`
		CountryCode    string          `json:"countryCode"`
		Address        string          `json:"address"`
		PostCode       string          `json:"postCode"`
		Phone          string          `json:"phone"`
		Access         string          `json:"access"`
		StayDaysMin    int             `json:"stayDaysMin"`
		StayDaysMax    int             `json:"stayDaysMax"`
		CheckInTime    string          `json:"checkInTime"`
		CheckOutTime   string          `json:"checkOutTime"`
		RoomCount      int             `json:"roomCount"`
		BedCount       int             `json:"bedCount"`
		Score          float64         `json:"score"`
		ReviewCount    int             `json:"reviewCount"`
		MinPrice       interface{}     `json:"minPrice"`
		MaxPrice       interface{}     `json:"maxPrice"`
		InnTypes       []InnType       `json:"innTypes"`
		AvailableCards []AvailableCard `json:"availableCards"`
		RoomTypes      []RoomType      `json:"roomTypes"`
		TrueTags       []InnTag        `json:"trueTags"`
		FalseTags      []InnTag        `json:"falseTags"`
	}

	Inns struct {
		Count     int `json:"count"`
		Rectangle struct {
			LatNE float64 `json:"latNE"`
			LngNE float64 `json:"lngNE"`
			LatSW float64 `json:"latSW"`
			LngSW float64 `json:"lngSW"`
		} `json:"rectangle"`
		Inns    []Inn `json:"inns"`
		HasNext bool  `json:"hasNext"`
	}
)

func (inns Inns) IDs() []int {
	res := make([]int, len(inns.Inns))
	for i, inn := range inns.Inns {
		res[i] = inn.ID
	}

	return res
}
