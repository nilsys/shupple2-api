package repository

import (
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/entity"
	"github.com/stayway-corp/stayway-media-api/pkg/util"
)

var _ = Describe("CategoryRepositoryImpl", func() {
	var (
		command *CategoryCommandRepositoryImpl
		query   *CategoryQueryRepositoryImpl
	)

	BeforeEach(func() {
		command = tests.CategoryCommandRepositoryImpl
		query = tests.CategoryQueryRepositoryImpl

		truncate(tests.DB)
		Expect(db.Save(newUser(userID)).Error).To(Succeed())
	})

	base := newCategory(categoryID)
	baseChanged := newCategory(categoryID)
	baseChanged.Name = "changed"

	DescribeTable("Saveは引数のcategoryを作成するか、その状態になるように更新する",
		func(before *entity.Category, saved *entity.Category) {
			if before != nil {
				Expect(command.Store(before)).To(Succeed())
			}

			Expect(command.Store(saved)).To(Succeed())
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
	)

	Describe("Metasearch系のIDは更新はされないが取得はできる", func() {
		const (
			metasearchAreaID       = 123
			metasearchSubAreaID    = 456
			metasearchSubSubAreaID = 789
		)
		BeforeEach(func() {
			category := newCategory(categoryID)
			Expect(command.Store(category)).To(Succeed())

			updateResult := tests.DB.Exec(
				"UPDATE category SET metasearch_area_id = ?, metasearch_sub_area_id = ?, metasearch_sub_sub_area_id = ? WHERE id = ?",
				metasearchAreaID, metasearchSubAreaID, metasearchSubSubAreaID, category.ID,
			)
			Expect(updateResult.Error).To(Succeed())
		})

		It("更新後に取得", func() {
			updatedCategory := newCategory(categoryID)
			updatedCategory.Name = "updated"
			updatedCategory.MetasearchAreaID = 1
			updatedCategory.MetasearchSubAreaID = 1
			updatedCategory.MetasearchSubSubAreaID = 1
			Expect(command.Store(updatedCategory)).To(Succeed())

			actual, err := query.FindByID(updatedCategory.ID)
			Expect(err).To(Succeed())

			Expect(actual.Name).To(Equal(updatedCategory.Name))
			Expect(actual.MetasearchAreaID).To(Equal(metasearchAreaID))
			Expect(actual.MetasearchSubAreaID).To(Equal(metasearchSubAreaID))
			Expect(actual.MetasearchSubSubAreaID).To(Equal(metasearchSubSubAreaID))
		})
	})
})

func newCategory(id int) *entity.Category {
	category := &entity.Category{ID: id}
	util.FillDymmyString(category, id)
	return category
}
