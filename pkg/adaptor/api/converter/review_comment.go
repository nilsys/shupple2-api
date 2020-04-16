package converter

import (
	"github.com/stayway-corp/stayway-media-api/pkg/adaptor/api/input"
	"github.com/stayway-corp/stayway-media-api/pkg/adaptor/api/output"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/entity"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/model"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/model/command"
)

func ConvertCreateReviewCommentReplyParamToCommand(param *input.CreateReviewCommentReply) *command.CreateReviewCommentReply {
	return &command.CreateReviewCommentReply{
		ReviewCommentID: param.ReviewCommentID,
		Body:            param.Body,
	}
}

func ConvertReviewCommentReplyListToOutput(r []*entity.ReviewCommentReply) []*output.ReviewCommentReply {
	response := make([]*output.ReviewCommentReply, len(r))
	for i, reply := range r {
		response[i] = ConvertReviewCommentReplyToOutput(reply)
	}
	return response
}

func ConvertReviewCommentReplyToOutput(r *entity.ReviewCommentReply) *output.ReviewCommentReply {
	user := output.UserSummary{
		ID:      r.User.ID,
		Name:    r.User.Name,
		IconURL: r.User.IconURL(),
	}
	return &output.ReviewCommentReply{
		ID:          r.ID,
		UserSummary: &user,
		Body:        r.Body,
		CreatedAt:   model.TimeResponse(r.CreatedAt),
	}
}
