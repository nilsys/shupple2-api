package converter

import (
	"github.com/stayway-corp/stayway-media-api/pkg/adaptor/api/param"
	"github.com/stayway-corp/stayway-media-api/pkg/adaptor/api/response"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/entity"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/model"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/model/query"
)

// i/oの構造体からレポジトリで使用するクエリ構造体へconvert
func ConvertFindPostListParamToQuery(param *param.ListPostParam) *query.FindPostListQuery {
	sortBy, _ := model.ParseSortBy(param.SortBy)
	return &query.FindPostListQuery{
		AreaID:       param.AreaID,
		SubAreaID:    param.SubAreaID,
		SubSubAreaID: param.SubSubAreaID,
		ThemeID:      param.ThemeID,
		HashTag:      param.HashTag,
		SortBy:       sortBy,
		Limit:        param.GetLimit(),
		OffSet:       param.GetOffSet(),
	}
}

// outputの構造体へconvert
func convertPostToOutput(queryPost *entity.QueryPost) *response.Post {
	var areaCategories []response.Category
	var themeCategories []response.Category

	for _, category := range queryPost.Categories {
		if category.Type == model.CategoryTypeArea || category.Type == model.CategoryTypeSubArea || category.Type == model.CategoryTypeSubSubArea {
			areaCategories = append(areaCategories, response.NewCategory(category.ID, category.Name))
		}
		if category.Type == model.CategoryTypeTheme {
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

// ConvertPostToOutput()のスライスバージョン
func ConvertPostToOutput(queryPostList []*entity.QueryPost) []*response.Post {
	// MEMO: 代入しないと0件の時にフロントにnullが返る
	responsePosts := []*response.Post{}

	for _, queryPost := range queryPostList {
		responsePosts = append(responsePosts, convertPostToOutput(queryPost))
	}

	return responsePosts
}
