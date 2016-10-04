build: bindata-assetfs
	go build

bindata-assetfs:
	go-bindata-assetfs assets

run:
	go-bindata-assetfs -debug assets
	go run *.go

install: bindata-assetfs
	go install

osx: bindata-assetfs
	env GOOS=darwin GOARCH=amd64 go build -o build/bellinghamcodes-darwin_amd64
	env GOOS=darwin GOARCH=386   go build -o build/bellinghamcodes-darwin_386

linux: bindata-assetfs
	env GOOS=linux GOARCH=amd64 go build -o build/bellinghamcodes-linux_amd64
	env GOOS=linux GOARCH=386   go build -o build/bellinghamcodes-linux_386

freebsd: bindata-assetfs
	env GOOS=freebsd GOARCH=amd64 go build -o build/bellinghamcodes-freebsd_amd64
	env GOOS=freebsd GOARCH=386   go build -o build/bellinghamcodes-freebsd_386

all: osx linux freebsd