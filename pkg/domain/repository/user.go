package repository

import "github.com/stayway-corp/stayway-media-api/pkg/domain/entity"

type (
	UserCommandRepository interface {
		Store(user *entity.User) error
		StoreWithAvatar(user *entity.User, avatar []byte) error
	}

	UserQueryRepository interface {
		FindByID(id int) (*entity.User, error)
		FindByWordpressID(id int) (*entity.User, error)
	}
)
