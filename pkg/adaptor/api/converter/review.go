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
		OwnerID:                param.OwnerID,
		InnID:                  param.InnID,
		TouristSpotID:          param.TouristSpotID,
		HashTag:                param.HashTag,
		AreaID:                 param.AreaID,
		SubAreaID:              param.SubAreaID,
		SubSubAreaID:           param.SubSubAreaID,
		MetasearchAreaID:       param.MetasearchAreaID,
		MetasearchSubAreaID:    param.MetasearchSubAreaID,
		MetasearchSubSubAreaID: param.MetasearchSubSubAreaID,
		TargetType:             param.TargetType,
		SortBy:                 param.SortBy,
		Keyword:                param.Keyward,
		ExcludeID:              param.ExcludeID,
		Limit:                  param.GetLimit(),
		OffSet:                 param.GetOffset(),
	}
}

func (c Converters) ConvertListFavoriteReviewParamToQuery(param *input.ListFavoriteReviewParam) *query.FindListPaginationQuery {
	return &query.FindListPaginationQuery{
		Limit:  param.GetLimit(),
		Offset: param.GetOffset(),
	}
}

func (c Converters) ConvertListFeedReviewInputToQuery(i *input.PaginationQuery) *query.FindListPaginationQuery {
	return &query.FindListPaginationQuery{
		Limit:  i.GetReviewLimit(),
		Offset: i.GetReviewOffset(),
	}
}

func (c Converters) ConvertReviewCommentToOutput(reviewComment *entity.ReviewCommentTiny) *output.ReviewComment {
	return output.NewReviewComment(
		c.NewUserSummaryFromUser(reviewComment.User),
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
	return output.NewReviewComment(
		c.NewUserSummaryFromUser(reviewComment.User),
		reviewComment.Body,
		model.TimeResponse(reviewComment.CreatedAt),
		reviewComment.ID,
		reviewComment.ReplyCount,
		reviewComment.FavoriteCount,
		reviewComment.IsFavorite,
	)
}

func (c Converters) ConvertQueryReviewDetailWithIsFavoriteListToOutput(reviews *entity.ReviewDetailWithIsFavoriteList, idRelationFlgMap *entity.UserRelationFlgMap) output.ReviewDetailIsFavoriteList {
	responses := make([]*output.Review, len(reviews.List))
	for i, queryReview := range reviews.List {
		responses[i] = c.ConvertQueryReviewDetailWithIsFavoriteToOutput(queryReview, idRelationFlgMap)
	}
	return output.ReviewDetailIsFavoriteList{
		TotalNumber: reviews.TotalNumber,
		Reviews:     responses,
	}
}

func (c Converters) ConvertQueryReviewDetailWithIsFavoriteToOutput(queryReview *entity.ReviewDetailWithIsFavorite, idRelationFlgMap *entity.UserRelationFlgMap) *output.Review {
	hashtags := make([]output.Hashtag, len(queryReview.Hashtag))
	for i, hashtag := range queryReview.Hashtag {
		hashtags[i] = output.Hashtag{
			ID:   hashtag.ID,
			Name: hashtag.Name,
		}
	}

	return &output.Review{
		ID:            queryReview.ID,
		InnID:         queryReview.InnID.Int64,
		TouristSpotID: queryReview.TouristSpotID.Int64,
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
		Creator:       c.NewCreatorFromUser(queryReview.User, idRelationFlgMap.IsFollowByUserID(queryReview.UserID), idRelationFlgMap.IsBlockingByUserID(queryReview.UserID)),
		IsFavorite:    queryReview.IsFavorite,
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
		ID:           param.ID,
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
