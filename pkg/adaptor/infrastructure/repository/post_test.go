package repository

import (
	"context"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/entity"
	"github.com/stayway-corp/stayway-media-api/pkg/util"
)

var _ = Describe("PostRepositoryImpl", func() {
	var (
		command *PostCommandRepositoryImpl
		query   *PostQueryRepositoryImpl
	)

	BeforeEach(func() {
		command = tests.PostCommandRepositoryImpl
		query = tests.PostQueryRepositoryImpl

		truncate(db)
		Expect(db.Save(newUser(userID)).Error).To(Succeed())
		for _, cat := range append(categoryIDs, addedCategoryID) {
			Expect(db.Save(newCategory(cat)).Error).To(Succeed())
		}

		for _, hashtag := range append(hashtagIDs, addedHashtagID) {
			Expect(db.Save(newHashtag(hashtag)).Error).To(Succeed())
		}
	})

	base := newPost(postID, bodies, categoryIDs, hashtagIDs)
	baseChanged := newPost(postID, bodies, categoryIDs, hashtagIDs)
	baseChanged.Title = "changed"

	DescribeTable("Storeは引数のpostを作成するか、その状態になるように更新する",
		func(before *entity.Post, saved *entity.Post) {
			if before != nil {
				Expect(command.Store(context.TODO(), before)).To(Succeed())
			}

			Expect(command.Store(context.TODO(), saved)).To(Succeed())
			actual, err := query.FindByID(saved.ID)
			Expect(err).To(Succeed())

			Expect(actual).To(Equal(saved))
		},
		Entry("新規作成", nil, base),
		Entry("フィールドに変更がある場合", base, baseChanged),
		Entry("Bodyが追加される場合", base, newPost(postID, append(bodies, addedBody), categoryIDs, hashtagIDs)),
		Entry("categoryIDが追加される場合", base, newPost(postID, bodies, append(categoryIDs, addedCategoryID), hashtagIDs)),
		Entry("hashtagIDが追加される場合", base, newPost(postID, bodies, categoryIDs, append(hashtagIDs, addedHashtagID))),
		Entry("Bodyが削除される場合", base, newPost(postID, bodies[:1], categoryIDs, hashtagIDs)),
		Entry("categoryIDが削除される場合", base, newPost(postID, bodies, categoryIDs[:1], hashtagIDs)),
		Entry("hashtagIDが削除される場合", base, newPost(postID, bodies, categoryIDs, hashtagIDs[:1])),
	)
})

func newPost(id int, bodies []string, categoryIDs []int, hashtagIDs []int) *entity.Post {
	post := entity.PostTiny{
		ID:            id,
		UserID:        userID,
		FavoriteCount: id,
		FacebookCount: id,
		CreatedAt:     sampleTime,
		UpdatedAt:     sampleTime,
	}
	util.FillDymmyString(&post, id)
	p := entity.NewPost(post, bodies, categoryIDs, hashtagIDs)
	return &p
}
