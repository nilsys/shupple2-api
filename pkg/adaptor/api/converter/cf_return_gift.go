package converter

import (
	"github.com/stayway-corp/stayway-media-api/pkg/adaptor/api/output"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/entity"
)

func (c *Converters) ConvertCfReturnGiftListToOutput(gifts *entity.CfReturnGiftList) []*output.CfReturnGift {
	response := make([]*output.CfReturnGift, len(gifts.List))
	for i, gift := range gifts.List {
		response[i] = c.convertCfReturnGiftToOutput(gift)
	}
	return response
}

func (c *Converters) convertCfReturnGiftToOutput(gift *entity.CfReturnGift) *output.CfReturnGift {
	return &output.CfReturnGift{
		ID:         gift.ID,
		SnapshotID: gift.Snapshot.ID,
		Thumbnail:  gift.Thumbnail,
		GiftType:   gift.GiftType,
		Body:       gift.Snapshot.Body,
		Price:      gift.Snapshot.Price,
	}
}
