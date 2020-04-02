package converter

import (
	"github.com/stayway-corp/stayway-media-api/pkg/adaptor/api/param"
	"github.com/stayway-corp/stayway-media-api/pkg/adaptor/api/response"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/entity"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/model"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/model/query"
)

func ConvertShowFeatureListParamToQuery(param *param.ShowFeatureListParam) *query.FindListPaginationQuery {
	return &query.FindListPaginationQuery{
		Limit:  param.GetLimit(),
		Offset: param.GetOffSet(),
	}
}

func ConvertFeatureListToOutput(features *entity.FeatureList) *response.FeatureList {
	responseFeatures := make([]*response.Feature, len(features.Features))

	for i, feature := range features.Features {
		responseFeatures[i] = convertFeatureToOutput(feature)
	}

	return &response.FeatureList{
		TotalNumber: features.TotalNumber,
		Features:    responseFeatures,
	}
}

func ConvertQueryFeatureToOutput(feature *entity.QueryFeature) *response.ShowFeature {
	relationPosts := make([]*response.RelationPost, len(feature.Posts))

	for i, post := range feature.Posts {
		relationPosts[i] = convertPostToRelationPost(post)
	}

	return &response.ShowFeature{
		ID:            feature.ID,
		Slug:          feature.Slug,
		Thumbnail:     feature.Thumbnail,
		Title:         feature.Title,
		FacebookCount: feature.FacebookCount,
		TwitterCount:  feature.TwitterCount,
		Views:         feature.Views,
		Creator:       response.NewCreatorFromUser(feature.User),
		CreatedAt:     model.TimeResponse(feature.CreatedAt),
		UpdatedAt:     model.TimeResponse(feature.UpdatedAt),
		RelationPosts: relationPosts,
	}
}

func convertPostToRelationPost(post *entity.Post) *response.RelationPost {
	return &response.RelationPost{
		ID:        post.ID,
		Title:     post.Title,
		Thumbnail: post.Thumbnail,
		Slug:      post.Slug,
	}
}

func convertFeatureToOutput(feature *entity.Feature) *response.Feature {
	return &response.Feature{
		ID:        feature.ID,
		Title:     feature.Title,
		Slug:      feature.Slug,
		Thumbnail: feature.Thumbnail,
	}
}
