package main

import (
	"context"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/pkg/errors"

	"github.com/stayway-corp/stayway-media-api/pkg/application/service"
	"github.com/stayway-corp/stayway-media-api/pkg/config"
)

const (
	envNameAction = "action"
	envNameEnv    = "env"

	actionPersistMedia = "persist_media"
)

type Lambda struct {
	Config              *config.Config
	MediaCommandService service.MediaCommandService
}

func main() {
	lambda.Start(func(ctx context.Context, sqsEvent events.SQSEvent) error {
		envStr, _ := os.LookupEnv(envNameEnv)

		lambda, err := InitializeLambda(EnvString(envStr))
		if err != nil {
			return err
		}

		action, _ := os.LookupEnv(envNameAction)
		return lambda.Run(ctx, action, sqsEvent)
	})
}

func (l *Lambda) Run(ctx context.Context, action string, sqsEvent events.SQSEvent) error {
	switch action {
	case actionPersistMedia:
		return l.persistMedia(ctx, sqsEvent.Records)
	}

	return errors.Errorf("unknown action:  %s", action)
}
