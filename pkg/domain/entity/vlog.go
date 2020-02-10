package entity

import "time"

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
		TouristSpotIDs          []*VlogTouristSpot
	}

	VlogCategory struct {
		VlogID     int `gorm:"primary_key"`
		CategoryID int `gorm:"primary_key"`
	}

	VlogTouristSpot struct {
		VlogID     int `gorm:"primary_key"`
		TouristSpotID int `gorm:"primary_key"`
	}
)

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
			VlogID:     tiny.ID,
			TouristSpotID: l,
		}
	}

	return Vlog{
		tiny,
		vlogCategories,
		vlogTouristSpots,
	}
}
