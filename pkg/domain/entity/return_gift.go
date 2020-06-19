package entity

import "github.com/stayway-corp/stayway-media-api/pkg/domain/model"

type (
	CfReturnGiftTable struct {
		ID               int `gorm:"primary_key"`
		CfProjectID      int
		LatestSnapshotID int
		Thumbnail        string
		SortOrder        int
		GiftType         model.GiftType
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
		Summary *CfReturnGiftSnapshotTable `gorm:"foreignkey:LatestSnapshotID"`
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

func (r *CfReturnGift) TableName() string {
	return "return_gift"
}

// id:CfReturnGiftDetailのmapへ変換
func (r *CfReturnGiftList) ToIDMap() map[int]*CfReturnGift {
	idMap := make(map[int]*CfReturnGift, len(r.List))
	for _, summary := range r.List {
		idMap[summary.ID] = summary
	}
	return idMap
}

func (r *CfReturnGiftList) CfProjectIDs() []int {
	ids := make([]int, len(r.List))
	for i, summary := range r.List {
		ids[i] = summary.CfProjectID
	}
	return ids
}

// id:sold_countのmapへ変換
func (r *CfReturnGiftSoldCountList) ToIDSoldCountMap() map[int]int {
	result := make(map[int]int, len(r.List))
	for _, summary := range r.List {
		result[summary.ReturnGiftID] = summary.SoldCount
	}
	return result
}

func (c *CfReturnGiftTable) TableName() string {
	return "cf_return_gift"
}

func (c *CfReturnGiftSnapshotTable) TableName() string {
	return "cf_return_gift_snapshot"
}
