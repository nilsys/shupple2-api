package repository

import (
	"context"

	"github.com/stayway-corp/stayway-media-api/pkg/domain/entity"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/model/query"
)

type (
	VlogCommandRepository interface {
		Lock(c context.Context, id int) (*entity.Vlog, error)
		Store(c context.Context, vlog *entity.Vlog) error
		UndeleteByID(c context.Context, id int) error
		DeleteByID(id int) error
		UpdateViewsByID(id, views int) error
		IncrementFavoriteCount(c context.Context, vlogID int) error
		DecrementFavoriteCount(c context.Context, vlogID int) error
		UpdateMonthlyViewsByID(id, views int) error
		UpdateWeeklyViewsByID(id, views int) error
		UpdateFacebookCountByID(id, count int) error
		UpdateTwitterCountByID(id, count int) error
	}

	VlogQueryRepository interface {
		FindAll() ([]*entity.Vlog, error)
		FindByID(id int) (*entity.Vlog, error)
		FindByLastID(lastID, limit int) ([]*entity.Vlog, error)
		FindListByParams(query *query.FindVlogListQuery) (*entity.VlogList, error)
		FindWithIsFavoriteListByParams(query *query.FindVlogListQuery, userID int) (*entity.VlogList, error)
		FindDetailByID(id int) (*entity.VlogDetail, error)
		FindDetailWithIsFavoriteByID(id, userID int) (*entity.VlogDetail, error)
		IsExist(id int) (bool, error)
	}

	VlogFavoriteCommandRepository interface {
		Store(c context.Context, favorite *entity.UserFavoriteVlog) error
		Delete(c context.Context, unfavorite *entity.UserFavoriteVlog) error
	}

	VlogFavoriteQueryRepository interface {
		IsExist(userID, vlogID int) (bool, error)
	}
)
