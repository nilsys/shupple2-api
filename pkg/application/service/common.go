package service

import (
	"github.com/google/wire"
	"github.com/stayway-corp/stayway-media-api/pkg/application/service/helper"
)

var ServiceTestSet = wire.NewSet(
	CfInnReserveRequestCommandServiceSet,
	ChargeCommandServiceSet,
	helper.InquiryCodeGeneratorForTestSet,
)
