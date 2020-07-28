package converter

import (
	"time"

	"github.com/stayway-corp/stayway-media-api/pkg/domain/model/command"

	"github.com/stayway-corp/stayway-media-api/pkg/adaptor/api/input"
	"github.com/stayway-corp/stayway-media-api/pkg/adaptor/api/output"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/entity"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/model"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/model/query"
)

/*
 * o -> i
 */
func (c Converters) ConvertListRankinUserParamToQuery(param *input.ListUserRanking) *query.FindUserRankingListQuery {
	return &query.FindUserRankingListQuery{
		AreaID:       param.AreaID,
		SubAreaID:    param.SubAreaID,
		SubSubAreaID: param.SubSubAreaID,
		SortBy:       param.SortBy,
		FromDate:     time.Time(param.FromDate),
		ToDate:       time.Time(param.ToDate),
		Limit:        param.GetLimit(),
		Offset:       param.GetOffset(),
	}
}

func (c Converters) ConvertListFollowUserParamToQuery(param *input.ListFollowUser) *query.FindFollowUser {
	return &query.FindFollowUser{
		ID:     param.ID,
		Limit:  param.GetLimit(),
		Offset: param.GetOffset(),
	}
}

func (c Converters) ConvertStoreUserParamToEntity(param *input.StoreUser) *entity.User {
	interests := make([]*entity.UserInterest, len(param.Interests))
	for i, interest := range param.Interests {
		interests[i] = &entity.UserInterest{InterestID: interest}
	}
	return &entity.User{
		CognitoUserName: param.CognitoUserName,
		UID:             param.UID,
		Name:            param.Name,
		Email:           param.Email,
		Birthdate:       time.Time(param.BirthDate),
		Gender:          param.Gender,
		Profile:         param.Profile,
		URL:             param.URL,
		FacebookURL:     param.FacebookURL,
		InstagramURL:    param.InstagramURL,
		TwitterURL:      param.TwitterURL,
		YoutubeURL:      param.YoutubeURL,
		LivingArea:      param.LivingArea,
		UserInterests:   interests,
	}
}

func (c Converters) ConvertUpdateUserParamToCmd(i *input.UpdateUser) *command.UpdateUser {
	interests := make([]*entity.UserInterest, len(i.Interests))
	for i, interest := range i.Interests {
		interests[i] = &entity.UserInterest{InterestID: interest}
	}
	return &command.UpdateUser{
		Name:         i.Name,
		Email:        i.Email,
		BirthDate:    i.BirthDate,
		Gender:       i.Gender,
		Profile:      i.Profile,
		IconUUID:     i.IconUUID,
		HeaderUUID:   i.HeaderUUID,
		URL:          i.URL,
		FacebookURL:  i.FacebookURL,
		InstagramURL: i.InstagramURL,
		TwitterURL:   i.TwitterURL,
		YoutubeURL:   i.YoutubeURL,
		LivingArea:   i.LivingArea,
		Interests:    interests,
	}
}

func (c Converters) ConvertListFavoriteMediaUserToQuery(param *input.ListFavoriteMediaUser) *query.FindListPaginationQuery {
	return &query.FindListPaginationQuery{
		Limit:  param.GetLimit(),
		Offset: param.GetOffset(),
	}
}

/*
 * i -> o
 */

func (c Converters) ConvertUserTableListToOutput(users []*entity.UserTiny) []*output.UserSummary {
	res := make([]*output.UserSummary, len(users))

	for i, user := range users {
		res[i] = c.convertUserTableToOutput(user)
	}

	return res
}
func (c Converters) ConvertUserRankingToOutput(users []*entity.UserDetail) []*output.RankinUser {
	userRanking := make([]*output.RankinUser, len(users))

	for i, user := range users {
		userRanking[i] = c.ConvertUserDetailToOutput(user)
	}

	return userRanking
}

func (c Converters) ConvertUserTinyWithIsFavoriteListToUserSummaryList(users []*entity.UserTinyWithIsFollow) []*output.UserSummaryWithIsFollow {
	followUsers := make([]*output.UserSummaryWithIsFollow, len(users))
	for i, user := range users {
		followUsers[i] = c.UserTinyWithIsFollowToOutput(user)
	}
	return followUsers
}

func (c Converters) convertUserTableToOutput(user *entity.UserTiny) *output.UserSummary {
	return &output.UserSummary{
		ID:      user.ID,
		UID:     user.UID,
		Name:    user.Name,
		IconURL: user.AvatarURL(c.filesURL()),
	}
}

// UserDetailをランキング一覧で返す型にコンバート
func (c Converters) ConvertUserDetailToOutput(user *entity.UserDetail) *output.RankinUser {
	interests := make([]output.Interest, len(user.Interests))
	for i, interest := range user.Interests {
		interests[i] = c.InterestToOutput(interest)
	}

	return &output.RankinUser{
		ID:        user.ID,
		UID:       user.UID,
		Name:      user.Name,
		Profile:   user.Profile,
		Thumbnail: user.AvatarURL(c.filesURL()),
		Interests: interests,
	}
}

func (c Converters) ConvertUserDetailWithCountToOutPut(user *entity.UserDetailWithMediaCount) *output.MyPageUser {
	interests := make([]output.Interest, len(user.Interests))
	for i, interest := range user.Interests {
		interests[i] = c.InterestToOutput(interest)
	}

	return &output.MyPageUser{
		ID:             user.ID,
		UID:            user.UID,
		Name:           user.Name,
		Profile:        user.Profile,
		Birthdate:      model.Date(user.Birthdate),
		Email:          user.Email,
		Gender:         user.Gender,
		Icon:           user.AvatarURL(c.filesURL()),
		Header:         user.HeaderURL(c.filesURL()),
		FacebookURL:    user.FacebookURL,
		InstagramURL:   user.InstagramURL,
		TwitterURL:     user.TwitterURL,
		YoutubeURL:     user.YoutubeURL,
		URL:            user.URL,
		LivingArea:     user.LivingArea,
		PostCount:      user.PostCount + user.ReviewCount + user.VlogCount,
		FollowingCount: user.FollowingCount,
		FollowedCount:  user.FollowedCount,
		Interests:      interests,
		IsFollow:       user.IsFollow,
	}
}

func (c Converters) InterestToOutput(interest *entity.Interest) output.Interest {
	return output.Interest{
		ID:            interest.ID,
		Name:          interest.Name,
		InterestGroup: interest.InterestGroup,
	}
}

func (c Converters) NewCreatorFromUser(user *entity.User) output.Creator {
	return output.NewCreator(
		user.ID, user.UID, user.AvatarURL(c.filesURL()), user.Name, user.Profile,
		user.FacebookURL, user.InstagramURL, user.TwitterURL, user.YoutubeURL, user.URL,
	)
}

func (c Converters) NewUserSummaryFromUser(user *entity.User) *output.UserSummary {
	return output.NewUserSummary(user.ID, user.UID, user.Name, user.AvatarURL(c.filesURL()))
}

func (c Converters) UserTinyWithIsFollowToOutput(user *entity.UserTinyWithIsFollow) *output.UserSummaryWithIsFollow {
	return &output.UserSummaryWithIsFollow{
		UserSummary: *output.NewUserSummary(user.ID, user.UID, user.Name, user.AvatarURL(c.filesURL())),
		IsFollow:    user.IsFollow,
	}
}
