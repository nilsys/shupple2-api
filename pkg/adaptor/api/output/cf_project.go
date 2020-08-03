package output

import "github.com/stayway-corp/stayway-media-api/pkg/domain/model"

type (
	CfProjectSupportComment struct {
		ID        int                `json:"id"`
		Body      string             `json:"body"`
		User      *UserSummary       `json:"user"`
		CreatedAt model.TimeResponse `json:"createdAt"`
	}

	CfProject struct {
		ID              int                   `json:"id"`
		SnapshotID      int                   `json:"snapshotId"`
		Title           string                `json:"title"`
		Summary         string                `json:"summary"`
		Thumbnail       string                `json:"thumbnail"`
		Body            string                `json:"body"`
		GoalPrice       int                   `json:"goalPrice"`
		AchievedPrice   int                   `json:"achievedPrice"`
		SupporterCount  int                   `json:"supporterCount"`
		FavoriteCount   int                   `json:"favoriteCount"`
		FacebookCount   int                   `json:"facebookCount"`
		TwitterCount    int                   `json:"twitterCount"`
		Creator         Creator               `json:"creator"`
		Thumbnails      []*CfProjectThumbnail `json:"thumbnails"`
		AreaCategories  []*AreaCategory       `json:"areaCategories"`
		ThemeCategories []*ThemeCategory      `json:"themeCategories"`
		DeadLine        model.TimeResponse    `json:"deadLine"`
		CreatedAt       model.TimeResponse    `json:"createdAt"`
		EditedAt        model.TimeResponse    `json:"editedAt"`
	}

	CfProjectThumbnail struct {
		Priority  int    `json:"priority"`
		Thumbnail string `json:"thumbnail"`
	}
)
