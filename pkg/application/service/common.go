package service

import (
	"github.com/google/wire"
)

var ServiceTestSet = wire.NewSet(
	PaymentCommandServiceSet,
)
