package repository

import (
	"github.com/stayway-corp/stayway-media-api/pkg/domain/entity"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/model/query"
)

type (
	UserCommandRepository interface {
		Store(user *entity.User) error
		StoreWithAvatar(user *entity.User, avatar []byte) error
	}

	UserQueryRepository interface {
		FindByID(id int) (*entity.User, error)
		FindByCognitoID(cognitoID string) (*entity.User, error)
		FindByWordpressID(id int) (*entity.User, error)
		FindUserRankingListByParams(query *query.FindUserRankingListQuery) ([]*entity.QueryRankingUser, error)
		IsExistByUID(uid string) (bool, error)
		// name部分一致検索
		SearchByName(name string) ([]*entity.User, error)
		FindFolloweeByID(query *query.FindFollowUser) ([]*entity.User, error)
		FindFollowerByID(query *query.FindFollowUser) ([]*entity.User, error)
	}
)
