package repository

import (
	"net/http"
	"time"

	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/entity/wordpress"
	"github.com/stayway-corp/stayway-media-api/pkg/mock"
	"github.com/stayway-corp/stayway-media-api/pkg/util"
)

var _ = Describe("WordpressRepositoryImpl", func() {
	var (
		query    *WordpressQueryRepositoryImpl
		mockCtrl *gomock.Controller
		httpMock *mock.MockRoundTripper
	)

	BeforeEach(func() {
		query = tests.WordpressQueryRepositoryImpl

		mockCtrl = gomock.NewController(GinkgoT())
		httpMock = mock.NewMockRoundTripper(mockCtrl)
		query.Client.Transport = httpMock
	})

	AfterEach(func() {
		mockCtrl.Finish()
	})

	Describe("FindPostsByIDs", func() {
		var (
			postIDs = []int{135485, 134737, 135005}
			posts   []*wordpress.Post
			err     error
		)

		JustBeforeEach(func() {
			posts, err = query.FindPostsByIDs(postIDs)
		})

		Describe("正常系", func() {
			BeforeEach(func() {
				httpMock.EXPECT().RoundTrip(gomock.Any()).DoAndReturn(func(req *http.Request) (*http.Response, error) {
					params := req.URL.Query()
					Expect(params).To(HaveKeyWithValue("include", []string{"135485,134737,135005"}))
					return responseTestdata("posts.json")
				})
			})

			It("指定したデータがドメインが置換された状態で正常に取得できる", func() {
				Expect(posts).To(HaveLen(len(postIDs)))
				Expect(posts[0]).To(Equal(&wordpress.Post{
					ID:      135485,
					Date:    wordpress.Time(time.Date(2020, 3, 5, 21, 57, 6, 0, util.JSTLoc)),
					DateGmt: wordpress.Time(time.Date(2020, 3, 5, 12, 57, 6, 0, util.JSTLoc)),
					GUID: wordpress.Text{
						Rendered: "https://stg.stayway.jp/tourism/?p=135485",
					},
					Modified:    wordpress.Time(time.Date(2020, 3, 5, 22, 0, 56, 0, util.JSTLoc)),
					ModifiedGmt: wordpress.Time(time.Date(2020, 3, 5, 13, 0, 56, 0, util.JSTLoc)),
					Slug:        "%e3%83%86%e3%82%b9%e3%83%88",
					Status:      wordpress.StatusPublish,
					Type:        "post",
					Title: wordpress.Text{
						Rendered: "テスト",
					},
					Content: wordpress.ProtectableText{
						Rendered:  "<p><strong>本文</strong></p>\nhttps://stg.stayway.jp/tourism/posts/135485\n<img src='https://stg-files.stayway.jp/wp-content/uploads/2020/07/dummy-hoge.jpg'>",
						Protected: false,
					},
					Author:        260,
					FeaturedMedia: 135463,
					Categories:    []int{1},
					Tags:          []int{},
				}))
			})

			It("errorはnil", func() {
				Expect(err).To(Succeed())
			})
		})
	})
})
