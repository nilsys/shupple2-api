package response

import "github.com/stayway-corp/stayway-media-api/pkg/domain/model"

type (
	Vlog struct {
		ID                      int              `json:"id"`
		Thumbnail               string           `json:"thumbnail"`
		Title                   string           `json:"title"`
		AreaCategories          []*AreaCategory  `json:"areaCategories"`
		ThemeCategoryCategories []*ThemeCategory `json:"themeCategoryCategories"`
	}

	VlogList struct {
		TotalNumber int     `json:"totalNumber"`
		Vlogs       []*Vlog `json:"vlogs"`
	}

	VlogDetail struct {
		ID                      int                `json:"id"`
		Thumbnail               string             `json:"thumbnail"`
		Title                   string             `json:"title"`
		Body                    string             `json:"body"`
		Series                  string             `json:"series"`
		Creator                 Creator            `json:"creator"`
		CreatorSNS              string             `json:"creatorSns"`
		EditorName              string             `json:"editorName"`
		YoutubeURL              string             `json:"youtubeUrl"`
		Views                   int                `json:"views"`
		PhotoYearMonth          string             `json:"photoYearMonth"`
		PlayTime                string             `json:"playTime"`
		Timeline                string             `json:"timeline"`
		FacebookCount           int                `json:"facebookCount"`
		TwitterCount            int                `json:"twitterCount"`
		AreaCategories          []*AreaCategory    `json:"areaCategories"`
		ThemeCategoryCategories []*ThemeCategory   `json:"themeCategoryCategories"`
		CreatedAt               model.TimeResponse `json:"createdAt"`
		UpdatedAt               model.TimeResponse `json:"updatedAt"`
		TouristSpot             []*TouristSpot     `json:"touristSpots"`
	}
)
