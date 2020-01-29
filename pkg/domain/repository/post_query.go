package repository

import "github.com/stayway-corp/stayway-media-api/pkg/domain/entity"

type PostQueryRepository interface {
	FindByID(id int) (*entity.Post, error)
}
