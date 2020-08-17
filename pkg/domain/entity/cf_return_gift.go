package entity

import (
	"fmt"
	"time"

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
		Title          string
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

// お問い合わせ番号とタイトルを並べる
func (r *CfReturnGiftList) TitlesOnEmail(idInquiryCodeMap map[int]string) string {
	var titles string
	for _, gift := range r.List {
		if gift.Snapshot.Deadline.Valid {
			titles += fmt.Sprintf("<br>%s<br>お問い合わせ番号: %s<br>有効期限: %s<br>", gift.Snapshot.Title, idInquiryCodeMap[gift.ID], model.TimeFront(gift.Snapshot.Deadline.Time).ToString())
		} else {
			titles += fmt.Sprintf("<br>%s<br>お問い合わせ番号: %s<br>", gift.Snapshot.Title, idInquiryCodeMap[gift.ID])
		}
	}
	return titles
}

func (r *CfReturnGiftSnapshotTiny) IsExpired() bool {
	if deadline := r.Deadline; deadline.Valid && deadline.Time.Before(time.Now()) {
		return true
	}
	return false
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

// 全てのCfReturnGiftのGiftTypeが引数に取ったものと等しいか
func (r *CfReturnGiftList) IsAllCfReturnGiftTypeEqualArg(giftType model.CfReturnGiftType) bool {
	for _, pCfReturnGift := range r.List {
		if pCfReturnGift.CfReturnGiftTiny.GiftType != giftType {
			return false
		}
	}

	return true
}

/**************************
         TableName
***************************/

func (c *CfReturnGiftTiny) TableName() string {
	return "cf_return_gift"
}

func (r *CfReturnGiftSnapshotTiny) TableName() string {
	return "cf_return_gift_snapshot"
}

func (c *CfReturnGiftWithCount) TableName() string {
	return "cf_return_gift"
}
