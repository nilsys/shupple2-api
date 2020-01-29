package repository

import "github.com/stayway-corp/stayway-media-api/pkg/domain/entity"

type PostCommandRepository interface {
	Store(post *entity.Post) error
}
