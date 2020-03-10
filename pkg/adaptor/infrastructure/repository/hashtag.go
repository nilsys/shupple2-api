package repository

import (
	"strconv"

	"github.com/stayway-corp/stayway-media-api/pkg/domain/entity"
)

func newHashtag(id int) *entity.Hashtag {
	return &entity.Hashtag{ID: id, Name: strconv.Itoa(id)}
}
