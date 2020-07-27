package converter

import (
	"sort"

	"github.com/stayway-corp/stayway-media-api/pkg/adaptor/api/output"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/entity"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/model"
)

func (c *Converters) ConvertCfReturnGiftWithCountListToOutput(gifts *entity.CfReturnGiftWithCountList) []*output.CfReturnGift {
	sort.Slice(gifts.List, func(i, j int) bool { return gifts.List[i].Snapshot.SortOrder < gifts.List[j].Snapshot.SortOrder })
	response := make([]*output.CfReturnGift, len(gifts.List))
	for i, gift := range gifts.List {
		response[i] = c.convertCfReturnGiftWithCountToOutput(gift)
	}
	return response
}

func (c *Converters) convertCfReturnGiftWithCountToOutput(gift *entity.CfReturnGiftWithCount) *output.CfReturnGift {
	var deadline *model.Date
	if gift.Snapshot.Deadline.Valid {
		deadline = (*model.Date)(&gift.Snapshot.Deadline.Time)
	}

	return &output.CfReturnGift{
		ID:             gift.ID,
		SnapshotID:     gift.Snapshot.SnapshotID,
		CfProjectID:    gift.CfProjectID,
		Thumbnail:      gift.Snapshot.Thumbnail,
		GiftType:       gift.GiftType,
		Body:           gift.Snapshot.Body,
		Price:          gift.Snapshot.Price,
		IsCancelable:   gift.Snapshot.IsCancelable,
		Deadline:       deadline,
		SoldCount:      gift.SoldCount,
		SupporterCount: gift.SupporterCount,
		FullAmount:     gift.Snapshot.FullAmount,
		DeliveryDate:   gift.Snapshot.DeliveryDate,
	}
}

func (c *Converters) convertCfReturnGiftToOutput(gift *entity.CfReturnGift) *output.CfReturnGift {
	var deadline *model.Date
	if gift.Snapshot.Deadline.Valid {
		deadline = (*model.Date)(&gift.Snapshot.Deadline.Time)
	}

	return &output.CfReturnGift{
		ID:           gift.ID,
		SnapshotID:   gift.Snapshot.SnapshotID,
		CfProjectID:  gift.CfProjectID,
		Thumbnail:    gift.Snapshot.Thumbnail,
		GiftType:     gift.GiftType,
		Body:         gift.Snapshot.Body,
		Price:        gift.Snapshot.Price,
		IsCancelable: gift.Snapshot.IsCancelable,
		Deadline:     deadline,
		FullAmount:   gift.Snapshot.FullAmount,
		DeliveryDate: gift.Snapshot.DeliveryDate,
	}
}
