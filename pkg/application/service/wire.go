//+build wireinject

package service

import (
	"github.com/golang/mock/gomock"
	"github.com/google/wire"
	"github.com/onsi/ginkgo"
	"github.com/stayway-corp/stayway-media-api/pkg/config"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/repository"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/repository/payjp"
	"github.com/stayway-corp/stayway-media-api/pkg/mock"
)

type Test struct {
	*CfInnReserveRequestCommandServiceImpl
	*ChargeCommandServiceImpl
}

func InitializeTest(filePath config.FilePath) (*Test, error) {
	wire.Build(
		wire.Struct(new(Test), "*"),
		ProvideTestReporter,
		config.GetConfig,
		wire.FieldsOf(new(*config.Config), "CfProject"),
		gomock.NewController,
		ProvideMockPaymentQueryRepo,
		ProvideMockPaymentCmdRepo,
		ProvideMockMailCmdRepo,
		ProvideMockTransactionService,
		ProvideCardQueryRepo,
		ProvideCfProjectQueryRepo,
		ProvideChargeCmdRepo,
		ProvideCfReturnGiftQueryRepo,
		ProvideUserQueryRepo,
		ProvideCfReturnGiftCmdRepo,
		ProvideShippingQueryRepo,
		ProvideCfProjectCmdRepo,
		ProvideUserSalesHistoryRepo,
		ProvideCfInnReserveRequestCmdRepo,
		ServiceTestSet,
	)

	return new(Test), nil
}

func ProvideTestReporter() gomock.TestReporter {
	return ginkgo.GinkgoT()
}

func ProvideMockPaymentQueryRepo(ctrl *gomock.Controller) repository.PaymentQueryRepository {
	return mock.NewMockPaymentQueryRepository(ctrl)
}

func ProvideMockPaymentCmdRepo(ctrl *gomock.Controller) repository.PaymentCommandRepository {
	return mock.NewMockPaymentCommandRepository(ctrl)
}

func ProvideMockMailCmdRepo(ctrl *gomock.Controller) repository.MailCommandRepository {
	return mock.NewMockMailCommandRepository(ctrl)
}

func ProvideMockTransactionService() TransactionService {
	return TransactionServiceForTest{}
}

func ProvideCardQueryRepo(ctrl *gomock.Controller) repository.CardQueryRepository {
	return mock.NewMockCardQueryRepository(ctrl)
}

func ProvideCfProjectQueryRepo(ctrl *gomock.Controller) repository.CfProjectQueryRepository {
	return mock.NewMockCfProjectQueryRepository(ctrl)
}

func ProvideChargeCmdRepo(ctrl *gomock.Controller) payjp.ChargeCommandRepository {
	return mock.NewMockChargeCommandRepository(ctrl)
}

func ProvideCfReturnGiftQueryRepo(ctrl *gomock.Controller) repository.CfReturnGiftQueryRepository {
	return mock.NewMockCfReturnGiftQueryRepository(ctrl)
}

func ProvideUserQueryRepo(ctrl *gomock.Controller) repository.UserQueryRepository {
	return mock.NewMockUserQueryRepository(ctrl)
}

func ProvideCfReturnGiftCmdRepo(ctrl *gomock.Controller) repository.CfReturnGiftCommandRepository {
	return mock.NewMockCfReturnGiftCommandRepository(ctrl)
}

func ProvideShippingQueryRepo(ctrl *gomock.Controller) repository.ShippingQueryRepository {
	return mock.NewMockShippingQueryRepository(ctrl)
}

func ProvideCfProjectCmdRepo(ctrl *gomock.Controller) repository.CfProjectCommandRepository {
	return mock.NewMockCfProjectCommandRepository(ctrl)
}

func ProvideUserSalesHistoryRepo(ctrl *gomock.Controller) repository.UserSalesHistoryCommandRepository {
	return mock.NewMockUserSalesHistoryCommandRepository(ctrl)
}

func ProvideCfInnReserveRequestCmdRepo(ctrl *gomock.Controller) repository.CfInnReserveRequestCommandRepository {
	return mock.NewMockCfInnReserveRequestCommandRepository(ctrl)
}
