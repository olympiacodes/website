PKG := github.com/bellinghamcodes/website
COMMIT := $(strip $(shell git rev-parse --short HEAD))
VERSION := $(shell git describe --always --dirty)
.PHONY: help docker update-ca bindata-assetfs
.DEFAULT_GOAL := help

run: ## Run in development mode
	go-bindata-assetfs -debug assets assets/*/**
	go run *.go

docker: bindata-assetfs ## Builds docker image
	docker build \
		--build-arg VERSION=$(VERSION) \
		--build-arg VCS_REF=$(COMMIT) \
		--build-arg BUILD_DATE=`date -u +"%Y-%m-%dT%H:%M:%SZ"` \
		-t tantalic/bellinghamcodes-website:$(VERSION) .

vet: ## Run tests
	go vet $(shell go list ${PKG}/... | grep -v /vendor/)

update-ca: ## Fetches latest root certificates 
	curl --time-cond certs/ca-certificates.crt -o certs/ca-certificates.crt https://curl.haxx.se/ca/cacert.pem 

bindata-assetfs:
	go-bindata-assetfs assets assets/*/**

help: ## List available make commands
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

