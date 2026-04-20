# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Overview

`article-service` is a Connect-RPC / gRPC microservice in the `blogapi.miyamo.today` monorepo (siblings: `blogging-event-service`, `tag-service`, `federator`, `read-model-updater`, `core`). It serves read-side article queries (`GetArticleById`, `GetAllArticles`, `GetNextArticles`, `GetPrevArticles`) backed by CockroachDB/PostgreSQL. Connect handlers are mounted on Echo running in H2C mode.

## Common commands

Run from `article-service/`.

```bash
# Build (CI runs `go mod tidy` first, then `go build -v ./...`)
go build ./...

# Lint — CI uses staticcheck (honnef.co/go/tools/cmd/staticcheck@latest)
staticcheck ./...

# Full test run — matches CI's package filter (excludes generated code + DI)
go test $(go list ./... | grep -v "mock" | grep -v "infra/fw/gqlgen" | grep -v "configs" | grep -v "infra/grpc" | grep -v "infra/rdb/sqlc") -v

# Single test / suite
go test ./internal/if-adapter/controller/pb -run Test_ArticleServiceServerTestSuite -v
go test ./internal/app/usecase -run TestGetByID -v

# Run server locally (loads .env via godotenv; default PORT=8080)
go run ./cmd/main.go

# Regenerate SQL code from internal/infra/rdb/{schema,query}.sql
sqlc generate

# Regenerate protobuf/Connect stubs (requires .proto submodule populated)
go generate -tags protogen ./...   # see protogen.go; invokes `buf generate`

# Regenerate Wire DI after editing wire.go or provider/*
go run github.com/google/wire/cmd/wire ./internal/configs/di
```

Required runtime env: `COCKROACHDB_DSN`, `NEW_RELIC_CONFIG_APP_NAME`, `NEW_RELIC_CONFIG_LICENSE`; optional `PORT`.

## Architecture

Clean Architecture with Wire-driven DI. When adding a new RPC, all four layers must be updated together — the compile-time `var _ I = (*Impl)(nil)` checks in `internal/configs/di/provider/*.go` will break the build if any binding is missing.

Request flow (illustrated by `GetArticleById`):

1. **Controller** — `internal/if-adapter/controller/pb/pb.go` (`ArticleServiceServer`) implements `grpcconnect.ArticleServiceHandler`. Each method opens a New Relic segment, maps the Connect request to a DTO, calls the usecase, runs the output through a converter, and wraps errors via `nrpkgerrors.Wrap`.
2. **Usecase** — `internal/app/usecase/*.go` (`GetByID`, `ListAll`, `ListAfter`, `ListBefore`) depend on the `query.Queries` interface. Shared helpers in `share.go`: `tagDtoFromQueryModel` and the generic `getPage`.
3. **Query port** — `internal/app/usecase/query/query.go` declares the interface; it's satisfied by `*sqlc.Queries`, constructed via `sqlc.Prepare` at startup in `provider/query.go`.
4. **Infra** — `internal/infra/rdb/sqlc/**` is sqlc-generated from `internal/infra/rdb/{schema,query}.sql`. Presenter converters in `internal/if-adapter/presenter/pb/convert/` implement interfaces declared in `controller/pb/presenter/convert/`, mapping DTOs into the protobuf types under `internal/infra/grpc/`.

**Interface / implementation split.** The controller package owns the interfaces for both usecases (`controller/pb/usecase/*`) and presenters (`controller/pb/presenter/convert/*`). Concrete implementations live in `internal/app/usecase` and `internal/if-adapter/presenter/pb/convert` respectively, and are wired in `provider/usecase.go` / `provider/presenter.go`.

**Pagination convention.** List usecases over-fetch by `first+1` (or `last+1`) rows, then `getPage` trims to the requested page; `hasNext` / `stillExists` is derived from whether the extra row came back. Page size is clamped with `min(max(n, 1), 100)` — the 100 is an inline constant marked `// TODO: config`.

**Tag aggregation.** `query.sql` collapses tags into one JSON array per article via `jsonb_agg(...) FILTER(WHERE t.id IS NOT NULL)`. The `pg_catalog.json` column is mapped (via `sqlc.yaml` overrides) to `internal/infra/rdb/types.Tags`, which implements `sql.Scanner` using `goccy/go-json`. Timestamps are `synchro.Time[tz.UTC]` (aliased as `types.UTCTime`), **not** stdlib `time.Time`.

**Observability & errors.** Every usecase and controller method wraps its body in `defer newrelic.FromContext(ctx).StartSegment(...).End()`. The logger is pulled from context via `altnrslog.FromContext`, installed by `middlewares.SetLoggerToContext` in `configs/di/echo.go`. Use `github.com/cockroachdb/errors` (`WithStack`, `Is`, `WithMessage`) — not stdlib `errors` or `pkg/errors` — so stack traces survive New Relic wrapping.

## Generated code (do not hand-edit)

CI excludes these paths from tests:

- `internal/infra/grpc/**` — buf-generated from the `.proto/` git submodule (points at `proto.miyamo.today`). `buf.gen.yaml` pins `go_package` to the grpc package.
- `internal/infra/rdb/sqlc/**` — `sqlc generate` output; configured in `sqlc.yaml`. To change DB access, edit `internal/infra/rdb/query.sql` / `schema.sql` and regenerate.
- `internal/configs/di/wire_gen.go` — Wire output. Hand-edit `wire.go` (build tag `wireinject`) and the provider sets in `internal/configs/di/provider/`, then rerun Wire.

## Testing conventions

Tests use `stretchr/testify/suite` with `github.com/ovechkin-dm/mockio/v2/mock` (dot-imported as `. "github.com/ovechkin-dm/mockio/v2/mock"` so `Mock[T]`, `WhenDouble`, `AnyContext`, `Equal` are unqualified). `go.uber.org/mock` is only an indirect dep — prefer mockio. Use `s.T().Context()` for the suite context and construct times with `synchro.New[tz.UTC](...)`.

## Deployment

`.build/package/Dockerfile` targets `linux/arm64` (Go 1.25.1 alpine). CI runs on `ubuntu-24.04-arm`. Keep new code ARM64-compatible and avoid cgo-only deps.
