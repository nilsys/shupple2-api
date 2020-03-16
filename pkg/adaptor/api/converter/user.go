package converter

import (
	"time"

	"github.com/stayway-corp/stayway-media-api/pkg/adaptor/api/param"
	"github.com/stayway-corp/stayway-media-api/pkg/adaptor/api/response"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/entity"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/model"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/model/query"
)

/*
 * o -> i
 */
func ConvertListRankinUserParamToQuery(param *param.ListUserRanking) *query.FindUserRankingListQuery {
	var categoryID int
	fromDate, _ := model.ParseTimeFromFrontStr(param.FromDate)
	toDate, _ := model.ParseTimeFromFrontStr(param.ToDate)
	if param.AreaID != 0 {
		categoryID = param.AreaID
	}
	if param.SubAreaID != 0 {
		categoryID = param.SubAreaID
	}
	if param.SubSubAreaID != 0 {
		categoryID = param.SubSubAreaID
	}

	return &query.FindUserRankingListQuery{
		CategoryID: categoryID,
		SortBy:     param.SortBy,
		FromDate:   fromDate,
		ToDate:     toDate,
		Limit:      param.GetLimit(),
		Offset:     param.GetOffset(),
	}
}

func ConvertListFollowUserParamToQuery(param *param.ListFollowUser) *query.FindFollowUser {
	return &query.FindFollowUser{
		ID:     param.ID,
		Limit:  param.GetLimit(),
		Offset: param.GetOffset(),
	}
}

func ConvertStoreUserParamToEntity(param *param.StoreUser, cognitoID string) *entity.User {
	interests := make([]*entity.UserInterest, len(param.Interests))
	for i, interest := range param.Interests {
		interests[i] = &entity.UserInterest{InterestID: interest}
	}
	return &entity.User{
		UID:        param.UID,
		CognitoID:  cognitoID,
		Name:       param.Name,
		Email:      param.Email,
		Birthdate:  time.Time(param.BirthDate),
		Gender:     param.Gender,
		Profile:    param.Profile,
		AvatarUUID: param.IconUUID,
		Interests:  interests,
	}
}

/*
 * i -> o
 */
func ConvertUserRankingToOutput(users []*entity.QueryRankingUser) []*response.RankinUser {
	userRanking := make([]*response.RankinUser, len(users))

	for i, user := range users {
		userRanking[i] = convertRankingUserToOutput(user)
	}

	return userRanking
}

func ConvertUsersToFollowUsers(users []*entity.User) []*response.FollowUser {
	followUsers := make([]*response.FollowUser, len(users))
	for i, user := range users {
		followUsers[i] = convertUserToFollowUser(user)
	}
	return followUsers
}

func convertUserToFollowUser(user *entity.User) *response.FollowUser {
	return &response.FollowUser{
		ID:        user.ID,
		Name:      user.Name,
		Thumbnail: user.GenerateThumbnailURL(),
	}
}

// QueryRankingUserをランキング一覧で返す型にコンバート
func convertRankingUserToOutput(user *entity.QueryRankingUser) *response.RankinUser {
	interests := make([]string, len(user.Interests))
	for i, interest := range user.Interests {
		interests[i] = interest.Name
	}

	return &response.RankinUser{
		ID:        user.ID,
		Name:      user.Name,
		Profile:   user.Profile,
		Thumbnail: user.GenerateThumbnailURL(),
		Interest:  interests,
	}
}
