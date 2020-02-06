package repository

import "github.com/stayway-corp/stayway-media-api/pkg/domain/entity"

type UserQueryRepository interface {
	FindByID(id int) (*entity.User, error)
}
