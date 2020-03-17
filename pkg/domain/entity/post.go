package entity

import (
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
		Bodies      []*PostBody
		CategoryIDs []*PostCategory
		HashtagIDs  []*PostHashtag
	}

	PostBody struct {
		PostID int `gorm:"primary_key"`
		Page   int `gorm:"primary_key"`
		Body   string
	}

	PostCategory struct {
		PostID     int `gorm:"primary_key"`
		CategoryID int `gorm:"primary_key"`
	}

	PostHashtag struct {
		PostID    int `gorm:"primary_key"`
		HashtagID int `gorm:"primary_key"`
	}

	// 参照用Post
	PostDetail struct {
		PostTiny
		Bodies     []*PostBody `gorm:"foreignkey:PostID"`
		User       *User       `gorm:"foreignkey:UserID"`
		Categories []*Category `gorm:"many2many:post_category;jointable_foreignkey:post_id;"`
	}

	PostDetailList struct {
		TotalNumber int
		Posts       []*PostDetail
	}

	// 参照用Post詳細
	PostDetailWithHashtag struct {
		PostTiny
		Bodies     []*PostBody `gorm:"foreignkey:PostID"`
		User       *User       `gorm:"foreignkey:UserID"`
		Categories []*Category `gorm:"many2many:post_category;jointable_foreignkey:post_id;"`
		Hashtag    []*Hashtag  `gorm:"many2many:post_hashtag;jointable_foreignkey:post_id;"`
	}
)

// Postが持つCategoryID(int)を配列で返す
func (post *Post) GetCategoryIDs() []int {
	var ids []int
	for _, postCategory := range post.CategoryIDs {
		ids = append(ids, postCategory.CategoryID)
	}
	return ids
}

func (post *PostDetail) TableName() string {
	return "post"
}

func (post *PostDetailWithHashtag) TableName() string {
	return "post"
}

func NewPost(tiny PostTiny, bodies []string, categoryIDs []int, hashtagIDs []int) Post {
	postBodies := make([]*PostBody, len(bodies))
	for i, body := range bodies {
		postBodies[i] = &PostBody{
			PostID: tiny.ID,
			Page:   i + 1,
			Body:   body,
		}
	}

	postCategoryIDs := make([]*PostCategory, len(categoryIDs))
	for i, c := range categoryIDs {
		postCategoryIDs[i] = &PostCategory{
			PostID:     tiny.ID,
			CategoryID: c,
		}
	}

	postHashtagIDs := make([]*PostHashtag, len(hashtagIDs))
	for i, h := range hashtagIDs {
		postHashtagIDs[i] = &PostHashtag{
			PostID:    tiny.ID,
			HashtagID: h,
		}
	}

	return Post{
		tiny,
		postBodies,
		postCategoryIDs,
		postHashtagIDs,
	}
}
