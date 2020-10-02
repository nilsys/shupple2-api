package service

import (
	"github.com/google/wire"
	"github.com/pkg/errors"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/entity"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/model/query"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/repository"
)

type (
	UserQueryService interface {
		ShowByUID(uid string, ouser entity.OptionalUser) (*entity.UserDetailWithMediaCount, error)
		ShowByID(id int) (*entity.UserDetailWithMediaCount, error)
		ShowByMigrationCode(code string) (*entity.UserDetailWithMediaCount, error)
		ShowUserRanking(query *query.FindUserRankingListQuery) ([]*entity.UserDetail, error)
		ListRecommendFollowUser(interestIDs []int) ([]*entity.UserTiny, error)
		ListFollowing(query *query.FindFollowUser, ouser *entity.OptionalUser) ([]*entity.UserTinyWithIsFollow, error)
		ListFollowed(query *query.FindFollowUser, ouser *entity.OptionalUser) ([]*entity.UserTinyWithIsFollow, error)
		ListFavoritePostUser(postID int, user *entity.OptionalUser, query *query.FindListPaginationQuery) ([]*entity.UserTinyWithIsFollow, error)
		ListFavoriteReviewUser(reviewID int, user *entity.OptionalUser, query *query.FindListPaginationQuery) ([]*entity.UserTinyWithIsFollow, error)
		ListFavoriteComicUser(comicID int, ouser *entity.OptionalUser, query *query.FindListPaginationQuery) ([]*entity.UserTinyWithIsFollow, error)
		ListFavoriteVlogUser(vlogID int, ouser *entity.OptionalUser, query *query.FindListPaginationQuery) ([]*entity.UserTinyWithIsFollow, error)
		IsExistByPhoneNumber(number string) (bool, error)
		RelationFlgMaps(baseUserID int, targetUserIDs []int) (*entity.UserRelationFlgMap, error)
	}

	UserQueryServiceImpl struct {
		repository.UserQueryRepository
	}
)

var UserQueryServiceSet = wire.NewSet(
	wire.Struct(new(UserQueryServiceImpl), "*"),
	wire.Bind(new(UserQueryService), new(*UserQueryServiceImpl)),
)

func (s *UserQueryServiceImpl) ShowByUID(uid string, ouser entity.OptionalUser) (*entity.UserDetailWithMediaCount, error) {
	showTargetUser, err := s.UserQueryRepository.FindByUID(uid)
	if err != nil {
		return nil, errors.Wrap(err, "failed find user by uid")
	}

	if ouser.Authenticated {
		user, err := s.UserQueryRepository.FindUserDetailWithCountByID(showTargetUser.ID)
		if err != nil {
			return nil, errors.Wrap(err, "failed find user by id")
		}
		idRelationFlgMap, err := s.RelationFlgMaps(ouser.ID, []int{showTargetUser.ID})
		if err != nil {
			return nil, errors.Wrap(err, "failed find user_following")
		}

		user.IsFollow = idRelationFlgMap.IsFollowByUserID(showTargetUser.ID)
		user.IsBlocking = idRelationFlgMap.IsBlockingByUserID(showTargetUser.ID)

		return user, nil
	}

	return s.UserQueryRepository.FindUserDetailWithCountByID(showTargetUser.ID)
}

func (s *UserQueryServiceImpl) ShowByID(id int) (*entity.UserDetailWithMediaCount, error) {
	return s.UserQueryRepository.FindUserDetailWithCountByID(id)
}

func (s *UserQueryServiceImpl) ShowByMigrationCode(code string) (*entity.UserDetailWithMediaCount, error) {
	// 非ログインユーザーの場合論理削除されている為、論理削除されているUserも対象に含める
	userTable, err := s.UserQueryRepository.UnscopedFindByMigrationCode(code)
	if err != nil {
		return nil, errors.Wrap(err, "failed find user by migration_code")
	}
	// 非ログインユーザーの場合論理削除されている為、論理削除されているUserも対象に含める
	return s.UserQueryRepository.UnscopedFindUserDetailWithCountByID(userTable.ID)
}

func (s *UserQueryServiceImpl) ShowUserRanking(query *query.FindUserRankingListQuery) ([]*entity.UserDetail, error) {
	return s.UserQueryRepository.FindUserRankingListByParams(query)
}

func (s *UserQueryServiceImpl) ListRecommendFollowUser(interestIDs []int) ([]*entity.UserTiny, error) {
	return s.UserQueryRepository.FindRecommendFollowUserList(interestIDs)
}

func (s *UserQueryServiceImpl) ListFollowing(query *query.FindFollowUser, ouser *entity.OptionalUser) ([]*entity.UserTinyWithIsFollow, error) {
	if ouser.IsAuthorized() {
		return s.UserQueryRepository.FindFollowingWithIsFollowByID(ouser.ID, query)
	}
	return s.UserQueryRepository.FindFollowingByID(query)
}

func (s *UserQueryServiceImpl) ListFollowed(query *query.FindFollowUser, ouser *entity.OptionalUser) ([]*entity.UserTinyWithIsFollow, error) {
	if ouser.IsAuthorized() {
		return s.UserQueryRepository.FindFollowedWithIsFollowByID(ouser.ID, query)
	}
	return s.UserQueryRepository.FindFollowedByID(query)
}

func (s *UserQueryServiceImpl) ListFavoritePostUser(postID int, user *entity.OptionalUser, query *query.FindListPaginationQuery) ([]*entity.UserTinyWithIsFollow, error) {
	if user.IsAuthorized() {
		return s.UserQueryRepository.FindFavoritePostUserByUserID(postID, user.ID, query)
	}

	return s.UserQueryRepository.FindFavoritePostUser(postID, query)
}

func (s *UserQueryServiceImpl) ListFavoriteReviewUser(reviewID int, user *entity.OptionalUser, query *query.FindListPaginationQuery) ([]*entity.UserTinyWithIsFollow, error) {
	if user.IsAuthorized() {
		return s.UserQueryRepository.FindFavoriteReviewUserByUserID(reviewID, user.ID, query)
	}

	return s.UserQueryRepository.FindFavoriteReviewUser(reviewID, query)
}

func (s *UserQueryServiceImpl) ListFavoriteComicUser(comicID int, ouser *entity.OptionalUser, query *query.FindListPaginationQuery) ([]*entity.UserTinyWithIsFollow, error) {
	if ouser.IsAuthorized() {
		return s.UserQueryRepository.FindFavoriteComicUserByUserID(comicID, ouser.ID, query)
	}
	return s.UserQueryRepository.FindFavoriteComicUser(comicID, query)
}

func (s *UserQueryServiceImpl) ListFavoriteVlogUser(vlogID int, ouser *entity.OptionalUser, query *query.FindListPaginationQuery) ([]*entity.UserTinyWithIsFollow, error) {
	if ouser.IsAuthorized() {
		return s.UserQueryRepository.FindFavoriteVlogUserByUserID(vlogID, ouser.ID, query)
	}
	return s.UserQueryRepository.FindFavoriteVlogUser(vlogID, query)
}

func (s *UserQueryServiceImpl) IsExistByPhoneNumber(number string) (bool, error) {
	cognitoUsers, err := s.UserQueryRepository.FindConfirmedUserTypeByPhoneNumberFromCognito(number)
	if err != nil {
		return false, errors.Wrap(err, "failed find from cognito")
	}
	if len(cognitoUsers) == 0 {
		return false, nil
	}

	cognitoUserNames := make([]string, len(cognitoUsers))
	for i, user := range cognitoUsers {
		cognitoUserNames[i] = *user.Username
	}

	// stayway側に登録されているか
	user, err := s.UserQueryRepository.FindByCognitoUserName(cognitoUserNames)
	if err != nil {
		return false, errors.Wrap(err, "failed find by cognito_user_name")
	}

	return len(user) > 0, nil
}

/*
MEMO: IsFollow, IsBlockingのmapをを持つentityを返す
      他にIsXXXのフラグが必要になった場合はここに足す事
*/
func (s *UserQueryServiceImpl) RelationFlgMaps(baseUserID int, targetUserIDs []int) (*entity.UserRelationFlgMap, error) {
	idIsFollowMap, err := s.UserQueryRepository.IsFollowing(baseUserID, targetUserIDs)
	if err != nil {
		return nil, errors.Wrap(err, "failed find user_following")
	}
	idIsBlockingMap, err := s.UserQueryRepository.IsBlocking(baseUserID, targetUserIDs)
	if err != nil {
		return nil, errors.Wrap(err, "failed find user_blocking")
	}

	return &entity.UserRelationFlgMap{
		IDIsFollowMap:   idIsFollowMap,
		IDIsBlockingMap: idIsBlockingMap,
	}, nil
}
