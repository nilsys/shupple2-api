package repository

import (
	"github.com/stayway-corp/stayway-media-api/pkg/domain/entity"
	"github.com/stayway-corp/stayway-media-api/pkg/util"
)

func newHashtag(id int) *entity.Hashtag {
	hashtag := &entity.Hashtag{
		ID: id,
	}
	util.FillDymmyString(hashtag, id)
	return hashtag
}
