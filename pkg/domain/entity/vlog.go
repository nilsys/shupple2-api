package entity

import (
	"github.com/stayway-corp/stayway-media-api/pkg/util"
)

type (
	VlogTiny struct {
		ID            int `gorm:"primary_key"`
		UserID        int
		Slug          string
		Thumbnail     string
		Title         string
		Body          string
		YoutubeURL    string
		Series        string
		YearMonth     string `gorm:"column:yearmonth"`
		PlayTime      string
		Timeline      string
		FacebookCount int
		TwitterCount  int
		FavoriteCount int
		Views         int
		Times
	}

	Vlog struct {
		VlogTiny
		AreaCategoryIDs  []*VlogAreaCategory
		ThemeCategoryIDs []*VlogThemeCategory
		TouristSpotIDs   []*VlogTouristSpot
		Editors          []*VlogEditor
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

	VlogEditor struct {
		VlogID int `gorm:"primary_key"`
		UserID int `gorm:"primary_key"`
	}

	VlogForList struct {
		VlogTiny
		IsFavorite      bool
		AreaCategories  []*AreaCategory  `gorm:"many2many:vlog_area_category;jointable_foreignkey:vlog_id;"`
		ThemeCategories []*ThemeCategory `gorm:"many2many:vlog_theme_category;jointable_foreignkey:vlog_id;"`
	}

	VlogList struct {
		TotalNumber int
		Vlogs       []*VlogForList
	}

	VlogDetail struct {
		VlogTiny
		AreaCategories  []*AreaCategory  `gorm:"many2many:vlog_area_category;jointable_foreignkey:vlog_id;"`
		ThemeCategories []*ThemeCategory `gorm:"many2many:vlog_theme_category;jointable_foreignkey:vlog_id;"`
		TouristSpots    []*TouristSpot   `gorm:"many2many:vlog_tourist_spot;jointable_foreignkey:vlog_id;"`
		User            *User            `gorm:"foreignkey:UserID"`
		Editors         []*User          `gorm:"many2many:vlog_editor;jointable_foreignkey:vlog_id;"`
		IsFavorite      bool
	}

	UserFavoriteVlog struct {
		UserID int
		VlogID int
	}
)

func NewVlog(tiny VlogTiny, areaCategoryIDs, themeCategoryIDs, touristSpotIDs, editors []int) Vlog {
	vlog := Vlog{VlogTiny: tiny}
	vlog.SetAreaCategories(areaCategoryIDs)
	vlog.SetThemeCategories(themeCategoryIDs)
	vlog.SetTouristSpots(touristSpotIDs)
	vlog.SetEditors(editors)

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

func (vlog *Vlog) SetEditors(editors []int) {
	vlog.Editors = make([]*VlogEditor, len(editors))
	for i, e := range editors {
		vlog.Editors[i] = &VlogEditor{
			VlogID: vlog.ID,
			UserID: e,
		}
	}
}

func (queryVlog *VlogTiny) TableName() string {
	return "vlog"
}

func NewUserFavoriteVlog(userID, vlogID int) *UserFavoriteVlog {
	return &UserFavoriteVlog{
		UserID: userID,
		VlogID: vlogID,
	}
}

func (v *VlogList) AreaCategoryIDs() []int {
	ids := make([]int, 0)

	for _, vlog := range v.Vlogs {
		for _, area := range vlog.AreaCategories {
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

func (v *VlogList) ThemeCategoryIDs() []int {
	ids := make([]int, 0)

	for _, vlog := range v.Vlogs {
		for _, theme := range vlog.ThemeCategories {
			ids = append(ids, theme.ThemeID)

			if theme.SubThemeID.Valid {
				ids = append(ids, int(theme.SubThemeID.Int64))
			}
		}
	}

	return util.RemoveDuplicatesAndZeroFromIntSlice(ids)
}

func (v *VlogDetail) AreaCategoryIDs() []int {
	ids := make([]int, 0)

	for _, area := range v.AreaCategories {
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

func (v *VlogDetail) ThemeCategoryIDs() []int {
	ids := make([]int, 0)

	for _, theme := range v.ThemeCategories {
		ids = append(ids, theme.ThemeID)

		if theme.SubThemeID.Valid {
			ids = append(ids, int(theme.SubThemeID.Int64))
		}
	}

	return util.RemoveDuplicatesAndZeroFromIntSlice(ids)
}
