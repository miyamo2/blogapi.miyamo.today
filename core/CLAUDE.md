# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Overview

`blogapi.miyamo.today/core` is the **shared library** consumed by every Go microservice in the `blogapi.miyamo.today` monorepo (`article-service`, `blogging-event-service`, `tag-service`, `federator`, `read-model-updater`). It has no `main.go` — it is released as a versioned library. The consumers import it via `blogapi.miyamo.today/core`, `.../core/echo`, `.../core/graphql`, `.../core/grpc`.

## Four independent modules, not one

This directory contains **four separate Go modules**, each with its own `go.mod`, `go.sum`, `CHANGELOG.md`, CI workflow, and release tag:

| Module path                      | Import path                         | Release tag prefix  |
| -------------------------------- | ----------------------------------- | ------------------- |
| `./`                             | `blogapi.miyamo.today/core`         | `core/vX.Y.Z`       |
| `./echo`                         | `blogapi.miyamo.today/core/echo`    | `core/echo/vX.Y.Z`  |
| `./graphql`                      | `blogapi.miyamo.today/core/graphql` | `core/graphql/vX.Y.Z` |
| `./grpc`                         | `blogapi.miyamo.today/core/grpc`    | `core/grpc/vX.Y.Z`  |

The framework sub-modules (`echo`, `graphql`, `grpc`) depend on the root module through `require blogapi.miyamo.today/core vX.Y.Z` — not a replace directive. Consequence: a change to the root API that the sub-modules consume cannot be tested in-repo by the sub-modules until the root is released and the sub-modules bump. When editing root types used by sub-modules, change the sub-module together and bump both CHANGELOGs.

CI (`.github/workflows/ci_core*.yaml`) watches path filters so that `ci_core` fires only on non-sub-module changes (`core/**` excluding `core/echo/**`, `core/graphql/**`, `core/grpc/**`). Release workflows fire only on `*/CHANGELOG.md` edits on `main`; the version is parsed from the first `## X.Y.Z` heading.

## Common commands

Run from the specific module directory (root, `echo/`, `graphql/`, or `grpc/`) — `go` commands are scoped per-module.

```bash
# Build + tidy (CI step)
go mod tidy
go build -v ./...

# Lint (CI step)
go install honnef.co/go/tools/cmd/staticcheck@latest
staticcheck ./...

# Full test with coverage — matches CI's package filter
go test $(go list ./... | grep -v "mock") -v -coverprofile=coverage.out

# Root module release test uses -p 1 (serialize packages)
# Required because db/gorm uses global singletons (conn.Instance, dial.Instance)
# that races would corrupt across parallel packages.
go test $(go list ./... | grep -v "mock") -v -p 1 -coverprofile=coverage.out

# Regenerate mocks for sibling services (see `go:generate` directives in db/*.go)
go install go.uber.org/mock/mockgen@latest
go generate ./...

# Run a single test or subtest
go test ./db/gorm -run TestTransaction_Commit -v
go test ./db -run TestMultipleStatementResult_Get/happy_path -v
```

CI runs on `ubuntu-24.04-arm`; keep code ARM64-compatible and avoid cgo-only deps.

## Architecture

### The request-context pipeline

Every framework sub-module (`echo`, `graphql`, `grpc`) exposes the same three-middleware pipeline, applied in this order:

1. **`SetBlogAPIContextToContext`** — reads the incoming request headers, pulls `x-request-id` (gRPC uses `request_id` metadata) or mints a ULID, and stores a `blogapictx.BlogAPIContext` into `context.Context`. The `BlogAPIContext` carries request ID, service name, path, headers, and body — and is the source of truth consumed by the log handler.
2. **`SetLoggerToContext(app)`** — constructs a `*slog.Logger` wired with `altnrslog.NewTransactionalHandler` so that log output is associated with the current New Relic transaction, then stashes it with `altnrslog.StoreToContext`.
3. **`NRConnect` / `StartNewRelicTransaction`** — binds a New Relic transaction onto the context. (gRPC doesn't have this explicitly; gRPC instrumentation is handled by the New Relic gRPC integration upstream.)

Downstream code retrieves the logger via `altnrslog.FromContext(ctx)` with a fallback to `log.DefaultLogger()`. This pattern (`logger, err := altnrslog.FromContext(ctx); if err != nil { logger = log.DefaultLogger() }`) appears throughout and is the convention — do not replace it with `slog.Default()`.

### Log handler chain

`log.New()` returns `slog.New(BlogAPILogHandler{handler: slog.NewJSONHandler(...)})`. `BlogAPILogHandler.Handle` reads `BlogAPIContext` from ctx and injects `request_id`, `in_request`, and (if set) `out_request` into every record. `WithAltNRSlogTransactionalHandler(app, nrtx)` wraps that handler with altnrslog's transactional handler by swapping the inner JSON handler through `altnrslog.WithInnerHandlerProvider`. Effectively: `altnrslog → BlogAPILogHandler → JSONHandler`. The JSON handler options (`internal.JSONHandlerOption`) pin `AddSource: true` and strip paths from source entries to basenames.

### `db` and `db/gorm` — transactions as goroutines

`db.Transaction` is an interface, but the `db/gorm.Transaction` implementation is unusual: `GetAndStart` launches a goroutine (`Transaction.process`) that owns the `*gorm.DB` and services four channels — `stmtQueue`, `commit`, `rollback`, `errQueue`. `ExecuteStatement` ships the statement through `stmtQueue` with a per-call error channel; `Commit`/`Rollback` send to the respective channel and the goroutine breaks its loop. The goroutine runs on a New-Relic-cloned context (`nrtx.NewGoroutine()`) with its own logger, so transaction work shows up as a separate NR segment tree.

Consequence: a `Transaction` must receive exactly one terminal signal (`Commit` or `Rollback`); after either, all channels are closed and further calls will panic on send. `SubscribeError()` returns `errQueue` — consumers should drain it, not just check the return of `ExecuteStatement`.

`db/gorm` holds two package-global singletons: `conn.Instance` (the `*gorm.DB`) and `dial.Instance` (the `*gorm.Dialector`) each guarded by its own `sync.RWMutex`. `Initialize` / `InitializeDialector` are one-shot setters (second call logs a warning and no-ops); `Invalidate*` clears them. Tests set/reset these directly — that's why the root module is tested with `-p 1`.

Optional `gorm.io/plugin/dbresolver` source selection is plumbed through `db.GetAndStartWithDBSource(source)` → `conn.Clauses(dbresolver.Use(source))` inside `Transaction.process`.

### `db.Statement` generic results

Statements are functions `func(ctx, *gorm.DB, db.StatementResult) error` wrapped by `gorm.NewStatement`. Results flow through `SingleStatementResult[T]` / `MultipleStatementResult[T]` — generic wrappers that also expose a type-safe `StrictGet() T` / `StrictGet() []T`. `Set` is idempotent (second call is a silent no-op), so re-executing a statement won't clobber prior results. `Execute` itself enforces once-only semantics via `executed` flag (`ErrAlreadyExecuted`).

### Mockgen fan-out

`db/transaction.go` and `db/statement.go` carry `//go:generate mockgen -source=$GOFILE -destination ../../{service}/internal/mock/core/db/...` directives — running `go generate ./...` in this module regenerates mocks that live inside the sibling service directories (`tag-service`, `article-service`, `blogging-event-service`). Releasing the root module won't regenerate consumer mocks automatically; those services must re-run `go generate` after bumping the `core` dependency.

### Sub-module specifics

- **`echo/middlewares`** — adds `Auth(verifier Verifier)` which expects a `Bearer` header and stores the parsed `jwt.Token` under `JWTContextKey{}`. JWT parsing uses `lestrrat-go/jwx/v3` (v3 API; do not confuse with v2 — the module was bumped at echo v0.6.0). `NRConnect` sets `newrelic.WebRequest{Type: "ConnectRPC"}` — this is tuned for Connect-RPC over Echo/h2c, not plain REST.
- **`echo/s11n`** — generic `JSONSerializer[E, D]` is parameterized over encoder/decoder interfaces so callers can swap in `goccy/go-json` without rewriting Echo plumbing.
- **`graphql/model`** — `MarshalTime` / `UnmarshalTime` for `synchro.Time[tz.UTC]` (Code-Hex/synchro). All timestamps across the monorepo are `synchro.Time[tz.UTC]`, not stdlib `time.Time`; wire these up in `gqlgen.yml` `models:` mappings.
- **`grpc/interceptor`** — `SetBlogAPIContextToContext` and `SetLoggerToContext(app)` match the echo/graphql equivalents; chain them as `UnaryInterceptor`s. Request ID lookup uses metadata key `"request_id"` (snake_case), not `"x-request-id"`.

## Conventions

- **Errors:** use `github.com/cockroachdb/errors` (`errors.Wrap`, `errors.New`) throughout — not stdlib `errors` or `pkg/errors`. Stack traces propagate through New Relic's `nrpkgerrors` wrapper in downstream services.
- **New Relic segments:** wrap notable operations with `defer newrelic.FromContext(ctx).StartSegment("BlogAPICore: <op>").End()`. The `"BlogAPICore: "` prefix is a monorepo-wide naming convention that dashboards filter on.
- **Logger:** never call `slog.Default()` in library code — always go through `altnrslog.FromContext` with `log.DefaultLogger()` fallback.
- **Context:** when spawning goroutines that outlive the request, clone the New Relic transaction (`nrtx.NewGoroutine()`) and reinstall a fresh altnrslog logger (see `manager.GetAndStart` for the pattern).
