package service

import (
	"context"
	"time"

	"github.com/stayway-corp/stayway-media-api/pkg/config"

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
		Capture(user *entity.User, cmd *command.PaymentList) (*entity.CaptureResult, error)
		Refund(user *entity.User, paymentID, cfReturnGiftID int) error
	}

	ChargeCommandServiceImpl struct {
		repository.PaymentCommandRepository
		repository.PaymentQueryRepository
		repository.CardQueryRepository
		repository.CfProjectQueryRepository
		payjp.ChargeCommandRepository
		repository.CfReturnGiftQueryRepository
		repository.UserQueryRepository
		repository.CfReturnGiftCommandRepository
		repository.ShippingQueryRepository
		repository.CfProjectCommandRepository
		repository.MailCommandRepository
		TransactionService
		CfProjectConfig config.CfProject
	}
)

var ChargeCommandServiceSet = wire.NewSet(
	wire.Struct(new(ChargeCommandServiceImpl), "*"),
	wire.Bind(new(ChargeCommandService), new(*ChargeCommandServiceImpl)),
)

const (
	// 24h * 180days = 4230h
	// https://pay.jp/docs/api/#%E6%94%AF%E6%89%95%E3%81%84%E6%83%85%E5%A0%B1%E3%82%92%E6%9B%B4%E6%96%B0
	chargeRefundExpiredHour = 4320

	// PaymentCfReturnGift.InquiryCode(お問い合わせ番号)の桁数
	paymentCfReturnGiftInquiryCodeLength = 7
)

// 決済確定
// 登録されている最新のカードで決済が行われる
func (s *ChargeCommandServiceImpl) Capture(user *entity.User, cmd *command.PaymentList) (*entity.CaptureResult, error) {
	resolve := &entity.CaptureResult{}

	err := s.TransactionService.Do(func(c context.Context) error {
		address, err := s.ShippingQueryRepository.FindLatestShippingAddressByUserID(c, user.ID)
		if err != nil {
			return errors.Wrap(err, "failed find latest")
		}

		gifts, err := s.CfReturnGiftCommandRepository.LockByIDs(c, cmd.ReturnIDs())
		if err != nil {
			return errors.Wrap(err, "failed lock return_gift")
		}

		projectID, isUnique := gifts.UniqueCfProjectID()
		// 複数のprojectのギフトを同時購入しようとした時
		if !isUnique {
			return serror.New(nil, serror.CodeInvalidParam, "need gift's project is unique")
		}

		project, err := s.CfProjectCommandRepository.Lock(c, projectID)
		if err != nil {
			return errors.Wrap(err, "failed lock cf_project")
		}

		projectOwner, err := s.UserQueryRepository.FindByID(project.UserID)
		if err != nil {
			return errors.Wrap(err, "failed to get project owner")
		}

		soldCountList, err := s.CfReturnGiftQueryRepository.FindSoldCountByReturnGiftIDs(c, cmd.ReturnIDs())
		if err != nil {
			return errors.Wrap(err, "failed find sold_count list")
		}

		giftIDMap := gifts.ToIDMap()

		paymentReturnGifts := make([]*entity.PaymentCfReturnGiftTiny, len(cmd.List))
		idInquiryCodeMap := make(map[int]string, len(cmd.List))
		var price int

		for i, payment := range cmd.List {
			// 商品情報が更新されている時
			if int64(payment.ReturnGiftSnapshotID) != giftIDMap[payment.ReturnGiftID].LatestSnapshotID.Int64 {
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

			// お問い合わせ番号生成
			inquiryCode, err := model.RandomStr(paymentCfReturnGiftInquiryCodeLength)
			if err != nil {
				return errors.Wrap(err, "failed random str")
			}

			idInquiryCodeMap[gift.ID] = inquiryCode

			if gift.CfReturnGiftTiny.GiftType == model.CfReturnGiftTypeReservedTicket {
				paymentReturnGifts[i] = entity.NewPaymentReturnGiftForReservedTicket(payment.ReturnGiftID, int(gift.LatestSnapshotID.Int64), gift.CfProjectID, int(project.LatestSnapshotID.Int64), payment.Amount, inquiryCode)
				continue
			}
			paymentReturnGifts[i] = entity.NewPaymentReturnGiftForOther(payment.ReturnGiftID, int(gift.LatestSnapshotID.Int64), gift.CfProjectID, int(project.LatestSnapshotID.Int64), payment.Amount, inquiryCode)
		}

		// 最新のカードidを取得
		card, err := s.CardQueryRepository.FindLatestByUserID(c, user.ID)
		if err != nil {
			return errors.Wrap(err, "failed find latest credit_card")
		}

		// 手数料込の金額
		includeCommissionPrice := price + s.CfProjectConfig.SystemFee

		// 支払いを作成
		charge, err := s.ChargeCommandRepository.Create(user.PayjpCustomerID(), card.CardID, includeCommissionPrice)
		if err != nil {
			return errors.Wrap(err, "failed create pay")
		}

		// 金額の確保等、認証に失敗した場合
		if !charge.Paid {
			return serror.New(nil, serror.CodePayAgentError, "credit card authorization error")
		}

		// 支払い情報を保存
		payment := entity.NewPaymentTiny(user.ID, project.UserID, card.ID, address.ID, charge.ID, price, s.CfProjectConfig.SystemFee)
		if err := s.PaymentCommandRepository.Store(c, payment); err != nil {
			return errors.Wrap(err, "failed store payment")
		}
		if err := s.PaymentCommandRepository.StorePaymentReturnGiftList(c, paymentReturnGifts, payment.ID); err != nil {
			return errors.Wrap(err, "failed store payment_return_gift")
		}

		// 応援コメントを保存
		comment := entity.NewCfProjectSupportTiny(user.ID, project.ID, cmd.Body)
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
		template := entity.NewThanksPurchaseTemplate(projectOwner.Name, gifts.TitlesOnEmail(idInquiryCodeMap), charge.ID, util.WithComma(s.CfProjectConfig.SystemFee), util.WithComma(includeCommissionPrice), address.Email, address.FullAddress(), address.PhoneNumber, user.Name)
		if err := s.MailCommandRepository.SendTemplateMail([]string{address.Email}, template); err != nil {
			return errors.Wrap(err, "failed send email from ses")
		}

		// TODO: 楽観的？
		resolve.CfProjectID = project.ID
		resolve.SupporterCount = project.SupportCommentCount + 1
		resolve.AchievedPrice = project.AchievedPrice + price

		return nil
	})
	if err != nil {
		return nil, errors.Wrap(err, "failed capture charge")
	}

	return resolve, nil
}

func (s *ChargeCommandServiceImpl) Refund(user *entity.User, paymentID, cfReturnGiftID int) error {
	payment, err := s.PaymentQueryRepository.FindByID(paymentID)
	if err != nil {
		return errors.Wrap(err, "failed find payment")
	}
	if !user.IsSelfID(payment.UserID) {
		return serror.New(nil, serror.CodeForbidden, "forbidden")
	}

	return s.TransactionService.Do(func(ctx context.Context) error {
		paymentCfReturnGift, err := s.PaymentQueryRepository.LockPaymentCfReturnGift(ctx, payment.ID, cfReturnGiftID)
		if err != nil {
			return errors.Wrap(err, "failed find payment_cf_return_gift")
		}

		if !paymentCfReturnGift.CfReturnGiftSnapshot.IsCancelable || !paymentCfReturnGift.ResolveGiftTypeOtherStatus().CanTransit(model.PaymentCfReturnGiftOtherTypeStatusCanceled) {
			return serror.New(nil, serror.CodeInvalidParam, "can't cancel")
		}
		if paymentCfReturnGift.CreatedAt.Add(chargeRefundExpiredHour * time.Hour).After(time.Now()) {
			return serror.New(nil, serror.CodeInvalidParam, "expired")
		}

		if err := s.PaymentCommandRepository.MarkPaymentCfReturnGiftAsCancel(ctx, payment.ID, cfReturnGiftID); err != nil {
			return errors.Wrap(err, "failed mark as cancel")
		}

		// projectの達成金額から減算
		if err := s.CfProjectCommandRepository.DecrementAchievedPrice(ctx, payment.ID, paymentCfReturnGift.TotalPrice()); err != nil {
			return errors.Wrap(err, "failed decrement achieved_price")
		}

		// 返金
		if err := s.ChargeCommandRepository.Refund(payment.ChargeID, paymentCfReturnGift.TotalPrice()); err != nil {
			return errors.Wrap(err, "failed refund charge")
		}

		return nil
	})
}
