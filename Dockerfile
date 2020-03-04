FROM golang:1.13-alpine3.10 as build

RUN apk add --update --no-cache git tzdata && \
    mkdir /app

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

WORKDIR /app/cmd

RUN go build

#FROM alpine:3.10

#RUN apk add --update --no-cache ca-certificates
#WORKDIR /app
#COPY --from=build /usr/share/zoneinfo /usr/share/zoneinfo
#COPY --from=build /app/cmd /app/cmd
#COPY migrations /app/migrations

EXPOSE 3000

RUN ["./cmd"]

# vim: set ft=dockerfile: