// +build tools
package tools

import (
	_ "github.com/cosmtrek/air"
	_ "github.com/golang-migrate/migrate/v4/cmd/migrate"
	_ "github.com/golang/mock/mockgen"
	_ "github.com/golangci/golangci-lint/cmd/golangci-lint"
	_ "github.com/google/wire/cmd/wire"
	_ "github.com/onsi/ginkgo/ginkgo"
)
