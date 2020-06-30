package output

import (
	"github.com/stayway-corp/stayway-media-api/pkg/domain/entity"
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

	TouristSpotTiny struct {
		ID           int        `json:"id"`
		Slug         string     `json:"slug"`
		Name         string     `json:"name"`
		Thumbnail    string     `json:"thumbnail"`
		URL          string     `json:"url"`
		City         string     `json:"city"`
		Address      string     `json:"address"`
		Latitude     null.Float `json:"latitude"`
		Longitude    null.Float `json:"longitude"`
		AccessCar    string     `json:"accessCar"`
		AccessTrain  string     `json:"accessTrain"`
		AccessBus    string     `json:"accessBus"`
		Tel          string     `json:"tel"`
		Price        string     `json:"price"`
		InstagramURL string     `json:"instagramUrl"`
		SearchInnURL string     `json:"searchInnUrl"`
		OpeningHours string     `json:"openingHours"`
		Rate         float64    `json:"rate"`
		VendorRate   float64    `json:"vendorRate"`
		ReviewCount  int        `json:"reviewCount"`
	}
)

// TODO: review_countをtourist_spotテーブルに追加して、Review投稿時にIncrementする様にする、その際にscriptを書いて既存のReviewの数を含める
func NewTouristSpotTinyFromEntity(touristSpot *entity.TouristSpot, reviewCount int) *TouristSpotTiny {
	return &TouristSpotTiny{
		ID:           touristSpot.ID,
		Slug:         touristSpot.Slug,
		Name:         touristSpot.Name,
		Thumbnail:    touristSpot.Thumbnail,
		URL:          touristSpot.WebsiteURL,
		City:         touristSpot.City,
		Address:      touristSpot.Address,
		Latitude:     touristSpot.Lat,
		Longitude:    touristSpot.Lng,
		AccessCar:    touristSpot.AccessCar,
		AccessTrain:  touristSpot.AccessTrain,
		AccessBus:    touristSpot.AccessBus,
		Tel:          touristSpot.TEL,
		Price:        touristSpot.Price,
		InstagramURL: touristSpot.InstagramURL,
		SearchInnURL: touristSpot.SearchInnURL,
		OpeningHours: touristSpot.OpeningHours,
		Rate:         touristSpot.Rate,
		VendorRate:   touristSpot.VendorRate,
		ReviewCount:  reviewCount,
	}
}
