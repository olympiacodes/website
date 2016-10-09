build: bindata-assetfs
	go build

bindata-assetfs:
	go-bindata-assetfs assets assets/*/**

run:
	go-bindata-assetfs -debug assets assets/*/**
	go run *.go

install: bindata-assetfs
	go install

linux: bindata-assetfs
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -v -a -tags netgo -ldflags '-w' -o build/bellinghamcodes-linux-amd64

docker: linux
	docker build -t tantalic/bellinghamcodes-website:latest .

update-ca:
	curl --time-cond certs/ca-certificates.crt -o certs/ca-certificates.crt https://curl.haxx.se/ca/cacert.pem 
