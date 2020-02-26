package converter

import (
	"github.com/stayway-corp/stayway-media-api/pkg/adaptor/api/param"
	"github.com/stayway-corp/stayway-media-api/pkg/adaptor/api/response"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/entity"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/model"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/model/query"
)

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

func ConvertUserRankingToOutput(users []*entity.QueryRankingUser) []*response.RankinUser {
	userRanking := make([]*response.RankinUser, len(users))

	for i, user := range users {
		userRanking[i] = convertRankingUserToOutput(user)
	}

	return userRanking
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
