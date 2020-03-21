package repository

import (
	"context"
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"
	"github.com/stayway-corp/stayway-media-api/pkg/adaptor/api/converter"
	"github.com/stayway-corp/stayway-media-api/pkg/adaptor/api/param"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/entity"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/model"
	"github.com/stayway-corp/stayway-media-api/pkg/util"
)

var _ = Describe("ReviewRepositoryTest", func() {
	var (
		query   *ReviewQueryRepositoryImpl
		command *ReviewCommandRepositoryImpl
		hashtag *entity.Hashtag = newHashtag(hashtagID)
	)

	Describe("ShowReviewListByParamsのテスト", func() {
		BeforeEach(func() {
			query = &ReviewQueryRepositoryImpl{DB: db}

			truncate(db)
			Expect(db.Save(hashtag).Error).To(Succeed())
			Expect(db.Save(newTouristSpot(touristSpotID, nil, nil)))
			Expect(db.Save(newUser(userID)).Error).To(Succeed())
			Expect(db.Save(newReview(reviewID, userID, touristSpotID, innID)).Error).To(Succeed())
			Expect(db.Exec("INSERT INTO review_hashtag(review_id,hashtag_id) VALUES (?,?)", reviewID, hashtag.ID).Error).To(Succeed())
		})

		DescribeTable("ShowReviewListByParams",
			func(param *param.ListReviewParams) {
				queryStruct := converter.ConvertFindReviewListParamToQuery(param)
				actual, err := query.ShowReviewListByParams(queryStruct)
				Expect(err).To(Succeed())

				for _, result := range actual {
					Expect(result.CreatedAt).NotTo(BeZero())
					Expect(result.UpdatedAt).NotTo(BeZero())
					Expect(result.User.CreatedAt).NotTo(BeZero())
					Expect(result.User.UpdatedAt).NotTo(BeZero())
					result.CreatedAt = time.Time{}
					result.UpdatedAt = time.Time{}
					result.User.CreatedAt = time.Time{}
					result.User.UpdatedAt = time.Time{}
				}
				Expect(actual).To(Equal([]*entity.QueryReview{newQueryReview(hashtag.Name, hashtag.ID)}))
			},
			Entry("正常系_全条件検索", newShowReviewListParam(userID, innID, touristSpotID, hashtag.Name)),
			Entry("正常系_UserID検索", newShowReviewListParam(userID, 0, 0, "")),
			Entry("正常系_InnID検索", newShowReviewListParam(0, innID, 0, "")),
			Entry("正常系_SpotID検索", newShowReviewListParam(0, 0, touristSpotID, "")),
			Entry("正常系_HashTag検索", newShowReviewListParam(0, 0, 0, hashtag.Name)),
		)
	})

	Describe("FindReviewCommentListByReviewID", func() {
		BeforeEach(func() {
			query = &ReviewQueryRepositoryImpl{DB: db}

			truncate(db)
			Expect(db.Save(newTouristSpot(touristSpotID, nil, nil)))

			Expect(db.Save(newUser(1)).Error).To(Succeed())
			Expect(db.Save(newUser(2)).Error).To(Succeed())
			// コメントの存在する投稿
			Expect(db.Save(newReview(1, 1, touristSpotID, innID)).Error).To(Succeed())
			// コメントの存在しない投稿
			Expect(db.Save(newReview(2, 2, touristSpotID, innID)).Error).To(Succeed())

			// 投稿日の違う2件の投稿
			Expect(db.Exec("INSERT INTO review_comment(id, user_id, review_id, body, created_at, updated_at) VALUES (1, 1, 1, 'dummy 1', '2020-01-01 10:10:10', '2020-01-01 10:10:10');").Error).To(Succeed())
			Expect(db.Exec("INSERT INTO review_comment(id, user_id, review_id, body, created_at, updated_at) VALUES (2, 1, 1, 'dummy 2', '2020-02-01 10:10:10', '2020-02-01 10:10:10');").Error).To(Succeed())

		})

		DescribeTable("コメントの存在する投稿の場合",
			func(id int) {
				actual, err := query.FindReviewCommentListByReviewID(id, 2)
				Expect(err).To(Succeed())

				// 新しい順になっている
				Expect(actual[0].ID).To(Equal(2))
				Expect(actual[1].ID).To(Equal(1))

				// 内容が正しいか
				Expect(actual[0].Body).To(Equal("dummy 2"))
				Expect(actual[1].Body).To(Equal("dummy 1"))
				Expect(actual[0].User.ID).To(Equal(newUser(1).ID))
				Expect(actual[1].User.ID).To(Equal(newUser(1).ID))
			},
			Entry("正常系_全件取得", 1),
		)

		DescribeTable("コメントの存在する投稿の場合",
			func(id int) {
				actual, err := query.FindReviewCommentListByReviewID(id, 1)
				Expect(err).To(Succeed())

				// 取得は一件だけ
				Expect(len(actual)).To(Equal(1))
			},
			Entry("正常系_リミット指定", 1),
		)

		DescribeTable("コメントの存在しない投稿の場合",
			func(id int) {
				actual, err := query.FindReviewCommentListByReviewID(id, 2)
				Expect(err).To(Succeed())

				// コメントが0件であること
				Expect(len(actual)).To(Equal(0))
			},
			Entry("正常系_0件取得", 2),
		)
	})

	Describe("CreateReviewCommentのテスト",
		func() {
			BeforeEach(func() {
				command = &ReviewCommandRepositoryImpl{DAO: DAO{DB_: db}}

				truncate(db)
				Expect(db.Save(newTouristSpot(touristSpotID, nil, nil)))
				Expect(db.Save(newUser(1)).Error).To(Succeed())
				Expect(db.Save(newReview(1, 1, touristSpotID, innID)).Error).To(Succeed())
			})

			DescribeTable("コメントを新規追加",
				func() {
					reviewComment := entity.NewReviewComment(1, 1, "dummy body")
					err := command.CreateReviewComment(context.TODO(), reviewComment)
					Expect(err).To(Succeed())

					// コメントが一件増えているか
					actualCount := 0
					db.Table("review_comment").Count(&actualCount)
					Expect(actualCount).To(Equal(1))
				},
				Entry("正常系"),
			)
		})

	Describe("IncrementReviewCommentCountのテスト",
		func() {
			BeforeEach(func() {
				command = &ReviewCommandRepositoryImpl{DAO: DAO{DB_: db}}

				truncate(db)
				Expect(db.Save(newTouristSpot(touristSpotID, nil, nil)))

				Expect(db.Save(newUser(1)).Error).To(Succeed())
				Expect(db.Save(newReview(1, 1, touristSpotID, innID)).Error).To(Succeed())
			})

			DescribeTable("コメントを追加",
				func() {
					err := command.IncrementReviewCommentCount(context.TODO(), 1)
					Expect(err).To(Succeed())

					// コメントが一件増えているか
					review := &entity.Review{}
					err = db.
						Where("id=?", 1).
						Find(review).
						Error
					Expect(err).To(Succeed())
					Expect(review.CommentCount).To(Equal(1))
				},
				Entry("正常系"),
			)
		})

	Describe("DecrementReviewCommentCountのテスト",
		func() {
			BeforeEach(func() {
				command = &ReviewCommandRepositoryImpl{DAO: DAO{DB_: db}}

				truncate(db)
				Expect(db.Save(newTouristSpot(touristSpotID, nil, nil)))

				Expect(db.Save(newUser(1)).Error).To(Succeed())
				Expect(db.Save(newReview(1, 1, touristSpotID, innID)).Error).To(Succeed())

				// コメント数を1にしておく
				err := command.IncrementReviewCommentCount(context.TODO(), 1)
				Expect(err).To(Succeed())
			})

			DescribeTable("コメントを追加",
				func() {
					err := command.DecrementReviewCommentCount(context.TODO(), 1)
					Expect(err).To(Succeed())

					// コメントが一件減っているか
					review := &entity.Review{}
					err = db.
						Where("id=?", 1).
						Find(review).
						Error
					Expect(err).To(Succeed())
					Expect(review.CommentCount).To(Equal(0))
				},
				Entry("正常系"),
			)
		})

	Describe("ShowReviewCommentのテスト", func() {
		BeforeEach(func() {
			command = &ReviewCommandRepositoryImpl{DAO: DAO{DB_: db}}

			truncate(db)
			Expect(db.Save(newTouristSpot(touristSpotID, nil, nil)))
			Expect(db.Save(newUser(userID)).Error).To(Succeed())
			Expect(db.Save(newReview(reviewID, userID, touristSpotID, innID)).Error).To(Succeed())

			Expect(db.Exec("INSERT INTO review_comment(id, user_id, review_id, body, created_at, updated_at) VALUES (1, ?, ?, 'dummy 1', '2020-01-01 10:10:10', '2020-01-01 10:10:10');", userID, reviewID).Error).To(Succeed())
		})

		DescribeTable("コメントを取得_正常系",
			func() {
				comment, err := command.ShowReviewComment(context.TODO(), 1)
				Expect(err).To(Succeed())

				Expect(comment.ID).To(Equal(1))
			},
			Entry("正常系"),
		)

		DescribeTable("コメントを取得_異常系_コメントが存在しない場合",
			func() {
				_, err := command.ShowReviewComment(context.TODO(), 2)

				// エラーがnilじゃない
				Expect(err).NotTo(Succeed())
			},
			Entry("異常系"),
		)
	})

	Describe("DeleteReviewCommentのテスト", func() {
		BeforeEach(func() {
			command = &ReviewCommandRepositoryImpl{DAO: DAO{DB_: db}}

			truncate(db)
			Expect(db.Save(newTouristSpot(touristSpotID, nil, nil)))
			Expect(db.Save(newUser(userID)).Error).To(Succeed())
			Expect(db.Save(newReview(reviewID, userID, touristSpotID, innID)).Error).To(Succeed())
			Expect(db.Exec("INSERT INTO review_comment(id, user_id, review_id, body, created_at, updated_at) VALUES (1, ?, ?, 'dummy 1', '2020-01-01 10:10:10', '2020-01-01 10:10:10');", userID, reviewID).Error).To(Succeed())
		})

		DescribeTable("コメントを削除_正常系",
			func() {
				comment, err := command.ShowReviewComment(context.TODO(), 1)
				Expect(err).To(Succeed())
				err = command.DeleteReviewCommentByID(context.TODO(), comment)
				Expect(err).To(Succeed())

				// 削除されているか
				var count int
				db.Table("review_comment").Where("deleted_at IS NULL").Count(&count)
				Expect(count).To(Equal(0))
			},
			Entry("正常系"),
		)
	})

})

func newReview(id, userID, touristSpotID, innID int) *entity.Review {
	// 全てのパラメータは仮置き
	review := &entity.Review{
		ID:            id,
		UserID:        userID,
		TouristSpotID: touristSpotID,
		InnID:         innID,
		Score:         id,
		MediaCount:    id,
		Body:          "dummy",
		FavoriteCount: id,
		TravelDate:    time.Date(2020, 7, 7, 0, 0, 0, 0, time.Local),
		Accompanying:  model.AccompanyingTypeBUISINESS,
		Medias:        []*entity.ReviewMedia{},
		HashtagIDs:    []*entity.ReviewHashtag{},
	}
	// ここで全てのパラメータにダミーデータが挿入される
	util.FillDymmyString(review, id)
	return review
}

func newQueryReview(hashtagName string, hashtagID int) *entity.QueryReview {
	review := newReview(reviewID, userID, touristSpotID, innID)
	queryReview := &entity.QueryReview{
		Review: entity.Review{
			ID:            review.ID,
			UserID:        review.UserID,
			TouristSpotID: review.TouristSpotID,
			InnID:         review.InnID,
			Score:         review.Score,
			MediaCount:    review.MediaCount,
			Body:          review.Body,
			FavoriteCount: review.FavoriteCount,
			TravelDate:    review.TravelDate,
			Accompanying:  review.Accompanying,
			Medias:        review.Medias,
			HashtagIDs:    []*entity.ReviewHashtag{{ReviewID: reviewID, HashtagID: hashtagID}},
		},
		User:    newUser(userID),
		Hashtag: []*entity.Hashtag{&entity.Hashtag{ID: hashtagID, Name: hashtagName}},
	}
	return queryReview
}

func newShowReviewListParam(userID, innID, touristSpotID int, hashtag string) *param.ListReviewParams {
	return &param.ListReviewParams{
		UserID:        userID,
		InnID:         innID,
		TouristSpotID: touristSpotID,
		HashTag:       hashtag,
		PerPage:       mockReviewPerPage,
		Page:          mockReviewPage,
	}
}
