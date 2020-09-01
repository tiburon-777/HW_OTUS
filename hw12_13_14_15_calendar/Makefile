build:
	go build -o ./bin/calendar ./src

test:
	go test -race ./src/...

install-lint-deps:
	(which golangci-lint > /dev/null) || curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin v1.30.0

lint: install-lint-deps
	golangci-lint run ./...

.PHONY: build test lint
