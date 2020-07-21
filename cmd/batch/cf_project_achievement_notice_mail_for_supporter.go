package main

import (
	"time"

	"github.com/stayway-corp/stayway-media-api/pkg/adaptor/logger"

	"github.com/urfave/cli"
)

const (
	tickTimeMinutes = 10 * time.Minute
)

func (b *Batch) cliCfProjectAchievementMail() cli.Command {
	return cli.Command{
		Name:   "cf_project_achievement_notice_mail_for_supporter",
		Usage:  "達成金額に達しているプロジェクトをサポート（購入）したユーザーに通知のメールを送る",
		Action: b.cfProjectAchievementMailTickWrapper,
	}
}

func (b *Batch) cfProjectAchievementMailTickWrapper(c *cli.Context) {
	for range time.Tick(tickTimeMinutes) {
		if err := b.CfProjectFacade.SendAchievementMail(); err != nil {
			logger.Error(err.Error())
		}
	}
}
