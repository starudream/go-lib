PROJECT ?= $(shell basename $(CURDIR))
VERSION ?= $(shell git describe --tags 2>/dev/null)

GO      ?= go
BITTAGS :=
LDFLAGS := -s -w
LDFLAGS += -X "github.com/starudream/go-lib/core/v2/config/version.gitVersion=$(VERSION)"

.PHONY: init
init:
	git status -b -s

.PHONY: test-%
test-%: init
	@if command -v richgo >/dev/null; then echo $(eval GO := richgo) >/dev/null; fi
	@mkdir -p cover
	cd $* && go mod tidy && CGO_ENABLED=1 $(GO) test -race -count 1 -cover -failfast -coverprofile ../cover/$*.out -v ./...
	cd $* && go tool cover -html ../cover/$*.out -o ../cover/$*.html

.PHONY: bin-%
bin-%: init
	cd example && CGO_ENABLED=0 $(GO) build -race -tags '$(BITTAGS)' -ldflags '$(LDFLAGS)' -o ../bin/example-$* github.com/starudream/go-lib/example/v2/$*

.PHONY: run-%
run-%: bin-%
	DEBUG=true bin/example-$* $(ARGS)
