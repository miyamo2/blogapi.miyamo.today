# syntax=docker/dockerfile:1
ARG GO_VERSION=1.24.0

FROM --platform=linux/amd64 golang:${GO_VERSION}-alpine3.20 as build

ARG GRPC_HEALTH_PROBE_VERSION=v0.4.24

WORKDIR /go/src/github.com/miyamo2/blogapi.miyamo.today/article-service

RUN apk add --no-cache git

COPY . .
RUN go mod download

RUN go build -ldflags="-s -w" -trimpath -o ./bin/article-service ./cmd/main.go

FROM alpine:3.21

WORKDIR /app

COPY --from=build /go/src/github.com/miyamo2/blogapi.miyamo.today/article-service/bin/article-service /app/article-service

CMD ["/app/article-service"]
