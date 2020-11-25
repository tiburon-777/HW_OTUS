FROM golang:1.14 as builder
RUN mkdir -p /app
WORKDIR /app
COPY . .
RUN go get -d ./cmd/scheduler/.
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o scheduler ./cmd/scheduler/.

FROM alpine:latest
ENV APP_RABBITMQ_ADDRESS=rabbitmq
ENV APP_RABBITMQ_PORT=5672
ENV APP_RABBITMQ_LOGIN=rabbit
ENV APP_RABBITMQ_PASS=rabbit
ENV APP_RABBITMQ_EXCHANGE=calendar
ENV APP_RABBITMQ_QUEUE=notifications
ENV APP_RABBITMQ_KEY=events
ENV APP_API_ADDRESS=calendar
ENV APP_API_PORT=50053
ENV APP_LOGGER_FILE=/calendar.log
ENV APP_LOGGER_LEVEL=INFO
WORKDIR /
COPY --from=builder /app/scheduler ./sbin
CMD ["scheduler"]
