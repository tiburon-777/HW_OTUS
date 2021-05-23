FROM golang:1.14 as builder
RUN mkdir -p /app
WORKDIR /app
COPY . .
RUN go get -d ./cmd/sender/.
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -o sender ./cmd/sender/.

FROM alpine:latest
WORKDIR /
COPY --from=builder /app/sender ./sbin
ENTRYPOINT ["sender"]
