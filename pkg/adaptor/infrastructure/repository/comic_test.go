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
	baseQuery := newQueryComic(comicID)
	baseQueryChanged := newQueryComic(comicID)
	baseQueryChanged.Title = "changed"

	DescribeTable("Saveは引数のcomicを作成するか、その状態になるように更新する",
		func(before *entity.Comic, saved *entity.Comic, querySaved *entity.QueryComic) {
			if before != nil {
				Expect(command.Store(context.Background(), before)).To(Succeed())
			}

			Expect(command.Store(context.Background(), saved)).To(Succeed())
			actual, err := query.FindByID(saved.ID)
			Expect(err).To(Succeed())

			Expect(actual.User.CreatedAt).NotTo(BeZero())
			Expect(actual.User.UpdatedAt).NotTo(BeZero())
			actual.User.CreatedAt = time.Time{}
			actual.User.UpdatedAt = time.Time{}
			Expect(actual).To(Equal(querySaved))
		},
		Entry("新規作成", nil, &base, baseQuery),
		Entry("フィールドに変更がある場合", &base, &baseChanged, baseQueryChanged),
	)
})

func newComic(id int) entity.Comic {
	comic := entity.Comic{
		ID:        id,
		UserID:    userID,
		CreatedAt: sampleTime,
		UpdatedAt: sampleTime,
	}
	util.FillDummyString(&comic, id)

	return comic
}

func newQueryComic(id int) *entity.QueryComic {
	return &entity.QueryComic{
		Comic: newComic(id),
		User:  newUser(userID),
	}
}
