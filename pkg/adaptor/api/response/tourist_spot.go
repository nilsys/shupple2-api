package response

import "github.com/stayway-corp/stayway-media-api/pkg/domain/model"

type (
	TouristSpot struct {
		ID        int    `json:"id"`
		Name      string `json:"name"`
		Thumbnail string `json:"thumbnail"`
		URL       string `json:"url,omitempty"`
	}

	ShowTouristSpot struct {
		ID              int                `json:"id"`
		Slug            string             `json:"slug"`
		Name            string             `json:"name"`
		WebsiteURL      string             `json:"websiteUrl"`
		City            string             `json:"city"`
		Address         string             `json:"address"`
		Latitude        float64            `json:"latitude"`
		Longitude       float64            `json:"longitude"`
		AccessCar       string             `json:"accessCar"`
		AccessTrain     string             `json:"accessTrain"`
		AccessBus       string             `json:"accessBus"`
		OpeningHours    string             `json:"openingHours"`
		Tel             string             `json:"tel"`
		Price           string             `json:"price"`
		InstagramURL    string             `json:"instagramUrl,omitempty"`
		SearchInnURL    string             `json:"searchInnUrl,omitempty"`
		Rate            float64            `json:"rate"`
		VendorRate      float64            `json:"vendorRate"`
		AreaCategories  []Category         `json:"areaCategories"`
		ThemeCategories []Category         `json:"themeCategories"`
		Lcategories     []Lcategory        `json:"lcategories"`
		CreatedAt       model.TimeResponse `json:"createdAt"`
		UpdatedAt       model.TimeResponse `json:"updatedAt"`
	}
)
