package repository

import (
	"context"
	"time"

	"github.com/stayway-corp/stayway-media-api/pkg/util"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/entity"
)

var _ = Describe("ShippingRepositoryTest", func() {
	var (
		queryRepo   *ShippingQueryRepositoryImpl
		commandRepo *ShippingCommandRepositoryImpl
	)
	BeforeEach(func() {
		queryRepo = tests.ShippingQueryRepositoryImpl
		commandRepo = tests.ShippingCommandRepositoryImpl
		truncate(db)
		Expect(db.Save(newUser(userID)).Error).To(Succeed())
	})

	base := newShippingAddress(shippingAddressID, userID)
	baseChanged := newShippingAddress(shippingAddressID, userID)
	baseChanged.FirstName = "changed"

	It("StoreShippingAddress: 引数のshippingAddressを作成or更新", func() {
		test := func(before *entity.ShippingAddress, saved *entity.ShippingAddress) {
			if before != nil {
				Expect(commandRepo.StoreShippingAddress(before)).To(Succeed())
			}

			Expect(commandRepo.StoreShippingAddress(saved)).To(Succeed())
			actual, err := queryRepo.FindLatestShippingAddressByUserID(context.Background(), saved.UserID)
			Expect(err).To(Succeed())

			Expect(actual.CreatedAt).NotTo(BeZero())
			Expect(actual.UpdatedAt).NotTo(BeZero())
			actual.CreatedAt = time.Time{}
			actual.UpdatedAt = time.Time{}

			Expect(actual).To(Equal(saved))
		}

		// 新規作成
		test(nil, base)
		// 更新
		test(base, baseChanged)
	})
})

func newShippingAddress(id, userID int) *entity.ShippingAddress {
	address := &entity.ShippingAddress{
		ID:     id,
		UserID: userID,
		Times:  entity.Times{},
	}

	util.FillDummyString(address, id)
	return address
}
