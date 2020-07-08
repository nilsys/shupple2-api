package service

import (
	"context"

	"github.com/stayway-corp/stayway-media-api/pkg/domain/model"

	"github.com/stayway-corp/stayway-media-api/pkg/util"

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
		repository.CfReturnGiftQueryRepository
		repository.ShippingQueryRepository
		repository.CfProjectCommandRepository
		repository.MailCommandRepository
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

		paymentReturnGifts := make([]*entity.PaymentCfReturnGift, len(cmd.List))

		gifts, err := s.CfReturnGiftQueryRepository.LockCfReturnGiftList(c, cmd.ReturnIDs())
		if err != nil {
			return errors.Wrap(err, "failed lock return_gift")
		}

		projectID, isUnique := gifts.UniqueCfProjectID()
		// 複数のprojectのギフトを同時購入しようとした時
		if !isUnique {
			return serror.New(nil, serror.CodeInvalidParam, "need gift's project is unique")
		}

		project, err := s.CfProjectQueryRepository.Lock(c, projectID)
		if err != nil {
			return errors.Wrap(err, "failed lock cf_project")
		}

		soldCountList, err := s.CfReturnGiftQueryRepository.FindSoldCountByReturnGiftIDs(c, cmd.ReturnIDs())
		if err != nil {
			return errors.Wrap(err, "failed find sold_count list")
		}

		giftIDMap := gifts.ToIDMap()

		var price int
		for i, payment := range cmd.List {
			// 商品情報が更新されている時
			if payment.ReturnGiftSnapshotID != giftIDMap[payment.ReturnGiftID].LatestCfReturnGiftSnapshotID {
				return serror.New(nil, serror.CodeInvalidParam, "return_gift updated")
			}

			// 残り在庫数確認
			if giftIDMap[payment.ReturnGiftID].Snapshot.FullAmount < soldCountList.GetSoldCount(payment.ReturnGiftID)+payment.Amount {
				// 在庫が確保できなかった場合
				return serror.New(nil, serror.CodeInvalidParam, "failed stock acquisition")
			}

			// 金額x数量
			price += giftIDMap[payment.ReturnGiftID].Snapshot.Price * payment.Amount
			gift := giftIDMap[payment.ReturnGiftID]
			if gift.CfReturnGiftTable.GiftType == model.GiftTypeReservedTicket {
				paymentReturnGifts[i] = entity.NewPaymentReturnGiftForReservedTicket(payment.ReturnGiftID, gift.LatestCfReturnGiftSnapshotID, gift.CfProjectID, int(project.LatestSnapshotID.Int64), payment.Amount)
				continue
			}
			paymentReturnGifts[i] = entity.NewPaymentReturnGift(payment.ReturnGiftID, gift.LatestCfReturnGiftSnapshotID, gift.CfProjectID, int(project.LatestSnapshotID.Int64), payment.Amount)
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
		payment := entity.NewPayment(user.ID, project.UserID, card.ID, address.ID, charge.ID, price)
		if err := s.PaymentCommandRepository.Store(c, payment); err != nil {
			return errors.Wrap(err, "failed store payment")
		}
		if err := s.PaymentCommandRepository.StorePaymentReturnGiftList(c, paymentReturnGifts, payment.ID); err != nil {
			return errors.Wrap(err, "failed store payment_return_gift")
		}

		// 応援コメントを保存
		comment := entity.NewCfProjectSupportTable(user.ID, project.ID, cmd.Body)
		if err := s.CfProjectCommandRepository.StoreSupportComment(c, comment); err != nil {
			return errors.Wrap(err, "failed store support_comment")
		}
		if err := s.CfProjectCommandRepository.IncrementSupportCommentCount(c, project.ID); err != nil {
			return errors.Wrap(err, "failed increment support_comment_count")
		}

		// projectの達成金額に加算
		if err := s.CfProjectCommandRepository.IncrementAchievedPrice(c, project.ID, price); err != nil {
			return errors.Wrap(err, "failed increment project_snapshot.achieved_price")
		}

		// 決済確定
		if err := s.ChargeCommandRepository.Capture(charge.ID); err != nil {
			return errors.Wrap(err, "failed capture charge")
		}

		// 決済確定メール送信
		template := entity.NewThanksPurchaseTemplate(project.User.Name, gifts.OnEmailDescription(), charge.ID, util.WithComma(price), address.Email, address.FullAddress(), user.Name)
		if err := s.MailCommandRepository.SendTemplateMail(address.Email, template); err != nil {
			return errors.Wrap(err, "failed send email from ses")
		}

		return nil
	})
}
