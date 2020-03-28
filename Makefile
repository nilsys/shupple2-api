BIN=$(abspath ./bin)
PKG=github.com/stayway-corp/stayway
LDFLAGS="-X $(PKG)/config.Version=$(shell git rev-parse --short HEAD)"
BUILD_OPTS = -ldflags $(LDFLAGS)

MYSQL_HOST ?= mysql

air=$(BIN)/air
migrate=$(BIN)/migrate
wire=$(BIN)/wire
ginkgo=$(BIN)/ginkgo
mockgen=$(BIN)/mockgen
linter=$(BIN)/golangci-lint
go-enum=$(BIN)/go-enum

build: build-app build-batch

build-app:
	go build $(BUILD_OPTS) -o bin/stayway-media-api ./cmd/app
build-batch:
	go build $(BUILD_OPTS) -o bin/stayway-media-batch ./cmd/batch

install-cli: $(air) $(migrate) $(wire) $(ginkgo) $(mockgen) $(linter) $(go-enum)
$(air):
	GOBIN=$(BIN) go install github.com/cosmtrek/air
$(migrate):
	GOBIN=$(BIN) go install -tags mysql github.com/golang-migrate/migrate/v4/cmd/migrate
$(wire):
	GOBIN=$(BIN) go install github.com/google/wire/cmd/wire
$(ginkgo):
	GOBIN=$(BIN) go install github.com/onsi/ginkgo/ginkgo
$(mockgen):
	GOBIN=$(BIN) go install github.com/golang/mock/mockgen
$(linter):
	GOBIN=$(BIN) go install github.com/golangci/golangci-lint/cmd/golangci-lint
$(go-enum):
	GOBIN=$(BIN) go install github.com/abice/go-enum

wait-mysql:
	dockerize -wait tcp://$(MYSQL_HOST):3306 -timeout 30s

migrate-up: $(migrate) $(wait-mysql)
	$(migrate) -database 'mysql://root:@tcp($(MYSQL_HOST):3306)/stayway' -path ./migrations up $(NUM)

migrate-down: $(migrate) $(wait-mysql)
	$(migrate) -database 'mysql://root:@tcp($(MYSQL_HOST):3306)/stayway' -path ./migrations down $(NUM)

migrate-cmd: $(migrate)
	$(migrate) -database 'mysql://root:@tcp($(MYSQL_HOST):3306)/stayway' -path ./migrations $(CMD)

run: start-dev
start-dev: $(air) $(wire) migrate-up
	$(air)

TARGET ?= ./pkg/...
lint: $(linter)
	$(linter) run $(TARGET)
test: $(ginkgo)
	ENV=test $(ginkgo) $(TARGET)

mysql-cli:
	mysql --port $(shell docker inspect stayway-media-api-mysql | jq -r '.[0].NetworkSettings.Ports."3306/tcp"[0].HostPort') \
	--host 127.0.0.1 --user root stayway

generate: $(mockgen) $(wire) $(go-enum)
	PATH=$(PATH):$(BIN) go generate ./...
