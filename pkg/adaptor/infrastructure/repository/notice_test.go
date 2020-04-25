package repository

import (
	"context"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/entity"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/model"
)

var _ = Describe("ReviewRepositoryTest", func() {
	var (
		command *NoticeCommandRepositoryImpl
		query   *NoticeQueryRepositoryImpl
	)

	Describe("StoreNoticeのテスト", func() {
		targetUserID := userID
		triggeredUserID := userID + 1
		actionTargetID := 1

		BeforeEach(func() {
			command = &NoticeCommandRepositoryImpl{DAO: DAO{UnderlyingDB: db}}

			truncate(db)
			Expect(db.Save(newTouristSpot(touristSpotID, nil, nil, nil)))

			Expect(db.Save(newUser(targetUserID)).Error).To(Succeed())
			Expect(db.Save(newUser(triggeredUserID)).Error).To(Succeed())
			Expect(db.Save(newReview(reviewID, targetUserID, touristSpotID, innID)).Error).To(Succeed())
			Expect(db.Save(newReviewComment(triggeredUserID, reviewID)))
		})

		DescribeTable("正常系",
			func() {
				notice := entity.NewNotice(
					triggeredUserID,
					targetUserID,
					model.NoticeActionTypeCOMMENT,
					model.NoticeActionTargetTypeREVIEW,
					actionTargetID,
				)

				err := command.StoreNotice(context.TODO(), notice)
				Expect(err).To(Succeed())

				actual := &entity.Notice{}
				err = db.
					Where("id = ?", notice.ID).
					Find(actual).
					Error
				Expect(err).To(Succeed())
				Expect(actual.UserID).To(Equal(triggeredUserID))
				Expect(actual.TriggeredUserID).To(Equal(targetUserID))
				Expect(actual.ActionTargetType).To(Equal(model.NoticeActionTargetTypeREVIEW))
				Expect(actual.ActionType).To(Equal(model.NoticeActionTypeCOMMENT))
				Expect(actual.ActionTargetID).To(Equal(actionTargetID))
			},
			Entry("正常系"),
		)
	})

	Describe("MarkedAsReadのテスト", func() {
		targetUserID := userID
		triggeredUserID := userID + 1

		BeforeEach(func() {
			command = &NoticeCommandRepositoryImpl{DAO: DAO{UnderlyingDB: db}}

			truncate(db)
			Expect(db.Save(newTouristSpot(touristSpotID, nil, nil, nil)))

			Expect(db.Save(newUser(targetUserID)).Error).To(Succeed())
			Expect(db.Save(newUser(triggeredUserID)).Error).To(Succeed())
			Expect(db.Save(newReview(reviewID, targetUserID, touristSpotID, innID)).Error).To(Succeed())
			Expect(db.Save(newReviewComment(triggeredUserID, reviewID)))
			Expect(db.Save(newNotice(targetUserID, triggeredUserID, reviewID)))
			Expect(db.Save(newNotice(targetUserID, triggeredUserID, reviewID)))
			Expect(db.Save(newNotice(targetUserID, triggeredUserID, reviewID)))
			Expect(db.Save(newNotice(targetUserID, triggeredUserID, reviewID)))
		})

		DescribeTable("正常系",
			func() {
				var ids []int
				err := db.Table("notice").Select("id").Find(&ids).Error
				Expect(err).To(Succeed())

				err = command.MarkAsRead(ids)
				Expect(err).To(Succeed())

				var actual int
				err = db.
					Table("notice").
					Where("is_read = ?", true).
					Count(&actual).
					Error

				// 全部既読になっているか
				Expect(err).To(Succeed())
				Expect(actual).To(Equal(0))
			},
			Entry("正常系"),
		)
	})

	Describe("ListNoticeのテスト", func() {
		targetUserID := userID
		triggeredUserID := userID + 1

		BeforeEach(func() {
			query = &NoticeQueryRepositoryImpl{DB: db}

			truncate(db)
			Expect(db.Save(newTouristSpot(touristSpotID, nil, nil, nil)))

			Expect(db.Save(newUser(targetUserID)).Error).To(Succeed())
			Expect(db.Save(newUser(triggeredUserID)).Error).To(Succeed())
			Expect(db.Save(newReview(reviewID, targetUserID, touristSpotID, innID)).Error).To(Succeed())
			Expect(db.Save(newReviewComment(triggeredUserID, reviewID)))
			Expect(db.Save(newNotice(targetUserID, triggeredUserID, reviewID)))
			Expect(db.Save(newNotice(targetUserID, triggeredUserID, reviewID)))
			Expect(db.Save(newNotice(targetUserID, triggeredUserID, reviewID)))
			Expect(db.Save(newNotice(targetUserID, triggeredUserID, reviewID)))
		})

		DescribeTable("正常系",
			func() {
				actual, err := query.ListNotice(targetUserID, 2)
				Expect(err).To(Succeed())

				Expect(len(actual)).To(Equal(2))
			},
			Entry("正常系"),
		)
	})
})

func newNotice(userID, triggeredUser int, actionTargetID int) *entity.Notice {
	return entity.NewNotice(userID, triggeredUser, model.NoticeActionTypeTAGGED, model.NoticeActionTargetTypeREVIEW, actionTargetID)
}
