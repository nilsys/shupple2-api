package service

import (
	"context"

	"github.com/stayway-corp/stayway-media-api/pkg/domain/repository/payjp"

	"github.com/stayway-corp/stayway-media-api/pkg/domain/model/serror"

	"github.com/stayway-corp/stayway-media-api/pkg/domain/model/command"

	"github.com/google/wire"
	"github.com/pkg/errors"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/entity"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/repository"
)

type (
	ChargeCommandService interface {
		CaptureCharge(user *entity.User, cmd *command.PaymentList) error
	}

	ChargeCommandServiceImpl struct {
		repository.PaymentCommandRepository
		repository.CardQueryRepository
		repository.CfProjectQueryRepository
		payjp.ChargeCommandRepository
		repository.ReturnGiftQueryRepository
		repository.ShippingQueryRepository
		TransactionService
	}
)

var ChargeCommandServiceSet = wire.NewSet(
	wire.Struct(new(ChargeCommandServiceImpl), "*"),
	wire.Bind(new(ChargeCommandService), new(*ChargeCommandServiceImpl)),
)

// 決済確定
// 登録されている最新のカードで決済が行われる
func (s *ChargeCommandServiceImpl) CaptureCharge(user *entity.User, cmd *command.PaymentList) error {
	return s.TransactionService.Do(func(c context.Context) error {
		address, err := s.ShippingQueryRepository.FindLatestShippingAddressByUserID(c, user.ID)
		if err != nil {
			return errors.Wrap(err, "failed find latest")
		}

		paymentReturnGifts := make([]*entity.PaymentReturnGift, len(cmd.Payments))
		gifts, err := s.ReturnGiftQueryRepository.LockReturnGiftsWitLatestSummary(c, cmd.ReturnIDs())
		if err != nil {
			return errors.Wrap(err, "failed lock return_gift")
		}

		projects, err := s.CfProjectQueryRepository.LockCfProjectListByIDs(c, gifts.CfProjectIDs())
		if err != nil {
			return errors.Wrap(err, "failed lock cf_project")
		}

		soldCountList, err := s.ReturnGiftQueryRepository.FindSoldCountByReturnGiftIDs(c, cmd.ReturnIDs())
		if err != nil {
			return errors.Wrap(err, "failed find sold_count list")
		}

		giftIDMap := gifts.ToIDMap()
		giftIDSoldCountMap := soldCountList.ToIDSoldCountMap()
		projectIDMap := projects.ToIDMap()

		var price int
		for i, payment := range cmd.Payments {
			// 商品情報が更新されている時
			if payment.ReturnGiftSummaryID != giftIDMap[payment.ReturnGiftID].LatestSnapshotID {
				return serror.New(nil, serror.CodeInvalidParam, "modify return_gift_summary")
			}

			// 残り在庫数確認
			if giftIDMap[payment.ReturnGiftID].Summary.FullAmount < giftIDSoldCountMap[payment.ReturnGiftID]+payment.Amount {
				// 在庫が確保できなかった場合
				return serror.New(nil, serror.CodeInvalidParam, "failed stock acquisition")
			}

			// 金額x数量
			price += giftIDMap[payment.ReturnGiftID].Summary.Price * payment.Amount
			gift := giftIDMap[payment.ReturnGiftID]
			project := projectIDMap[gift.CfProjectID]
			paymentReturnGifts[i] = entity.NewPaymentReturnGift(payment.ReturnGiftID, gift.LatestSnapshotID, gift.CfProjectID, project.LatestSnapshotID, payment.Amount)
		}

		// 最新のカードidを取得
		card, err := s.CardQueryRepository.FindLatestByUserID(c, user.ID)
		if err != nil {
			return errors.Wrap(err, "failed find latest credit_card")
		}

		// 支払いを作成
		charge, err := s.ChargeCommandRepository.Create(user.PayjpCustomerID(), card.CardID, price)
		if err != nil {
			return errors.Wrap(err, "failed create pay")
		}

		// 金額の確保等、認証に失敗した場合
		if !charge.Paid {
			return serror.New(nil, serror.CodePayAgentError, "credit card authorization error")
		}

		// 支払い情報を保存
		payment := entity.NewPayment(user.ID, card.ID, address.ID, charge.ID)
		if err := s.PaymentCommandRepository.Store(c, payment); err != nil {
			return errors.Wrap(err, "failed store payment")
		}
		if err := s.PaymentCommandRepository.StorePaymentReturnGiftList(c, paymentReturnGifts, payment.ID); err != nil {
			return errors.Wrap(err, "failed store payment_return_gift")
		}

		// 決済確定
		if err := s.ChargeCommandRepository.Capture(charge.ID); err != nil {
			return errors.Wrap(err, "failed capture charge")
		}

		return nil
	})
}
