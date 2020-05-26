package entity

import (
	"github.com/stayway-corp/stayway-media-api/pkg/util"

	"gopkg.in/guregu/null.v3"
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
		Lat          null.Float
		Lng          null.Float
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
		Times
	}

	TouristSpotList struct {
		TotalNumber  int
		TouristSpots []*TouristSpotDetail
	}

	TouristSpot struct {
		TouristSpotTiny
		AreaCategoryIDs  []*TouristSpotAreaCategory  `gorm:"foreignkey:TouristSpotID"`
		ThemeCategoryIDs []*TouristSpotThemeCategory `gorm:"foreignkey:TouristSpotID"`
		SpotCategoryIDs  []*TouristSpotSpotCategory  `gorm:"foreignkey:TouristSpotID"`
	}

	TouristSpotAreaCategory struct {
		TouristSpotID  int `gorm:"primary_key"`
		AreaCategoryID int `gorm:"primary_key"`
	}

	TouristSpotThemeCategory struct {
		TouristSpotID   int `gorm:"primary_key"`
		ThemeCategoryID int `gorm:"primary_key"`
	}

	TouristSpotSpotCategory struct {
		TouristSpotID  int `gorm:"primary_key"`
		SpotCategoryID int `gorm:"primary_key"`
	}

	TouristSpotDetail struct {
		TouristSpotTiny
		ReviewCount     int
		AreaCategories  []*AreaCategory  `gorm:"many2many:tourist_spot_area_category;jointable_foreignkey:tourist_spot_id;"`
		ThemeCategories []*ThemeCategory `gorm:"many2many:tourist_spot_theme_category;jointable_foreignkey:tourist_spot_id;"`
		SpotCategories  []*SpotCategory  `gorm:"many2many:tourist_spot_spotCategory;jointable_foreignkey:tourist_spot_id;"`
	}
)

func NewTouristSpot(tiny TouristSpotTiny, areaCategoryIDs, themeCategoryIDs, spotCategoryIDs []int) TouristSpot {
	touristSpot := TouristSpot{TouristSpotTiny: tiny}
	touristSpot.SetAreaCategories(areaCategoryIDs)
	touristSpot.SetThemeCategories(themeCategoryIDs)
	touristSpot.SetSpotCategories(spotCategoryIDs)

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

func (ts *TouristSpot) SetSpotCategories(spotCategoryIDs []int) {
	ts.SpotCategoryIDs = make([]*TouristSpotSpotCategory, len(spotCategoryIDs))
	for i, c := range spotCategoryIDs {
		ts.SpotCategoryIDs[i] = &TouristSpotSpotCategory{
			TouristSpotID:  ts.ID,
			SpotCategoryID: c,
		}
	}
}

// テーブル名
func (t *TouristSpotDetail) TableName() string {
	return "tourist_spot"
}

func (t *TouristSpotList) AreaCategoryIDs() []int {
	ids := make([]int, 0)

	for _, touristSpot := range t.TouristSpots {
		for _, area := range touristSpot.AreaCategories {
			ids = append(ids, area.AreaID)

			if area.SubAreaID.Valid {
				ids = append(ids, int(area.SubAreaID.Int64))
			}

			if area.SubSubAreaID.Valid {
				ids = append(ids, int(area.SubSubAreaID.Int64))
			}
		}
	}

	return util.RemoveDuplicatesAndZeroFromIntSlice(ids)
}

func (t *TouristSpotList) ThemeCategoryIDs() []int {
	ids := make([]int, 0)

	for _, touristSpot := range t.TouristSpots {
		for _, theme := range touristSpot.ThemeCategories {
			ids = append(ids, theme.ThemeID)

			if theme.SubThemeID.Valid {
				ids = append(ids, int(theme.SubThemeID.Int64))
			}
		}
	}

	return util.RemoveDuplicatesAndZeroFromIntSlice(ids)
}

func (t *TouristSpotDetail) AreaCategoryIDs() []int {
	ids := make([]int, 0)

	for _, area := range t.AreaCategories {
		ids = append(ids, area.AreaID)

		if area.SubAreaID.Valid {
			ids = append(ids, int(area.SubAreaID.Int64))
		}

		if area.SubSubAreaID.Valid {
			ids = append(ids, int(area.SubSubAreaID.Int64))
		}
	}

	return util.RemoveDuplicatesAndZeroFromIntSlice(ids)
}

func (t *TouristSpotDetail) ThemeCategoryIDs() []int {
	ids := make([]int, 0)

	for _, theme := range t.ThemeCategories {
		ids = append(ids, theme.ThemeID)

		if theme.SubThemeID.Valid {
			ids = append(ids, int(theme.SubThemeID.Int64))
		}
	}

	return util.RemoveDuplicatesAndZeroFromIntSlice(ids)
}

func (t *TouristSpotTiny) TableName() string {
	return "tourist_spot"
}
