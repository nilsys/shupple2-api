package main

import (
	"context"
	"encoding/json"

	"github.com/aws/aws-lambda-go/events"
	"github.com/pkg/errors"
	"github.com/stayway-corp/stayway-media-api/pkg/adaptor/logger"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/model"
	"go.uber.org/zap"
)

func (l *Lambda) persistMedia(ctx context.Context, msgs []events.SQSMessage) error {
	// NOTE: lambdaの設定で1つずつとるようにしている
	if len(msgs) != 1 {
		return errors.Errorf("strange messages count(%d)", len(msgs))
	}
	message := msgs[0]

	var req model.PersistMediaRequest
	if err := json.Unmarshal([]byte(message.Body), &req); err != nil {
		return errors.Wrapf(err, "failed to unmarshal request(messageId: %s)", message.MessageId)
	}

	logger.Info("persit_media", zap.Reflect("request", req))

	if err := l.MediaCommandService.Persist(&req); err != nil {
		return errors.Wrapf(err, "failed to persist media(uuid:%s)", req.UUID)
	}

	return nil
}
