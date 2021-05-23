FROM golang:1.14 as builder
RUN mkdir -p /app
WORKDIR /app
COPY . .
RUN go get -d ./cmd/scheduler/.
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -o scheduler ./cmd/scheduler/.

FROM alpine:latest
WORKDIR /
COPY --from=builder /app/scheduler ./sbin
ENTRYPOINT ["scheduler"]
