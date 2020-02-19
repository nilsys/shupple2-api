package repository

import (
	"github.com/stayway-corp/stayway-media-api/pkg/domain/entity"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/model/query"
)

type (
	PostCommandRepository interface {
		Store(post *entity.Post) error
	}

	PostQueryRepository interface {
		FindByID(id int) (*entity.Post, error)
		FindQueryByID(id int) (*entity.QueryPost, error)
		FindListByParams(query *query.FindPostListQuery) ([]*entity.QueryPost, error)
		FindFeedListByUserID(userID int, query *query.FindListPaginationQuery) ([]*entity.QueryPost, error)
	}
)
