package helper

import (
	"github.com/google/wire"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/model"
)

type (
	InquiryCodeGenerator interface {
		Gen() (string, error)
	}

	InquiryCodeGeneratorImpl struct {
	}

	InquiryCodeGeneratorImplForTest struct {
	}
)

const (
	// PaymentCfReturnGift.InquiryCode(お問い合わせ番号)の桁数
	paymentCfReturnGiftInquiryCodeLength = 7
	// テスト使用時のダミー
	testInquiryCode = "dummy"
)

var InquiryCodeGeneratorSet = wire.NewSet(
	wire.Struct(new(InquiryCodeGeneratorImpl), "*"),
	wire.Bind(new(InquiryCodeGenerator), new(*InquiryCodeGeneratorImpl)),
)

var InquiryCodeGeneratorForTestSet = wire.NewSet(
	wire.Struct(new(InquiryCodeGeneratorImplForTest), "*"),
	wire.Bind(new(InquiryCodeGenerator), new(*InquiryCodeGeneratorImplForTest)),
)

func (g *InquiryCodeGeneratorImpl) Gen() (string, error) {
	return model.RandomStr(paymentCfReturnGiftInquiryCodeLength)
}

func (g *InquiryCodeGeneratorImplForTest) Gen() (string, error) {
	return testInquiryCode, nil
}
