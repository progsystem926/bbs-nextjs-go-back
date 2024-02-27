FROM golang:1.22.0-alpine3.19 as dev

ENV CGO_ENABLED 0
ENV GO111MODULE auto

RUN apk --no-cache add git

WORKDIR /app
COPY . /app

RUN go install github.com/go-delve/delve/cmd/dlv@v1.22.1 && \
    go install github.com/pressly/goose/v3/cmd/goose@v3.18.0 && \
    go install go.uber.org/mock/mockgen@v0.4.0 && \
    go install github.com/cosmtrek/air@v1.51.0

FROM golang:1.22.0-alpine3.19 as builder

RUN apk --no-cache add gcc musl-dev

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -trimpath -ldflags "-w -s" -o ./main cmd/main.go

FROM alpine:3.19.1 as runner

RUN apk update && \
    apk add --no-cache shadow && \
    useradd -m appuser && \
    rm -f /usr/bin/gpasswd /usr/bin/passwd /usr/bin/chfn /sbin/unix_chkpwd /usr/bin/expiry /usr/bin/chage /usr/bin/chsh && \
    rm -rf /var/cache/apk/*

USER appuser

WORKDIR /app
COPY --from=builder /app/main .

CMD ["./main"]