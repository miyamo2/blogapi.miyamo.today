# syntax=docker/dockerfile:1
ARG GO_VERSION=1.21.3

FROM golang:${GO_VERSION}-alpine3.18 as build

ARG GITHUB_TOKEN
ARG GRPC_HEALTH_PROBE_VERSION=v0.4.24

WORKDIR /go/src/github.com/miyamo2/blogapi-tag

RUN apk add --no-cache git

COPY . .
RUN export GOPRIVATE=github.com/miyamo2
RUN git config --global url."https://${GITHUB_TOKEN}:x-oauth-basic@github.com/".insteadOf "https://github.com/"
RUN go mod download

RUN go build -o ./bin/blogapi-tag ./cmd/main.go

RUN wget -qO/bin/grpc_health_probe https://github.com/grpc-ecosystem/grpc-health-probe/releases/download/${GRPC_HEALTH_PROBE_VERSION}/grpc_health_probe-linux-amd64 && \
    chmod +x /bin/grpc_health_probe

FROM alpine:3.18

WORKDIR /app

COPY --from=build /go/src/github.com/miyamo2/blogapi-tag/bin/blogapi-tag /app/blogapi-tag
COPY --from=build /bin/grpc_health_probe /bin/grpc_health_probe

CMD ["/app/blogapi-tag"]
