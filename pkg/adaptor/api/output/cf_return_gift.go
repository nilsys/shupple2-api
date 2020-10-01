package output

import "github.com/stayway-corp/stayway-media-api/pkg/domain/model"

type (
	CfReturnGiftList struct {
		TotalNumber int             `json:"totalNumber"`
		List        []*CfReturnGift `json:"cfReturnGifts"`
	}

	CfReturnGift struct {
		ID             int                    `json:"id"`
		SnapshotID     int                    `json:"snapshotId"`
		CfProjectID    int                    `json:"cfProjectId"`
		Title          string                 `json:"title"`
		Thumbnail      string                 `json:"thumbnail"`
		GiftType       model.CfReturnGiftType `json:"giftType"`
		Body           string                 `json:"body"`
		Price          int                    `json:"price"`
		IsCancelable   bool                   `json:"isCancelable"`
		Deadline       *model.Date            `json:"deadline"`
		SoldCount      int                    `json:"soldCount"`
		SupporterCount int                    `json:"supporterCount"`
		FullAmount     int                    `json:"fullAmount"`
		DeliveryDate   string                 `json:"deliveryDate"`
	}
)
