package entity

import (
	"fmt"

	"github.com/stayway-corp/stayway-media-api/pkg/domain/model"
	"github.com/stayway-corp/stayway-media-api/pkg/util"
)

type (
	CfReturnGiftTable struct {
		ID                           int `gorm:"primary_key"`
		CfProjectID                  int
		LatestCfReturnGiftSnapshotID int
		Thumbnail                    string
		SortOrder                    int
		GiftType                     model.GiftType
		Times
	}

	CfReturnGiftSnapshotTable struct {
		ID             int `gorm:"primary_key"`
		CfReturnGiftID int
		Thumbnail      string
		Body           string
		Price          int
		FullAmount     int
		Times
	}

	CfReturnGift struct {
		CfReturnGiftTable
		Snapshot *CfReturnGiftSnapshotTable `gorm:"foreignkey:LatestCfReturnGiftSnapshotID"`
		Times
	}

	CfReturnGiftList struct {
		List []*CfReturnGift
	}

	CfReturnGiftSoldCount struct {
		ReturnGiftID int
		SoldCount    int
	}

	CfReturnGiftSoldCountList struct {
		List []*CfReturnGiftSoldCount
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
		if summary.ReturnGiftID == id {
			return summary.SoldCount
		}
	}
	return 0
}

/**************************
         TableName
***************************/

func (c *CfReturnGiftTable) TableName() string {
	return "cf_return_gift"
}

func (c *CfReturnGiftSnapshotTable) TableName() string {
	return "cf_return_gift_snapshot"
}
