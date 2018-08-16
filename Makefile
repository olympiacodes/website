PKG := github.com/bellinghamcodes/website
GO_VERSION := 1.10.3
DEP_VERSION := 0.5.0
COMMIT := $(strip $(shell git rev-parse --short HEAD))
VERSION := $(shell git describe --always --dirty)
SRC := "/go/src/$(PKG)"
.PHONY: help docker update-ca dev run generate dep
.DEFAULT_GOAL := help

# Run go in Docker container
DOCKER = docker run -it --rm \
		-p 3000:3000 \
		-v "$(PWD)":$(SRC) \
		golang:$(GO_VERSION) \
		bash -c 

dep: ## Install dependencies
	$(DOCKER) "curl -o /usr/local/bin/dep -L https://github.com/golang/dep/releases/download/v$(DEP_VERSION)/dep-linux-amd64 && chmod a+x /usr/local/bin/dep && cd $(SRC) && dep ensure"

dev: ## Run in development mode
	$(DOCKER) "cd $(SRC) && go build -tags=dev -o /tmp/bellingham-codes-website . && /tmp/bellingham-codes-website"

run: generate ## Run in production mode
	$(DOCKER) "cd $(SRC) && go build -o /tmp/bellingham-codes-website . && /tmp/bellingham-codes-website"

docker: ## Builds docker image
	docker build \
		--build-arg VERSION=$(VERSION) \
		--build-arg VCS_REF=$(COMMIT) \
		--build-arg BUILD_DATE=`date -u +"%Y-%m-%dT%H:%M:%SZ"` \
		-t tantalic/bellinghamcodes-website:$(VERSION) .

vet: ## Run tests
	$(DOCKER) "go vet $(shell go list ${PKG}/... | grep -v /vendor/)"

update-ca: ## Fetches latest root certificates 
	curl --time-cond certs/ca-certificates.crt -o certs/ca-certificates.crt https://curl.haxx.se/ca/cacert.pem 

generate: ## Create/update code generated files
	$(DOCKER) "cd $(SRC) && go generate"

help: ## List available make commands
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

