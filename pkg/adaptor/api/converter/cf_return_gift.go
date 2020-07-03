package converter

import (
	"github.com/stayway-corp/stayway-media-api/pkg/adaptor/api/output"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/entity"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/model"
	"github.com/stayway-corp/stayway-media-api/pkg/util"
)

func (c *Converters) ConvertCfReturnGiftListToOutput(gifts *entity.CfReturnGiftList, soldCountList *entity.CfReturnGiftSoldCountList) []*output.CfReturnGift {
	response := make([]*output.CfReturnGift, len(gifts.List))
	for i, gift := range gifts.List {
		response[i] = c.convertCfReturnGiftToOutput(gift, soldCountList.GetSoldCount(gift.ID))
	}
	return response
}

func (c *Converters) convertCfReturnGiftToOutput(gift *entity.CfReturnGift, soldCount int) *output.CfReturnGift {
	return &output.CfReturnGift{
		ID:           gift.ID,
		SnapshotID:   gift.Snapshot.ID,
		Thumbnail:    gift.Thumbnail,
		GiftType:     gift.GiftType,
		Body:         gift.Snapshot.Body,
		Price:        util.WithComma(gift.Snapshot.Price),
		IsCancelable: gift.Snapshot.IsCancelable,
		Deadline:     model.Date(gift.Snapshot.Deadline),
		SoldCount:    soldCount,
	}
}
