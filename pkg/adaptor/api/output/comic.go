package output

import "github.com/stayway-corp/stayway-media-api/pkg/domain/model"

type (
	// フロント返却用Comic
	Comic struct {
		ID            int    `json:"id"`
		Slug          string `json:"slug"`
		Title         string `json:"title"`
		Thumbnail     string `json:"thumbnail"`
		IsFavorite    bool   `json:"isFavorite"`
		FavoriteCount int    `json:"favoriteCount"`
	}

	ComicList struct {
		TotalNumber int      `json:"totalNumber"`
		Comics      []*Comic `json:"comics"`
	}

	// フロント返却用Comic詳細
	ShowComic struct {
		ID            int                `json:"id"`
		Slug          string             `json:"slug"`
		Title         string             `json:"title"`
		Thumbnail     string             `json:"thumbnail"`
		Body          string             `json:"body"`
		IsFavorite    bool               `json:"isFavorite"`
		FavoriteCount int                `json:"favoriteCount"`
		Creator       Creator            `json:"creator"`
		CreatedAt     model.TimeResponse `json:"createdAt"`
	}
)
