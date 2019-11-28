.PHONY: build
build:
	@echo "Building application..."
	go build -v ./cmd/api-server

.DEFAULT_GOAL := build