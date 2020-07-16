package repository

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/entity"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/model"
	"github.com/stayway-corp/stayway-media-api/pkg/util"
)

var _ = Describe("CfReturnGiftRepositoryImpl", func() {
	var (
		commandRepo *CfReturnGiftCommandRepositoryImpl
		base        *entity.CfReturnGift
	)
	findCfReturnGiftByID := func(id int) (*entity.CfReturnGift, error) {
		var cfReturnGift entity.CfReturnGift
		return &cfReturnGift, db.Find(&cfReturnGift, id).Error
	}

	BeforeEach(func() {
		commandRepo = tests.CfReturnGiftCommandRepositoryImpl
		truncate(db)
		Expect(db.Save(newUser(userID)).Error).To(Succeed())
		Expect(db.Save(newCfProject(cfProjectID, userID, nil, nil, nil)).Error).To(Succeed())

		base = newCfReturnGift(cfReturnGiftID, cfProjectID)
		Expect(db.Save(base).Error).To(Succeed())
	})

	Describe("Store", func() {
		Context("新規作成時", func() {
			It("cf_return_giftとcf_return_gift_snapshotが作成される", func() {
				const newInsertedID = cfReturnGiftID + 1
				newInserted := newCfReturnGift(newInsertedID, cfProjectID)
				Expect(commandRepo.Store(newInserted)).To(Succeed())

				var count int
				Expect(db.Model(&entity.CfReturnGiftTiny{}).Count(&count).Error).To(Succeed())
				Expect(count).To(Equal(2))
				Expect(db.Model(&entity.CfReturnGiftSnapshotTiny{}).Count(&count).Error).To(Succeed())
				Expect(count).To(Equal(2))

				actual, err := findCfReturnGiftByID(newInsertedID)
				Expect(err).To(Succeed())

				newInserted.LatestSnapshotID = actual.LatestSnapshotID
				Expect(actual).To(entity.EqualEntity(newInserted))
			})
		})

		Context("更新時", func() {
			It("cf_return_giftのカラムは更新されず、cf_return_gift_snapshotが挿入される。cf_return_gift.latest_snapshot_idが更新される。", func() {
				Expect(commandRepo.Store(base)).To(Succeed())

				var count int
				Expect(db.Model(&entity.CfReturnGiftTiny{}).Count(&count).Error).To(Succeed())
				Expect(count).To(Equal(1))
				Expect(db.Model(&entity.CfReturnGiftSnapshotTiny{}).Count(&count).Error).To(Succeed())
				Expect(count).To(Equal(2))

				actual, err := findCfReturnGiftByID(cfReturnGiftID)
				Expect(err).To(Succeed())

				var latestSnapshotID int64
				Expect(db.Raw("SELECT MAX(id) FROM cf_return_gift_snapshot").Row().Scan(&latestSnapshotID)).To(Succeed())
				Expect(actual.LatestSnapshotID.Int64).To(Equal(latestSnapshotID))

				base.LatestSnapshotID = actual.LatestSnapshotID
				base.Snapshot.SnapshotID = actual.Snapshot.SnapshotID
				Expect(actual).To(entity.EqualEntity(base))
			})
		})
	})
})

func newCfReturnGift(id, cfProjectID int) *entity.CfReturnGift {
	cfReturnGift := &entity.CfReturnGift{
		CfReturnGiftTiny: entity.CfReturnGiftTiny{
			ID:          id,
			CfProjectID: cfProjectID,
			GiftType:    model.CfReturnGiftTypeReservedTicket,
		},
		Snapshot: &entity.CfReturnGiftSnapshotTiny{
			CfReturnGiftID: id,
			SortOrder:      10,
			Price:          10000,
			FullAmount:     1000,
		},
	}
	util.FillDummyString(cfReturnGift, id)
	return cfReturnGift
}
