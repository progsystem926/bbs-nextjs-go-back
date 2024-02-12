FROM golang:alpine

ENV CGO_ENABLED 0
ENV GO111MODULE auto

RUN apk --no-cache add git

WORKDIR /app
COPY . /app
