# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Overview

`tag-service` is a Connect-RPC / gRPC microservice in the `blogapi.miyamo.today` monorepo (siblings: `article-service`, `blogging-event-service`, `federator`, `read-model-updater`, `core`). It serves read-side tag queries (`GetTagById`, `GetAllTags`, `GetNextTags`, `GetPrevTags`) backed by CockroachDB/PostgreSQL. Each tag is returned with its articles eagerly joined. Connect handlers are mounted on Echo running in H2C mode.

## Common commands

Run from `tag-service/`.

```bash
# Build (CI runs `go mod tidy` first, then `go build -v ./...`)
go build ./...

# Lint — CI uses staticcheck (honnef.co/go/tools/cmd/staticcheck@latest)
staticcheck ./...

# Full test run — matches CI's package filter (excludes generated code + DI)
go test $(go list ./... | grep -v "mock" | grep -v "infra/fw/gqlgen" | grep -v "configs" | grep -v "infra/grpc" | grep -v "infra/rdb/sqlc") -v

# Single test / suite
go test ./internal/if-adapter/controller/pb -run Test_TagServiceServerTestSuite -v
go test ./internal/app/usecase -run TestGetByIdTestSuite -v

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

The `.proto/` directory is a git submodule pointing at `proto.miyamo.today` — run `git submodule update --init --recursive` from the repo root before regenerating stubs.

## Architecture

Clean Architecture with Wire-driven DI. When adding a new RPC, all four layers must be updated together — the compile-time `var _ I = (*Impl)(nil)` checks in `internal/configs/di/provider/*.go` (one per provider file) will break the build if any binding is missing.

Request flow (illustrated by `GetTagById`):

1. **Controller** — `internal/if-adapter/controller/pb/pb.go` (`TagServiceServer`) implements `grpcconnect.TagServiceHandler`. Each method opens a New Relic segment, maps the Connect request to a DTO, calls the usecase, runs the output through a converter, and wraps errors via `nrpkgerrors.Wrap`. Conversion failures return sentinel errors like `ErrConversionToGetTagByIdFailed`.
2. **Usecase** — `internal/app/usecase/*.go` (`GetById`, `ListAll`, `ListAfter`, `ListBefore`) depend on the `query.Queries` interface. Shared helpers in `share.go`: `articleDtoFromQueryModel` and the generic `getPage`.
3. **Query port** — `internal/app/usecase/query/queries.go` declares the interface; it's satisfied by `*sqlc.Queries`, constructed via `sqlc.Prepare` at startup in `provider/query_service.go` (3s timeout, panics on failure).
4. **Infra** — `internal/infra/rdb/sqlc/**` is sqlc-generated from `internal/infra/rdb/{schema,query}.sql`. Presenter converters in `internal/if-adapter/presenter/pb/convert/` implement interfaces declared in `controller/pb/presenter/convert/`, mapping DTOs into the protobuf types under `internal/infra/grpc/`.

**Interface / implementation split.** The controller package owns the interfaces for both usecases (`controller/pb/usecase/*`) and presenters (`controller/pb/presenter/convert/*`). Concrete implementations live in `internal/app/usecase` and `internal/if-adapter/presenter/pb/convert` respectively, and are wired in `provider/usecase.go` / `provider/presenter.go`.

**Data model.** `tags` has a one-to-many to `articles` (FK `articles.tag_id → tags.id`, with `(id, tag_id)` as composite PK on `articles`). `query.sql` collapses the joined articles into one JSON array per tag via `jsonb_agg(json_build_object(...)) FILTER(WHERE a.id IS NOT NULL)`. The `pg_catalog.json` column is mapped (via `sqlc.yaml` overrides) to `internal/infra/rdb/types.Articles`, which implements `sql.Scanner` using `goccy/go-json`. `articles.thumbnail` is overridden to Go `string` (schema is `VARCHAR(524271)`, nullable, but the code treats it as non-null). Timestamps are `synchro.Time[tz.UTC]` (aliased as `types.UTCTime`), **not** stdlib `time.Time`.

**Pagination convention.** `ListAfter` / `ListBefore` over-fetch by `first+1` (or `last+1`) rows, then `getPage` trims to the requested page; `hasNext` / `hasPrevious` is derived from whether the extra row came back. Page size is clamped with `min(max(n, 1), 100)` — the 100 is an inline constant marked `// TODO: config`. `ListBefore` orders by `id DESC`, so cursors compare with `<` instead of `>`.

**`ListAll` reuses `ListAfter` (no limit).** `list_all.go` calls `u.queries.ListAfter(ctx)` — the unbounded variant — rather than a dedicated `ListAll` SQL. If you need a true "all tags" query distinct from cursor pagination, keep that reuse in mind before renaming or removing SQL entries.

**Observability & errors.** Every usecase and controller method wraps its body in `defer newrelic.FromContext(ctx).StartSegment(...).End()`. The logger is pulled from context via `altnrslog.FromContext`, installed by `middlewares.SetLoggerToContext` in `configs/di/dependencies.go` (this file is the Echo/Connect wire-up — not `echo.go`). Use `github.com/cockroachdb/errors` (`WithStack`, `Is`, `WithMessage`) — not stdlib `errors` or `pkg/errors` — so stack traces survive New Relic wrapping. The `HTTPErrorHandler` in `dependencies.go` funnels all errors through `nrpkgerrors.Wrap` + structured logging before delegating to Echo's default.

**DB driver.** The SQL DB is a `*sql.DB` opened from a `pgxpool.Pool` (pgx v5) via `stdlib.OpenDBFromPool`, with `nrpgx5.NewTracer()` attached to the pool config for New Relic. Do not swap to `database/sql`'s `sql.Open("pgx", ...)` path — it loses the pool + tracer wiring.

## Generated code (do not hand-edit)

CI excludes these paths from tests:

- `internal/infra/grpc/**` — buf-generated from the `.proto/` git submodule (points at `proto.miyamo.today`). `buf.gen.yaml` pins `go_package` to `blogapi.miyamo.today/tag-service/internal/infra/grpc`.
- `internal/infra/rdb/sqlc/**` — `sqlc generate` output; configured in `sqlc.yaml`. To change DB access, edit `internal/infra/rdb/query.sql` / `schema.sql` and regenerate.
- `internal/configs/di/wire_gen.go` — Wire output. Hand-edit `wire.go` (build tag `wireinject`) and the provider sets in `internal/configs/di/provider/`, then rerun Wire.

## Testing conventions

Tests use `stretchr/testify/suite` with `github.com/ovechkin-dm/mockio/v2/mock` (dot-imported as `. "github.com/ovechkin-dm/mockio/v2/mock"` so `Mock[T]`, `WhenDouble`, `AnyContext`, `Exact`, `Equal` are unqualified). `go.uber.org/mock` is only an indirect dep — prefer mockio. Use `s.T().Context()` for the suite context and construct times with `synchro.New[tz.UTC](...)`.

## Deployment

`.build/package/Dockerfile` targets `linux/arm64` (Go 1.25.1 alpine, two-stage build, `-ldflags="-s -w" -trimpath`). CI runs on `ubuntu-24.04-arm` via the shared composite action `.github/custom_actions/ci/action.yaml`. Keep new code ARM64-compatible and avoid cgo-only deps.
