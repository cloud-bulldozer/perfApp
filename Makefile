# vim: ft=make ts=4

.PHONY: all build test clean help

BUILD_DATE=$(shell date '+%Y-%m-%d-%H:%M:%S')
GIT_COMMIT=$(shell git rev-parse HEAD)
GIT_DIRTY=$(shell test -n "`git status --porcelain`" && echo "+CHANGES" || true)
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test -v
GO_BUILD_RECIPE=GOOS=linux CGO_ENABLED=0 go build
ENGINE=podman
REGISTRY=quay.io
PROJECT=rsevilla
IMAGE=perfapp
TAG=latest
SRC = $(shell find . -name *.go)

all: build buildContainer pushContainer

build: build/perfApp

build/perfApp: $(SRC)
	@echo Building perfApp
	mkdir -p build
	$(GO_BUILD_RECIPE) -o build/perfApp cmd/perfApp/perfApp.go

container: buildContainer pushContainer

buildContainer: build/perfApp Containerfile
	$(ENGINE) build --pull-always -t $(REGISTRY)/$(PROJECT)/$(IMAGE):$(TAG) .

pushContainer:
	$(ENGINE) push $(REGISTRY)/$(PROJECT)/$(IMAGE):$(TAG)


