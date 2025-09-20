# syntax=docker/dockerfile:1
ARG GO_VERSION=1.25.1

FROM golang:${GO_VERSION}-alpine as build

WORKDIR /go/src/github.com/miyamo2/blogapi.miyamo.today/blogging-event-service

RUN apk add --no-cache git

COPY . .
RUN go mod download

RUN go build -ldflags="-s -w" -trimpath -o ./bin/blogging-event-service ./cmd/main.go

FROM alpine

WORKDIR /app

COPY --from=build /go/src/github.com/miyamo2/blogapi.miyamo.today/blogging-event-service/bin/blogging-event-service /app/blogging-event-service

CMD ["/app/blogging-event-service"]
