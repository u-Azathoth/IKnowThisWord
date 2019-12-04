.PHONY: build
build:
	@echo "Building application..."
	go build -v ./cmd/api-server

.PHONY: test
test:
	@echo "Run tests"
	go test -v -race -cover -timeout 30s ./...

testci:
	@echo "Run CI tests"
	go test -v -race -coverprofile=coverage.txt -covermode=atomic ./...

.DEFAULT_GOAL := build