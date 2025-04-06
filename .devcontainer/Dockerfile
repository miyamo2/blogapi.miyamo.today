FROM --platform=linux/amd64 mcr.microsoft.com/devcontainers/go:1-1.24-bookworm

RUN apt-get update \
    && sudo apt install build-essential \
    && go install github.com/sqlc-dev/sqlc/cmd/sqlc@v1.28.0 \
    && go install github.com/99designs/gqlgen@v0.17.0 \
    && go install github.com/google/wire/cmd/wire@v0.6.0 \
    && go install github.com/joerdav/xc/cmd/xc@v0.8.5 \
    && go install honnef.co/go/tools/cmd/staticcheck@v0.6.1

WORKDIR /root/go/src/blogapi.miyamo.today