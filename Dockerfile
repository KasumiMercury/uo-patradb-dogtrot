FROM golang:1.21-alpine as builder

RUN apk update && apk --no-cache add gcc musl-dev git

WORKDIR /app

COPY go.mod ./
RUN go mod download

COPY . .

RUN go build -trimpath -ldflags "-w -s" -o ./main cmd/main.go

FROM alpine as runner

RUN apk update && \
    apk add --no-cache shadow && \
    useradd -m appuser && \
    rm -f /usr/bin/gpasswd /usr/bin/passwd /usr/bin/chfn /sbin/unix_chkpwd /usr/bin/expiry /usr/bin/chage /usr/bin/chsh && \
    rm -rf /var/cache/apk/*

USER appuser

WORKDIR /app
COPY --from=builder /app/main .

CMD ["./main"]

FROM golang:1.21-alpine as dev

ENV CGO_ENABLED 0
ENV GO111MODULE auto

RUN apk --no-cache add git bash

WORKDIR /app
COPY . ./app

RUN go install github.com/go-delve/delve/cmd/dlv@latest && \
    go install github.com/cosmtrek/air@latest