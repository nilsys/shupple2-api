package converter

import (
	"github.com/stayway-corp/stayway-media-api/pkg/adaptor/api/input"
	"github.com/stayway-corp/stayway-media-api/pkg/adaptor/api/output"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/entity"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/model"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/model/query"
)

func ConvertShowFeatureListParamToQuery(param *input.ShowFeatureListParam) *query.FindListPaginationQuery {
	return &query.FindListPaginationQuery{
		Limit:  param.GetLimit(),
		Offset: param.GetOffSet(),
	}
}

func ConvertFeatureListToOutput(features *entity.FeatureList) *output.FeatureList {
	responseFeatures := make([]*output.Feature, len(features.Features))

	for i, feature := range features.Features {
		responseFeatures[i] = convertFeatureToOutput(feature)
	}

	return &output.FeatureList{
		TotalNumber: features.TotalNumber,
		Features:    responseFeatures,
	}
}

func ConvertQueryFeatureToOutput(feature *entity.QueryFeature) *output.ShowFeature {
	relationPosts := make([]*output.RelationPost, len(feature.Posts))

	for i, post := range feature.Posts {
		relationPosts[i] = convertPostToRelationPost(post)
	}

	return &output.ShowFeature{
		ID:            feature.ID,
		Slug:          feature.Slug,
		Thumbnail:     feature.Thumbnail,
		Title:         feature.Title,
		FacebookCount: feature.FacebookCount,
		TwitterCount:  feature.TwitterCount,
		Views:         feature.Views,
		Creator:       output.NewCreatorFromUser(feature.User),
		CreatedAt:     model.TimeResponse(feature.CreatedAt),
		UpdatedAt:     model.TimeResponse(feature.UpdatedAt),
		RelationPosts: relationPosts,
	}
}

func convertPostToRelationPost(post *entity.Post) *output.RelationPost {
	return &output.RelationPost{
		ID:        post.ID,
		Title:     post.Title,
		Thumbnail: post.Thumbnail,
		Slug:      post.Slug,
	}
}

func convertFeatureToOutput(feature *entity.Feature) *output.Feature {
	return &output.Feature{
		ID:        feature.ID,
		Title:     feature.Title,
		Slug:      feature.Slug,
		Thumbnail: feature.Thumbnail,
	}
}
