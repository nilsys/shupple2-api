package output

import "github.com/stayway-corp/stayway-media-api/pkg/domain/model"

type (
	CfReturnGift struct {
		ID         int            `json:"id"`
		SnapshotID int            `json:"snapshotId"`
		Thumbnail  string         `json:"thumbnail"`
		GiftType   model.GiftType `json:"giftType"`
		Body       string         `json:"body"`
		Price      string         `json:"price"`
	}
)
