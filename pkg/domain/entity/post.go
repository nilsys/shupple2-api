package entity

import (
	"sort"
	"time"
)

type (
	PostTiny struct {
		ID             int `gorm:"primary_key"`
		UserID         int
		Slug           string
		Thumbnail      string
		Title          string
		TOC            string
		IsSticky       bool
		HideAds        bool
		FavoriteCount  int
		FacebookCount  int
		TwitterCount   int
		Views          int
		SEOTitle       string
		SEODescription string
		CreatedAt      time.Time `gorm:"default:current_timestamp"`
		UpdatedAt      time.Time `gorm:"default:current_timestamp"`
		DeletedAt      *time.Time
	}

	Post struct {
		PostTiny
		Bodies           []*PostBody
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

	// 参照用Post
	PostDetail struct {
		PostTiny
		Bodies          []*PostBody      `gorm:"foreignkey:PostID"`
		User            *User            `gorm:"foreignkey:UserID"`
		AreaCategories  []*AreaCategory  `gorm:"many2many:post_area_category;jointable_foreignkey:post_id;"`
		ThemeCategories []*ThemeCategory `gorm:"many2many:post_theme_category;jointable_foreignkey:post_id;"`
	}

	UserFavoritePost struct {
		UserID int
		PostID int
	}

	PostDetailList struct {
		TotalNumber int
		Posts       []*PostDetail
	}

	PostList struct {
		TotalNumber int
		Posts       []*PostListTiny
	}

	// 一覧用Post
	PostListTiny struct {
		PostTiny
		User            *User `gorm:"foreignkey:UserID"`
		IsFavorite      bool
		AreaCategories  []*AreaCategory  `gorm:"many2many:post_area_category;jointable_foreignkey:post_id;"`
		ThemeCategories []*ThemeCategory `gorm:"many2many:post_theme_category;jointable_foreignkey:post_id;"`
	}

	// 参照用Post詳細
	PostDetailWithHashtag struct {
		PostTiny
		Bodies          []*PostBody      `gorm:"foreignkey:PostID"`
		User            *User            `gorm:"foreignkey:UserID"`
		AreaCategories  []*AreaCategory  `gorm:"many2many:post_area_category;jointable_foreignkey:post_id;"`
		ThemeCategories []*ThemeCategory `gorm:"many2many:post_theme_category;jointable_foreignkey:post_id;"`
		Hashtag         []*Hashtag       `gorm:"many2many:post_hashtag;jointable_foreignkey:post_id;"`
	}

	PostDetailWithHashtagAndIsFavorite struct {
		PostTiny
		Bodies          []*PostBody `gorm:"foreignkey:PostID"`
		User            *User       `gorm:"foreignkey:UserID"`
		IsFavorite      bool
		AreaCategories  []*AreaCategory  `gorm:"many2many:post_area_category;jointable_foreignkey:post_id;"`
		ThemeCategories []*ThemeCategory `gorm:"many2many:post_theme_category;jointable_foreignkey:post_id;"`
		Hashtag         []*Hashtag       `gorm:"many2many:post_hashtag;jointable_foreignkey:post_id;"`
	}
)

func (post *PostDetail) TableName() string {
	return "post"
}

func (post *PostDetailWithHashtag) TableName() string {
	return "post"
}

func (post *PostListTiny) TableName() string {
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
				// MEMO: int64->intに変換しているので値が溢れる可能性がある
				ids = append(ids, int(area.SubAreaID.Int64))
			}

			if area.SubSubAreaID.Valid {
				// MEMO: int64->intに変換しているので値が溢れる可能性がある
				ids = append(ids, int(area.SubSubAreaID.Int64))
			}
		}
	}

	// 昇順
	sort.Slice(ids, func(i, j int) bool {
		return ids[i] < ids[j]
	})

	results := make([]int, 0)

	// 重複,0削除
	encountered := map[int]bool{}
	for i := 0; i < len(ids); i++ {
		if !encountered[ids[i]] && ids[i] != 0 {
			encountered[ids[i]] = true
			results = append(results, ids[i])
		}
	}

	return results
}
