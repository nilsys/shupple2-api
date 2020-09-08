package firebase

import (
	"context"
	"fmt"

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

func (r *CloudMessageRepositoryImpl) Send(token, body string, data map[string]string, badge int) error {
	// マストのdata
	data[androidDefaultDataKey] = androidDefaultDataVal

	msg := &messaging.Message{
		Token: token,
		Data:  data,
		Notification: &messaging.Notification{
			Title: "Stayway",
			Body:  body,
		},
		APNS: &messaging.APNSConfig{
			Payload: &messaging.APNSPayload{
				Aps: &messaging.Aps{
					Badge: &badge,
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

func (r *CloudMessageRepositoryForLocalImpl) Send(token, body string, data map[string]string, badge int) error {
	var dataStr []string

	for k, v := range data {
		dataStr = append(dataStr, fmt.Sprintf("%s: %s", k, v))
	}

	logger.Info("Push Notification", zap.String("Token", token), zap.String("Body", body), zap.Strings("Data", dataStr), zap.Int("Badge", badge))
	return nil
}
