FROM golang:1.24-alpine AS builder

WORKDIR /app

COPY . .

RUN go mod tidy && \
    go mod vendor

CMD go test -v ./...