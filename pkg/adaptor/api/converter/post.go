package converter

import (
	"github.com/stayway-corp/stayway-media-api/pkg/adaptor/api/input"
	"github.com/stayway-corp/stayway-media-api/pkg/adaptor/api/output"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/entity"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/model"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/model/query"
)

/*
 * o -> i
 */
func (c Converters) ConvertFindPostListParamToQuery(param *input.ListPostParam) *query.FindPostListQuery {
	return &query.FindPostListQuery{
		UserID:                 param.UserID,
		AreaID:                 param.AreaID,
		SubAreaID:              param.SubAreaID,
		SubSubAreaID:           param.SubSubAreaID,
		ChildAreaID:            param.ChildAreaID,
		ChildSubAreaID:         param.ChildSubAreaID,
		ChildSubSubAreaID:      param.ChildSubSubAreaID,
		MetasearchAreaID:       param.MetasearchAreaID,
		MetasearchSubAreaID:    param.MetasearchSubAreaID,
		MetasearchSubSubAreaID: param.MetasearchSubSubAreaID,
		InnTypeID:              param.InnTypeID,
		InnDiscerningType:      param.InnDiscerningType,
		ThemeID:                param.ThemeID,
		HashTag:                param.HashTag,
		SortBy:                 param.SortBy,
		Keyword:                param.Keyward,
		CfProjectID:            param.CfProjectID,
		Limit:                  param.GetLimit(),
		OffSet:                 param.GetOffSet(),
	}
}

func (c Converters) ConvertListFeedPostParamToQuery(param *input.ListFeedPostParam) *query.FindListPaginationQuery {
	return &query.FindListPaginationQuery{
		Limit:  param.GetLimit(),
		Offset: param.GetOffSet(),
	}
}

func (c Converters) ConvertPostDetailWithHashtagAndIsFavoriteToOutput(post *entity.PostDetailWithHashtagAndIsFavorite, areaCategories map[int]*entity.AreaCategory, themeCategories map[int]*entity.ThemeCategory, idIsFollowMap map[int]bool) *output.PostShow {
	var hashtags = make([]*output.Hashtag, len(post.Hashtag))
	var bodies = make([]*output.PostBody, len(post.Bodies))

	for i, hashtag := range post.Hashtag {
		hashtags[i] = output.NewHashtag(hashtag.ID, hashtag.Name)
	}
	for i, body := range post.Bodies {
		bodies[i] = output.NewPostBody(body.Page, body.Body)
	}

	areaCategoriesRes := make([]*output.AreaCategoryDetail, len(post.AreaCategories))
	for i, areaCate := range post.AreaCategories {
		areaCategoriesRes[i] = c.ConvertAreaCategoryDetailFromAreaCategory(areaCate, areaCategories)
	}

	themeCategoriesRes := make([]*output.ThemeCategoryDetail, len(post.ThemeCategories))
	for i, themeCate := range post.ThemeCategories {
		themeCategoriesRes[i] = c.ConvertThemeCategoryDetailFromThemeCategory(themeCate, themeCategories)
	}

	return &output.PostShow{
		PostTiny: output.PostTiny{
			ID:             post.ID,
			Thumbnail:      post.PostTiny.Thumbnail,
			Title:          post.Title,
			Beginning:      post.Beginning,
			Slug:           post.Slug,
			TOC:            post.TOC,
			FavoriteCount:  post.FavoriteCount,
			FacebookCount:  post.FacebookCount,
			TwitterCount:   post.TwitterCount,
			Views:          post.Views,
			SEOTitle:       post.SEOTitle,
			SEODescription: post.SEODescription,
			HideAds:        post.HideAds,
			IsSticky:       post.IsSticky,
			CreatedAt:      model.TimeResponse(post.CreatedAt),
			EditedAt:       model.TimeResponse(post.EditedAt),
		},
		IsFavorite:      post.IsFavorite,
		Creator:         c.NewCreatorFromUser(post.User, idIsFollowMap[post.UserID]),
		AreaCategories:  areaCategoriesRes,
		ThemeCategories: themeCategoriesRes,
		Hashtags:        hashtags,
	}
}

func (c Converters) ConvertPostListTinyWithCategoryDetailForListToOutput(posts *entity.PostList, areaCategories map[int]*entity.AreaCategory, themeCategories map[int]*entity.ThemeCategory, idIsFollowMap map[int]bool) *output.PostWithCategoryDetailList {
	postsRes := make([]*output.PostWithCategoryDetail, len(posts.Posts))

	for i, post := range posts.Posts {
		postsRes[i] = c.ConvertPostListTinyWithCategoryDetailToOutput(post, areaCategories, themeCategories, idIsFollowMap)
	}

	return &output.PostWithCategoryDetailList{
		TotalNumber: posts.TotalNumber,
		Posts:       postsRes,
	}
}

func (c Converters) ConvertPostListTinyWithCategoryDetailToOutput(post *entity.PostListTiny, areaCategories map[int]*entity.AreaCategory, themeCategories map[int]*entity.ThemeCategory, idIsFollowMap map[int]bool) *output.PostWithCategoryDetail {
	areaCategoriesRes := make([]*output.AreaCategoryDetail, len(post.AreaCategories))
	for i, areaCate := range post.AreaCategories {
		areaCategoriesRes[i] = c.ConvertAreaCategoryDetailFromAreaCategory(areaCate, areaCategories)
	}

	themeCategoriesRes := make([]*output.ThemeCategoryDetail, len(post.ThemeCategories))
	for i, themeCate := range post.ThemeCategories {
		themeCategoriesRes[i] = c.ConvertThemeCategoryDetailFromThemeCategory(themeCate, themeCategories)
	}

	return &output.PostWithCategoryDetail{
		PostTiny: output.PostTiny{
			ID:             post.ID,
			Slug:           post.Slug,
			Thumbnail:      post.Thumbnail,
			Title:          post.Title,
			Beginning:      post.Beginning,
			TOC:            post.TOC,
			IsSticky:       post.IsSticky,
			HideAds:        post.HideAds,
			FavoriteCount:  post.FavoriteCount,
			FacebookCount:  post.FacebookCount,
			TwitterCount:   post.TwitterCount,
			Views:          post.Views,
			SEOTitle:       post.SEOTitle,
			SEODescription: post.SEODescription,
			CreatedAt:      model.TimeResponse(post.CreatedAt),
			EditedAt:       model.TimeResponse(post.EditedAt),
		},
		IsFavorite:      post.IsFavorite,
		Creator:         c.NewCreatorFromUser(post.User, idIsFollowMap[post.UserID]),
		AreaCategories:  areaCategoriesRes,
		ThemeCategories: themeCategoriesRes,
	}
}
