package repository

import (
	"context"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/entity"
	"github.com/stayway-corp/stayway-media-api/pkg/util"
)

var _ = Describe("VlogRepositoryImpl", func() {
	var (
		command *VlogCommandRepositoryImpl
		query   *VlogQueryRepositoryImpl
	)

	BeforeEach(func() {
		command = tests.VlogCommandRepositoryImpl
		query = tests.VlogQueryRepositoryImpl

		truncate(db)
		Expect(db.Save(newUser(userID)).Error).To(Succeed())
		for _, cat := range append(categoryIDs, addedCategoryID) {
			Expect(db.Save(newCategory(cat)).Error).To(Succeed())
		}
		for _, loc := range append(touristSpotIDs, addedTouristSpotID) {
			Expect(db.Save(newTouristSpot(loc, nil, nil)).Error).To(Succeed())
		}
	})

	base := newVlog(vlogID, categoryIDs, touristSpotIDs)
	baseChanged := newVlog(vlogID, categoryIDs, touristSpotIDs)
	baseChanged.Title = "changed"

	DescribeTable("Saveは引数のvlogを作成するか、その状態になるように更新する",
		func(before *entity.Vlog, saved *entity.Vlog) {
			if before != nil {
				Expect(command.Store(context.Background(), before)).To(Succeed())
			}

			Expect(command.Store(context.Background(), saved)).To(Succeed())
			actual, err := query.FindByID(saved.ID)
			Expect(err).To(Succeed())

			Expect(actual).To(Equal(saved))
		},
		Entry("新規作成", nil, base),
		Entry("フィールドに変更がある場合", base, baseChanged),
		Entry("categoryが追加される場合", base, newVlog(vlogID, append(categoryIDs, addedCategoryID), touristSpotIDs)),
		Entry("touristSpotが追加される場合", base, newVlog(vlogID, categoryIDs, append(touristSpotIDs, addedTouristSpotID))),
		Entry("categoryが削除される場合", base, newVlog(vlogID, categoryIDs[:1], touristSpotIDs)),
		Entry("touristSpotが削除される場合", base, newVlog(vlogID, categoryIDs, touristSpotIDs[:1])),
	)
})

func newVlog(id int, categoryIDs, touristSpotIDs []int) *entity.Vlog {
	vlog := entity.VlogTiny{
		ID:        id,
		UserID:    userID,
		CreatedAt: sampleTime,
		UpdatedAt: sampleTime,
	}
	util.FillDymmyString(&vlog, id)

	v := entity.NewVlog(vlog, categoryIDs, touristSpotIDs)
	return &v
}
