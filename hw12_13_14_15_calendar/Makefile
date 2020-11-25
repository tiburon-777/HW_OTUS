cdir = $(shell pwd)

build-all:
	go build -o ./bin/calendar ./cmd/calendar/main.go
	go build -o ./bin/scheduler ./cmd/scheduler/main.go
	go build -o ./bin/sender ./cmd/sender/main.go

start: build-all rabbit-start calendar-start scheduler-start sender-start

stop: scheduler-stop sender-stop calendar-stop rabbit-stop

test:
	go test -race ./internal/... ./pkg/... --tags calendar

lint: install-lint-deps
	golangci-lint run .cmd/... ./internal/...  ./pkg/...

install-lint-deps:
	(which golangci-lint > /dev/null) || curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(shell go env GOPATH)/bin v1.30.0

generate:
	protoc  -I ./api/public  --go_out=plugins=grpc:./pkg/api/public --grpc-gateway_out=logtostderr=true:./pkg/api/public ./api/public/public.proto
	protoc  -I ./api/private  --go_out=plugins=grpc:./internal/api/private ./api/private/private.proto

compose-build:
	sudo -S docker-compose -f ./cicd/docker-compose.yml build

compose-up:
	sudo -S docker-compose -f ./cicd/docker-compose.yml up -d

compose-down:
	sudo -S docker-compose -f ./cicd/docker-compose.yml down

calendar-start:
	nohup  ./bin/calendar -config ./configs/calendar.conf > /dev/null 2>&1 &
calendar-stop:
	kill -9 `pgrep calendar`

scheduler-start:
	nohup  ./bin/scheduler -config ./configs/scheduler.conf > /dev/null 2>&1 &
scheduler-stop:
	kill -9 `pgrep scheduler`

sender-start:
	nohup  ./bin/sender -config ./configs/sender.conf > /dev/null 2>&1 &
sender-stop:
	kill -9 `pgrep sender`




.PHONY: build run test lint