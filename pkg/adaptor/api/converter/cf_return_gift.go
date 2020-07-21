package converter

import (
	"sort"

	"github.com/stayway-corp/stayway-media-api/pkg/adaptor/api/output"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/entity"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/model"
)

func (c *Converters) ConvertCfReturnGiftListToOutput(gifts *entity.CfReturnGiftList, soldCountList *entity.CfReturnGiftSoldCountList) []*output.CfReturnGift {
	sort.Slice(gifts.List, func(i, j int) bool { return gifts.List[i].Snapshot.SortOrder < gifts.List[j].Snapshot.SortOrder })
	response := make([]*output.CfReturnGift, len(gifts.List))
	for i, gift := range gifts.List {
		response[i] = c.convertCfReturnGiftToOutput(gift, soldCountList.GetSoldCount(gift.ID))
	}
	return response
}

func (c *Converters) convertCfReturnGiftToOutput(gift *entity.CfReturnGift, soldCount int) *output.CfReturnGift {
	var deadline *model.Date
	if gift.Snapshot.Deadline.Valid {
		deadline = (*model.Date)(&gift.Snapshot.Deadline.Time)
	}

	return &output.CfReturnGift{
		ID:           gift.ID,
		SnapshotID:   gift.Snapshot.SnapshotID,
		Thumbnail:    gift.Snapshot.Thumbnail,
		GiftType:     gift.GiftType,
		Body:         gift.Snapshot.Body,
		Price:        gift.Snapshot.Price,
		IsCancelable: gift.Snapshot.IsCancelable,
		Deadline:     deadline,
		SoldCount:    soldCount,
		DeliveryDate: gift.Snapshot.DeliveryDate,
	}
}
