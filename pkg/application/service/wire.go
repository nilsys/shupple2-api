//+build wireinject

package service

import (
	"github.com/golang/mock/gomock"
	"github.com/google/wire"
	"github.com/onsi/ginkgo"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/repository"
	"github.com/stayway-corp/stayway-media-api/pkg/mock"
)

type Test struct {
	*PaymentCommandServiceImpl
}

func InitializeTest() (*Test, error) {
	wire.Build(
		wire.Struct(new(Test), "*"),
		ProvideTestReporter,
		gomock.NewController,
		ProvideMockPaymentQueryRepo,
		ProvideMockPaymentCmdRepo,
		ProvideMockMailCmdRepo,
		ProvideMockTransactionService,
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
