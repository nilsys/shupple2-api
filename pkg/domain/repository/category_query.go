package repository

import "github.com/stayway-corp/stayway-media-api/pkg/domain/entity"

type CategoryQueryRepository interface {
	FindByIDs(ids ...int) ([]*entity.Category, error)
}
