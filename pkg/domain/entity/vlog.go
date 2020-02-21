package entity

import (
	"strconv"
	"time"
)

type (
	VlogTiny struct {
		ID         int `gorm:"primary_key"`
		UserID     int
		Slug       string
		Title      string
		Body       string
		YoutubeURL string
		Series     string
		UserSNS    string
		EditorName string
		YearMonth  string `gorm:"column:yearmonth"`
		PlayTime   string
		Timeline   string
		CreatedAt  time.Time `gorm:"default:current_timestamp"`
		UpdatedAt  time.Time `gorm:"default:current_timestamp"`
		DeletedAt  *time.Time
	}

	Vlog struct {
		VlogTiny
		WordpressCategoryIDs []*VlogCategory
		TouristSpotIDs       []*VlogTouristSpot
	}

	VlogCategory struct {
		VlogID     int `gorm:"primary_key"`
		CategoryID int `gorm:"primary_key"`
	}

	VlogTouristSpot struct {
		VlogID        int `gorm:"primary_key"`
		TouristSpotID int `gorm:"primary_key"`
	}

	QueryVlog struct {
		VlogTiny
		WordpressCategories []*Category    `gorm:"many2many:vlog_category;jointable_foreignkey:vlog_id;"`
		TouristSpots        []*TouristSpot `gorm:"many2many:vlog_tourist_spot;jointable_foreignkey:vlog_id;"`
	}
)

func (queryVlog *QueryVlog) GenerateThumbnailURL() string {
	return "https://files.stayway.jp/vlog/" + strconv.Itoa(queryVlog.VlogTiny.ID)
}

func NewVlog(tiny VlogTiny, categoryIDs, touristSpotIDs []int) Vlog {
	vlogCategories := make([]*VlogCategory, len(categoryIDs))
	for i, c := range categoryIDs {
		vlogCategories[i] = &VlogCategory{
			VlogID:     tiny.ID,
			CategoryID: c,
		}
	}

	vlogTouristSpots := make([]*VlogTouristSpot, len(touristSpotIDs))
	for i, l := range touristSpotIDs {
		vlogTouristSpots[i] = &VlogTouristSpot{
			VlogID:        tiny.ID,
			TouristSpotID: l,
		}
	}

	return Vlog{
		tiny,
		vlogCategories,
		vlogTouristSpots,
	}
}
