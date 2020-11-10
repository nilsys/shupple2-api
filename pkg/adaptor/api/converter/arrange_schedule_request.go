package converter

import (
	"time"

	"github.com/uma-co82/shupple2-api/pkg/adaptor/api/output"
	"github.com/uma-co82/shupple2-api/pkg/domain/entity"
	"github.com/uma-co82/shupple2-api/pkg/domain/model"

	"github.com/uma-co82/shupple2-api/pkg/adaptor/api/input"
	"github.com/uma-co82/shupple2-api/pkg/domain/model/command"
)

func (c Converters) ConvertStoreArrangeScheduleRequestInputToCmd(input *input.StoreArrangeScheduleRequest) *command.StoreArrangeScheduleRequest {
	return &command.StoreArrangeScheduleRequest{
		MatchingUserID: input.MatchingUserID.ID,
		Date:           time.Time(input.DateTime),
		Remark:         input.Remark,
		StartNow:       input.StartNow,
	}
}

func (c Converters) ConvertArrangeScheduleRequestList2Output(reqs []*entity.ArrangeScheduleRequest) []output.ArrangeScheduleRequest {
	resolve := make([]output.ArrangeScheduleRequest, len(reqs))

	for i, req := range reqs {
		resolve[i] = c.ConvertArrangeScheduleRequest2Output(req)
	}

	return resolve
}

func (c Converters) ConvertArrangeScheduleRequest2Output(req *entity.ArrangeScheduleRequest) output.ArrangeScheduleRequest {
	return output.ArrangeScheduleRequest{
		ID:                  req.ID,
		UserID:              req.UserID,
		MatchingUserID:      req.MatchingUserID,
		DateTime:            model.DateTime(req.DateTime),
		Remark:              req.Remark,
		StartNow:            req.StartNow.Bool,
		MatchingUserApprove: req.MatchingUserApprove,
		User:                c.ConvertUser2Output(req.User),
		MatchingUser:        c.ConvertUser2Output(req.MatchingUser),
	}
}
