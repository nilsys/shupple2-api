package converter

import (
	"encoding/json"

	"github.com/pkg/errors"

	"github.com/stayway-corp/stayway-media-api/pkg/adaptor/api/input"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/model/command"
)

func (c Converters) ConvertReportToCmd(p *input.Report) *command.Report {
	return &command.Report{
		TargetID:   p.TargetID,
		TargetType: p.TargetType,
		Reason:     p.Reason,
		Body:       p.Body,
	}
}

func (c Converters) ConvertSlackReportCallbackPayloadToCmd(p *input.SlackCallbackPayload) (*command.MarkAsReport, error) {
	src := input.SlackCallback{}
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
