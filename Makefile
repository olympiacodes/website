.PHONY: help
.DEFAULT_GOAL := help

build: bindata-assetfs ## Build for current platform
	go build

bindata-assetfs:
	go-bindata-assetfs assets assets/*/**

run: ## Run in development mode
	go-bindata-assetfs -debug assets assets/*/**
	go run *.go

install: bindata-assetfs ## Build and install on current machine
	go install

linux: bindata-assetfs ## Builds linux binary
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -v -a -tags netgo -ldflags '-w' -o build/bellinghamcodes-linux-amd64

docker: linux ## Builds docker image
	docker build -t tantalic/bellinghamcodes-website:latest .

update-ca: ## Fetches latest root certificates 
	curl --time-cond certs/ca-certificates.crt -o certs/ca-certificates.crt https://curl.haxx.se/ca/cacert.pem 

help: ## List available make commands
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
