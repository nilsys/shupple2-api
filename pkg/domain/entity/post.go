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

	UserFavoritePost struct {
		UserID int
		PostID int
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
func (p *Post) GetCategoryIDs() []int {
	var ids []int
	for _, postCategory := range p.CategoryIDs {
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
	post := Post{PostTiny: tiny}

	post.SetBodies(bodies)
	post.SetCategories(categoryIDs)
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

func (p *Post) SetCategories(categoryIDs []int) {
	p.CategoryIDs = make([]*PostCategory, len(categoryIDs))
	for i, c := range categoryIDs {
		p.CategoryIDs[i] = &PostCategory{
			PostID:     p.ID,
			CategoryID: c,
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
