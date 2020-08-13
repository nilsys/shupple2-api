package service

import (
	"testing"

	"github.com/pkg/errors"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var (
	tests *Test
)

const (
	configFilePath = "./../../../config.test.yaml"
)

func TestService(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Service Suite")
}

var _ = BeforeSuite(func() {
	Expect(beforeSuite()).To(Succeed())
})

func beforeSuite() error {
	var err error
	tests, err = InitializeTest(configFilePath)
	if err != nil {
		return errors.Wrap(err, "failed to initialize test")
	}
	return nil
}
