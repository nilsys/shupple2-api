package repository

import (
	"github.com/stayway-corp/stayway-media-api/pkg/domain/entity/slack"
)

type (
	SlackRepository interface {
		SendReport(report *slack.Report) error
	}
)
