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

var _ = Describe("FeatureRepositoryImpl", func() {
	var (
		command *FeatureCommandRepositoryImpl
		query   *FeatureQueryRepositoryImpl
	)

	BeforeEach(func() {
		command = tests.FeatureCommandRepositoryImpl
		query = tests.FeatureQueryRepositoryImpl

		truncate(db)
		Expect(db.Save(newUser(userID)).Error).To(Succeed())
		for _, post := range append(postIDs, addedPostID) {
			Expect(db.Save(newPost(post, nil, nil, nil, nil)).Error).To(Succeed())
		}
	})

	base := newFeature(featureID, postIDs)
	baseChanged := newFeature(featureID, postIDs)
	baseChanged.Title = "changed"

	DescribeTable("Saveは引数のfeatureを作成するか、その状態になるように更新する",
		func(before *entity.Feature, saved *entity.Feature) {
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
		Entry("postが追加される場合", base, newFeature(featureID, append(postIDs, addedPostID))),
		Entry("postが削除される場合", base, newFeature(featureID, postIDs[:1])),
	)
})

func newFeature(id int, postIDs []int) *entity.Feature {
	feature := entity.FeatureTiny{
		ID:       id,
		UserID:   userID,
		EditedAt: time.Now().Truncate(time.Second),
	}
	util.FillDummyString(&feature, id)

	f := entity.NewFeature(feature, postIDs)
	return &f
}
