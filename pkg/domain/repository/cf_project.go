package repository

import (
	"context"

	"github.com/stayway-corp/stayway-media-api/pkg/domain/entity"

	"github.com/stayway-corp/stayway-media-api/pkg/domain/model/query"
)

type (
	CfProjectCommandRepository interface {
		Store(*entity.CfProject) error
		Lock(c context.Context, id int) (*entity.CfProject, error)
		UndeleteByID(c context.Context, id int) error
		DeleteByID(id int) error
		StoreUserFavoriteCfProject(c context.Context, fav *entity.UserFavoriteCfProject) error
		DeleteUserFavoriteCfProject(c context.Context, fav *entity.UserFavoriteCfProject) error
		IncrementFavoriteCountByID(c context.Context, projectID int) error
		DecrementFavoriteCountByID(c context.Context, projectID int) error
		StoreSupportComment(c context.Context, comment *entity.CfProjectSupportCommentTiny) error
		IncrementSupportCommentCount(c context.Context, id int) error
		IncrementAchievedPrice(c context.Context, id, price int) error
		DecrementAchievedPrice(c context.Context, id, price int) error
		MarkAsIsSentAchievementNoticeMail(id int) error
	}

	CfProjectQueryRepository interface {
		FindByID(id int) (*entity.CfProjectDetail, error)
		FindListByQuery(query *query.FindCfProjectQuery) (*entity.CfProjectDetailList, error)
		FindSupportCommentListByCfProjectID(projectID, limit int) ([]*entity.CfProjectSupportComment, error)
		FindNotSentAchievementNoticeMailAndAchievedListByLastID(lastID, limit int) (*entity.CfProjectDetailList, error)
	}
)
