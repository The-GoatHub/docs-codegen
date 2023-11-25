BIN := $(shell pwd)/bin
GO?=$(shell which go)
export GOBIN := $(BIN)
export PATH := $(BIN):$(PATH)


$(BIN)/oapi-codegen: tools.go go.mod go.sum
	go get github.com/deepmap/oapi-codegen/cmd/oapi-codegen
	$(GO) install github.com/deepmap/oapi-codegen/cmd/oapi-codegen

.PHONY: api
api: $(BIN)/oapi-codegen
	$(BIN)/oapi-codegen -config ./api/config.yaml ./api/api.yaml > ./internal/api/api.gen.go

.PHONY: run
run:
	$(GO) run ./cmd/main.go