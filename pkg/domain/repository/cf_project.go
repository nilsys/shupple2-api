package repository

import (
	"context"

	"github.com/stayway-corp/stayway-media-api/pkg/domain/entity"

	"github.com/stayway-corp/stayway-media-api/pkg/domain/model/query"
)

type (
	CfProjectCommandRepository interface {
		StoreUserFavoriteCfProject(c context.Context, fav *entity.UserFavoriteCfProject) error
		DeleteUserFavoriteCfProject(c context.Context, fav *entity.UserFavoriteCfProject) error
		IncrementFavoriteCountByID(c context.Context, projectID int) error
		DecrementFavoriteCountByID(c context.Context, projectID int) error
		StoreSupportComment(c context.Context, comment *entity.CfProjectSupportCommentTable) error
		IncrementSupportCommentCount(c context.Context, id int) error
		IncrementAchievedPrice(c context.Context, id, price int) error
	}

	CfProjectQueryRepository interface {
		FindByID(id int) (*entity.CfProject, error)
		FindListByQuery(query *query.FindCfProjectQuery) (*entity.CfProjectList, error)
		Lock(c context.Context, id int) (*entity.CfProject, error)
		LockCfProjectListByIDs(c context.Context, ids []int) (*entity.CfProjectList, error)
		FindSupportCommentListByCfProjectID(projectID, limit int) ([]*entity.CfProjectSupportComment, error)
	}
)
