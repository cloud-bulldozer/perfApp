# vim: ft=make ts=4

.PHONY: all build test clean help

SRC=$(shell find . -name *.go)
BINARY=perfApp-$(ARCH)
GOCMD=go
ARCH ?= amd64
GOOS ?= linux
GO_BUILD_RECIPE:=GOOS=$(GOOS) CGO_ENABLED=0 GOARCH=$(ARCH) go build
ENGINE=podman
REGISTRY=quay.io
ORG?=cloud-bulldozer
IMAGE=perfapp

# Container versioning
VERSION?=latest
TAG?=$(VERSION)-$(ARCH)
CONTAINER_NAME=$(REGISTRY)/$(ORG)/$(IMAGE):$(TAG)

all: build buildContainer pushContainer

build: build/${BINARY}

go-deps:
	go mod tidy
	go mod vendor

build/${BINARY}: $(SRC)
	@echo -e "\033[2mBuilding ${BINARY} for $(GOOS)-$(ARCH)\033[0m"
	mkdir -p build
	$(GO_BUILD_RECIPE) -o build/${BINARY} cmd/perfApp/perfApp.go

container: buildContainer pushContainer

buildContainer: build/$(BINARY) Containerfile
	@echo -e "\n\033[2mBuilding container $(CONTAINER_NAME)\033[0m"
	$(ENGINE) build --pull-always --arch=$(ARCH) --build-arg=ARCH=$(ARCH) -t $(CONTAINER_NAME) .

pushContainer:
	@echo -e "\n\033[2mPushing container $(CONTAINER_NAME)\033[0m"
	$(ENGINE) push $(CONTAINER_NAME)

run: build
	./build/${BINARY}
