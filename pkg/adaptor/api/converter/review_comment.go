package converter

import (
	"github.com/stayway-corp/stayway-media-api/pkg/adaptor/api/input"
	"github.com/stayway-corp/stayway-media-api/pkg/adaptor/api/output"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/entity"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/model"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/model/command"
)

func (c Converters) ConvertCreateReviewCommentReplyParamToCommand(param *input.CreateReviewCommentReply) *command.CreateReviewCommentReply {
	return &command.CreateReviewCommentReply{
		ReviewCommentID: param.ReviewCommentID,
		Body:            param.Body,
	}
}

func (c Converters) ConvertReviewCommentReplyListToOutput(r []*entity.ReviewCommentReplyWithIsFavorite) []*output.ReviewCommentReply {
	response := make([]*output.ReviewCommentReply, len(r))
	for i, reply := range r {
		response[i] = c.ConvertReviewCommentReplyWithIsFavoriteToOutput(reply)
	}
	return response
}

func (c Converters) ConvertReviewCommentReplyWithIsFavoriteToOutput(r *entity.ReviewCommentReplyWithIsFavorite) *output.ReviewCommentReply {
	user := output.UserSummary{
		ID:      r.User.ID,
		Name:    r.User.Name,
		IconURL: r.User.AvatarURL(c.filesURL()),
	}
	return &output.ReviewCommentReply{
		ID:            r.ID,
		UserSummary:   &user,
		Body:          r.Body,
		FavoriteCount: r.FavoriteCount,
		IsFavorite:    r.IsFavorite,
		CreatedAt:     model.TimeResponse(r.CreatedAt),
	}
}

func (c Converters) ConvertReviewCommentReplyToOutput(r *entity.ReviewCommentReply) *output.ReviewCommentReply {
	user := output.UserSummary{
		ID:      r.User.ID,
		Name:    r.User.Name,
		IconURL: r.User.AvatarURL(c.filesURL()),
	}
	return &output.ReviewCommentReply{
		ID:            r.ID,
		UserSummary:   &user,
		Body:          r.Body,
		FavoriteCount: r.FavoriteCount,
		CreatedAt:     model.TimeResponse(r.CreatedAt),
	}
}
