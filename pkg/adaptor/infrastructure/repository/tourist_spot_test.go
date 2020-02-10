package repository

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/entity"
	"github.com/stayway-corp/stayway-media-api/pkg/util"
)

var _ = Describe("TouristSpotRepositoryImpl", func() {
	var (
		command *TouristSpotCommandRepositoryImpl
		query   *TouristSpotQueryRepositoryImpl
	)

	BeforeEach(func() {
		command = tests.TouristSpotCommandRepositoryImpl
		query = tests.TouristSpotQueryRepositoryImpl

		truncate(db)
		Expect(db.Save(newUser(userID)).Error).To(Succeed())
		for _, cat := range append(categoryIDs, addedCategoryID) {
			Expect(db.Save(newCategory(cat)).Error).To(Succeed())
		}
		for _, lcat := range append(lcategoryIDs, addedLcategoryID) {
			Expect(db.Save(newLcategory(lcat)).Error).To(Succeed())
		}
	})

	base := newTouristSpot(touristSpotID, categoryIDs, lcategoryIDs)
	baseChanged := newTouristSpot(touristSpotID, categoryIDs, lcategoryIDs)
	baseChanged.Name = "changed"

	DescribeTable("Saveは引数のtouristSpotを作成するか、その状態になるように更新する",
		func(before *entity.TouristSpot, saved *entity.TouristSpot) {
			if before != nil {
				Expect(command.Store(before)).To(Succeed())
			}

			Expect(command.Store(saved)).To(Succeed())
			actual, err := query.FindByID(saved.ID)
			Expect(err).To(Succeed())

			Expect(actual).To(Equal(saved))
		},
		Entry("新規作成", nil, base),
		Entry("フィールドに変更がある場合", base, baseChanged),
		Entry("categoryが追加される場合", base, newTouristSpot(touristSpotID, append(categoryIDs, addedCategoryID), lcategoryIDs)),
		Entry("lcategoryが追加される場合", base, newTouristSpot(touristSpotID, categoryIDs, append(lcategoryIDs, addedLcategoryID))),
		Entry("categoryが削除される場合", base, newTouristSpot(touristSpotID, categoryIDs[:1], lcategoryIDs)),
		Entry("lcategoryが削除される場合", base, newTouristSpot(touristSpotID, categoryIDs, lcategoryIDs[:1])),
	)
})

func newTouristSpot(id int, categoryIDs, lcategoryIDs []int) *entity.TouristSpot {
	touristSpot := entity.TouristSpotTiny{
		ID:        id,
		Lat:       float64(id),
		Lng:       float64(id * 10),
		CreatedAt: sampleTime,
		UpdatedAt: sampleTime,
	}
	util.FillDymmyString(&touristSpot, id)

	l := entity.NewTouristSpot(touristSpot, categoryIDs, lcategoryIDs)
	return &l
}
