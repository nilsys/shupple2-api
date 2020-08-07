package main

import (
	"time"

	"github.com/stayway-corp/stayway-media-api/pkg/adaptor/logger"

	"github.com/urfave/cli"
)

const (
	tickTimeMinutes = 5 * time.Minute
)

func (b *Batch) cliCfProjectNotice() cli.Command {
	return cli.Command{
		Name:  "cf_project_notice",
		Usage: "cf_project関連の通知メール、サブコマンドで指定",
		Subcommands: []cli.Command{
			b.cliCfProjectAchievementEmail(),
			b.cliCfProjectNewPostEmail(),
		},
	}
}

func (b *Batch) cliCfProjectAchievementEmail() cli.Command {
	return cli.Command{
		Name:   "achievement",
		Usage:  "達成金額に達しているプロジェクトをサポート（購入）したユーザーに通知のメールを送る",
		Action: b.cfProjectAchievementEmailTickWrapper,
	}
}

func (b *Batch) cliCfProjectNewPostEmail() cli.Command {
	return cli.Command{
		Name:   "new_post",
		Usage:  "新たに報告（Post）が投稿されたプロジェクトをサポート（購入）したユーザーに通知のメールを送る",
		Action: b.cfProjectNewPostEmailTickWrapper,
	}
}

// CfProject達成メール
func (b *Batch) cfProjectAchievementEmailTickWrapper(c *cli.Context) {
	for range time.Tick(tickTimeMinutes) {
		if err := b.CfProjectFacade.SendAchievementEmail(); err != nil {
			logger.Error(err.Error())
		}
	}
}

// CfProjectに新しい報告(post)が投稿された通知メール
func (b *Batch) cfProjectNewPostEmailTickWrapper(c *cli.Context) {
	for range time.Tick(tickTimeMinutes) {
		if err := b.CfProjectFacade.SendNewPostEmail(); err != nil {
			logger.Error(err.Error())
		}
	}
}
