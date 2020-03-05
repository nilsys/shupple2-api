package entity

import (
	"strconv"
	"time"
)

type (
	PostTiny struct {
		ID            int `gorm:"primary_key"`
		UserID        int
		Title         string
		TOC           string
		FavoriteCount int
		FacebookCount int
		TwitterCount  int
		Views         int
		Slug          string
		CreatedAt     time.Time `gorm:"default:current_timestamp"`
		UpdatedAt     time.Time `gorm:"default:current_timestamp"`
		DeletedAt     *time.Time
	}

	Post struct {
		PostTiny
		Bodies      []*PostBody     `gorm:"foreignkey:PostID"`
		CategoryIDs []*PostCategory `gorm:"foreignkey:PostID"`
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

	// 参照用Post
	PostDetail struct {
		PostTiny
		Bodies     []*PostBody `gorm:"foreignkey:PostID"`
		User       *User       `gorm:"foreignkey:UserID"`
		Categories []*Category `gorm:"many2many:post_category;jointable_foreignkey:post_id;"`
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

// MEMO: サムネイルロジック仮置き
func (post *PostTiny) GenerateThumbnailURL() string {
	return "https://files.stayway.jp/post/" + strconv.Itoa(post.ID)
}

func (post *PostDetail) TableName() string {
	return "post"
}

func (post *PostDetailWithHashtag) TableName() string {
	return "post"
}

func NewPost(tiny PostTiny, bodies []string, categoryIDs []int) Post {
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

	return Post{
		tiny,
		postBodies,
		postCategoryIDs,
	}
}
