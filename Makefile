PROJECT ?= $(shell basename $(CURDIR))
VERSION ?= $(shell git describe --tags 2>/dev/null)

BITTAGS :=
LDFLAGS := -s -w
LDFLAGS += -X "github.com/starudream/go-lib/core/v2/config/version.gitVersion=$(VERSION)"

.PHONY: init
init:
	git status -b -s

.PHONY: update-all
update-all: update-core update-cobra update-cron update-resty update-ntfy update-selfupdate update-service update-sqlite

.PHONY: update-%
update-%:
	cd $* && go get -v -u ./...
	@cd $* && if grep -q 'go-resty' go.mod; then go get github.com/go-resty/resty/v2@v2.9.1; fi
	@cd $* && go mod tidy

.PHONY: test-all
test-all: test-core test-cobra test-cron test-resty test-tablew

.PHONY: test-%
test-%: init
	@CGO_ENABLED=0 go install github.com/kyoh86/richgo@latest
	@mkdir -p cover
	cd $* && go mod tidy && CGO_ENABLED=1 richgo test -race -count 1 -cover -failfast -coverprofile ../cover/$*.out -v ./...
	cd $* && go tool cover -html ../cover/$*.out -o ../cover/$*.html

.PHONY: bin-all
bin-all: bin-app bin-release bin-service

.PHONY: bin-%
bin-%: init
	cd example && CGO_ENABLED=0 go build -tags '$(BITTAGS)' -ldflags '$(LDFLAGS)' -o ../bin/example-$* github.com/starudream/go-lib/example/v2/$*

.PHONY: run-%
run-%: bin-%
	DEBUG=true APP__LOG__FILE__ENABLED=true APP__LOG__FILE__LEVEL=debug bin/example-$* $(ARGS)

.PHONY: lint-all
lint-all: lint-core lint-cobra lint-cron lint-ntfy lint-resty lint-service lint-sqlite

.PHONY: lint-%
lint-%:
	cd $* && golangci-lint run --skip-dirs internal/driver
