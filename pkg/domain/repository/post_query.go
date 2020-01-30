package repository

import (
	"github.com/stayway-corp/stayway-media-api/pkg/domain/entity"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/model/query"
)

type PostQueryRepository interface {
	FindByID(id int) (*entity.Post, error)
	// TODO: 命名
	FindByParams(query *query.FindPostListQuery) ([]*entity.Post, error)
}
