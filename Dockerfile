FROM golang:1.13-alpine3.10 as build

RUN apk add --update --no-cache git tzdata && \
    mkdir /app

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

WORKDIR /app/cmd

RUN go build

EXPOSE 3000

ENTRYPOINT ["./cmd"]

# vim: set ft=dockerfile: