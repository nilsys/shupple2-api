package entity

import (
	"path"
	"time"

	"github.com/huandu/facebook/v2"

	"github.com/stayway-corp/stayway-media-api/pkg/config"

	facebookEntity "github.com/stayway-corp/stayway-media-api/pkg/domain/entity/facebook"
	"github.com/stayway-corp/stayway-media-api/pkg/util"
	"gopkg.in/guregu/null.v3"
)

type (
	PostTiny struct {
		ID             int `gorm:"primary_key"`
		UserID         int
		Slug           string
		Thumbnail      string
		Title          string
		Beginning      string
		TOC            string
		CfProjectID    null.Int
		IsSticky       bool
		HideAds        bool
		FavoriteCount  int
		FacebookCount  int
		TwitterCount   int
		Views          int
		SEOTitle       string
		SEODescription string
		EditedAt       time.Time
		Times
	}

	Post struct {
		PostTiny
		Bodies           []*PostBody `gorm:"foreignkey:PostID"`
		AreaCategoryIDs  []*PostAreaCategory
		ThemeCategoryIDs []*PostThemeCategory
		HashtagIDs       []*PostHashtag
	}

	PostBody struct {
		PostID int `gorm:"primary_key"`
		Page   int `gorm:"primary_key"`
		Body   string
	}

	PostAreaCategory struct {
		PostID         int `gorm:"primary_key"`
		AreaCategoryID int `gorm:"primary_key"`
	}

	PostThemeCategory struct {
		PostID          int `gorm:"primary_key"`
		ThemeCategoryID int `gorm:"primary_key"`
	}

	PostHashtag struct {
		PostID    int `gorm:"primary_key"`
		HashtagID int `gorm:"primary_key"`
	}

	PostDetail struct {
		PostTiny
		Bodies          []*PostBody      `gorm:"foreignkey:PostID"`
		User            *User            `gorm:"foreignkey:ID;association_foreignkey:UserID"`
		AreaCategories  []*AreaCategory  `gorm:"many2many:post_area_category;jointable_foreignkey:post_id;"`
		ThemeCategories []*ThemeCategory `gorm:"many2many:post_theme_category;jointable_foreignkey:post_id;"`
	}

	UserFavoritePost struct {
		UserID int
		PostID int
	}

	PostDetailList struct {
		List []*PostDetail
	}

	PostList struct {
		TotalNumber int
		Posts       []*PostListTiny
	}

	// 一覧用Post
	PostListTiny struct {
		PostTiny
		User            *User `gorm:"foreignkey:ID;association_foreignkey:UserID"`
		IsFavorite      bool
		AreaCategories  []*AreaCategory  `gorm:"many2many:post_area_category;jointable_foreignkey:post_id;"`
		ThemeCategories []*ThemeCategory `gorm:"many2many:post_theme_category;jointable_foreignkey:post_id;"`
	}

	PostDetailWithHashtagAndIsFavorite struct {
		PostTiny
		Bodies          []*PostBody `gorm:"foreignkey:PostID"`
		User            *User       `gorm:"foreignkey:ID;association_foreignkey:UserID"`
		IsFavorite      bool
		AreaCategories  []*AreaCategory  `gorm:"many2many:post_area_category;jointable_foreignkey:post_id;"`
		ThemeCategories []*ThemeCategory `gorm:"many2many:post_theme_category;jointable_foreignkey:post_id;"`
		Hashtag         []*Hashtag       `gorm:"many2many:post_hashtag;jointable_foreignkey:post_id;"`
		Features        []*Feature       `gorm:"many2many:feature_post;jointable_foreignkey:post_id"`
	}

	PostTinyList []*PostTiny
)

func (p *PostDetail) TableName() string {
	return "post"
}

func (p *PostListTiny) TableName() string {
	return "post"
}

func NewPost(tiny PostTiny, bodies []string, areaCategoryIDs, themeCategoryIDs, hashtagIDs []int) Post {
	post := Post{PostTiny: tiny}

	post.SetBodies(bodies)
	post.SetAreaCategories(areaCategoryIDs)
	post.SetThemeCategories(themeCategoryIDs)
	post.SetHashtags(hashtagIDs)

	return post
}

func (p *Post) SetBodies(bodies []string) {
	p.Bodies = make([]*PostBody, len(bodies))
	for i, body := range bodies {
		p.Bodies[i] = &PostBody{
			PostID: p.ID,
			Page:   i + 1,
			Body:   body,
		}
	}
}

func (p *Post) SetAreaCategories(areaCategoryIDs []int) {
	p.AreaCategoryIDs = make([]*PostAreaCategory, len(areaCategoryIDs))
	for i, c := range areaCategoryIDs {
		p.AreaCategoryIDs[i] = &PostAreaCategory{
			PostID:         p.ID,
			AreaCategoryID: c,
		}
	}
}

func (p *Post) SetThemeCategories(themeCategoryIDs []int) {
	p.ThemeCategoryIDs = make([]*PostThemeCategory, len(themeCategoryIDs))
	for i, c := range themeCategoryIDs {
		p.ThemeCategoryIDs[i] = &PostThemeCategory{
			PostID:          p.ID,
			ThemeCategoryID: c,
		}
	}
}

func (p *Post) SetHashtags(hashtagIDs []int) {
	p.HashtagIDs = make([]*PostHashtag, len(hashtagIDs))
	for i, h := range hashtagIDs {
		p.HashtagIDs[i] = &PostHashtag{
			PostID:    p.ID,
			HashtagID: h,
		}
	}
}

func NewUserFavoritePost(userID, postID int) *UserFavoritePost {
	return &UserFavoritePost{
		UserID: userID,
		PostID: postID,
	}
}

func (p *PostDetailWithHashtagAndIsFavorite) TableName() string {
	return "post"
}

func (p *PostList) AreaCategoryIDs() []int {
	ids := make([]int, 0)

	for _, post := range p.Posts {
		for _, area := range post.AreaCategories {
			ids = append(ids, area.AreaID)

			if area.SubAreaID.Valid {
				ids = append(ids, int(area.SubAreaID.Int64))
			}

			if area.SubSubAreaID.Valid {
				ids = append(ids, int(area.SubSubAreaID.Int64))
			}
		}
	}

	return util.RemoveDuplicatesAndZeroFromIntSlice(ids)
}

func (p *PostList) ThemeCategoryIDs() []int {
	ids := make([]int, 0)

	for _, post := range p.Posts {
		for _, theme := range post.ThemeCategories {
			ids = append(ids, theme.ThemeID)

			if theme.SubThemeID.Valid {
				ids = append(ids, int(theme.SubThemeID.Int64))
			}
		}
	}

	return util.RemoveDuplicatesAndZeroFromIntSlice(ids)
}

func (p *PostDetailWithHashtagAndIsFavorite) AreaCategoryIDs() []int {
	ids := make([]int, 0)

	for _, area := range p.AreaCategories {
		ids = append(ids, area.AreaID)

		if area.SubAreaID.Valid {
			ids = append(ids, int(area.SubAreaID.Int64))
		}

		if area.SubSubAreaID.Valid {
			ids = append(ids, int(area.SubSubAreaID.Int64))
		}
	}

	return util.RemoveDuplicatesAndZeroFromIntSlice(ids)
}

func (p *PostDetailWithHashtagAndIsFavorite) ThemeCategoryIDs() []int {
	ids := make([]int, 0)

	for _, theme := range p.ThemeCategories {
		ids = append(ids, theme.ThemeID)

		if theme.SubThemeID.Valid {
			ids = append(ids, int(theme.SubThemeID.Int64))
		}
	}

	return util.RemoveDuplicatesAndZeroFromIntSlice(ids)
}

func (p *PostList) UserIDs() []int {
	ids := make([]int, len(p.Posts))
	for i, tiny := range p.Posts {
		ids[i] = tiny.UserID
	}
	return ids
}

func (p *PostDetailList) ToCfProjectIDMap() map[int]*PostDetail {
	resolve := make(map[int]*PostDetail, len(p.List))
	for _, tiny := range p.List {
		resolve[int(tiny.CfProjectID.Int64)] = tiny
	}
	return resolve
}

func (p *PostTiny) MediaWebURL(baseURL config.URL) *config.URL {
	baseURL.Path = path.Join(baseURL.Path, p.Slug)
	return &baseURL
}

/*
import_facebook_share_countのバッチで使用する想定
Graph APIのバッチリクエストのbatchクエリへ整形する
トレイリングスラッシュを区別する為,1つのPostにつき以下の形で2つのリクエストを発行する
上限がバッチリクエストに含める事が出来る上限は50なので
事実上len(PostTinyList.List)==25の時のみ使用出来る
 https://developers.facebook.com/docs/graph-api/making-multiple-requests

[
    {
      "method":"GET",
      "relative_url":"?id=https://stayway.jp/tourism/asia-heritage13/&fields=engagement"
    },
    {
      "method":"GET",
      "relative_url":"?id=https://stayway.jp/tourism/asia-heritage13&fields=engagement"
    },
]
*/
func (p PostTinyList) ToGraphAPIBatchRequestQueryStr(baseURL config.URL) []facebook.Params {
	resolve := make([]facebook.Params, 0, len(p)*2)

	for _, post := range p {
		resolve = append(resolve, facebookEntity.GetRelativeURLParams(post.MediaWebURL(baseURL)), facebookEntity.GetRelativeTrailingSlashURLParams(post.MediaWebURL(baseURL)))
	}

	return resolve
}

func (p *PostTiny) TableName() string {
	return "post"
}
