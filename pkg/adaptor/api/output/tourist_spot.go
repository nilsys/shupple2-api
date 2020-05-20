package output

import (
	"gopkg.in/guregu/null.v3"

	"github.com/stayway-corp/stayway-media-api/pkg/domain/model"
)

type (
	TouristSpot struct {
		ID              int                    `json:"id"`
		Slug            string                 `json:"slug"`
		Name            string                 `json:"name"`
		Thumbnail       string                 `json:"thumbnail"`
		URL             string                 `json:"url"`
		City            string                 `json:"city"`
		Address         string                 `json:"address"`
		Latitude        null.Float             `json:"latitude"`
		Longitude       null.Float             `json:"longitude"`
		AccessCar       string                 `json:"accessCar"`
		AccessTrain     string                 `json:"accessTrain"`
		AccessBus       string                 `json:"accessBus"`
		Tel             string                 `json:"tel"`
		Price           string                 `json:"price"`
		InstagramURL    string                 `json:"instagramUrl"`
		SearchInnURL    string                 `json:"searchInnUrl"`
		OpeningHours    string                 `json:"openingHours"`
		Rate            float64                `json:"rate"`
		VendorRate      float64                `json:"vendorRate"`
		ReviewCount     int                    `json:"reviewCount"`
		AreaCategories  []*AreaCategoryDetail  `json:"areaCategories"`
		ThemeCategories []*ThemeCategoryDetail `json:"themeCategories"`
		SpotCategories  []*SpotCategory        `json:"spotCategories"`
		CreatedAt       model.TimeResponse     `json:"createdAt"`
		UpdatedAt       model.TimeResponse     `json:"updatedAt"`
	}

	TouristSpotList struct {
		TotalNumber  int            `json:"totalNumber"`
		TouristSpots []*TouristSpot `json:"touristSpots"`
	}
)

func NewTouristSpots(id int, name, thumbnail string) *TouristSpot {
	return &TouristSpot{
		ID:        id,
		Name:      name,
		Thumbnail: thumbnail,
	}
}
