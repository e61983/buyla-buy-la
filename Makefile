GO ?= go
GO_TEST = $(GO) test
GO_FMT = $(GO) fmt
GO_RUN = $(GO) run
GO_BUILD = $(GO) build

SRC ?= main.go

all: fmt test

.PHONY: fmt
fmt:
	@$(GO_FMT) ./...

.PHONY: test
test:
	@$(GO_TEST)

run: $(SRC)
	$(GO_RUN) github.com/e61983/buyla-buy-la
