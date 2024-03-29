PROJECT ?= $(shell basename $(CURDIR))
VERSION ?= $(shell git describe --tags 2>/dev/null)

BITTAGS :=
LDFLAGS := -s -w
LDFLAGS += -X "github.com/starudream/go-lib/core/v2/config/version.gitVersion=$(VERSION)"

.PHONY: init
init:
	git status -b -s

.PHONY: update-all
update-all: update-core update-cobra update-cron update-resty update-ntfy update-selfupdate update-server update-service update-sqlite update-example

.PHONY: update-%
update-%:
	cd $* && go get -v -u ./... && go mod tidy

.PHONY: test-all
test-all: test-core test-cobra test-cron test-resty update-selfupdate test-server test-tablew

.PHONY: test-%
test-%: init
	@CGO_ENABLED=0 go install github.com/kyoh86/richgo@latest
	@mkdir -p cover
	cd $* && go mod tidy && CGO_ENABLED=1 richgo test -race -count 1 -cover -failfast -coverprofile ../cover/$*.out -v ./...
	cd $* && go tool cover -html ../cover/$*.out -o ../cover/$*.html

.PHONY: bin-all
bin-all: bin-app bin-release bin-selfupdate bin-service bin-sqlite

.PHONY: bin-%
bin-%: init
	cd example && go mod tidy && CGO_ENABLED=0 go build -tags '$(BITTAGS)' -ldflags '$(LDFLAGS)' -o ../bin/example-$* github.com/starudream/go-lib/example/v2/$*

.PHONY: run-%
run-%: bin-%
	DEBUG=true APP__LOG__FILE__ENABLED=true APP__LOG__FILE__LEVEL=debug bin/example-$* $(ARGS)

.PHONY: lint-all
lint-all: lint-core lint-cobra lint-cron lint-resty lint-ntfy lint-selfupdate lint-server lint-service lint-sqlite

.PHONY: lint-sqlite
lint-sqlite:
	cd sqlite && golangci-lint run --sort-results --print-resources-usage --show-stats --skip-dirs internal/driver

.PHONY: lint-%
lint-%:
	cd $* && golangci-lint run --sort-results --print-resources-usage --show-stats
