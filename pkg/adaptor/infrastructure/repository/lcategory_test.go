package repository

import (
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/entity"
	"github.com/stayway-corp/stayway-media-api/pkg/util"
)

var _ = Describe("LcategoryRepositoryImpl", func() {
	var (
		command *LcategoryCommandRepositoryImpl
		query   *LcategoryQueryRepositoryImpl
	)

	BeforeEach(func() {
		command = &LcategoryCommandRepositoryImpl{DB: db}
		query = &LcategoryQueryRepositoryImpl{DB: db}

		truncate(db)
		Expect(db.Save(newUser(userID)).Error).To(Succeed())
	})

	base := newLcategory(lcategoryID)
	baseChanged := newLcategory(lcategoryID)
	baseChanged.Name = "changed"

	DescribeTable("Saveは引数のlcategoryを作成するか、その状態になるように更新する",
		func(before *entity.Lcategory, saved *entity.Lcategory) {
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
})

func newLcategory(id int) *entity.Lcategory {
	lcategory := entity.Lcategory{ID: id}
	util.FillDymmyString(&lcategory, id)

	return &lcategory
}
