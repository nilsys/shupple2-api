package repository

import (
	"context"
	"time"

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
		for _, cat := range append(areaCategoryIDs, addedAreaCategoryID) {
			Expect(db.Save(newAreaCategory(cat)).Error).To(Succeed())
		}
		for _, cat := range append(themeCategoryIDs, addedThemeCategoryID) {
			Expect(db.Save(newThemeCategory(cat)).Error).To(Succeed())
		}
		for _, loc := range append(touristSpotIDs, addedTouristSpotID) {
			Expect(db.Save(newTouristSpot(loc, nil, nil, nil)).Error).To(Succeed())
		}
		for _, user := range append(userIDs, addedUserID) {
			Expect(db.Save(newUser(user)).Error).To(Succeed())
		}
	})

	base := newVlog(vlogID, areaCategoryIDs, themeCategoryIDs, touristSpotIDs, userIDs)
	baseChanged := newVlog(vlogID, areaCategoryIDs, themeCategoryIDs, touristSpotIDs, userIDs)
	baseChanged.Title = "changed"

	DescribeTable("Saveは引数のvlogを作成するか、その状態になるように更新する",
		func(before *entity.Vlog, saved *entity.Vlog) {
			if before != nil {
				Expect(command.Store(context.Background(), before)).To(Succeed())
			}

			Expect(command.Store(context.Background(), saved)).To(Succeed())
			actual, err := query.FindByID(saved.ID)
			Expect(err).To(Succeed())

			Expect(actual.CreatedAt).NotTo(BeZero())
			Expect(actual.UpdatedAt).NotTo(BeZero())
			actual.CreatedAt = time.Time{}
			actual.UpdatedAt = time.Time{}

			Expect(actual).To(Equal(saved))
		},
		Entry("新規作成", nil, base),
		Entry("フィールドに変更がある場合", base, baseChanged),
		Entry("areaCategoryが追加される場合", base, newVlog(vlogID, append(areaCategoryIDs, addedAreaCategoryID), themeCategoryIDs, touristSpotIDs, userIDs)),
		Entry("themeCategoryが追加される場合", base, newVlog(vlogID, areaCategoryIDs, append(themeCategoryIDs, addedThemeCategoryID), touristSpotIDs, userIDs)),
		Entry("touristSpotが追加される場合", base, newVlog(vlogID, areaCategoryIDs, themeCategoryIDs, append(touristSpotIDs, addedTouristSpotID), userIDs)),
		Entry("editorが追加される場合", base, newVlog(vlogID, areaCategoryIDs, themeCategoryIDs, touristSpotIDs, append(userIDs, addedUserID))),
		Entry("areaCategoryが削除される場合", base, newVlog(vlogID, areaCategoryIDs[:1], themeCategoryIDs, touristSpotIDs, userIDs)),
		Entry("themeCategoryが削除される場合", base, newVlog(vlogID, areaCategoryIDs, themeCategoryIDs[:1], touristSpotIDs, userIDs)),
		Entry("touristSpotが削除される場合", base, newVlog(vlogID, areaCategoryIDs, themeCategoryIDs, touristSpotIDs[:1], userIDs)),
		Entry("editorが削除される場合", base, newVlog(vlogID, areaCategoryIDs, themeCategoryIDs, touristSpotIDs, userIDs[:1])),
	)

	Describe("FindDetailByID", func() {
		It("関連entityを含めてTouristSpotを取得する", func() {
			Expect(command.Store(context.Background(), base)).To(Succeed())

			actual, err := query.FindDetailByID(base.ID)
			Expect(err).To(Succeed())

			Expect(actual.CreatedAt).NotTo(BeZero())
			Expect(actual.UpdatedAt).NotTo(BeZero())
			actual.CreatedAt = time.Time{}
			actual.UpdatedAt = time.Time{}

			Expect(actual.VlogTiny).To(Equal(base.VlogTiny))

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

			Expect(actual.TouristSpots).To(HaveLen(len(base.TouristSpotIDs)))
			for i, c := range actual.TouristSpots {
				Expect(c.ID).To(Equal(base.TouristSpotIDs[i].TouristSpotID))
				Expect(c.Name).NotTo(BeEmpty())
			}

			Expect(actual.Editors).To(HaveLen(len(base.Editors)))
			for i, c := range actual.Editors {
				Expect(c.ID).To(Equal(base.Editors[i].UserID))
				Expect(c.Name).NotTo(BeEmpty())
			}
		})
	})
})

func newVlog(id int, areaCategoryIDs, themeCategoryIDs, touristSpotIDs, editors []int) *entity.Vlog {
	vlog := entity.VlogTiny{
		ID:       id,
		UserID:   userID,
		EditedAt: time.Now().Truncate(time.Second),
	}
	util.FillDummyString(&vlog, id)

	v := entity.NewVlog(vlog, areaCategoryIDs, themeCategoryIDs, touristSpotIDs, editors)
	return &v
}
