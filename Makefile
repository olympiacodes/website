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
	env GOOS=darwin GOARCH=amd64 go build -o build/hackbellingham-darwin_amd64
	env GOOS=darwin GOARCH=386   go build -o build/hackbellingham-darwin_386

linux: bindata-assetfs
	env GOOS=linux GOARCH=amd64 go build -o build/hackbellingham-linux_amd64
	env GOOS=linux GOARCH=386   go build -o build/hackbellingham-linux_386

freebsd: bindata-assetfs
	env GOOS=freebsd GOARCH=amd64 go build -o build/hackbellingham-freebsd_amd64
	env GOOS=freebsd GOARCH=386   go build -o build/hackbellingham-freebsd_386

all: osx linux freebsd