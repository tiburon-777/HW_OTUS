FROM golang:1.14 as builder
RUN mkdir -p /app
WORKDIR /app
COPY . .
RUN go get -d ./cmd/sender/.
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -o sender ./cmd/sender/.

FROM alpine:latest
ENV APP_RABBITMQ_ADDRESS=rabbitmq
ENV APP_RABBITMQ_PORT=5672
ENV APP_RABBITMQ_LOGIN=rabbit
ENV APP_RABBITMQ_PASS=rabbit
ENV APP_RABBITMQ_EXCHANGE=calendar
ENV APP_RABBITMQ_QUEUE=notifications
ENV APP_RABBITMQ_KEY=events
ENV APP_LOGGER_FILE=/logs/calendar.log
ENV APP_LOGGER_LEVEL=INFO
WORKDIR /
COPY --from=builder /app/sender ./sbin
ENTRYPOINT ["sender"]
