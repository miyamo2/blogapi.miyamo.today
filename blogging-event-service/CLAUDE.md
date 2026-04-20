# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Overview

`blogging-event-service` is a Connect-RPC / gRPC microservice in the `blogapi.miyamo.today` monorepo (siblings: `article-service`, `tag-service`, `federator`, `read-model-updater`, `core`). It is the **write side** of the blogging domain: it accepts commands (`CreateArticle`, `UpdateArticleTitle`, `UpdateArticleBody`, `UpdateArticleThumbnail`, `AttachTags`, `DetachTags`, `UploadImage`) and persists them as events in a single DynamoDB table (`BLOGGING_EVENTS_TABLE_NAME`). `UploadImage` is a **client-streaming** RPC that writes to S3 and returns a CDN URL. The read-side sibling services query the projections produced by `read-model-updater`. Connect handlers are mounted on Echo running in H2C mode.

## Common commands

Run from `blogging-event-service/`.

```bash
# Build (CI runs `go mod tidy` first, then `go build -v ./...`)
go build ./...

# Lint — CI uses staticcheck (honnef.co/go/tools/cmd/staticcheck@latest)
staticcheck ./...

# Full test run — matches CI's package filter (excludes generated code + DI)
go test $(go list ./... | grep -v "mock" | grep -v "configs" | grep -v "infra/grpc") -v

# Single test / suite
go test ./internal/if-adapter/controller/pb -run TestBloggingEventServiceServer_CreateArticle -v
go test ./internal/app/usecase -run TestCreateArticle -v

# Run server locally (loads .env via godotenv; default PORT=8080)
go run ./cmd/main.go

# Regenerate mockgen mocks under internal/mock/** (mockgen must be on PATH)
go generate ./...

# Regenerate protobuf/Connect stubs (requires .proto submodule populated)
go generate -tags protogen ./...   # see protogen.go; invokes `buf generate`

# Regenerate Wire DI after editing wire.go or provider/*
go run github.com/google/wire/cmd/wire ./internal/configs/di
```

Required runtime env: `BLOGGING_EVENTS_TABLE_NAME` (DynamoDB table), `S3_BUCKET`, `CDN_HOST` (used to build the returned image URL), `NEW_RELIC_CONFIG_APP_NAME`, `NEW_RELIC_CONFIG_LICENSE`, plus standard AWS credential env (loaded via `config.LoadDefaultConfig`). Optional: `PORT`.

The `.proto/` directory is a git submodule pointing at `proto.miyamo.today` — run `git submodule update --init --recursive` from the repo root before regenerating stubs.

## Architecture

Clean Architecture with Wire-driven DI. When adding a new RPC, all four layers must be updated together — the compile-time `var _ I = (*Impl)(nil)` checks in `internal/configs/di/provider/*.go` (one per provider file) will break the build if any binding is missing.

Request flow (illustrated by `CreateArticle`):

1. **Controller** — `internal/if-adapter/controller/pb/pb.go` (`BloggingEventServiceServer`, constructed via functional options `With*Usecase` / `With*Converter`) implements `grpcconnect.BloggingEventServiceHandler`. Each method opens a New Relic segment, maps the Connect request to a DTO, calls the usecase, runs the output through a converter, and wraps errors via `nrpkgerrors.Wrap`.
2. **Usecase** — `internal/app/usecase/*.go` (`CreateArticle`, `UpdateArticleTitle`, `UpdateArticleBody`, `UpdateArticleThumbnail`, `AttachTags`, `DetachTags`, `UploadImage`). Each builds a `model.*Event` from the DTO, calls `BloggingEventService` (for event persistence) or `storage.Uploader` (for S3), and returns an out-DTO holding the generated `eventID` / `articleID`.
3. **Command port** — `internal/app/usecase/command/blogging_event.go` declares `BloggingEventService`; it is satisfied by `*dynamo.BloggingEventCommandService`. Each method returns a `db.Statement` (`blogapi.miyamo.today/core/db` + `core/db/gorm`) whose body is run inside `Execute(ctx)`. `ulid.Make` generates both event IDs and (for `CreateArticle`) the new article ID; `NewBloggingEventCommandService(nil)` uses the default generator.
4. **Infra** — `internal/infra/dynamo/blogging_event.go` defines one unexported GORM model per event variant (`bloggingEventCreateArticle`, `bloggingEventUpdateArticleTitle`, …), each with `(EventID, ArticleID)` as the composite primary key and `TableName()` resolving `BLOGGING_EVENTS_TABLE_NAME` at call time. Tag-bearing events use `sqldav.Set[string]` so DynamoDB stores them as a native String Set. For `UploadImage`, `internal/infra/s3/s3.go` uploads via `PutObject` to `S3_BUCKET` and returns `CDN_HOST/<name>` as the public URL. The generated stubs in `internal/infra/grpc/` (and `grpc/grpcconnect/`) are buf output.

**DB layer.** `provider/gorm.go` opens `sql.Open("godynamo", "")` after calling `godynamo.RegisterAWSConfig(*awsConfig)`, then wraps it with `dynmgrm.New(dynmgrm.WithConnection(db))`. `main.go` finishes wiring by calling `gwrapper.InitializeDialector(dependencies.GORMDialector)` and `nraws.AppendMiddlewares(&dependencies.AWSConfig.APIOptions, nil)` — keep this order (Dialector init + NR AWS middleware) when changing startup.

**Interface / implementation split.** The controller package owns the interfaces for both usecases (`controller/pb/usecase/*`) and presenters (`controller/pb/presenters/converter.go`). Concrete implementations live in `internal/app/usecase` and `internal/if-adapter/presenter/pb` respectively, and are wired in `provider/usecase.go` / `provider/presenter.go`. The command / storage ports live on the **app** side (`app/usecase/command`, `app/usecase/storage`) because they are called by usecases, not by controllers.

**Streaming RPC.** `BloggingEventServiceServer.UploadImage` consumes `connect.ClientStream[grpcgen.UploadImageRequest]` in a `for stream.Receive()` loop. The proto's `oneof value { Meta meta | bytes data }` means clients must send exactly one `Meta` frame (filename + content type) and any number of `data` frames; the server concatenates into `bytes.Buffer` before calling `UploadImage.Execute`. Do not add blocking work inside the receive loop.

**Observability & errors.** Every usecase, controller, presenter, and DynamoDB statement wraps its body in `defer newrelic.FromContext(ctx).StartSegment(...).End()`. The logger is pulled from context via `altnrslog.FromContext`, installed by `middlewares.SetLoggerToContext` in `configs/di/provider/echo.go` (together with `SetBlogAPIContextToContext(RequestTypeGRPC)` and `NRConnect`). Use `github.com/cockroachdb/errors` (`WithStack`, `Is`, `WithMessage`) — not stdlib `errors` or `pkg/errors` — so stack traces survive `nrpkgerrors.Wrap`. The `HTTPErrorHandler` in `provider/echo.go` funnels all errors through `nrpkgerrors.Wrap` + structured logging before delegating to Echo's default. JSON uses `goccy/go-json` via `s11n.JSONSerializer`.

## Generated code (do not hand-edit)

CI excludes these paths from tests:

- `internal/infra/grpc/**` and `internal/infra/grpc/grpcconnect/**` — buf-generated from the `.proto/` git submodule (points at `proto.miyamo.today`). `buf.gen.yaml` pins `go_package` to `blogapi.miyamo.today/blogging-event-service/internal/infra/grpc`.
- `internal/configs/di/wire_gen.go` — Wire output. Hand-edit `wire.go` (build tag `wireinject`) and the provider sets in `internal/configs/di/provider/`, then rerun Wire.
- `internal/mock/**` — `mockgen` output. Each source interface declares a `//go:generate mockgen -source=$GOFILE -destination=...` directive; run `go generate ./...` to refresh after editing an interface.

## Testing conventions

Unlike `article-service` / `tag-service` (which use `ovechkin-dm/mockio`), this module uses **`go.uber.org/mock`** (gomock) with MockGen-generated mocks under `internal/mock/`. Import `go.uber.org/mock/gomock` and use `gomock.NewController(t)`; call `u.EXPECT().Method(gomock.Any(), ...)`. Table-driven tests (`tests := map[string]testCase{...}`) with `t.Run(name, ...)` are the dominant style. Use `github.com/google/go-cmp/cmp` with `cmpopts.IgnoreUnexported(...)` / `protocmp.Transform()` for DTO and protobuf comparisons.

## Deployment

`.build/package/Dockerfile` targets `linux/arm64` (Go 1.25.1 alpine, two-stage build, `-ldflags="-s -w" -trimpath`). CI runs on `ubuntu-24.04-arm` via the shared composite action `.github/custom_actions/ci/action.yaml`; deploy is triggered by `.github/workflows/deploy_blogging-event-service.yaml` on pushes to `main` under `blogging-event-service/**`. Keep new code ARM64-compatible and avoid cgo-only deps.
