package output

import "github.com/stayway-corp/stayway-media-api/pkg/domain/model"

type (
	Feature struct {
		ID        int    `json:"id"`
		Title     string `json:"title"`
		Slug      string `json:"slug"`
		Thumbnail string `json:"thumbnail"`
	}

	FeatureList struct {
		TotalNumber int        `json:"totalNumber"`
		Features    []*Feature `json:"featurePosts"`
	}

	// フロント返却用Feature詳細
	ShowFeature struct {
		ID            int                       `json:"id"`
		Slug          string                    `json:"slug"`
		Thumbnail     string                    `json:"thumbnail"`
		Title         string                    `json:"title"`
		Body          string                    `json:"body"`
		FacebookCount int                       `json:"facebookCount"`
		TwitterCount  int                       `json:"twitterCount"`
		Views         int                       `json:"views"`
		Creator       Creator                   `json:"creator"`
		EditedAt      model.TimeResponse        `json:"editedAt"`
		CreatedAt     model.TimeResponse        `json:"createdAt"`
		UpdatedAt     model.TimeResponse        `json:"updatedAt"`
		RelationPosts []*PostWithCategoryDetail `json:"relationPosts"`
	}
)
