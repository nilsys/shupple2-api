FROM golang:1.13-alpine3.10 as build

RUN apk add --update --no-cache git tzdata make && \
    mkdir /app

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

WORKDIR /app

RUN make

FROM alpine:3.10

RUN apk add --update --no-cache ca-certificates
WORKDIR /app
COPY --from=build /usr/share/zoneinfo /usr/share/zoneinfo
COPY --from=build /app/bin/stayway-media-api /app
COPY --from=build /app/bin/stayway-media-batch /app
COPY migrations /app/migrations

EXPOSE 3000

ENTRYPOINT ["./stayway-media-api"]

# vim: set ft=dockerfile:
