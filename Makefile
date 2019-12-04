.PHONY: build
build:
	@echo "Building application..."
	go build -v ./cmd/api-server

.PHONY: test
test:
	@echo "Run tests"
	go test -v -race -cover -timeout 30s ./...

.PHONY: lint
lint:
	@echo "Run lint"
	golint ./...

.DEFAULT_GOAL := build