PKG := github.com/bellinghamcodes/website
COMMIT := $(strip $(shell git rev-parse --short HEAD))
VERSION := $(shell git describe --always --dirty)
.PHONY: help docker update-ca dev run generate
.DEFAULT_GOAL := help

dev: ## Run in development mode
	go build -tags=dev -o /tmp/bellingham-codes-website . && /tmp/bellingham-codes-website

run: generate ## Run in production mode
	go build -o /tmp/bellingham-codes-website . && /tmp/bellingham-codes-website

docker: ## Builds docker image
	docker build \
		--build-arg VERSION=$(VERSION) \
		--build-arg VCS_REF=$(COMMIT) \
		--build-arg BUILD_DATE=`date -u +"%Y-%m-%dT%H:%M:%SZ"` \
		-t tantalic/bellinghamcodes-website:$(VERSION) .

vet: ## Run tests
	go vet $(shell go list ${PKG}/... | grep -v /vendor/)

update-ca: ## Fetches latest root certificates 
	curl --time-cond certs/ca-certificates.crt -o certs/ca-certificates.crt https://curl.haxx.se/ca/cacert.pem 

generate: ## Create/update code generated files
	go generate

help: ## List available make commands
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

