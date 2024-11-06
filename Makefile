BINARY_NAME=calltester

## help: Display available commands
.PHONY: help
help:
		@echo 'calltester Development:'
		@sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':' |  sed -e 's/^/ /'

## check: Code quality check
.PHONY: check
check:
		go mod verify
		go vet ./...

## tidy: Clean and tidy dependencies
.PHONY: tidy
tidy:
		go mod tidy

## test: Runs tests
.PHONY: test
test: 
		go test ./pkg/...

## build: Build the binary for the current platform
.PHONY: build
build:
	go build -o $(BINARY_NAME)

## build-all: Builds executable for all platforms
.PHONY: build-all
build-all: build-linux build-darwin build-windows

## build-linux: Build for Linux (x86-64)
.PHONY: build-linux
build-linux:
	GOOS=linux GOARCH=amd64 go build -o ./bin/$(BINARY_NAME)-linux-amd64

## build-darwin: Build for macOS (x86-64)
.PHONY: build-darwin
build-darwin:
	GOOS=darwin GOARCH=amd64 go build -o ./bin/$(BINARY_NAME)-darwin-amd64

## build-windows: Build for Windows (x86-64)
.PHONY: build-windows
build-windows:
	GOOS=windows GOARCH=amd64 go build -o ./bin/$(BINARY_NAME)-windows-amd64.exe
