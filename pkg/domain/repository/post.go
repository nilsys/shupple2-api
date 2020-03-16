package repository

import (
	"context"

	"github.com/stayway-corp/stayway-media-api/pkg/domain/entity"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/model/query"
)

type (
	PostCommandRepository interface {
		Store(c context.Context, post *entity.Post) error
		DeleteByID(id int) error
	}

	PostQueryRepository interface {
		FindByID(id int) (*entity.Post, error)
		FindQueryShowByID(id int) (*entity.PostDetailWithHashtag, error)
		FindListByParams(query *query.FindPostListQuery) (*entity.PostDetailList, error)
		FindFeedListByUserID(userID int, query *query.FindListPaginationQuery) ([]*entity.PostDetail, error)
	}
)
