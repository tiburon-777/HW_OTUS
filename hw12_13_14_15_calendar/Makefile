cdir = $(shell pwd)

integration-test:
	go test -count 10 -race ./cmd/calendar/... --config ./../../configs/calendar.conf

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

sheduler-up:
	docker-compose -f ./cicd/dc_sheduler.yml up -d --build
sheduler-down:
	docker-compose -f ./cicd/dc_sheduler.yml down

sup-up:
	docker-compose -f ./cicd/dc_sup.yml up -d --build
sup-down:
	docker-compose -f ./cicd/dc_sup.yml down

clean:
	sudo docker rmi $(sudo docker images | grep '<none>' | awk '{print $3}')

.PHONY: build run test lint