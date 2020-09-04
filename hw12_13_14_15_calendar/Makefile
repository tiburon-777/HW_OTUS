build:
	go build -o ./bin/calendar ./src

run:
	go run ./src/calendar/main.go -config ./configs/config.toml

test:
	go test -race ./src/...

lint: install-lint-deps
	golangci-lint run ./...

install-lint-deps:
	go install github.com/golangci/golangci-lint/cmd/golangci-lint

.PHONY: build run test lint