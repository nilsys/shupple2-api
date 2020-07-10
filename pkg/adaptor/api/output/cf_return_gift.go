package output

import "github.com/stayway-corp/stayway-media-api/pkg/domain/model"

type (
	CfReturnGift struct {
		ID           int                    `json:"id"`
		SnapshotID   int                    `json:"snapshotId"`
		Thumbnail    string                 `json:"thumbnail"`
		GiftType     model.CfReturnGiftType `json:"giftType"`
		Body         string                 `json:"body"`
		Price        string                 `json:"price"`
		IsCancelable bool                   `json:"isCancelable"`
		Deadline     model.Date             `json:"deadline"`
		SoldCount    int                    `json:"soldCount"`
	}
)
