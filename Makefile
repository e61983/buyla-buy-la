
GO ?= go
GO_RUN = $(GO) run

SRC ?= main.go

all: run

run: $(SRC)
	$(GO_RUN) my-local-test
