package converter

import (
	"github.com/stayway-corp/stayway-media-api/pkg/adaptor/api/param"
	"github.com/stayway-corp/stayway-media-api/pkg/adaptor/api/response"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/entity"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/model"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/model/query"
)

/*
 * o -> i
 */
func ConvertFindPostListParamToQuery(param *param.ListPostParam) *query.FindPostListQuery {
	return &query.FindPostListQuery{
		UserID:                 param.UserID,
		AreaID:                 param.AreaID,
		SubAreaID:              param.SubAreaID,
		SubSubAreaID:           param.SubSubAreaID,
		MetasearchAreaID:       param.MetasearchAreaID,
		MetasearchSubAreaID:    param.MetasearchSubAreaID,
		MetasearchSubSubAreaID: param.MetasearchSubSubAreaID,
		InnTypeID:              param.InnTypeID,
		InnDiscerningType:      param.InnDiscerningType,
		ThemeID:                param.ThemeID,
		HashTag:                param.HashTag,
		SortBy:                 param.SortBy,
		Keyward:                param.Keyward,
		Limit:                  param.GetLimit(),
		OffSet:                 param.GetOffSet(),
	}
}

// ConvertPostToOutput()のスライスバージョン
func ConvertPostToOutput(queryPostList []*entity.QueryPost) []*response.Post {
	responsePosts := make([]*response.Post, len(queryPostList))

	for i, queryPost := range queryPostList {
		responsePosts[i] = ConvertQueryPostToOutput(queryPost)
	}

	return responsePosts
}

func ConvertListFeedPostParamToQuery(param *param.ListFeedPostParam) *query.FindListPaginationQuery {
	return &query.FindListPaginationQuery{
		Limit:  param.GetLimit(),
		Offset: param.GetOffSet(),
	}
}

// outputの構造体へconvert
func ConvertQueryPostToOutput(queryPost *entity.QueryPost) *response.Post {
	var areaCategories []response.Category
	var themeCategories []response.Category

	for _, category := range queryPost.Categories {
		if category.Type.IsAreaKind() {
			areaCategories = append(areaCategories, response.NewCategory(category.ID, category.Name))
		} else {
			themeCategories = append(themeCategories, response.NewCategory(category.ID, category.Name))
		}
	}

	return &response.Post{
		ID:              queryPost.ID,
		Thumbnail:       queryPost.GenerateThumbnailURL(),
		AreaCategories:  areaCategories,
		ThemeCategories: themeCategories,
		Title:           queryPost.Title,
		Creator: response.Creator{
			Thumbnail: queryPost.User.GenerateThumbnailURL(),
			Name:      queryPost.User.Name,
		},
		LikeCount: queryPost.FavoriteCount,
		UpdatedAt: model.TimeFmtToFrontStr(queryPost.UpdatedAt),
	}
}
