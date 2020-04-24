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
func ConvertFindReviewListParamToQuery(param *input.ListReviewParams) *query.ShowReviewListQuery {
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
		Limit:                  param.GetLimit(),
		OffSet:                 param.GetOffset(),
	}
}

func ConvertListFeedReviewParamToQuery(param *input.ListFeedReviewParam) *query.FindListPaginationQuery {
	return &query.FindListPaginationQuery{
		Limit:  param.GetLimit(),
		Offset: param.GetOffset(),
	}
}

func ConvertReviewCommentToOutput(reviewComment *entity.ReviewComment) *output.ReviewComment {
	userSummary := output.NewUserSummary(reviewComment.User.ID, reviewComment.User.UID, reviewComment.User.Name, reviewComment.User.IconURL())
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

func ConvertReviewCommentWithIsFavoriteListToOutput(reviewComments []*entity.ReviewCommentWithIsFavorite) []*output.ReviewComment {
	reviewCommentOutputs := make([]*output.ReviewComment, len(reviewComments))
	for i, reviewComment := range reviewComments {
		reviewCommentOutputs[i] = ConvertReviewCommentWithIsFavoriteToOutput(reviewComment)
	}
	return reviewCommentOutputs
}

func ConvertReviewCommentWithIsFavoriteToOutput(reviewComment *entity.ReviewCommentWithIsFavorite) *output.ReviewComment {
	userSummary := output.NewUserSummary(reviewComment.User.ID, reviewComment.User.UID, reviewComment.User.Name, reviewComment.User.IconURL())
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

func ConvertQueryReviewListToOutput(queryReviews []*entity.ReviewDetail) []*output.Review {
	outputs := make([]*output.Review, len(queryReviews))
	for i, queryReview := range queryReviews {
		outputs[i] = convertQueryReviewToOutput(queryReview)
	}
	return outputs
}

func ConvertQueryReviewDetailWithIsFavoriteListToOutput(reviews []*entity.ReviewDetailWithIsFavorite) []*output.Review {
	responses := make([]*output.Review, len(reviews))
	for i, queryReview := range reviews {
		responses[i] = ConvertQueryReviewDetailWithIsFavoriteToOutput(queryReview)
	}
	return responses
}

func ConvertQueryReviewDetailWithIsFavoriteToOutput(queryReview *entity.ReviewDetailWithIsFavorite) *output.Review {
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
		Media:         ConvertReviewMediaList(queryReview.Medias),
		Views:         queryReview.Views,
		Accompanying:  queryReview.Accompanying.String(),
		UpdatedAt:     model.TimeFmtToFrontStr(queryReview.Review.UpdatedAt),
		TravelDate:    model.NewYearMonth(queryReview.TravelDate),
		Hashtag:       hashtags,
		CommentCount:  queryReview.CommentCount,
		Creator:       output.NewCreatorFromUser(queryReview.User),
		IsFavorited:   queryReview.IsFavorite,
	}
}

func convertQueryReviewToOutput(queryReview *entity.ReviewDetail) *output.Review {
	medias := make([]output.ReviewMedia, len(queryReview.Medias))
	hashtags := make([]output.Hashtag, len(queryReview.Hashtag))
	for i, media := range queryReview.Medias {
		medias[i] = output.ReviewMedia{
			UUID: media.ID,
			Mime: media.MimeType,
			URL:  media.GenerateURL(),
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
		Media:         ConvertReviewMediaList(queryReview.Medias),
		Views:         queryReview.Views,
		Accompanying:  queryReview.Accompanying.String(),
		UpdatedAt:     model.TimeFmtToFrontStr(queryReview.Review.UpdatedAt),
		TravelDate:    model.NewYearMonth(queryReview.TravelDate),
		CommentCount:  queryReview.CommentCount,
		Hashtag:       hashtags,
		Creator:       output.NewCreatorFromUser(queryReview.User),
	}
}

func ConvertCreateReviewParamToCommand(param *input.StoreReviewParam) *command.CreateReview {
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

func ConvertUpdateReviewParamToCommand(param *input.UpdateReviewParam) *command.UpdateReview {
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

func ConvertReviewMediaList(reviewMediaList entity.ReviewMediaList) []output.ReviewMedia {
	reviewMediaList.Sort()
	medias := make([]output.ReviewMedia, len(reviewMediaList))
	for i, media := range reviewMediaList {
		medias[i] = output.ReviewMedia{
			UUID:     media.ID,
			Mime:     media.MimeType,
			URL:      media.GenerateURL(),
			Priority: media.Priority,
		}
	}

	return medias
}
