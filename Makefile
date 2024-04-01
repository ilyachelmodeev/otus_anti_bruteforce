BIN := "./bin/anti-bruteforce"
BIN_CLI := "./bin/cli-anti-bruteforce"

GIT_HASH := $(shell git log --format="%h" -n 1)
LDFLAGS := -X main.release="develop" -X main.buildDate=$(shell date -u +%Y-%m-%dT%H:%M:%S) -X main.gitHash=$(GIT_HASH)

build:
	go build -v -o $(BIN) -ldflags "$(LDFLAGS)" ./cmd/anti-bruteforce
	go build -v -o $(BIN_CLI) -ldflags "$(LDFLAGS)" ./cmd/cli-anti-bruteforce

run: build
	ANTI_BRUTEFORCE_REDIS_URL=redis://localhost:6379 $(BIN)

test:
	go test -race -count 10 ./internal/... ./pkg/...

install-lint-deps:
	(which golangci-lint > /dev/null) || curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(shell go env GOPATH)/bin v1.57.2

lint: install-lint-deps
	golangci-lint run ./...

up:
	cd ./deployments && docker-compose up -d --build

down:
	cd ./deployments && docker-compose down

restart: down up

.PHONY: build run test lint