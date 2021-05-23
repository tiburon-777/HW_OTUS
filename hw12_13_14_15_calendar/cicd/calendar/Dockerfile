FROM golang:1.14 as builder
RUN mkdir -p /app
WORKDIR /app
COPY . .
RUN go get -d ./cmd/calendar/.
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -o calendar ./cmd/calendar/.

FROM alpine:latest
ENV APP_GRPC_PORT $APP_GRPC_PORT
ENV APP_HTTP_PORT $APP_HTTP_PORT
ENV APP_API_PORT $APP_API_PORT
WORKDIR /
COPY --from=builder /app/calendar ./sbin
EXPOSE ${APP_GRPC_PORT}
EXPOSE ${APP_HTTP_PORT}
EXPOSE ${APP_API_PORT}
ENTRYPOINT ["calendar"]
