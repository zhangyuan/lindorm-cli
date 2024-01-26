.PHONY: build
build:
	go build

.PHONY: clean
clean:
	rm -rf lindorm-cli
	rm -rf bin/lindorm-cli_*

.PHONY: build-macos
build-macos:
	GOOS=darwin GOARCH=amd64 go build -o bin/lindorm-cli_darwin-amd64
	GOOS=darwin GOARCH=arm64 go build -o bin/lindorm-cli_darwin-arm64

.PHONY: build-linux
build-linux:
	GOOS=linux GOARCH=amd64 go build -o bin/lindorm-cli_linux-amd64

.PHONE: build-all
build-all: clean build-macos build-linux

.PHONE: compress
compress:
	upx ./bin/lindorm-cli_*
