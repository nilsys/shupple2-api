package entity

import (
	"time"

	"github.com/stayway-corp/stayway-media-api/pkg/domain/model"
)

type (
	Report struct {
		ID         int `gorm:"primary_key"`
		UserID     int
		TargetID   int
		TargetType model.ReportTargetType
		Reason     model.ReportReasonType
		IsDone     bool
		CreatedAt  time.Time `gorm:"default:current_timestamp"`
	}
)

func NewReport(userID int, targetID int, targetType model.ReportTargetType, targetReason model.ReportReasonType) *Report {
	return &Report{
		UserID:     userID,
		TargetID:   targetID,
		TargetType: targetType,
		Reason:     targetReason,
		IsDone:     false,
	}
}
