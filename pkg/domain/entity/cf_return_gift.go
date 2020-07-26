package entity

import (
	"fmt"

	"github.com/stayway-corp/stayway-media-api/pkg/domain/model"
	"github.com/stayway-corp/stayway-media-api/pkg/util"
	"gopkg.in/guregu/null.v3"
)

type (
	CfReturnGiftTiny struct {
		ID               int `gorm:"primary_key"`
		CfProjectID      int
		LatestSnapshotID null.Int
		GiftType         model.CfReturnGiftType
		Times
	}

	CfReturnGiftSnapshotTiny struct {
		SnapshotID     int `gorm:"column:id;primary_key"`
		CfReturnGiftID int
		Thumbnail      string
		Body           string
		Price          int
		FullAmount     int
		DeliveryDate   string
		IsCancelable   bool
		Deadline       null.Time
		SortOrder      int
		Times
	}

	CfReturnGift struct {
		CfReturnGiftTiny
		Snapshot *CfReturnGiftSnapshotTiny // `gorm:"foreignkey:ID;association_foreignkey:LatestSnapshotID"`
	}

	CfReturnGiftList struct {
		List []*CfReturnGift
	}

	CfReturnGiftSoldCount struct {
		CfReturnGiftID int
		SoldCount      int
	}

	CfReturnGiftWithCount struct {
		CfReturnGiftTiny
		Snapshot       *CfReturnGiftSnapshotTiny `gorm:"foreignkey:ID;association_foreignkey:LatestSnapshotID"`
		SoldCount      int
		SupporterCount int
	}

	CfReturnGiftSoldCountList struct {
		List []*CfReturnGiftSoldCount
	}

	CfReturnGiftWithCountList struct {
		List []*CfReturnGiftWithCount
	}
)

// id:CfReturnGiftDetailのmapへ変換
func (r *CfReturnGiftList) ToIDMap() map[int]*CfReturnGift {
	idMap := make(map[int]*CfReturnGift, len(r.List))
	for _, summary := range r.List {
		idMap[summary.ID] = summary
	}
	return idMap
}

func (r *CfReturnGiftList) IDs() []int {
	ids := make([]int, len(r.List))
	for i, summary := range r.List {
		ids[i] = summary.ID
	}
	return util.RemoveDuplicatesAndZeroFromIntSlice(ids)
}

func (r *CfReturnGiftList) UniqueCfProjectID() (int, bool) {
	ids := make([]int, len(r.List))
	for i, summary := range r.List {
		ids[i] = summary.CfProjectID
	}
	if len(util.RemoveDuplicatesAndZeroFromIntSlice(ids)) == 1 {
		return ids[0], true
	}
	return 0, false
}

func (r *CfReturnGiftList) OnEmailDescription() string {
	var body string
	for _, gift := range r.List {
		body += fmt.Sprintf("<br>%s", gift.Snapshot.Body)
	}
	return body
}

func (r *CfReturnGiftSoldCountList) GetSoldCount(id int) int {
	for _, summary := range r.List {
		if summary.CfReturnGiftID == id {
			fmt.Println(summary.SoldCount)
			return summary.SoldCount
		}
	}
	return 0
}

/**************************
         TableName
***************************/

func (c *CfReturnGiftTiny) TableName() string {
	return "cf_return_gift"
}

func (c *CfReturnGiftSnapshotTiny) TableName() string {
	return "cf_return_gift_snapshot"
}

func (c *CfReturnGiftWithCount) TableName() string {
	return "cf_return_gift"
}
