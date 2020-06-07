package command

import "github.com/stayway-corp/stayway-media-api/pkg/domain/model"

type (
	Report struct {
		TargetID   int
		TargetType model.ReportTargetType
		Reason     model.ReportReasonType
		Body       string
	}

	MarkAsReport struct {
		UserID     int
		TargetID   int
		TargetType model.ReportTargetType
		IsApproved bool
	}
)
