package converter

import (
	"github.com/stayway-corp/stayway-media-api/pkg/adaptor/api/param"
	"github.com/stayway-corp/stayway-media-api/pkg/adaptor/api/response"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/entity"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/model"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/model/command"
)

func ConvertCreateReviewCommentReplyParamToCommand(param *param.CreateReviewCommentReply) *command.CreateReviewCommentReply {
	return &command.CreateReviewCommentReply{
		ReviewCommentID: param.ReviewCommentID,
		Body:            param.Body,
	}
}

func ConvertReviewCommentReplyListToOutput(r []*entity.ReviewCommentReply) []*response.ReviewCommentReply {
	response := make([]*response.ReviewCommentReply, len(r))
	for i, reply := range r {
		response[i] = convertReviewCommentReplyToOutput(reply)
	}
	return response
}

func convertReviewCommentReplyToOutput(r *entity.ReviewCommentReply) *response.ReviewCommentReply {
	user := response.UserSummary{
		ID:   r.User.ID,
		Name: r.User.Name,
		Icon: r.User.GenerateThumbnailURL(),
	}
	return &response.ReviewCommentReply{
		ID:          r.ID,
		UserSummary: &user,
		Body:        r.Body,
		CreatedAt:   model.TimeResponse(r.CreatedAt),
	}
}
