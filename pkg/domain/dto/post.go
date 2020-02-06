package dto

import "github.com/stayway-corp/stayway-media-api/pkg/domain/entity"

type PostDetail struct {
	Post       *entity.Post
	Categories []*entity.Category
	User       *entity.User
}
