build: bindata-assetfs
	go build

bindata-assetfs:
	go-bindata-assetfs assets

run:
	go-bindata-assetfs -debug assets
	go run *.go

install: bindata-assetfs
	go install

linux: bindata-assetfs
	env GOOS=linux GOARCH=amd64 go build -o build/bellinghamcodes-linux-amd64

