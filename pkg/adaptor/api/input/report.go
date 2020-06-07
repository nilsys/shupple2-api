package input

import "github.com/stayway-corp/stayway-media-api/pkg/domain/model"

type Report struct {
	TargetID   int                    `json:"targetId"`
	TargetType model.ReportTargetType `json:"targetType"`
	Reason     model.ReportReasonType `json:"reason"`
	Body       string                 `json:"body"`
}
