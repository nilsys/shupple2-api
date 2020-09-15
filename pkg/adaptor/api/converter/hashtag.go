package converter

import (
	"github.com/stayway-corp/stayway-media-api/pkg/adaptor/api/output"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/entity"
)

func (c Converters) ConvertHashtagListToOutput(hashtags *entity.Hashtags, isFollowMap map[int]bool) []*output.Hashtag {
	res := make([]*output.Hashtag, len([]*entity.Hashtag(*hashtags)))

	for i, hashtag := range []*entity.Hashtag(*hashtags) {
		res[i] = output.NewHashtag(hashtag.ID, hashtag.Name, isFollowMap[hashtag.ID])
	}

	return res
}
