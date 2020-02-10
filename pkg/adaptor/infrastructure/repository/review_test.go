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
	var query *ReviewQueryRepositoryImpl

	BeforeEach(func() {
		query = &ReviewQueryRepositoryImpl{DB: db}

		truncate(db)
		Expect(db.Exec("INSERT INTO tourist_spot(id,name,slug,city,address,lat,lng,access_car,access_train,opening_hours,price) VALUES (1,'dummy','dummy','dummy','dummy',39.09,391.00,'dummy','dummy','dummy','dummy');").Error).To(Succeed())
		Expect(db.Save(newUser(1)).Error).To(Succeed())
		Expect(db.Save(newReview(1)).Error).To(Succeed())
		Expect(db.Exec("INSERT INTO hashtag(id,name) VALUES (1,'dummy');").Error).To(Succeed())
		Expect(db.Exec("INSERT INTO review_hashtag(review_id,hashtag_id) VALUES (1,1);").Error).To(Succeed())
	})

	DescribeTable("ShowReviewListByParams",
		func(param *param.ListReviewParams) {
			queryStruct := converter.ConvertFindReviewListParamToQuery(param)
			actual, err := query.ShowReviewListByParams(queryStruct)
			Expect(err).To(Succeed())

			Expect(actual).To(Equal([]*entity.Review{newReview(1)}))
		},
		Entry("正常系_全条件検索", newShowReviewListParam(mockReviewUserID, mockReviewInnID, mockReviewTouristSpotID, mockReviewHashTag)),
		Entry("正常系_UserID検索", newShowReviewListParam(mockReviewUserID, 0, 0, "")),
		Entry("正常系_InnID検索", newShowReviewListParam(0, mockReviewInnID, 0, "")),
		Entry("正常系_SpotID検索", newShowReviewListParam(0, 0, mockReviewTouristSpotID, "")),
		Entry("正常系_HashTag検索", newShowReviewListParam(0, 0, 0, mockReviewHashTag)),
	)
})

func newReview(id int) *entity.Review {
	// 全てのパラメータは仮置き
	review := &entity.Review{
		ID:            id,
		UserID:        id,
		TouristSpotID: id,
		InnID:         id,
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

func newShowReviewListParam(userID, innID, touristSpotID int, hashTag string) *param.ListReviewParams {
	return &param.ListReviewParams{
		UserID:        userID,
		InnID:         innID,
		TouristSpotID: touristSpotID,
		HashTag:       hashTag,
		PerPage:       mockReviewPerPage,
		Page:          mockReviewPage,
	}
}
