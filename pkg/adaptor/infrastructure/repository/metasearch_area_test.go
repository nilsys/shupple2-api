package repository

import (
	"github.com/jinzhu/gorm"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/entity"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/model"
)

var _ = Describe("MetasearchAreaRepositoryImpl", func() {
	const (
		metasearchAreaID       = metasearchAreaID
		metasearchSubAreaID    = metasearchAreaID + 1
		metasearchSubSubAreaID = metasearchAreaID + 2
	)
	var (
		query *MetasearchAreaQueryRepositoryImpl
	)

	BeforeEach(func() {
		query = tests.MetasearchAreaQueryRepositoryImpl

		truncate(db)
		Expect(db.Create(newAreaCategory(areaCategoryID)).Error).To(Succeed())
		createMetasearchArea(db, metasearchAreaID, model.AreaCategoryTypeArea)
		createMetasearchArea(db, metasearchSubAreaID, model.AreaCategoryTypeSubArea)
		createMetasearchArea(db, metasearchSubSubAreaID, model.AreaCategoryTypeSubSubArea)
	})

	DescribeTable("FindByMetasearchAreaIDはメタサーチのエリアに相当するエリアカテゴリを一つ返す",
		func(metasearchAreaID int, metasearchAreaType model.AreaCategoryType, exists bool) {
			actual, err := query.FindByMetasearchAreaID(metasearchAreaID, metasearchAreaType)
			if exists {
				Expect(err).To(Succeed())
				Expect(actual.AreaCategoryID).To(Equal(areaCategoryID))
			} else {
				Expect(err).To(HaveOccurred())
				Expect(actual).To(BeNil())
			}
		},
		Entry("正常", metasearchAreaID, model.AreaCategoryTypeArea, true),
		Entry("存在しないメタサーチエリアIDの場合エラー", metasearchAreaID, model.AreaCategoryTypeSubArea, false),
	)

	DescribeTable("FindByAreaCategoryIDはエリアカテゴリに紐づくメタサーチのエリアの一覧を返す",
		func(areaCategoryID int, areaCategoryType model.AreaCategoryType, length int) {
			actual, err := query.FindByAreaCategoryID(areaCategoryID, areaCategoryType)
			Expect(actual).To(HaveLen(length))
			Expect(err).To(Succeed())
		},
		Entry("正常", areaCategoryID, model.AreaCategoryTypeArea, 3),
		Entry("存在しないもエラーにはならない", areaCategoryID*100, model.AreaCategoryTypeArea, 0),
	)
})

func createMetasearchArea(db *gorm.DB, metasearchAreaID int, metasearchAreaType model.AreaCategoryType) {
	area := &entity.MetasearchArea{
		MetasearchAreaID:   metasearchAreaID,
		MetasearchAreaType: metasearchAreaType,
		AreaCategoryID:     areaCategoryID,
	}
	Expect(db.Create(area).Error).To(Succeed())
}
