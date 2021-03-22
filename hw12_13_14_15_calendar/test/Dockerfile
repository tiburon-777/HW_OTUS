FROM golang:1.14
RUN mkdir -p /app
WORKDIR /app
COPY . .
RUN go get -d /app/test/.
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go test ./test/...