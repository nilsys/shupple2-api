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

var _ = Describe("SpotCategoryRepositoryImpl", func() {
	var (
		command *SpotCategoryCommandRepositoryImpl
		query   *SpotCategoryQueryRepositoryImpl
	)

	BeforeEach(func() {
		command = tests.SpotCategoryCommandRepositoryImpl
		query = tests.SpotCategoryQueryRepositoryImpl

		truncate(db)
		Expect(db.Save(newUser(userID)).Error).To(Succeed())
	})

	base := newSpotCategory(spotCategoryID)
	baseChanged := newSpotCategory(spotCategoryID)
	baseChanged.Name = "changed"

	DescribeTable("Saveは引数のspotCategoryを作成するか、その状態になるように更新する",
		func(before *entity.SpotCategory, saved *entity.SpotCategory) {
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

func newSpotCategory(id int) *entity.SpotCategory {
	spotCategory := entity.SpotCategory{}
	spotCategory.ID = id
	spotCategory.SpotCategoryID = id
	util.FillDummyString(&spotCategory, id)

	return &spotCategory
}
