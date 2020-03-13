package repository

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"
	"github.com/stayway-corp/stayway-media-api/pkg/adaptor/api/converter"
	"github.com/stayway-corp/stayway-media-api/pkg/adaptor/api/param"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/entity"
	"github.com/stayway-corp/stayway-media-api/pkg/util"
)

var _ = Describe("ReviewRepositoryTest", func() {
	var (
		query   *ReviewQueryRepositoryImpl
		hashtag *entity.Hashtag = newHashtag(hashtagID)
	)

	BeforeEach(func() {
		query = &ReviewQueryRepositoryImpl{DB: db}

		truncate(db)
		Expect(db.Save(newTouristSpot(touristSpotID, nil, nil)))
		Expect(db.Save(newUser(userID)).Error).To(Succeed())
		Expect(db.Save(newReview(reviewID, userID, touristSpotID, innID)).Error).To(Succeed())
		Expect(db.Save(hashtag).Error).To(Succeed())
		Expect(db.Exec("INSERT INTO review_hashtag(review_id,hashtag_id) VALUES (?,?)", reviewID, hashtagID).Error).To(Succeed())
	})

	DescribeTable("ShowReviewListByParams",
		func(param *param.ListReviewParams) {
			queryStruct := converter.ConvertFindReviewListParamToQuery(param)
			actual, err := query.ShowReviewListByParams(queryStruct)
			Expect(err).To(Succeed())

			Expect(actual).To(Equal([]*entity.Review{newReview(reviewID, userID, touristSpotID, innID)}))
		},
		Entry("正常系_全条件検索", newShowReviewListParam(userID, innID, touristSpotID, hashtag.Name)),
		Entry("正常系_UserID検索", newShowReviewListParam(userID, 0, 0, "")),
		Entry("正常系_InnID検索", newShowReviewListParam(0, innID, 0, "")),
		Entry("正常系_SpotID検索", newShowReviewListParam(0, 0, touristSpotID, "")),
		Entry("正常系_HashTag検索", newShowReviewListParam(0, 0, 0, hashtag.Name)),
	)
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
		Medias:        []*entity.ReviewMedia{},
	}
	// ここで全てのパラメータにダミーデータが挿入される
	util.FillDymmyString(review, id)
	return review
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
