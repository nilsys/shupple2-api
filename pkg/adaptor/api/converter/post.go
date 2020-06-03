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

/*
 * i -> o
 */
func (c Converters) ConvertPostDetailListToOutput(posts []*entity.PostDetail) []*output.Post {
	responsePosts := make([]*output.Post, len(posts))

	for i, queryPost := range posts {
		responsePosts[i] = c.ConvertQueryPostToOutput(queryPost)
	}

	return responsePosts
}

func (c Converters) ConvertQueryPostToOutput(queryPost *entity.PostDetail) *output.Post {
	return &output.Post{
		ID:              queryPost.ID,
		Thumbnail:       queryPost.Thumbnail,
		Title:           queryPost.Title,
		Slug:            queryPost.Slug,
		AreaCategories:  c.ConvertAreaCategoriesToOutput(queryPost.AreaCategories),
		ThemeCategories: c.ConvertThemeCategoriesToOutput(queryPost.ThemeCategories),
		Creator:         c.NewCreatorFromUser(queryPost.User),
		FavoriteCount:   queryPost.FavoriteCount,
		Views:           queryPost.Views,
		HideAds:         queryPost.HideAds,
		CreatedAt:       model.TimeResponse(queryPost.CreatedAt),
		UpdatedAt:       model.TimeResponse(queryPost.UpdatedAt),
	}
}

func (c Converters) ConvertPostDetailWithHashtagAndIsFavoriteToOutput(post *entity.PostDetailWithHashtagAndIsFavorite, areaCategories map[int]*entity.AreaCategory, themeCategories map[int]*entity.ThemeCategory) *output.PostShow {
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
		ID:              post.ID,
		Thumbnail:       post.PostTiny.Thumbnail,
		Title:           post.Title,
		Slug:            post.Slug,
		Body:            bodies,
		TOC:             post.TOC,
		FavoriteCount:   post.FavoriteCount,
		FacebookCount:   post.FacebookCount,
		TwitterCount:    post.TwitterCount,
		Views:           post.Views,
		SEOTitle:        post.SEOTitle,
		SEODescription:  post.SEODescription,
		HideAds:         post.HideAds,
		IsFavorite:      post.IsFavorite,
		Creator:         c.NewCreatorFromUser(post.User),
		AreaCategories:  areaCategoriesRes,
		ThemeCategories: themeCategoriesRes,
		Hashtags:        hashtags,
		CreatedAt:       model.TimeResponse(post.CreatedAt),
		UpdatedAt:       model.TimeResponse(post.UpdatedAt),
	}
}

func (c Converters) ConvertPostListToOutput(list *entity.PostList) *output.PostList {
	posts := c.ConvertPostListTiniesToOutput(list.Posts)

	return &output.PostList{
		TotalNumber: list.TotalNumber,
		Posts:       posts,
	}
}

func (c Converters) ConvertPostListTiniesToOutput(list []*entity.PostListTiny) []*output.Post {
	res := make([]*output.Post, len(list))

	for i, tiny := range list {
		res[i] = c.ConvertPostListTinyToOutput(tiny)
	}

	return res
}

func (c Converters) ConvertPostListTinyToOutput(post *entity.PostListTiny) *output.Post {
	return &output.Post{
		ID:              post.ID,
		Thumbnail:       post.Thumbnail,
		Title:           post.Title,
		Slug:            post.Slug,
		AreaCategories:  c.ConvertAreaCategoriesToOutput(post.AreaCategories),
		ThemeCategories: c.ConvertThemeCategoriesToOutput(post.ThemeCategories),
		Creator:         c.NewCreatorFromUser(post.User),
		FavoriteCount:   post.FavoriteCount,
		Views:           post.Views,
		HideAds:         post.HideAds,
		IsFavorite:      post.IsFavorite,
		CreatedAt:       model.TimeResponse(post.CreatedAt),
		UpdatedAt:       model.TimeResponse(post.UpdatedAt),
	}
}

func (c Converters) ConvertPostListTinyWithCategoryDetailForListToOutput(posts *entity.PostList, areaCategories map[int]*entity.AreaCategory, themeCategories map[int]*entity.ThemeCategory) *output.PostWithCategoryDetailList {
	postsRes := make([]*output.PostWithCategoryDetail, len(posts.Posts))

	for i, post := range posts.Posts {
		postsRes[i] = c.ConvertPostListTinyWithCategoryDetailToOutput(post, areaCategories, themeCategories)
	}

	return &output.PostWithCategoryDetailList{
		TotalNumber: posts.TotalNumber,
		Posts:       postsRes,
	}
}

func (c Converters) ConvertPostListTinyWithCategoryDetailToOutput(post *entity.PostListTiny, areaCategories map[int]*entity.AreaCategory, themeCategories map[int]*entity.ThemeCategory) *output.PostWithCategoryDetail {
	areaCategoriesRes := make([]*output.AreaCategoryDetail, len(post.AreaCategories))
	for i, areaCate := range post.AreaCategories {
		areaCategoriesRes[i] = c.ConvertAreaCategoryDetailFromAreaCategory(areaCate, areaCategories)
	}

	themeCategoriesRes := make([]*output.ThemeCategoryDetail, len(post.ThemeCategories))
	for i, themeCate := range post.ThemeCategories {
		themeCategoriesRes[i] = c.ConvertThemeCategoryDetailFromThemeCategory(themeCate, themeCategories)
	}

	return &output.PostWithCategoryDetail{
		ID:              post.ID,
		Thumbnail:       post.Thumbnail,
		AreaCategories:  areaCategoriesRes,
		ThemeCategories: themeCategoriesRes,
		Title:           post.Title,
		Slug:            post.Slug,
		Creator:         c.NewCreatorFromUser(post.User),
		FavoriteCount:   post.FavoriteCount,
		Views:           post.Views,
		HideAds:         post.HideAds,
		IsFavorite:      post.IsFavorite,
		CreatedAt:       model.TimeResponse(post.CreatedAt),
		UpdatedAt:       model.TimeResponse(post.UpdatedAt),
	}
}
