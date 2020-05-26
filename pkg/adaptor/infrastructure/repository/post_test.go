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

		for _, cat := range append(areaCategoryIDs, addedAreaCategoryID) {
			Expect(db.Save(newAreaCategory(cat)).Error).To(Succeed())
		}

		for _, cat := range append(themeCategoryIDs, addedThemeCategoryID) {
			Expect(db.Save(newThemeCategory(cat)).Error).To(Succeed())
		}

		for _, hashtag := range append(hashtagIDs, addedHashtagID) {
			Expect(db.Save(newHashtag(hashtag)).Error).To(Succeed())
		}
	})

	base := newPost(postID, bodies, areaCategoryIDs, themeCategoryIDs, hashtagIDs)
	baseChanged := newPost(postID, bodies, areaCategoryIDs, themeCategoryIDs, hashtagIDs)
	baseChanged.Title = "changed"

	DescribeTable("Storeは引数のpostを作成するか、その状態になるように更新する",
		func(before *entity.Post, saved *entity.Post) {
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
		Entry("Bodyが追加される場合", base, newPost(postID, append(bodies, addedBody), areaCategoryIDs, themeCategoryIDs, hashtagIDs)),
		Entry("themeCategoryIDが追加される場合", base, newPost(postID, bodies, append(areaCategoryIDs, addedAreaCategoryID), themeCategoryIDs, hashtagIDs)),
		Entry("areaCategoryIDが追加される場合", base, newPost(postID, bodies, areaCategoryIDs, append(themeCategoryIDs, addedThemeCategoryID), hashtagIDs)),
		Entry("hashtagIDが追加される場合", base, newPost(postID, bodies, areaCategoryIDs, themeCategoryIDs, append(hashtagIDs, addedHashtagID))),
		Entry("Bodyが削除される場合", base, newPost(postID, bodies[:1], areaCategoryIDs, themeCategoryIDs, hashtagIDs)),
		Entry("areaCategoryIDが削除される場合", base, newPost(postID, bodies, areaCategoryIDs[:1], themeCategoryIDs, hashtagIDs)),
		Entry("themeCategoryIDが削除される場合", base, newPost(postID, bodies, areaCategoryIDs, themeCategoryIDs[1:], hashtagIDs)),
		Entry("hashtagIDが削除される場合", base, newPost(postID, bodies, areaCategoryIDs, themeCategoryIDs, hashtagIDs[:1])),
	)
})

func newPost(id int, bodies []string, areaCategoryIDs []int, themeCategoryIDs []int, hashtagIDs []int) *entity.Post {
	post := entity.PostTiny{
		ID:            id,
		UserID:        userID,
		FavoriteCount: id,
		FacebookCount: id,
	}
	util.FillDummyString(&post, id)
	p := entity.NewPost(post, bodies, areaCategoryIDs, themeCategoryIDs, hashtagIDs)
	return &p
}
