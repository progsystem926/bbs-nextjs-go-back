FROM golang:alpine

ENV CGO_ENABLED 0
ENV GO111MODULE auto

RUN apk --no-cache add git

WORKDIR /app
COPY . /app

RUN go install github.com/go-delve/delve/cmd/dlv@latest && \
    go install github.com/pressly/goose/v3/cmd/goose@latest && \
    go install github.com/golang/mock/mockgen@latest && \
    go install github.com/cosmtrek/air@latest