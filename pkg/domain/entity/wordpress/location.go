package wordpress

import (
	"encoding/json"
	"strconv"

	"github.com/pkg/errors"
)

type Location struct {
	ID            int                `json:"id"`
	Date          Time               `json:"date"`
	DateGmt       Time               `json:"date_gmt"`
	GUID          Text               `json:"guid"`
	Modified      Time               `json:"modified"`
	ModifiedGmt   Time               `json:"modified_gmt"`
	Slug          string             `json:"slug"`
	Status        Status             `json:"status"`
	Title         Text               `json:"title"`
	FeaturedMedia int                `json:"featured_media"`
	Categories    []int              `json:"categories"`
	LocationCat   []int              `json:"location__cat"`
	Attributes    LocationAttributes `json:"acf"`
}

type LocationAttributes struct {
	OfficialURL  string          `json:"officialurl"`
	City         string          `json:"spotplace2"`
	Address      string          `json:"address"`
	Map          json.RawMessage `json:"spotmapaddress"`
	AccessTrain  string          `json:"access"`
	AccessBus    string          `json:"bus_access"`
	AccessCar    string          `json:"car_access"`
	OpeningHours string          `json:"time"`
	TEL          string          `json:"spottell"`
	Price        string          `json:"spotprice"`
	Instagram    string          `json:"instagram"`
	Inn          string          `json:"inn"`
}

type LocationMapString struct {
	Lat string `json:"lat"`
	Lng string `json:"lng"`
}

type LocationMapFloat struct {
	Lat float64 `json:"lat"`
	Lng float64 `json:"lng"`
}

func (attrs LocationAttributes) LatLang() (lat float64, lng float64, err error) {
	var locationMapFloat LocationMapFloat
	if err := json.Unmarshal([]byte(attrs.Map), &locationMapFloat); err == nil {
		return locationMapFloat.Lat, locationMapFloat.Lng, nil
	}

	var locationMapString LocationMapString
	if err := json.Unmarshal([]byte(attrs.Map), &locationMapString); err == nil {
		lat, err = strconv.ParseFloat(locationMapString.Lat, 64)
		if err != nil {
			return 0, 0, errors.Wrap(err, "failed to parse locationMap lat")
		}

		lng, err = strconv.ParseFloat(locationMapString.Lng, 64)
		if err != nil {
			return 0, 0, errors.Wrap(err, "failed to parse locationMap lng")
		}

		return lat, lng, nil
	}

	return 0, 0, errors.New("failed to parse json as LocationMap")
}
