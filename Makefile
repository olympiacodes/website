PKG := github.com/bellinghamcodes/website
GOVERSION := 1.9.0
VERSION := $(shell git describe --always --dirty)
.PHONY: help docker build linux install update-ca bindata-assetfs
.DEFAULT_GOAL := help

run: ## Run in development mode
	go-bindata-assetfs -debug assets assets/*/**
	go run *.go

docker: linux ## Builds docker image
	docker build -t tantalic/bellinghamcodes-website:$(VERSION) .

vet: ## Run tests
	go vet $(shell go list ${PKG}/... | grep -v /vendor/)

build: bindata-assetfs ## Build for current platform
	go build

linux: bindata-assetfs
	docker run --rm -v "${PWD}":/go/src/$(PKG) -w /go/src/$(PKG) -e "CGO_ENABLED=0" -e "GOOS=linux" -e "GOARCH=amd64" golang:$(GOVERSION) go build -v -a -tags netgo -ldflags '-w -X main.version=$(VERSION)' -o build/bellinghamcodes-linux-amd64

install: bindata-assetfs ## Build and install on current machine
	go install

update-ca: ## Fetches latest root certificates 
	curl --time-cond certs/ca-certificates.crt -o certs/ca-certificates.crt https://curl.haxx.se/ca/cacert.pem 

bindata-assetfs:
	go-bindata-assetfs assets assets/*/**

help: ## List available make commands
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

