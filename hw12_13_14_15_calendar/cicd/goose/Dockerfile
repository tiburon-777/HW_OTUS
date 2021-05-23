FROM golang:1.14
WORKDIR /
RUN go get -u github.com/pressly/goose/cmd/goose
COPY . .
ENTRYPOINT ["/go/bin/goose", "-dir", "migrations", "up"]
