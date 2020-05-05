package converter

import (
	"github.com/stayway-corp/stayway-media-api/pkg/adaptor/api/input"
	"github.com/stayway-corp/stayway-media-api/pkg/adaptor/api/output"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/entity"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/model"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/model/command"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/model/query"
)

// i/oの構造体からレポジトリで使用するクエリ発行用構造体へコンバート
func (c Converters) ConvertFindReviewListParamToQuery(param *input.ListReviewParams) *query.ShowReviewListQuery {
	return &query.ShowReviewListQuery{
		UserID:                 param.UserID,
		InnID:                  param.InnID,
		TouristSpotID:          param.TouristSpotID,
		HashTag:                param.HashTag,
		AreaID:                 param.AreaID,
		SubAreaID:              param.SubAreaID,
		SubSubAreaID:           param.SubSubAreaID,
		MetasearchAreaID:       param.MetasearchAreaID,
		MetasearchSubAreaID:    param.MetasearchSubAreaID,
		MetasearchSubSubAreaID: param.MetasearchSubSubAreaID,
		SortBy:                 param.SortBy,
		Keyword:                param.Keyward,
		ExcludeID:              param.ExcludeID,
		Limit:                  param.GetLimit(),
		OffSet:                 param.GetOffset(),
	}
}

func (c Converters) ConvertListFeedReviewParamToQuery(param *input.ListFeedReviewParam) *query.FindListPaginationQuery {
	return &query.FindListPaginationQuery{
		Limit:  param.GetLimit(),
		Offset: param.GetOffset(),
	}
}

func (c Converters) ConvertReviewCommentToOutput(reviewComment *entity.ReviewComment) *output.ReviewComment {
	userSummary := output.NewUserSummary(reviewComment.User.ID, reviewComment.User.UID, reviewComment.User.Name, reviewComment.User.AvatarURL(c.filesURL()))
	return output.NewReviewComment(
		userSummary,
		reviewComment.Body,
		model.TimeResponse(reviewComment.CreatedAt),
		reviewComment.ID,
		reviewComment.ReplyCount,
		reviewComment.FavoriteCount,
		false,
	)
}

func (c Converters) ConvertReviewCommentWithIsFavoriteListToOutput(reviewComments []*entity.ReviewCommentWithIsFavorite) []*output.ReviewComment {
	reviewCommentOutputs := make([]*output.ReviewComment, len(reviewComments))
	for i, reviewComment := range reviewComments {
		reviewCommentOutputs[i] = c.ConvertReviewCommentWithIsFavoriteToOutput(reviewComment)
	}
	return reviewCommentOutputs
}

func (c Converters) ConvertReviewCommentWithIsFavoriteToOutput(reviewComment *entity.ReviewCommentWithIsFavorite) *output.ReviewComment {
	userSummary := output.NewUserSummary(reviewComment.User.ID, reviewComment.User.UID, reviewComment.User.Name, reviewComment.User.AvatarURL(c.filesURL()))
	return output.NewReviewComment(
		userSummary,
		reviewComment.Body,
		model.TimeResponse(reviewComment.CreatedAt),
		reviewComment.ID,
		reviewComment.ReplyCount,
		reviewComment.FavoriteCount,
		reviewComment.IsFavorited,
	)
}

func (c Converters) ConvertQueryReviewListToOutput(queryReviews []*entity.ReviewDetail) []*output.Review {
	outputs := make([]*output.Review, len(queryReviews))
	for i, queryReview := range queryReviews {
		outputs[i] = c.convertQueryReviewToOutput(queryReview)
	}
	return outputs
}

func (c Converters) ConvertQueryReviewDetailWithIsFavoriteListToOutput(reviews *entity.ReviewDetailWithIsFavoriteList) output.ReviewDetailIsFavoriteList {
	responses := make([]*output.Review, len(reviews.Reviews))
	for i, queryReview := range reviews.Reviews {
		responses[i] = c.ConvertQueryReviewDetailWithIsFavoriteToOutput(queryReview)
	}
	return output.ReviewDetailIsFavoriteList{
		TotalNumber: reviews.TotalNumber,
		Reviews:     responses,
	}
}

func (c Converters) ConvertQueryReviewDetailWithIsFavoriteToOutput(queryReview *entity.ReviewDetailWithIsFavorite) *output.Review {
	hashtags := make([]output.Hashtag, len(queryReview.Hashtag))
	for i, hashtag := range queryReview.Hashtag {
		hashtags[i] = output.Hashtag{
			ID:   hashtag.ID,
			Name: hashtag.Name,
		}
	}

	return &output.Review{
		ID:            queryReview.ID,
		InnID:         queryReview.InnID,
		TouristSpotID: queryReview.TouristSpotID,
		Score:         queryReview.Score,
		Body:          queryReview.Body,
		FavoriteCount: queryReview.FavoriteCount,
		Media:         c.ConvertReviewMediaList(queryReview.Medias),
		Views:         queryReview.Views,
		Accompanying:  queryReview.Accompanying.String(),
		CreatedAt:     model.TimeResponse(queryReview.Review.CreatedAt),
		UpdatedAt:     model.TimeResponse(queryReview.Review.UpdatedAt),
		TravelDate:    model.NewYearMonth(queryReview.TravelDate),
		Hashtag:       hashtags,
		CommentCount:  queryReview.CommentCount,
		Creator:       c.NewCreatorFromUser(queryReview.User),
		IsFavorited:   queryReview.IsFavorite,
	}
}

func (c Converters) convertQueryReviewToOutput(queryReview *entity.ReviewDetail) *output.Review {
	medias := make([]output.ReviewMedia, len(queryReview.Medias))
	hashtags := make([]output.Hashtag, len(queryReview.Hashtag))
	for i, media := range queryReview.Medias {
		medias[i] = output.ReviewMedia{
			UUID: media.ID,
			Mime: media.MimeType,
			URL:  media.URL(c.filesURL()),
		}
	}
	for i, hashtag := range queryReview.Hashtag {
		hashtags[i] = output.Hashtag{
			ID:   hashtag.ID,
			Name: hashtag.Name,
		}
	}

	return &output.Review{
		ID:            queryReview.ID,
		InnID:         queryReview.InnID,
		TouristSpotID: queryReview.TouristSpotID,
		Score:         queryReview.Score,
		Body:          queryReview.Body,
		FavoriteCount: queryReview.FavoriteCount,
		Media:         c.ConvertReviewMediaList(queryReview.Medias),
		Views:         queryReview.Views,
		Accompanying:  queryReview.Accompanying.String(),
		CreatedAt:     model.TimeResponse(queryReview.Review.CreatedAt),
		UpdatedAt:     model.TimeResponse(queryReview.Review.UpdatedAt),
		TravelDate:    model.NewYearMonth(queryReview.TravelDate),
		CommentCount:  queryReview.CommentCount,
		Hashtag:       hashtags,
		Creator:       c.NewCreatorFromUser(queryReview.User),
	}
}

func (c Converters) ConvertCreateReviewParamToCommand(param *input.StoreReviewParam) *command.CreateReview {
	uuids := make([]*command.CreateReviewMedia, len(param.MediaUUIDs))
	for i, media := range param.MediaUUIDs {
		uuids[i] = &command.CreateReviewMedia{
			UUID:     media.UUID,
			MimeType: media.MimeType,
		}
	}

	return &command.CreateReview{
		TravelDate:    param.TravelDate,
		Accompanying:  param.Accompanying,
		TouristSpotID: param.TouristSpotID,
		InnID:         param.InnID,
		Score:         param.Score,
		Body:          param.Body,
		MediaUUIDs:    uuids,
	}
}

func (c Converters) ConvertUpdateReviewParamToCommand(param *input.UpdateReviewParam) *command.UpdateReview {
	uuids := make([]*command.CreateReviewMedia, len(param.MediaUUIDs))
	for i, media := range param.MediaUUIDs {
		uuids[i] = &command.CreateReviewMedia{
			UUID:     media.UUID,
			MimeType: media.MimeType,
		}
	}

	return &command.UpdateReview{
		TravelDate:   param.TravelDate,
		Accompanying: param.Accompanying,
		Score:        param.Score,
		Body:         param.Body,
		MediaUUIDs:   uuids,
	}
}

func (c Converters) ConvertReviewMediaList(reviewMediaList entity.ReviewMediaList) []output.ReviewMedia {
	reviewMediaList.Sort()
	medias := make([]output.ReviewMedia, len(reviewMediaList))
	for i, media := range reviewMediaList {
		medias[i] = output.ReviewMedia{
			UUID:     media.ID,
			Mime:     media.MimeType,
			URL:      media.URL(c.filesURL()),
			Priority: media.Priority,
		}
	}

	return medias
}
