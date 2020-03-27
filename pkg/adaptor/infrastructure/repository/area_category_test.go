package repository

import (
	"context"
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/entity"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/model"
	"github.com/stayway-corp/stayway-media-api/pkg/util"
)

var _ = Describe("AreaCategoryRepositoryImpl", func() {
	var (
		command *AreaCategoryCommandRepositoryImpl
		query   *AreaCategoryQueryRepositoryImpl
	)

	BeforeEach(func() {
		command = tests.AreaCategoryCommandRepositoryImpl
		query = tests.AreaCategoryQueryRepositoryImpl

		truncate(tests.DB)
		Expect(db.Save(newUser(userID)).Error).To(Succeed())
	})

	base := newAreaCategory(areaCategoryID)
	baseChanged := newAreaCategory(areaCategoryID)
	baseChanged.Name = "changed"

	DescribeTable("Saveは引数のareaCategoryを作成するか、その状態になるように更新する",
		func(before *entity.AreaCategory, saved *entity.AreaCategory) {
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

func newAreaCategory(id int) *entity.AreaCategory {
	areaCategory := &entity.AreaCategory{}
	areaCategory.ID = id
	areaCategory.AreaID = id
	areaCategory.Type = model.AreaCategoryTypeArea
	util.FillDummyString(areaCategory, id)
	return areaCategory
}
