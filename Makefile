.PHONY: build-linux
build-linux:
	GOOS=linux GOARCH=amd64 go build -o bin/lindorm-linux_amd64
