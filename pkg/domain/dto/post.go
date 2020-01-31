package dto

import "github.com/stayway-corp/stayway-media-api/pkg/domain/entity"

type PostAndCategories struct {
	Post       *entity.Post
	Categories []*entity.Category
	// TODO: ここにUserも入る
}
