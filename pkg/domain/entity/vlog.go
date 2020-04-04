package entity

import (
	"time"
)

type (
	VlogTiny struct {
		ID            int `gorm:"primary_key"`
		UserID        int
		EditorID      int
		Slug          string
		Thumbnail     string
		Title         string
		Body          string
		YoutubeURL    string
		Series        string
		UserSNS       string
		YearMonth     string `gorm:"column:yearmonth"`
		PlayTime      string
		Timeline      string
		FacebookCount int
		TwitterCount  int
		Views         int
		CreatedAt     time.Time `gorm:"default:current_timestamp"`
		UpdatedAt     time.Time `gorm:"default:current_timestamp"`
		DeletedAt     *time.Time
	}

	Vlog struct {
		VlogTiny
		AreaCategoryIDs  []*VlogAreaCategory
		ThemeCategoryIDs []*VlogThemeCategory
		TouristSpotIDs   []*VlogTouristSpot
	}

	VlogAreaCategory struct {
		VlogID         int `gorm:"primary_key"`
		AreaCategoryID int `gorm:"primary_key"`
	}

	VlogThemeCategory struct {
		VlogID          int `gorm:"primary_key"`
		ThemeCategoryID int `gorm:"primary_key"`
	}

	VlogTouristSpot struct {
		VlogID        int `gorm:"primary_key"`
		TouristSpotID int `gorm:"primary_key"`
	}

	VlogDetail struct {
		VlogTiny
		AreaCategories  []*AreaCategory  `gorm:"many2many:vlog_area_category;jointable_foreignkey:vlog_id;"`
		ThemeCategories []*ThemeCategory `gorm:"many2many:vlog_theme_category;jointable_foreignkey:vlog_id;"`
	}

	VlogDetailList struct {
		TotalNumber int
		Vlogs       []*VlogDetail
	}

	VlogDetailWithTouristSpots struct {
		VlogTiny
		User            *User            `gorm:"foreignkey:UserID"`
		Editor          *User            `gorm:"foreignkey:EditorID"`
		AreaCategories  []*AreaCategory  `gorm:"many2many:vlog_area_category;jointable_foreignkey:vlog_id;"`
		ThemeCategories []*ThemeCategory `gorm:"many2many:vlog_theme_category;jointable_foreignkey:vlog_id;"`
		TouristSpots    []*TouristSpot   `gorm:"many2many:vlog_tourist_spot;jointable_foreignkey:vlog_id;"`
	}
)

func NewVlog(tiny VlogTiny, areaCategoryIDs, themeCategoryIDs, touristSpotIDs []int) Vlog {
	vlog := Vlog{VlogTiny: tiny}
	vlog.SetAreaCategories(areaCategoryIDs)
	vlog.SetThemeCategories(themeCategoryIDs)
	vlog.SetTouristSpots(touristSpotIDs)

	return vlog
}

func (vlog *Vlog) SetAreaCategories(areaCategoryIDs []int) {
	vlog.AreaCategoryIDs = make([]*VlogAreaCategory, len(areaCategoryIDs))
	for i, c := range areaCategoryIDs {
		vlog.AreaCategoryIDs[i] = &VlogAreaCategory{
			VlogID:         vlog.ID,
			AreaCategoryID: c,
		}
	}
}

func (vlog *Vlog) SetThemeCategories(themeCategoryIDs []int) {
	vlog.ThemeCategoryIDs = make([]*VlogThemeCategory, len(themeCategoryIDs))
	for i, c := range themeCategoryIDs {
		vlog.ThemeCategoryIDs[i] = &VlogThemeCategory{
			VlogID:          vlog.ID,
			ThemeCategoryID: c,
		}
	}
}

func (vlog *Vlog) SetTouristSpots(touristSpotIDs []int) {
	vlog.TouristSpotIDs = make([]*VlogTouristSpot, len(touristSpotIDs))
	for i, l := range touristSpotIDs {
		vlog.TouristSpotIDs[i] = &VlogTouristSpot{
			VlogID:        vlog.ID,
			TouristSpotID: l,
		}
	}
}

func (queryVlog *VlogTiny) TableName() string {
	return "vlog"
}
