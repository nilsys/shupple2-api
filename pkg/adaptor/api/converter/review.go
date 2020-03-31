package converter

import (
	"github.com/stayway-corp/stayway-media-api/pkg/adaptor/api/param"
	"github.com/stayway-corp/stayway-media-api/pkg/adaptor/api/response"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/entity"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/model"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/model/command"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/model/query"
)

// i/oの構造体からレポジトリで使用するクエリ発行用構造体へコンバート
func ConvertFindReviewListParamToQuery(param *param.ListReviewParams) *query.ShowReviewListQuery {
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
		ChildID:                param.ChildID,
		SortBy:                 param.SortBy,
		Limit:                  param.GetLimit(),
		OffSet:                 param.GetOffset(),
	}
}

func ConvertListFeedReviewParamToQuery(param *param.ListFeedReviewParam) *query.FindListPaginationQuery {
	return &query.FindListPaginationQuery{
		Limit:  param.GetLimit(),
		Offset: param.GetOffset(),
	}
}

func ConvertReviewCommentListToOutput(reviewComments []*entity.ReviewComment) []*response.ReviewComment {
	reviewCommentOutputs := make([]*response.ReviewComment, len(reviewComments))
	for i, reviewComment := range reviewComments {
		reviewCommentOutputs[i] = convertReviewCommentToOutput(reviewComment)
	}
	return reviewCommentOutputs
}

func convertReviewCommentToOutput(reviewComment *entity.ReviewComment) *response.ReviewComment {
	userSummary := response.NewUserSummary(reviewComment.User.ID, reviewComment.User.UID, reviewComment.User.Name, reviewComment.User.GenerateThumbnailURL())
	return response.NewReviewComment(
		userSummary,
		reviewComment.Body,
		model.TimeResponse(reviewComment.CreatedAt),
		reviewComment.ID,
		reviewComment.ReplyCount,
		reviewComment.FavoriteCount,
	)
}

func ConvertQueryReviewListToOutput(queryReviews []*entity.QueryReview) []*response.Review {
	responses := make([]*response.Review, len(queryReviews))
	for i, queryReview := range queryReviews {
		responses[i] = convertQueryReviewToOutput(queryReview)
	}
	return responses
}

func convertQueryReviewToOutput(queryReview *entity.QueryReview) *response.Review {
	medias := make([]response.ReviewMedia, queryReview.MediaCount)
	hashtags := make([]response.Hashtag, len(queryReview.Hashtag))
	for i, media := range queryReview.Medias {
		medias[i] = response.ReviewMedia{
			UUID: media.ID,
			Mime: media.MimeType,
			URL:  media.GenerateURL(),
		}
	}
	for i, hashtag := range queryReview.Hashtag {
		hashtags[i] = response.Hashtag{
			ID:   hashtag.ID,
			Name: hashtag.Name,
		}
	}

	return &response.Review{
		ID:            queryReview.ID,
		InnID:         queryReview.InnID,
		TouristSpotID: queryReview.TouristSpotID,
		Score:         queryReview.Score,
		Body:          queryReview.Body,
		FavoriteCount: queryReview.FavoriteCount,
		Media:         medias,
		Views:         queryReview.Views,
		Accompanying:  queryReview.Accompanying.String(),
		UpdatedAt:     model.TimeFmtToFrontStr(queryReview.Review.UpdatedAt),
		TravelDate:    model.NewYearMonth(queryReview.TravelDate),
		Hashtag:       hashtags,
		CommentCount:  queryReview.CommentCount,
		Creator:       response.NewCreator(queryReview.User.ID, queryReview.User.UID, queryReview.User.GenerateThumbnailURL(), queryReview.User.Name, queryReview.Body),
	}
}

func ConvertQueryReviewShowToOutput(r *entity.QueryReview) *response.Review {
	medias := make([]response.ReviewMedia, r.MediaCount)
	hashtags := make([]response.Hashtag, len(r.Hashtag))
	for i, media := range r.Medias {
		medias[i] = response.ReviewMedia{
			UUID: media.ID,
			Mime: media.MimeType,
			URL:  media.GenerateURL(),
		}
	}
	for i, hashtag := range r.Hashtag {
		hashtags[i] = response.Hashtag{
			ID:   hashtag.ID,
			Name: hashtag.Name,
		}
	}

	return &response.Review{
		ID:            r.ID,
		InnID:         r.InnID,
		TouristSpotID: r.TouristSpotID,
		Score:         r.Score,
		Body:          r.Body,
		FavoriteCount: r.FavoriteCount,
		Media:         medias,
		Views:         r.Views,
		Accompanying:  r.Accompanying.String(),
		UpdatedAt:     model.TimeFmtToFrontStr(r.Review.UpdatedAt),
		TravelDate:    model.NewYearMonth(r.TravelDate),
		CommentCount:  r.CommentCount,
		Hashtag:       hashtags,
		Creator:       response.NewCreator(r.User.ID, r.User.UID, r.User.GenerateThumbnailURL(), r.User.Name, r.Body),
	}
}

func ConvertCreateReviewParamToCommand(param *param.StoreReviewParam) *command.CreateReview {
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

func ConvertUpdateReviewParamToCommand(param *param.UpdateReviewParam) *command.UpdateReview {
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
