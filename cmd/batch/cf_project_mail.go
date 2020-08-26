package main

import (
	"time"

	"github.com/stayway-corp/stayway-media-api/pkg/adaptor/logger"

	"github.com/urfave/cli"
)

const (
	fifthTickTimeMinutes  = 5 * time.Minute
	eighthTickTimeMinutes = 8 * time.Minute
)

func (b *Batch) cliCfProjectNotice() cli.Command {
	return cli.Command{
		Name:   "cf_project_notice",
		Usage:  "cf_project関連の通知メールを全て実行",
		Action: b.cfProjectNoticeFullWrapper,
	}
}

// fargateタスク数を減らす為に一度に動かす
func (b *Batch) cfProjectNoticeFullWrapper(c *cli.Context) {
	go b.cfProjectAchievementEmailTickWrapper(c)
	b.cfProjectNewPostEmailTickWrapper(c)
}

// CfProject達成メール
func (b *Batch) cfProjectAchievementEmailTickWrapper(c *cli.Context) {
	for range time.Tick(fifthTickTimeMinutes) {
		if err := b.CfProjectFacade.SendAchievementEmail(); err != nil {
			logger.Error(err.Error())
		}
	}
}

// CfProjectに新しい報告(post)が投稿された通知メール
func (b *Batch) cfProjectNewPostEmailTickWrapper(c *cli.Context) {
	for range time.Tick(eighthTickTimeMinutes) {
		if err := b.CfProjectFacade.SendNewPostEmail(); err != nil {
			logger.Error(err.Error())
		}
	}
}
