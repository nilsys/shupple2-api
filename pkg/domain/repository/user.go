package repository

import (
	"context"
	"io"

	"github.com/aws/aws-sdk-go/service/cognitoidentityprovider"

	"github.com/stayway-corp/stayway-media-api/pkg/domain/model/command"

	"github.com/stayway-corp/stayway-media-api/pkg/domain/entity"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/model/query"
)

type (
	UserCommandRepository interface {
		Store(user *entity.User) error
		Update(user *entity.User) error
		StoreWithAvatar(user *entity.User, avatar io.Reader, contentType string) error
		UpdateWordpressID(userID, wordpressUserID int) error
		StoreFollow(c context.Context, following *entity.UserFollowing, followed *entity.UserFollowed) error
		DeleteFollow(userID, targetID int) error
		PersistUserImage(cmd *command.UpdateUser) error
	}

	UserQueryRepository interface {
		FindByID(id int) (*entity.User, error)
		FindByUIDs(uIDs []string) ([]*entity.User, error)
		FindByCognitoID(cognitoID string) (*entity.User, error)
		FindByWordpressID(id int) (*entity.User, error)
		FindByMigrationCode(code string) (*entity.UserTable, error)
		FindUserRankingListByParams(query *query.FindUserRankingListQuery) ([]*entity.UserDetail, error)
		FindByUID(uid string) (*entity.UserTable, error)
		FindUserDetailWithCountByID(id int) (*entity.UserDetailWithMediaCount, error)
		IsFollow(targetID, userID int) (bool, error)
		FindRecommendFollowUserList(interestIDs []int) ([]*entity.UserTable, error)
		IsExistByUID(uid string) (bool, error)
		FindByCognitoUserName(cognitoUserName []string) ([]*entity.UserTable, error)
		// name部分一致検索
		SearchByName(name string) ([]*entity.User, error)
		FindFollowingByID(query *query.FindFollowUser) ([]*entity.User, error)
		FindFollowedByID(query *query.FindFollowUser) ([]*entity.User, error)
		FindFavoritePostUser(postID int, query *query.FindListPaginationQuery) ([]*entity.User, error)
		FindFavoriteReviewUser(reviewID int, query *query.FindListPaginationQuery) ([]*entity.User, error)
		FindFavoritePostUserByUserID(postID, userID int, query *query.FindListPaginationQuery) ([]*entity.User, error)
		FindFavoriteReviewUserByUserID(reviewID, userID int, query *query.FindListPaginationQuery) ([]*entity.User, error)
		// TODO: ここにあっていいか？
		FindConfirmedUserTypeByPhoneNumberFromCognito(number string) ([]*cognitoidentityprovider.UserType, error)
	}
)
