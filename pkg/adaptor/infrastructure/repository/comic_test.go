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

var _ = Describe("ComicRepositoryImpl", func() {
	var (
		command *ComicCommandRepositoryImpl
		query   *ComicQueryRepositoryImpl
	)

	BeforeEach(func() {
		command = tests.ComicCommandRepositoryImpl
		query = tests.ComicQueryRepositoryImpl

		truncate(tests.DB)
		Expect(db.Save(newUser(userID)).Error).To(Succeed())
	})

	base := newComic(comicID)
	baseChanged := newComic(comicID)
	baseChanged.Title = "changed"

	DescribeTable("Storeは引数のcomicを作成するか、その状態になるように更新する",
		func(before *entity.Comic, saved *entity.Comic) {
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
			Expect(&actual.Comic).To(Equal(saved))
		},
		Entry("新規作成", nil, &base),
		Entry("フィールドに変更がある場合", &base, &baseChanged),
	)
})

func newComic(id int) entity.Comic {
	comic := entity.Comic{
		ID:       id,
		UserID:   userID,
		EditedAt: time.Now().Truncate(time.Second),
	}
	util.FillDummyString(&comic, id)

	return comic
}
