package converter

import (
	"github.com/stayway-corp/stayway-media-api/pkg/adaptor/api/output"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/entity"
)

func (c Converters) ConvertHashtagListToOutput(hashtags []*entity.Hashtag) []*output.Hashtag {
	res := make([]*output.Hashtag, len(hashtags))

	for i, hashtag := range hashtags {
		res[i] = output.NewHashtag(hashtag.ID, hashtag.Name)
	}

	return res
}
