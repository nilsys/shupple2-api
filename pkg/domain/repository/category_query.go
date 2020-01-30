package repository

import "github.com/stayway-corp/stayway-media-api/pkg/domain/entity"

type CategoryQueryRepository interface {
	FindAreaByIDs(ids ...int) ([]*entity.Category, error)
	FindThemeByIDs(ids ...int) ([]*entity.Category, error)
}
