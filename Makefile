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

## build: Builds executable for Linux
.PHONY: build
build: 
		go build -o bin/calltester main.go 
