PROJECT ?= $(shell basename $(CURDIR))
VERSION ?= $(shell git describe --tags 2>/dev/null)

GO      ?= go
MODULE  ?= github.com/starudream/go-lib
MODDIRS := core cobra resty example
BITTAGS :=
LDFLAGS := -s -w
LDFLAGS += -X "$(MODULE)/core/v2/config/version.gitVersion=$(VERSION)"

.PHONY: init
init:
	@if [ ! -f "go.work" ]; then $(GO) work init $(MODDIRS); fi
	$(GO) work sync

.PHONY: test-%
test-%: init
	@if command -v richgo >/dev/null; then echo $(eval GO := richgo) >/dev/null; fi
	@mkdir -p cover
	cd $* && go mod tidy && CGO_ENABLED=1 $(GO) test -race -count 1 -cover -failfast -coverprofile ../cover/$*.out -v ./...
	go tool cover -html cover/$*.out -o cover/$*.html

.PHONY: bin-%
bin-%: init
	CGO_ENABLED=0 $(GO) build -race -tags '$(BITTAGS)' -ldflags '$(LDFLAGS)' -o bin/example-$* $(MODULE)/v2/example/$*

.PHONY: run-%
run-%: bin-%
	DEBUG=true bin/example-$* $(ARGS)
