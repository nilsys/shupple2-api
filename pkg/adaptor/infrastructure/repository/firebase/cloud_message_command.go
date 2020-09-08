package firebase

import (
	"context"
	"fmt"

	firebaseEntity "github.com/stayway-corp/stayway-media-api/pkg/domain/entity/firebase"

	"go.uber.org/zap"

	"firebase.google.com/go/messaging"
	"github.com/pkg/errors"
	"github.com/stayway-corp/stayway-media-api/pkg/adaptor/logger"
)

type (
	CloudMessageRepositoryImpl struct {
		Client *messaging.Client
	}

	CloudMessageRepositoryForLocalImpl struct {
	}
)

const (
	androidDefaultDataKey = "click_action"
	androidDefaultDataVal = "FLUTTER_NOTIFICATION_CLICK"
)

func (r *CloudMessageRepositoryImpl) Send(cloudMsgData *firebaseEntity.CloudMessageData) error {
	// マストのdata
	cloudMsgData.AddData(androidDefaultDataKey, androidDefaultDataVal)

	msg := &messaging.Message{
		Token: cloudMsgData.Token,
		Data:  cloudMsgData.Data,
		Notification: &messaging.Notification{
			Title: "Stayway",
			Body:  cloudMsgData.Body,
		},
		APNS: &messaging.APNSConfig{
			Payload: &messaging.APNSPayload{
				Aps: &messaging.Aps{
					Badge: &cloudMsgData.Badge,
				},
			},
		},
	}

	_, err := r.Client.Send(context.Background(), msg)
	if err != nil {
		return errors.Wrap(err, "failed send fcm")
	}

	return nil
}

func (r *CloudMessageRepositoryForLocalImpl) Send(cloudMsgData *firebaseEntity.CloudMessageData) error {
	var dataStr []string

	cloudMsgData.AddData(androidDefaultDataKey, androidDefaultDataVal)

	for k, v := range cloudMsgData.Data {
		dataStr = append(dataStr, fmt.Sprintf("%s: %s", k, v))
	}

	logger.Info("Push Notification", zap.String("Token", cloudMsgData.Token), zap.String("Body", cloudMsgData.Body), zap.Strings("Data", dataStr), zap.Int("Badge", cloudMsgData.Badge))
	return nil
}
