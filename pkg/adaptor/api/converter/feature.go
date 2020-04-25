package converter

import (
	"github.com/stayway-corp/stayway-media-api/pkg/adaptor/api/input"
	"github.com/stayway-corp/stayway-media-api/pkg/adaptor/api/output"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/entity"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/model"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/model/query"
)

func (c Converters) ConvertShowFeatureListParamToQuery(param *input.ShowFeatureListParam) *query.FindListPaginationQuery {
	return &query.FindListPaginationQuery{
		Limit:  param.GetLimit(),
		Offset: param.GetOffSet(),
	}
}

func (c Converters) ConvertFeatureListToOutput(features *entity.FeatureList) *output.FeatureList {
	responseFeatures := make([]*output.Feature, len(features.Features))

	for i, feature := range features.Features {
		responseFeatures[i] = c.convertFeatureToOutput(feature)
	}

	return &output.FeatureList{
		TotalNumber: features.TotalNumber,
		Features:    responseFeatures,
	}
}

func (c Converters) ConvertFeatureDetailPostsToOutput(feature *entity.FeatureDetailWithPosts) *output.ShowFeature {
	relationPosts := c.ConvertPostListTiniesToOutput(feature.Posts)

	return &output.ShowFeature{
		ID:            feature.ID,
		Slug:          feature.Slug,
		Thumbnail:     feature.Thumbnail,
		Title:         feature.Title,
		Body:          feature.Body,
		FacebookCount: feature.FacebookCount,
		TwitterCount:  feature.TwitterCount,
		Views:         feature.Views,
		Creator:       c.NewCreatorFromUser(feature.User),
		CreatedAt:     model.TimeResponse(feature.CreatedAt),
		UpdatedAt:     model.TimeResponse(feature.UpdatedAt),
		RelationPosts: relationPosts,
	}
}

func (c Converters) convertFeatureToOutput(feature *entity.Feature) *output.Feature {
	return &output.Feature{
		ID:        feature.ID,
		Title:     feature.Title,
		Slug:      feature.Slug,
		Thumbnail: feature.Thumbnail,
	}
}
