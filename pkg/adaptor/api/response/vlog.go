package response

import "github.com/stayway-corp/stayway-media-api/pkg/domain/model"

type (
	Vlog struct {
		ID              int              `json:"id"`
		Thumbnail       string           `json:"thumbnail"`
		Title           string           `json:"title"`
		AreaCategories  []*AreaCategory  `json:"areaCategories"`
		ThemeCategories []*ThemeCategory `json:"themeCategories"`
	}

	VlogList struct {
		TotalNumber int     `json:"totalNumber"`
		Vlogs       []*Vlog `json:"vlogs"`
	}

	VlogDetail struct {
		ID              int                `json:"id"`
		Thumbnail       string             `json:"thumbnail"`
		Title           string             `json:"title"`
		Body            string             `json:"body"`
		Series          string             `json:"series"`
		Creator         Creator            `json:"creator"`
		Editor          Creator            `json:"editor"`
		YoutubeURL      string             `json:"youtubeUrl"`
		Views           int                `json:"views"`
		ShootingDate    string             `json:"shootingDate"`
		PlayTime        string             `json:"playTime"`
		Timeline        string             `json:"timeline"`
		FacebookCount   int                `json:"facebookCount"`
		TwitterCount    int                `json:"twitterCount"`
		AreaCategories  []*AreaCategory    `json:"areaCategories"`
		ThemeCategories []*ThemeCategory   `json:"themeCategories"`
		CreatedAt       model.TimeResponse `json:"createdAt"`
		UpdatedAt       model.TimeResponse `json:"updatedAt"`
		TouristSpot     []*TouristSpot     `json:"touristSpots"`
	}
)
