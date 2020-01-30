package entity

import "time"

type (
	PostTiny struct {
		ID            int `gorm:"primary_key"`
		UserID        int
		Title         string
		TOC           string
		FavoriteCount int
		FacebookCount int
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
)

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
