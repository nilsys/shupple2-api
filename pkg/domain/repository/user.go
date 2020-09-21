package repository

import (
	"context"
	"io"

	"github.com/aws/aws-sdk-go/service/cognitoidentityprovider"

	"github.com/stayway-corp/stayway-media-api/pkg/domain/entity"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/model/query"
)

type (
	UserCommandRepository interface {
		Store(ctx context.Context, user *entity.User) error
		UnscopedStore(ctx context.Context, user *entity.User) error
		Update(user *entity.User) error
		UpdateDeviceTokenByID(id int, deviceToken string) error
		StoreWithAvatar(user *entity.User, avatar io.Reader, contentType string) error
		UpdateWordpressID(userID, wordpressUserID int) error
		StoreFollow(c context.Context, following *entity.UserFollowing, followed *entity.UserFollowed) error
		DeleteFollow(userID, targetID int) error
	}

	UserQueryRepository interface {
		FindByID(id int) (*entity.User, error)
		FindByUIDs(uIDs []string) ([]*entity.User, error)
		FindByCognitoID(cognitoID string) (*entity.User, error)
		FindByWordpressID(id int) (*entity.User, error)
		FindByAssociateID(associateID string) (*entity.UserTiny, error)
		UnscopedFindByMigrationCode(code string) (*entity.UserTiny, error)
		FindByMigrationCode(code string) (*entity.UserTiny, error)
		FindUserRankingListByParams(query *query.FindUserRankingListQuery) ([]*entity.UserDetail, error)
		FindByUID(uid string) (*entity.UserTiny, error)
		UnscopedFindUserDetailWithCountByID(id int) (*entity.UserDetailWithMediaCount, error)
		FindUserDetailWithCountByID(id int) (*entity.UserDetailWithMediaCount, error)
		IsFollowing(baseUserID int, userIDs []int) (map[int]bool, error)
		FindRecommendFollowUserList(interestIDs []int) ([]*entity.UserTiny, error)
		IsExistByUID(uid string) (bool, error)
		FindByCognitoUserName(cognitoUserName []string) ([]*entity.UserTiny, error)
		SearchByNameOrUID(name string) ([]*entity.UserTiny, error)
		FindFollowingByID(query *query.FindFollowUser) ([]*entity.UserTinyWithIsFollow, error)
		FindFollowedByID(query *query.FindFollowUser) ([]*entity.UserTinyWithIsFollow, error)
		FindFollowingWithIsFollowByID(userID int, query *query.FindFollowUser) ([]*entity.UserTinyWithIsFollow, error)
		FindFollowedWithIsFollowByID(userID int, query *query.FindFollowUser) ([]*entity.UserTinyWithIsFollow, error)
		FindFavoritePostUser(postID int, query *query.FindListPaginationQuery) ([]*entity.UserTinyWithIsFollow, error)
		FindFavoriteReviewUser(reviewID int, query *query.FindListPaginationQuery) ([]*entity.UserTinyWithIsFollow, error)
		FindFavoriteComicUser(comicID int, query *query.FindListPaginationQuery) ([]*entity.UserTinyWithIsFollow, error)
		FindFavoriteVlogUser(vlogID int, query *query.FindListPaginationQuery) ([]*entity.UserTinyWithIsFollow, error)
		FindFavoritePostUserByUserID(postID, userID int, query *query.FindListPaginationQuery) ([]*entity.UserTinyWithIsFollow, error)
		FindFavoriteReviewUserByUserID(reviewID, userID int, query *query.FindListPaginationQuery) ([]*entity.UserTinyWithIsFollow, error)
		FindFavoriteVlogUserByUserID(vlogID, userID int, query *query.FindListPaginationQuery) ([]*entity.UserTinyWithIsFollow, error)
		FindFavoriteComicUserByUserID(comicID, userID int, query *query.FindListPaginationQuery) ([]*entity.UserTinyWithIsFollow, error)
		FindCfProjectSupporterByCfProjectID(cfProjectID int) (*entity.UserTinyList, error)
		// TODO: ここにあっていいか？
		FindConfirmedUserTypeByPhoneNumberFromCognito(number string) ([]*cognitoidentityprovider.UserType, error)
	}
)
