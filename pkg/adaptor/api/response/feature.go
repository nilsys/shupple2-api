package response

import "github.com/stayway-corp/stayway-media-api/pkg/domain/model"

type (
	Feature struct {
		ID        int    `json:"id"`
		Title     string `json:"title"`
		Slug      string `json:"slug"`
		Thumbnail string `json:"thumbnail"`
	}

	// フロント返却用Feature詳細
	ShowFeature struct {
		ID            int    `json:"id"`
		Slug          string `json:"slug"`
		Thumbnail     string `json:"thumbnail"`
		Title         string `json:"title"`
		Creator       `json:"creator"`
		UpdatedAt     model.TimeResponse `json:"updatedAt"`
		RelationPosts []*RelationPost    `json:"relationPosts"`
	}

	// Featureに関する記事
	RelationPost struct {
		ID        int    `json:"id"`
		Title     string `json:"title"`
		Thumbnail string `json:"thumbnail"`
		Slug      string `json:"slug"`
	}
)
