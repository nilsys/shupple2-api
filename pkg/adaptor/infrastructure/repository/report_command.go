package repository

import (
	"context"

	"github.com/stayway-corp/stayway-media-api/pkg/domain/model/command"

	"github.com/google/wire"
	"github.com/pkg/errors"
	"github.com/stayway-corp/stayway-media-api/pkg/adaptor/infrastructure/client"
	"github.com/stayway-corp/stayway-media-api/pkg/config"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/entity"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/repository"
)

type ReportCommandRepositoryImpl struct {
	DAO
	SlackConfig config.Slack
	Client      client.Client
}

var ReportCommandRepositorySet = wire.NewSet(
	wire.Struct(new(ReportCommandRepositoryImpl), "*"),
	wire.Bind(new(repository.ReportCommandRepository), new(*ReportCommandRepositoryImpl)),
)

func (r *ReportCommandRepositoryImpl) Store(c context.Context, report *entity.Report) error {
	if err := r.DB(c).Save(report).Error; err != nil {
		return errors.Wrap(err, "failed to store report")
	}
	return nil
}

func (r *ReportCommandRepositoryImpl) MarkAsDone(c context.Context, cmd *command.MarkAsReport) error {
	if err := r.DB(c).Table("report").Where("user_id = ?", cmd.UserID).Where("target_id = ?", cmd.TargetID).Where("target_type = ?", cmd.TargetType).Updates(map[string]interface{}{"is_done": true}).Error; err != nil {
		return errors.Wrap(err, "failed to mark as done")
	}
	return nil
}
