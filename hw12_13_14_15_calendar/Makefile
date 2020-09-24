build:
	go build -o ./bin/calendar ./src

run:
	go run ./cmd/calendar/main.go -config ./configs/config.toml

test:
	go test -race ./internal/...

lint: install-lint-deps
	golangci-lint run .cmd/... ./internal/...

install-lint-deps:
	(which golangci-lint > /dev/null) || curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(shell go env GOPATH)/bin v1.30.0

.PHONY: build run test lint