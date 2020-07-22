package payjp

import (
	"github.com/payjp/payjp-go/v1"
)

type (
	ChargeCommandRepository interface {
		// 支払いを作成
		Create(customerID string, cardID string, amount int) (*payjp.ChargeResponse, error)
		// 支払いを確定
		Capture(chargeID string) error
		// 部分返金
		Refund(chargeID string, amount int) error
	}
)
