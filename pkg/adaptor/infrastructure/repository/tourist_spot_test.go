package repository

import (
	"context"

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
		for _, cat := range append(areaCategoryIDs, addedAreaCategoryID) {
			Expect(db.Save(newAreaCategory(cat)).Error).To(Succeed())
		}
		for _, cat := range append(themeCategoryIDs, addedThemeCategoryID) {
			Expect(db.Save(newThemeCategory(cat)).Error).To(Succeed())
		}
		for _, lcat := range append(lcategoryIDs, addedLcategoryID) {
			Expect(db.Save(newLcategory(lcat)).Error).To(Succeed())
		}
	})

	base := newTouristSpot(touristSpotID, areaCategoryIDs, themeCategoryIDs, lcategoryIDs)
	baseChanged := newTouristSpot(touristSpotID, areaCategoryIDs, themeCategoryIDs, lcategoryIDs)
	baseChanged.Name = "changed"

	DescribeTable("Saveは引数のtouristSpotを作成するか、その状態になるように更新する",
		func(before *entity.TouristSpot, saved *entity.TouristSpot) {
			if before != nil {
				Expect(command.Store(context.Background(), before)).To(Succeed())
			}

			Expect(command.Store(context.Background(), saved)).To(Succeed())
			actual, err := query.FindByID(saved.ID)
			Expect(err).To(Succeed())

			Expect(actual.CreatedAt).NotTo(BeZero())
			Expect(actual.UpdatedAt).NotTo(BeZero())
			Expect(actual).To(Equal(saved))
		},
		Entry("新規作成", nil, base),
		Entry("フィールドに変更がある場合", base, baseChanged),
		Entry("area_categoryが追加される場合", base, newTouristSpot(touristSpotID, append(areaCategoryIDs, addedAreaCategoryID), themeCategoryIDs, lcategoryIDs)),
		Entry("theme_categoryが追加される場合", base, newTouristSpot(touristSpotID, areaCategoryIDs, append(themeCategoryIDs, addedThemeCategoryID), lcategoryIDs)),
		Entry("lcategoryが追加される場合", base, newTouristSpot(touristSpotID, areaCategoryIDs, themeCategoryIDs, append(lcategoryIDs, addedLcategoryID))),
		Entry("area_categoryが削除される場合", base, newTouristSpot(touristSpotID, areaCategoryIDs[:1], themeCategoryIDs, lcategoryIDs)),
		Entry("theme_categoryが削除される場合", base, newTouristSpot(touristSpotID, areaCategoryIDs, themeCategoryIDs[:1], lcategoryIDs)),
		Entry("lcategoryが削除される場合", base, newTouristSpot(touristSpotID, areaCategoryIDs, themeCategoryIDs, lcategoryIDs[:1])),
	)

	Describe("FindDetailByID", func() {
		It("関連entityを含めてTouristSpotを取得する", func() {
			Expect(command.Store(context.Background(), base)).To(Succeed())

			actual, err := query.FindDetailByID(base.ID)
			Expect(err).To(Succeed())

			Expect(actual.CreatedAt).NotTo(BeZero())
			Expect(actual.UpdatedAt).NotTo(BeZero())

			Expect(actual.TouristSpotTiny).To(Equal(base.TouristSpotTiny))

			Expect(actual.AreaCategories).To(HaveLen(len(base.AreaCategoryIDs)))
			for i, c := range actual.AreaCategories {
				Expect(c.ID).To(Equal(base.AreaCategoryIDs[i].AreaCategoryID))
				Expect(c.Name).NotTo(BeEmpty())
			}

			Expect(actual.ThemeCategories).To(HaveLen(len(base.ThemeCategoryIDs)))
			for i, c := range actual.ThemeCategories {
				Expect(c.ID).To(Equal(base.ThemeCategoryIDs[i].ThemeCategoryID))
				Expect(c.Name).NotTo(BeEmpty())
			}

			Expect(actual.Lcategories).To(HaveLen(len(base.LcategoryIDs)))
			for i, c := range actual.Lcategories {
				Expect(c.ID).To(Equal(base.LcategoryIDs[i].LcategoryID))
				Expect(c.Name).NotTo(BeEmpty())
			}
		})
	})
})

func newTouristSpot(id int, areaCategoryIDs, themeCategoryIDs, lcategoryIDs []int) *entity.TouristSpot {
	touristSpot := entity.TouristSpotTiny{
		ID:        id,
		Lat:       float64(id),
		Lng:       float64(id * 10),
		CreatedAt: sampleTime,
		UpdatedAt: sampleTime,
	}
	util.FillDummyString(&touristSpot, id)

	l := entity.NewTouristSpot(touristSpot, areaCategoryIDs, themeCategoryIDs, lcategoryIDs)
	return &l
}
