.PHONY: test

BUILD_NAME		= go-slim
BUILD_DATE		= $(shell date -u '+%Y%m%d%I%M%S')

GIT_COUNT 	    = $(shell git rev-list --all --count)
GIT_HASH        = $(shell git rev-parse --short HEAD)

RELEASE_VER     = v1.0.0
RELEASE_TAG     = $(RELEASE_VER)-$(GIT_COUNT)-$(GIT_HASH)
RELEASE_DIR 	= deploy

all: version

run:
	@go run ./cmd/server/main.go

build:
	@wire gen ./internal/build/
	@cp -rf configs $(RELEASE_DIR)
	@mkdir -p $(RELEASE_DIR)/web && cp -rf web/templates $(RELEASE_DIR)/web && cp -rf web/assets $(RELEASE_DIR)/web
	@set CGO_ENABLED=0 && set GOOS=linux && set GOARCH=amd64 && go build -ldflags "-w -s -X 'go-slim/internal/app.BuildTag=$(RELEASE_TAG)' -X 'go-slim/internal/app.BuildDate=$(BUILD_DATE)'" -o $(RELEASE_DIR)/server ./cmd/server/main.go

winbuild:
	@wire gen ./internal/build/
	@cp -rf configs $(RELEASE_DIR)
	@mkdir -p $(RELEASE_DIR)/web && cp -rf web/templates $(RELEASE_DIR)/web && cp -rf web/assets $(RELEASE_DIR)/web
	@export CGO_ENABLED=0 && export GOOS=linux && export GOARCH=amd64 && go build -ldflags "-w -s -X 'go-slim/internal/app.BuildTag=$(RELEASE_TAG)' -X 'go-slim/internal/app.BuildDate=$(BUILD_DATE)'" -o $(RELEASE_DIR)/server ./cmd/server/main.go

buildwin:
	@wire gen ./internal/build/
	@cp -rf configs $(RELEASE_DIR)
	@mkdir -p $(RELEASE_DIR)/web && cp -rf web/templates $(RELEASE_DIR)/web && cp -rf web/assets $(RELEASE_DIR)/web
	@export CGO_ENABLED=0 && export GOOS=windows && export GOARCH=amd64 && go build -ldflags "-w -s -X 'go-slim/internal/app.BuildTag=$(RELEASE_TAG)' -X 'go-slim/internal/app.BuildDate=$(BUILD_DATE)'" -o $(RELEASE_DIR)/server.exe ./cmd/server/main.go

wire:
	@wire gen ./internal/build/

swagger:
	@swag init -g ./cmd/server/main.go

errs:
	@go generate ./internal/errs/

air:
	@air

test:
	@cd ./test && go test -v

version:
	@echo "Version: $(RELEASE_TAG)"

gormgen:
	@go run ./cmd/cli/main.go gormgen