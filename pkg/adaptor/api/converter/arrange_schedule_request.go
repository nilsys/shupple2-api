package converter

import (
	"time"

	"github.com/uma-co82/shupple2-api/pkg/adaptor/api/input"
	"github.com/uma-co82/shupple2-api/pkg/domain/model/command"
)

func (c Converters) ConvertStoreArrangeScheduleRequestInputToCmd(input *input.StoreArrangeScheduleRequest) *command.StoreArrangeScheduleRequest {
	return &command.StoreArrangeScheduleRequest{
		MatchingUserID: input.MatchingUserID.ID,
		Date:           time.Time(input.DateTime),
		Remark:         input.Remark,
	}
}
