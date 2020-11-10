package output

import "github.com/uma-co82/shupple2-api/pkg/domain/model"

type (
	ArrangeScheduleRequest struct {
		ID                  int            `json:"id"`
		UserID              int            `json:"userId"`
		MatchingUserID      int            `json:"matchingUserId"`
		DateTime            model.DateTime `json:"dateTime"`
		Remark              string         `json:"remark"`
		StartNow            bool           `json:"startNow"`
		MatchingUserApprove bool           `json:"matchingUserApprove"`
		User                User           `json:"user"`
		MatchingUser        User           `json:"matchingUser"`
	}
)
