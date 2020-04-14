package repository

import (
	"github.com/pkg/errors"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/entity/slack"

	"github.com/google/wire"
	"github.com/stayway-corp/stayway-media-api/pkg/adaptor/infrastructure/client"
	"github.com/stayway-corp/stayway-media-api/pkg/config"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/repository"
)

type SlackRepositoryImpl struct {
	SlackConfig config.Slack
	EnvConfig   config.Env
	Client      client.Client
}

const (
	slackPostMsgURL = "https://slack.com/api/chat.postMessage"
)

var SlackRepositorySet = wire.NewSet(
	wire.Struct(new(SlackRepositoryImpl), "*"),
	wire.Bind(new(repository.SlackRepository), new(*SlackRepositoryImpl)),
)

func (r *SlackRepositoryImpl) SendReport(report *slack.Report) error {
	opts := &client.Option{
		QueryParams: map[string][]string{},
	}

	// TODO: 暫定対応
	var res slack.Response

	opts.QueryParams.Add("blocks", report.ToSlackFmt())
	opts.QueryParams.Add("token", r.SlackConfig.Token)
	opts.QueryParams.Add("channel", r.SlackConfig.ReportChannel)

	if err := r.Client.GetJSON(slackPostMsgURL, opts, &res); err != nil {
		return errors.Wrapf(err, "failed to post slack")
	}

	return nil
}
