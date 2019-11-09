.PHONY: build
build:
	go build -v ./cmd/clean-architecture-rest-api/main.go

.PHONY: test
test:
	go test -v -race -timeout 30s ./...

.PHONY: test-bench
test-bench:
	go test -bench=. ./...

.PHONY: test-cover
test-cover:
	go test -cover ./...

.DEFAULT_GOAL := build