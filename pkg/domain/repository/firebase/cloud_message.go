package firebase

import firebaseEntity "github.com/stayway-corp/stayway-media-api/pkg/domain/entity/firebase"

type (
	CloudMessageCommandRepository interface {
		Send(cloudMsgData *firebaseEntity.CloudMessageData) error
	}
)
