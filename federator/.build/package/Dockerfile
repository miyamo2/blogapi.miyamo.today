# syntax=docker/dockerfile:1
ARG GO_VERSION=1.23.1

FROM golang:${GO_VERSION}-alpine3.20 as build

ARG GITHUB_TOKEN

WORKDIR /go/src/github.com/miyamo2/blogapi.miyamo.today/federator

RUN apk update \
    && apk add --no-cache git

COPY . .
RUN export GOPRIVATE=github.com/miyamo2
RUN go mod download

RUN go build -ldflags="-s -w" -trimpath -o ./bin/federator ./cmd/main.go

FROM alpine:3.21

WORKDIR /app

RUN apk update \
    && apk add --upgrade curl

COPY --from=build /go/src/github.com/miyamo2/blogapi.miyamo.today/federator/bin/federator /app/federator

CMD ["/app/federator"]
