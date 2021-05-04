cdir = $(shell pwd)

inplace-integration-test:
	go test -count 10 -race ./cmd/calendar/... --config ./../../configs/calendar.conf

bdd:
	docker build -t bdd --network host -f test/Dockerfile .

integration-test:
	set -e ;\
	docker-compose -f ./cicd/dc_sup.yml up -d --build ;\
	docker-compose -f ./cicd/dc_calendar.yml up -d --build ;\
	docker-compose -f ./cicd/dc_sender.yml up -d --build ;\
	docker-compose -f ./cicd/dc_sheduler.yml up -d --build ;\
	test_status_code=0 ;\
	docker build -t bdd --network host -f test/Dockerfile . || test_status_code=$$? ;\
	docker-compose -f ./cicd/dc_sup.yml down --rmi local --volumes --remove-orphans --timeout 60 ;\
	docker-compose -f ./cicd/dc_calendar.yml down --rmi local --volumes --remove-orphans --timeout 60 ;\
	docker-compose -f ./cicd/dc_sender.yml down --rmi local --volumes --remove-orphans --timeout 60 ;\
	docker-compose -f ./cicd/dc_sheduler.yml down --rmi local --volumes --remove-orphans --timeout 60 ;\
	exit $$test_status_code

test:
	go test -race ./internal/... ./pkg/...

lint: install-lint-deps
	golangci-lint run .cmd/... ./internal/...  ./pkg/...

install-lint-deps:
	(which golangci-lint > /dev/null) || curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(shell go env GOPATH)/bin v1.30.0

generate:
	protoc  -I ./api/public  --go_out=plugins=grpc:./pkg/api/public ./api/public/public.proto
	protoc  -I ./api/private  --go_out=plugins=grpc:./internal/api/private ./api/private/private.proto

up: sup-up calendar-up sender-up sheduler-up

down: calendar-down sender-down sheduler-down sup-down clean

calendar-up:
	docker-compose -f ./cicd/dc_calendar.yml up -d --build
calendar-down:
	docker-compose -f ./cicd/dc_calendar.yml down

sender-up:
	docker-compose -f ./cicd/dc_sender.yml up -d --build
sender-down:
	docker-compose -f ./cicd/dc_sender.yml down

scheduler-up:
	docker-compose -f ./cicd/dc_sheduler.yml up -d --build
scheduler-down:
	docker-compose -f ./cicd/dc_sheduler.yml down

sup-up:
	docker-compose -f ./cicd/dc_sup.yml up -d --build
sup-down:
	docker-compose -f ./cicd/dc_sup.yml down --rmi local -v

.PHONY: build run test lint