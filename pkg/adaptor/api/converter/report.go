package converter

import (
	"encoding/json"

	"github.com/pkg/errors"

	"github.com/stayway-corp/stayway-media-api/pkg/adaptor/api/param"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/model/command"
)

func ConvertReportToCmd(p *param.Report) *command.Report {
	return &command.Report{
		TargetID:   p.TargetID,
		TargetType: p.TargetType,
		Reason:     p.Reason,
	}
}

func ConvertSlackReportCallbackPayloadToCmd(p *param.SlackCallbackPayload) (*command.MarkAsReport, error) {
	src := param.SlackCallback{}
	if err := json.Unmarshal([]byte(p.Payload), &src); err != nil {
		return nil, errors.Wrap(err, "invalid slack report callback response type")
	}

	return &command.MarkAsReport{
		UserID:     src.ReportUserID(),
		TargetID:   src.TargetID(),
		TargetType: src.TargetType(),
		IsApproved: src.IsApproved(),
	}, nil
}
