package entity

import (
	"time"
)

type (
	TouristSpotTiny struct {
		ID           int `gorm:"primary_key"`
		Slug         string
		Thumbnail    string
		Name         string
		WebsiteURL   string
		City         string
		Address      string
		Lat          float64
		Lng          float64
		AccessCar    string
		AccessTrain  string
		AccessBus    string
		OpeningHours string
		TEL          string
		Price        string
		InstagramURL string
		SearchInnURL string
		Rate         float64
		VendorRate   float64
		CreatedAt    time.Time  `gorm:"default:current_timestamp"`
		UpdatedAt    time.Time  `gorm:"default:current_timestamp"`
		DeletedAt    *time.Time `json:"-"`
	}

	TouristSpot struct {
		TouristSpotTiny
		CategoryIDs  []*TouristSpotCategory  `gorm:"foreignkey:TouristSpotID"`
		LcategoryIDs []*TouristSpotLcategory `gorm:"foreignkey:TouristSpotID"`
	}

	TouristSpotCategory struct {
		TouristSpotID int `gorm:"primary_key"`
		CategoryID    int `gorm:"primary_key"`
	}

	TouristSpotLcategory struct {
		TouristSpotID int `gorm:"primary_key"`
		LcategoryID   int `gorm:"primary_key"`
	}

	QueryTouristSpot struct {
		TouristSpotTiny
		Categories  []*Category  `gorm:"many2many:tourist_spot_category;jointable_foreignkey:tourist_spot_id;"`
		Lcategories []*Lcategory `gorm:"many2many:tourist_spot_lcategory;jointable_foreignkey:tourist_spot_id;"`
	}
)

func NewTouristSpot(tiny TouristSpotTiny, categoryIDs, lcategoryIDs []int) TouristSpot {
	touristSpotCategoryIDs := make([]*TouristSpotCategory, len(categoryIDs))
	for i, c := range categoryIDs {
		touristSpotCategoryIDs[i] = &TouristSpotCategory{
			TouristSpotID: tiny.ID,
			CategoryID:    c,
		}
	}

	touristSpotLcategoryIDs := make([]*TouristSpotLcategory, len(lcategoryIDs))
	for i, c := range lcategoryIDs {
		touristSpotLcategoryIDs[i] = &TouristSpotLcategory{
			TouristSpotID: tiny.ID,
			LcategoryID:   c,
		}
	}

	return TouristSpot{
		tiny,
		touristSpotCategoryIDs,
		touristSpotLcategoryIDs,
	}
}

// テーブル名
func (queryTouristSpot *QueryTouristSpot) TableName() string {
	return "tourist_spot"
}
