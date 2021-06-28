PACKAGES=$(shell go list ./... | grep -v '/simulation')

VERSION := $(shell echo $(shell git describe --tags) | sed 's/^v//')
COMMIT := $(shell git log -1 --format='%H')

ldflags = -X github.com/cosmos/cosmos-sdk/version.Name=BlockchainTest \
	-X github.com/cosmos/cosmos-sdk/version.ServerName=blockchain-testd \
	-X github.com/cosmos/cosmos-sdk/version.ClientName=blockchain-testcli \
	-X github.com/cosmos/cosmos-sdk/version.Version=$(VERSION) \
	-X github.com/cosmos/cosmos-sdk/version.Commit=$(COMMIT)

BUILD_FLAGS := -ldflags '$(ldflags)'

include Makefile.ledger
all: install

install: go.sum
		@echo "--> Installing blockchain-testd"
		@go1.15.5 install -mod=readonly $(BUILD_FLAGS) ./cmd/blockchain-testd
		@echo "--> Installing blockchain-testcli"
		@go1.15.5 install -mod=readonly $(BUILD_FLAGS) ./cmd/blockchain-testcli

go.sum: go.mod
		@echo "--> Ensure dependencies have not been modified"
		GO111MODULE=on go mod verify

test:
	@go test -mod=readonly $(PACKAGES)
