package repository

import (
	"bytes"
	"context"
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

		Expect(clearBucket(tests.AWS, tests.Config.AWS.FilesBucket)).To(Succeed())
	})

	DescribeTable("Storeは引数のuserを作成するか、その状態になるように更新する",
		func(before *entity.User, saved *entity.User) {
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
		Entry("新規作成", nil, newUser(0)),
		Entry("フィールドに変更がある場合", newUser(userID), newUser(userID, "changed")),
	)

	Describe("FindByUIDs", func() {
		BeforeEach(func() {
			query = &UserQueryRepositoryImpl{DB: db}
			truncate(db)

			Expect(db.Save(newUser(1)).Error).To(Succeed())
			Expect(db.Save(newUser(2)).Error).To(Succeed())
			Expect(db.Save(newUser(3)).Error).To(Succeed())
		})

		DescribeTable("ユーザを取得できた場合", func() {
			users, err := query.FindByUIDs([]string{"UID1", "UID2"})

			Expect(err).To(Succeed())
			Expect(len(users)).To(Equal(2))
			Expect(users[0].ID).To(Equal(1))
			Expect(users[1].ID).To(Equal(2))
		},
			Entry("正常系"),
		)

		DescribeTable("ユーザを取得できなかった場合", func() {
			users, err := query.FindByUIDs([]string{"InvalidUID"})

			// エラーにならない
			Expect(err).To(Succeed())
			Expect(len(users)).To(Equal(0))
		},
			Entry("正常系"),
		)
	})

	Describe("StoreWithAvatar", func() {
		const contentType = "image/jpeg"
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
			err = command.StoreWithAvatar(base, bytes.NewReader(baseAvatar), contentType)
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
					user, err := query.FindByID(base.ID)
					Expect(err).To(Succeed())
					size, err := getS3ObjectSize(model.UserS3Path(user.AvatarUUID))
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
					Expect(command.StoreWithAvatar(&existing, bytes.NewReader(baseAvatar[:1]), contentType)).To(Succeed())
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
		UserTiny: entity.UserTiny{
			ID:             id,
			Birthdate:      sampleTime,
			UserAttributes: []*entity.UserAttribute{},
		},
		UserInterests: []*entity.UserInterest{},
	}
	util.FillDummyString(user, id)
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
