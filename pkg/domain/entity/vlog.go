package entity

import (
	"fmt"
	"path"
	"time"

	"github.com/huandu/facebook/v2"

	"github.com/stayway-corp/stayway-media-api/pkg/config"
	facebookEntity "github.com/stayway-corp/stayway-media-api/pkg/domain/entity/facebook"

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
		EditedAt      time.Time
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

	VlogTinyList []*VlogTiny
)

func NewVlog(tiny VlogTiny, areaCategoryIDs, themeCategoryIDs, touristSpotIDs, editors []int) Vlog {
	vlog := Vlog{VlogTiny: tiny}
	vlog.SetAreaCategories(areaCategoryIDs)
	vlog.SetThemeCategories(themeCategoryIDs)
	vlog.SetTouristSpots(touristSpotIDs)
	vlog.SetEditors(editors)

	return vlog
}

func (v *Vlog) SetAreaCategories(areaCategoryIDs []int) {
	v.AreaCategoryIDs = make([]*VlogAreaCategory, len(areaCategoryIDs))
	for i, c := range areaCategoryIDs {
		v.AreaCategoryIDs[i] = &VlogAreaCategory{
			VlogID:         v.ID,
			AreaCategoryID: c,
		}
	}
}

func (v *Vlog) SetThemeCategories(themeCategoryIDs []int) {
	v.ThemeCategoryIDs = make([]*VlogThemeCategory, len(themeCategoryIDs))
	for i, c := range themeCategoryIDs {
		v.ThemeCategoryIDs[i] = &VlogThemeCategory{
			VlogID:          v.ID,
			ThemeCategoryID: c,
		}
	}
}

func (v *Vlog) SetTouristSpots(touristSpotIDs []int) {
	v.TouristSpotIDs = make([]*VlogTouristSpot, len(touristSpotIDs))
	for i, l := range touristSpotIDs {
		v.TouristSpotIDs[i] = &VlogTouristSpot{
			VlogID:        v.ID,
			TouristSpotID: l,
		}
	}
}

func (v *Vlog) SetEditors(editors []int) {
	v.Editors = make([]*VlogEditor, len(editors))
	for i, e := range editors {
		v.Editors[i] = &VlogEditor{
			VlogID: v.ID,
			UserID: e,
		}
	}
}

func (v *VlogTiny) TableName() string {
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

func (v *VlogDetail) TouristSpotIDs() []int {
	ids := make([]int, len(v.TouristSpots))

	for i, touristSpot := range v.TouristSpots {
		ids[i] = touristSpot.ID
	}

	return util.RemoveDuplicatesAndZeroFromIntSlice(ids)
}

// Creator(User)とEditorsのIDを返す
func (v *VlogDetail) UserIDs() []int {
	ids := make([]int, len(v.Editors)+1)
	for i, tiny := range v.Editors {
		ids[i] = tiny.ID
	}
	ids[len(ids)-1] = v.UserID
	return ids
}

func (v *VlogTiny) MediaWebURL(baseURL config.URL) *config.URL {
	baseURL.Path = path.Join(baseURL.Path, fmt.Sprintf("/movie/%d", v.ID))
	return &baseURL
}

func (v VlogTinyList) ToGraphAPIBatchRequestQueryStr(baseURL config.URL) []facebook.Params {
	resolve := make([]facebook.Params, 0, len(v)*2)

	for _, vlog := range v {
		resolve = append(resolve, facebookEntity.GetRelativeURLParams(vlog.MediaWebURL(baseURL)), facebookEntity.GetRelativeTrailingSlashURLParams(vlog.MediaWebURL(baseURL)))
	}

	return resolve
}
