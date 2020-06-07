package entity

import (
	"github.com/stayway-corp/stayway-media-api/pkg/domain/model"
	"gopkg.in/guregu/null.v3"
)

type (
	Report struct {
		ID         int `gorm:"primary_key"`
		UserID     int
		TargetID   int
		TargetType model.ReportTargetType
		Reason     model.ReportReasonType
		Body       null.String
		IsDone     bool
		Times
	}
)

func NewReport(userID int, targetID int, targetType model.ReportTargetType, targetReason model.ReportReasonType, body string) *Report {
	return &Report{
		UserID:     userID,
		TargetID:   targetID,
		TargetType: targetType,
		Body:       null.StringFrom(body),
		Reason:     targetReason,
		IsDone:     false,
	}
}
