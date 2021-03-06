FROM golang:1.13-alpine3.10 as build
ARG version=unknown

RUN apk add --update --no-cache git tzdata make && \
    mkdir /app

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

WORKDIR /app

RUN make VERSION=$version

FROM alpine:3.10

RUN apk add --update --no-cache ca-certificates
WORKDIR /app
COPY --from=build /usr/share/zoneinfo /usr/share/zoneinfo
COPY --from=build /app/bin/shupple2-api /app
COPY migrations /app/migrations

EXPOSE 3000

ENTRYPOINT ["./shupple2-api"]

# vim: set ft=dockerfile:
