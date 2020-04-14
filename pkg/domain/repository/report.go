package repository

import (
	"context"

	"github.com/stayway-corp/stayway-media-api/pkg/domain/model/command"

	"github.com/stayway-corp/stayway-media-api/pkg/domain/model"

	"github.com/stayway-corp/stayway-media-api/pkg/domain/entity"
)

type (
	ReportCommandRepository interface {
		Store(c context.Context, report *entity.Report) error
		MarkAsDone(c context.Context, cmd *command.MarkAsReport) error
	}

	ReportQueryRepository interface {
		IsExist(userID, targetID int, targetType model.ReportTargetType) (bool, error)
	}
)
