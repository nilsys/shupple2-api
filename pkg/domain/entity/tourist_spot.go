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
		AreaCategoryIDs  []*TouristSpotAreaCategory  `gorm:"foreignkey:TouristSpotID"`
		ThemeCategoryIDs []*TouristSpotThemeCategory `gorm:"foreignkey:TouristSpotID"`
		LcategoryIDs     []*TouristSpotLcategory     `gorm:"foreignkey:TouristSpotID"`
	}

	TouristSpotAreaCategory struct {
		TouristSpotID  int `gorm:"primary_key"`
		AreaCategoryID int `gorm:"primary_key"`
	}

	TouristSpotThemeCategory struct {
		TouristSpotID   int `gorm:"primary_key"`
		ThemeCategoryID int `gorm:"primary_key"`
	}

	TouristSpotLcategory struct {
		TouristSpotID int `gorm:"primary_key"`
		LcategoryID   int `gorm:"primary_key"`
	}

	QueryTouristSpot struct {
		TouristSpotTiny
		AreaCategories  []*AreaCategory  `gorm:"many2many:tourist_spot_area_category;jointable_foreignkey:tourist_spot_id;"`
		ThemeCategories []*ThemeCategory `gorm:"many2many:tourist_spot_theme_category;jointable_foreignkey:tourist_spot_id;"`
		Lcategories     []*Lcategory     `gorm:"many2many:tourist_spot_lcategory;jointable_foreignkey:tourist_spot_id;"`
	}
)

func NewTouristSpot(tiny TouristSpotTiny, areaCategoryIDs, themeCategoryIDs, lcategoryIDs []int) TouristSpot {
	touristSpot := TouristSpot{TouristSpotTiny: tiny}
	touristSpot.SetAreaCategories(areaCategoryIDs)
	touristSpot.SetThemeCategories(themeCategoryIDs)
	touristSpot.SetLcategories(lcategoryIDs)

	return touristSpot
}

func (ts *TouristSpot) SetAreaCategories(areaCategoryIDs []int) {
	ts.AreaCategoryIDs = make([]*TouristSpotAreaCategory, len(areaCategoryIDs))
	for i, c := range areaCategoryIDs {
		ts.AreaCategoryIDs[i] = &TouristSpotAreaCategory{
			TouristSpotID:  ts.ID,
			AreaCategoryID: c,
		}
	}
}

func (ts *TouristSpot) SetThemeCategories(themeCategoryIDs []int) {
	ts.ThemeCategoryIDs = make([]*TouristSpotThemeCategory, len(themeCategoryIDs))
	for i, c := range themeCategoryIDs {
		ts.ThemeCategoryIDs[i] = &TouristSpotThemeCategory{
			TouristSpotID:   ts.ID,
			ThemeCategoryID: c,
		}
	}
}

func (ts *TouristSpot) SetLcategories(lcategoryIDs []int) {
	ts.LcategoryIDs = make([]*TouristSpotLcategory, len(lcategoryIDs))
	for i, c := range lcategoryIDs {
		ts.LcategoryIDs[i] = &TouristSpotLcategory{
			TouristSpotID: ts.ID,
			LcategoryID:   c,
		}
	}

}

// テーブル名
func (queryTouristSpot *QueryTouristSpot) TableName() string {
	return "tourist_spot"
}
