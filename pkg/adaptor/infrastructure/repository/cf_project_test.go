package repository

import (
	"context"
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/entity"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/model"
	"github.com/stayway-corp/stayway-media-api/pkg/util"
)

var _ = Describe("CfProjectRepositoryImpl", func() {
	var (
		commandRepo *CfProjectCommandRepositoryImpl
		base        *entity.CfProject
	)
	findCfProjectByID := func(id int) (*entity.CfProject, error) {
		var cfProject entity.CfProject
		return &cfProject, db.Find(&cfProject, id).Error
	}

	BeforeEach(func() {
		commandRepo = tests.CfProjectCommandRepositoryImpl
		truncate(db)
		Expect(db.Save(newUser(userID)).Error).To(Succeed())

		for _, cat := range append(areaCategoryIDs, addedAreaCategoryID) {
			Expect(db.Save(newAreaCategory(cat)).Error).To(Succeed())
		}

		for _, cat := range append(themeCategoryIDs, addedThemeCategoryID) {
			Expect(db.Save(newThemeCategory(cat)).Error).To(Succeed())
		}

		base = newCfProject(cfProjectID, userID, thumbnails, areaCategoryIDs, themeCategoryIDs)
		Expect(db.Save(base).Error).To(Succeed())
	})

	Describe("Store", func() {
		Context("新規作成時", func() {
			It("cf_projectとcf_project_snapshotが作成される", func() {
				const newInsertedID = cfProjectID + 1
				newInserted := newCfProject(newInsertedID, userID, nil, nil, nil)
				Expect(commandRepo.Store(newInserted)).To(Succeed())

				var count int
				Expect(db.Model(&entity.CfProjectTiny{}).Count(&count).Error).To(Succeed())
				Expect(count).To(Equal(2))
				Expect(db.Model(&entity.CfProjectSnapshotTiny{}).Count(&count).Error).To(Succeed())
				Expect(count).To(Equal(2))

				actual, err := findCfProjectByID(newInsertedID)
				Expect(err).To(Succeed())

				newInserted.LatestSnapshotID = actual.LatestSnapshotID
				Expect(actual).To(entity.EqualEntity(newInserted))
			})
		})

		Context("更新時", func() {
			const achievedPrice = 100
			BeforeEach(func() {
				Expect(commandRepo.IncrementSupportCommentCount(context.Background(), cfProjectID)).To(Succeed())
				Expect(commandRepo.IncrementAchievedPrice(context.Background(), cfProjectID, achievedPrice)).To(Succeed())
				Expect(commandRepo.IncrementFavoriteCountByID(context.Background(), cfProjectID)).To(Succeed())
			})

			It("cf_projectのカラムは更新されず、cf_project_snapshotが挿入される。cf_project.latest_snapshot_idが更新される。", func() {
				Expect(commandRepo.Store(base)).To(Succeed())

				var count int
				Expect(db.Model(&entity.CfProjectTiny{}).Count(&count).Error).To(Succeed())
				Expect(count).To(Equal(1))
				Expect(db.Model(&entity.CfProjectSnapshotTiny{}).Count(&count).Error).To(Succeed())
				Expect(count).To(Equal(2))

				actual, err := findCfProjectByID(cfProjectID)
				Expect(err).To(Succeed())

				var latestSnapshotID int
				Expect(db.Raw("SELECT MAX(id) FROM cf_project_snapshot").Row().Scan(&latestSnapshotID)).To(Succeed())
				Expect(int(actual.LatestSnapshotID.Int64)).To(Equal(latestSnapshotID))
				Expect(actual.Snapshot.SnapshotID).To(Equal(latestSnapshotID))

				base.FavoriteCount = 1
				base.SupportCommentCount = 1
				base.AchievedPrice = achievedPrice
				base.LatestSnapshotID = actual.LatestSnapshotID
				base.Snapshot.SnapshotID = actual.Snapshot.SnapshotID
				Expect(actual).To(entity.EqualEntity(base))
			})
		})
	})

	Describe("Lock", func() {
		It("正常系", func() {
			actual, err := commandRepo.Lock(context.WithValue(context.Background(), model.ContextKeyTransaction, db), cfProjectID)
			Expect(err).To(Succeed())
			Expect(actual).To(entity.EqualEntity(base))
		})
	})

	Describe("StoreSupportComment", func() {
		It("正常系", func() {
			err := commandRepo.StoreSupportComment(context.Background(), newSupportComment(cfProjectCommentID, userID, cfProjectID))
			Expect(err).To(Succeed())

			expect := newSupportComment(cfProjectCommentID, userID, cfProjectID)

			var actual entity.CfProjectSupportCommentTiny
			err = db.Find(&actual, cfProjectCommentID).Error
			Expect(err).To(Succeed())

			Expect(&actual).To(entity.EqualEntity(expect))
		})
	})

	Describe("IncrementSupportCommentCount", func() {
		It("正常系", func() {
			err := commandRepo.IncrementSupportCommentCount(context.Background(), cfProjectID)
			Expect(err).To(Succeed())

			base.SupportCommentCount++

			actual, err := findCfProjectByID(cfProjectID)
			Expect(err).To(Succeed())

			Expect(actual).To(entity.EqualEntity(base))
		})
	})

	Describe("IncrementAchievedPrice", func() {
		It("正常系", func() {
			err := commandRepo.IncrementAchievedPrice(context.Background(), cfProjectID, 100)
			Expect(err).To(Succeed())

			base.AchievedPrice += 100

			actual, err := findCfProjectByID(cfProjectID)
			Expect(err).To(Succeed())

			Expect(actual).To(entity.EqualEntity(base))
		})
	})
})

func newCfProject(id, userID int, thumbnails []string, areaCategoryIDs []int, themeCategoryIDs []int) *entity.CfProject {
	cfProject := &entity.CfProject{
		CfProjectTiny: entity.CfProjectTiny{
			ID:     id,
			UserID: userID,
		},
		Snapshot: entity.CfProjectSnapshot{
			CfProjectSnapshotTiny: entity.CfProjectSnapshotTiny{
				CfProjectID: id,
				Deadline:    sampleTime,
			},
		},
	}
	util.FillDummyString(cfProject, id)
	cfProject.Snapshot.SetThumbnails(thumbnails)
	cfProject.Snapshot.SetAreaCategories(areaCategoryIDs)
	cfProject.Snapshot.SetThemeCategories(themeCategoryIDs)
	return cfProject
}

func newCfProjectSnapshotTiny() *entity.CfProjectSnapshotTiny {
	snapshot := &entity.CfProjectSnapshotTiny{
		SnapshotID:  cfProjectSnapshotID,
		CfProjectID: cfProjectID,
		UserID:      userID,
		Deadline:    time.Date(2020, 7, 7, 0, 0, 0, 0, time.Local),
	}
	util.FillDummyString(snapshot, cfProjectSnapshotID)
	return snapshot
}

func newSupportComment(id, userID, projectID int) *entity.CfProjectSupportCommentTiny {
	return &entity.CfProjectSupportCommentTiny{
		ID:          id,
		UserID:      userID,
		CfProjectID: projectID,
	}
}
