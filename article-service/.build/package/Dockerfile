# syntax=docker/dockerfile:1
ARG GO_VERSION=1.25.1

FROM --platform=linux/arm64 golang:${GO_VERSION}-alpine as build

WORKDIR /go/src/github.com/miyamo2/blogapi.miyamo.today/article-service

RUN apk add --no-cache git

COPY . .
RUN go mod download

RUN go build -ldflags="-s -w" -trimpath -o ./bin/article-service ./cmd/main.go

FROM --platform=linux/arm64 alpine

WORKDIR /app

COPY --from=build /go/src/github.com/miyamo2/blogapi.miyamo.today/article-service/bin/article-service /app/article-service

CMD ["/app/article-service"]
