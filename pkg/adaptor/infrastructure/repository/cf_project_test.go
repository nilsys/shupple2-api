package repository

import (
	"context"
	"time"

	"github.com/stayway-corp/stayway-media-api/pkg/domain/model"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/entity"
	"github.com/stayway-corp/stayway-media-api/pkg/util"
)

var _ = Describe("", func() {
	var (
		queryRepo   *CfProjectQueryRepositoryImpl
		commandRepo *CfProjectCommandRepositoryImpl
		saved       *entity.CfProject
	)
	BeforeEach(func() {
		queryRepo = tests.CfProjectQueryRepositoryImpl
		commandRepo = tests.CfProjectCommandRepositoryImpl
		truncate(db)
		Expect(db.Save(newUser(userID)).Error).To(Succeed())
		Expect(db.Save(newCfProjectTable(cfProjectID, userID)).Error).To(Succeed())
		saved = newCfProject(cfProjectID, userID)
	})

	It("Lock: 正常系", func() {
		actual, err := queryRepo.Lock(context.WithValue(context.Background(), model.ContextKeyTransaction, db), cfProjectID)
		Expect(err).To(Succeed())

		Expect(actual.CreatedAt).NotTo(BeZero())
		Expect(actual.UpdatedAt).NotTo(BeZero())
		Expect(actual.User.CreatedAt).NotTo(BeZero())
		Expect(actual.User.UpdatedAt).NotTo(BeZero())
		actual.CreatedAt = time.Time{}
		actual.UpdatedAt = time.Time{}
		actual.User.CreatedAt = time.Time{}
		actual.User.UpdatedAt = time.Time{}

		Expect(actual).To(Equal(saved))
	})

	It("StoreSupportComment: 正常系", func() {
		err := commandRepo.StoreSupportComment(context.Background(), newSupportComment(cfProjectCommentID, userID, cfProjectID))
		Expect(err).To(Succeed())

		expect := newSupportComment(cfProjectCommentID, userID, cfProjectID)

		var actual entity.CfProjectSupportCommentTable
		err = db.Find(&actual, cfProjectCommentID).Error
		Expect(err).To(Succeed())

		Expect(actual.CreatedAt).NotTo(BeZero())
		Expect(actual.UpdatedAt).NotTo(BeZero())
		actual.CreatedAt = time.Time{}
		actual.UpdatedAt = time.Time{}

		Expect(&actual).To(Equal(expect))
	})

	It("IncrementSupportCommentCount: 正常系", func() {
		err := commandRepo.IncrementSupportCommentCount(context.Background(), cfProjectID)
		Expect(err).To(Succeed())

		saved.SupportCommentCount++

		actual, err := queryRepo.Lock(context.WithValue(context.Background(), model.ContextKeyTransaction, db), cfProjectID)
		Expect(err).To(Succeed())

		Expect(actual.CreatedAt).NotTo(BeZero())
		Expect(actual.UpdatedAt).NotTo(BeZero())
		Expect(actual.User.CreatedAt).NotTo(BeZero())
		Expect(actual.User.UpdatedAt).NotTo(BeZero())
		actual.CreatedAt = time.Time{}
		actual.UpdatedAt = time.Time{}
		actual.User.CreatedAt = time.Time{}
		actual.User.UpdatedAt = time.Time{}

		Expect(actual).To(Equal(saved))
	})
})

func newCfProjectTable(id, userID int) entity.CfProjectTable {
	return entity.CfProjectTable{
		ID:     id,
		UserID: userID,
	}
}

func newCfProject(id, userID int) *entity.CfProject {
	cfProject := &entity.CfProject{
		CfProjectTable: newCfProjectTable(id, userID),
		User:           newUser(userID),
	}
	util.FillDummyString(cfProject, id)
	return cfProject
}

func newSupportComment(id, userID, projectID int) *entity.CfProjectSupportCommentTable {
	return &entity.CfProjectSupportCommentTable{
		ID:          id,
		UserID:      userID,
		CfProjectID: projectID,
	}
}
