package output

import "github.com/stayway-corp/stayway-media-api/pkg/domain/model"

type (
	// フロント返却用Comic
	Comic struct {
		ID          int    `json:"id"`
		Slug        string `json:"slug"`
		Title       string `json:"title"`
		Thumbnail   string `json:"thumbnail"`
		IsFavorited bool   `json:"isFavorited"`
	}

	ComicList struct {
		TotalNumber int      `json:"totalNumber"`
		Comics      []*Comic `json:"comics"`
	}

	// フロント返却用Comic詳細
	ShowComic struct {
		ID          int                `json:"id"`
		Slug        string             `json:"slug"`
		Title       string             `json:"title"`
		Thumbnail   string             `json:"thumbnail"`
		Body        string             `json:"body"`
		IsFavorited bool               `json:"isFavorited"`
		Creator     Creator            `json:"creator"`
		CreatedAt   model.TimeResponse `json:"createdAt"`
	}
)
