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
	)
})

func newCategory(id int) *entity.Category {
	category := &entity.Category{ID: id}
	util.FillDymmyString(category, id)
	return category
}
