package repository

import "github.com/stayway-corp/stayway-media-api/pkg/domain/model"

type (
	MediaQueryRepository interface {
		GetUploadedMedia(uuid string) (*model.MediaBody, error)
	}

	MediaCommandRepository interface {
		SavePersistRequest(*model.PersistMediaRequest) error
		Save(mediaBody *model.MediaBody, destination string) error
	}
)
