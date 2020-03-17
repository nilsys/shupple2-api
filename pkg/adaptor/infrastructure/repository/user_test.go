package repository

import (
	"fmt"
	"time"

	"github.com/stayway-corp/stayway-media-api/pkg/domain/model"

	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3iface"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"
	"github.com/pkg/errors"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/entity"
	"github.com/stayway-corp/stayway-media-api/pkg/util"
)

var _ = Describe("UserRepositoryImpl", func() {
	var (
		command *UserCommandRepositoryImpl
		query   *UserQueryRepositoryImpl
	)

	BeforeEach(func() {
		command = tests.UserCommandRepositoryImpl
		query = tests.UserQueryRepositoryImpl

		truncate(db)

		Expect(prepareBucket(tests.AWS, tests.Config.AWS.FilesBucket)).To(Succeed())
	})

	DescribeTable("Storeは引数のuserを作成するか、その状態になるように更新する",
		func(before *entity.User, saved *entity.User) {
			if before != nil {
				Expect(command.Store(before)).To(Succeed())
			}

			Expect(command.Store(saved)).To(Succeed())
			actual, err := query.FindByID(saved.ID)
			Expect(err).To(Succeed())

			Expect(actual.CreatedAt).NotTo(BeZero())
			Expect(actual.UpdatedAt).NotTo(BeZero())
			actual.CreatedAt = time.Time{}
			actual.UpdatedAt = time.Time{}
			Expect(actual).To(Equal(saved))
		},
		Entry("新規作成", nil, newUser(0)),
		Entry("フィールドに変更がある場合", newUser(userID), newUser(userID, "changed")),
	)

	Describe("StoreWithAvatar", func() {
		var (
			base       *entity.User
			baseAvatar []byte
			err        error
		)

		BeforeEach(func() {
			base = newUser(0)
			baseAvatar = []byte{0x01, 0x02}
		})

		JustBeforeEach(func() {
			err = command.StoreWithAvatar(base, baseAvatar)
		})

		Describe("avatarの保存をした上でStoreと同様の動作をする", func() {
			checker := func() {
				It("Error無し", func() {
					Expect(err).To(Succeed())
				})
				It("Userが正常に保存される", func() {
					user, err := query.FindByID(base.ID)
					Expect(err).To(Succeed())

					Expect(user.CreatedAt).NotTo(BeZero())
					Expect(user.UpdatedAt).NotTo(BeZero())
					user.CreatedAt = time.Time{}
					user.UpdatedAt = time.Time{}
					Expect(user).To(Equal(base))
				})
				It("Avatarが正常に保存される", func() {
					size, err := getS3ObjectSize(fmt.Sprintf(avatarKeyFormat, base.ID))
					Expect(err).To(Succeed())
					Expect(size).To(Equal(len(baseAvatar)))
				})
			}

			Context("新規作成", func() {
				checker()
			})

			Context("更新", func() {
				BeforeEach(func() {
					existing := *base
					existing.Name = "changed"
					Expect(command.StoreWithAvatar(&existing, baseAvatar[:1])).To(Succeed())
					base.ID = existing.ID
				})

				checker()
			})
		})

		Context("avatarの保存に失敗した場合", func() {
			errUnexpected := errors.New("failed to save avatar")
			BeforeEach(func() {
				tests.Uploader.S3 = s3mock{errUnexpected, tests.Uploader.S3}
			})
			AfterEach(func() {
				tests.Uploader.S3 = tests.Uploader.S3.(s3mock).S3API
			})

			It("Errorが返る", func() {
				Expect(errors.Cause(err)).To(Equal(errUnexpected))
			})

			It("Userが保存されない。Avatarの状態は不定", func() {
				var count int
				Expect(db.Model(&entity.User{}).Count(&count).Error).To(Succeed())
				Expect(count).To(BeZero())
			})
		})
	})
})

func getS3ObjectSize(key string) (int, error) {
	s3c := s3.New(tests.AWS)
	out, err := s3c.HeadObject(&s3.HeadObjectInput{
		Bucket: &tests.Config.AWS.FilesBucket,
		Key:    &key,
	})
	if err != nil {
		return 0, err
	}

	return int(*out.ContentLength), nil
}

func newUser(id int, name ...string) *entity.User {
	user := &entity.User{
		ID:        id,
		Birthdate: time.Date(2020, 7, 7, 0, 0, 0, 0, time.Local),
		Interests: []*entity.UserInterest{},
		Gender:    model.GenderMale,
	}
	util.FillDymmyString(user, id)
	if len(name) > 0 {
		user.Name = name[0]
	}
	return user
}

type (
	s3mock struct {
		error
		s3iface.S3API
	}
)

func (m s3mock) PutObjectRequest(input *s3.PutObjectInput) (*request.Request, *s3.PutObjectOutput) {
	req, output := m.S3API.PutObjectRequest(input)
	if m.error != nil {
		req.Error = m.error
	}

	return req, output
}
