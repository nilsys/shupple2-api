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

var _ = Describe("ThemeCategoryRepositoryImpl", func() {
	var (
		command *ThemeCategoryCommandRepositoryImpl
		query   *ThemeCategoryQueryRepositoryImpl
	)

	BeforeEach(func() {
		command = tests.ThemeCategoryCommandRepositoryImpl
		query = tests.ThemeCategoryQueryRepositoryImpl

		truncate(tests.DB)
		Expect(db.Save(newUser(userID)).Error).To(Succeed())
	})

	base := newThemeCategory(themeCategoryID)
	baseChanged := newThemeCategory(themeCategoryID)
	baseChanged.Name = "changed"

	DescribeTable("Saveは引数のthemeCategoryを作成するか、その状態になるように更新する",
		func(before *entity.ThemeCategory, saved *entity.ThemeCategory) {
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

func newThemeCategory(id int) *entity.ThemeCategory {
	themeCategory := &entity.ThemeCategory{}
	themeCategory.ID = id
	themeCategory.ThemeID = id
	themeCategory.Type = model.ThemeCategoryTypeTheme
	util.FillDummyString(themeCategory, id)
	return themeCategory
}
