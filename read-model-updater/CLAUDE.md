# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Overview

`read-model-updater` is an SQS-polling worker in the `blogapi.miyamo.today` monorepo (siblings: `article-service`, `blogging-event-service`, `tag-service`, `federator`, `core`). It consumes DynamoDB stream events forwarded into SQS, replays the full event history for the affected `article_id`, and rewrites the CQRS read models in **two independent CockroachDB instances** (article-DB and tag-DB). After a successful sync, it fires a GitHub `repository_dispatch` (`event_type: sync-read-model`) to rebuild the static blog.

This is a long-running process, not an HTTP server — there is no `PORT`, no routing layer.

## Common commands

Run from `read-model-updater/`.

```bash
# Build (CI runs `go mod tidy` first, then `go build -v ./...`)
go build ./...

# Lint — CI uses staticcheck (honnef.co/go/tools/cmd/staticcheck@latest)
staticcheck ./...

# Full test run — matches CI's package filter
go test $(go list ./... | grep -v "mock" | grep -v "infra/fw/gqlgen" | grep -v "configs" | grep -v "infra/grpc" | grep -v "infra/rdb/sqlc") -v

# Run the worker locally
go run ./cmd/main.go

# Regenerate sqlc output for BOTH article and tag packages (sqlc.yaml has two sql entries)
sqlc generate

# Regenerate Wire DI after editing wire.go or provider.go
go run github.com/google/wire/cmd/wire ./internal/configs/di
```

Required runtime env: `QUEUE_URL`, `COCKROACHDB_DSN_ARTICLE`, `COCKROACHDB_DSN_TAG`, `BLOGGING_EVENTS_TABLE_NAME`, `BLOG_PUBLISH_ENDPOINT`, `GITHUB_TOKEN`, `NEW_RELIC_CONFIG_APP_NAME`, `NEW_RELIC_CONFIG_LICENSE`. AWS credentials are resolved via the default chain; region is hardcoded to `ap-northeast-1` in `provideAWSConfig`.

## Architecture

Clean Architecture + Wire DI. Layering: `cmd/main.go` → `if-adapters/handler` → `if-adapters/converter` → `app/usecase` → `infra/{dynamo,rdb,githubactions,queue}`.

### Runtime loop (`cmd/main.go`)

Two goroutines share typed channels:

1. `polling` — calls `sqs.ReceiveMessage` every 1s with `MaxNumberOfMessages: 10` and `MessageAttributeNames: ["dynamodb.NewImage"]`, pushes each message onto `messageCh`. Errors go to `errCh` but do not stop the loop.
2. `work` — for each message, starts a New Relic transaction (`stream-worker`), parses the SQS `SentTimestamp` into a `synchro.Time[tz.UTC]` (this is the canonical `eventAt`, **not** the DynamoDB stream timestamp), wraps `SyncHandler.Invoke` and `sqs.DeleteMessage` in `retry.Do`, and runs them in a per-message goroutine spawned via `wg.Go`. A failure in `Invoke` skips `DeleteMessage` so the message returns to the queue.

### Sync usecase (`internal/app/usecase/sync.go`)

`Sync.SyncBlogSnapshotWithEvents` is the core operation. The input DTO carries fields from a single DynamoDB stream event, but the usecase **ignores most of them** — it reuses only `ArticleID` and `EventAt`, then re-reads the full history:

1. `BloggingEventQueryService.ListEventsByArticleID` queries DynamoDB via the `article_id_event_id-Index` GSI (see `infra/dynamo/blogging_event.go`), sorts rows by ULID event_id, and returns `[]model.BloggingEvent`.
2. `model.ArticleCommandFromBloggingEvents` folds events chronologically into one `ArticleCommand`: `title/content/thumbnail` take the last non-nil value; tags accumulate `attach_tags ∪ tags` and subtract `detach_tags` per event.
3. Tag IDs are **derived deterministically** from the tag name: `base64.URLEncoding.EncodeToString([]byte("tag:" + name))` (`model/article.go:newArticleTagCommands`). Do not mint random tag IDs.
4. Two `pgx` transactions are opened concurrently via `errgroup` (`IsoLevel: ReadCommitted`, 30s `context.WithTimeout`):
   - **article-DB tx:** `CreateTempTagsTable` → `PutArticle` (upsert) → `PreAttachTags` (COPY into `tmp_tags`) → `AttachTags` (upsert into `tags` and DELETE rows not in the new set).
   - **tag-DB tx:** `CreateTempArticlesTable` + `CreateTempTagsTable` → `PrePutTags`/`PutTags` → `PrePutArticle`/`PutArticle` (same staging + DELETE-missing pattern, keyed by `(id, tag_id)`).
5. `CreatedAt` is **not** `eventAt`; it comes from `ulid.Parse(articleCommand.ID()).Timestamp()` — the article's ID is treated as the authoritative creation timestamp. `UpdatedAt = eventAt`.
6. Only after both commits succeed does `BlogPublisher.Publish` POST to GitHub with a fresh ULID as `client_payload.event_id`.

### Two databases, two sqlc packages

The read model is denormalized across two CockroachDB instances for independent scaling of article and tag reads. `sqlc.yaml` declares two entries producing `internal/infra/rdb/sqlc/article` and `internal/infra/rdb/sqlc/tag`. They are **not interchangeable** — the `articles` table in each has a different primary key (`id` vs `(id, tag_id)`) and different columns. Schemas live at `internal/infra/rdb/schema.{article,tag}.sql`; queries at `query.{article,tag}.sql`.

### Staging via CockroachDB temp tables

All upserts use a `CREATE TEMP TABLE ... ON COMMIT PRESERVE ROWS` + `COPY FROM` + `INSERT ... SELECT DISTINCT ON ... ON CONFLICT DO UPDATE` pattern (see `query.article.sql:AttachTags`, `query.tag.sql:PutArticle`). This requires `experimental_enable_temp_tables=on`, which is set on both `pgxpool` configs in `provider.go:provideArticleDBPool` / `provideTagDBPool`. If you add new queries that use temp tables, they must run in the same transaction as their staging inserts (the temp table is per-connection).

### Conventions inherited from the monorepo

- Timestamps are `synchro.Time[tz.UTC]` (aliased as `internal/infra/rdb/types.UTCTime` via `sqlc.yaml` overrides). Do not use stdlib `time.Time` in DB-facing code.
- Errors: `github.com/cockroachdb/errors` (`WithStack`, `Wrap`). New Relic integration wraps via `nrpkgerrors.Wrap` at the boundary (`cmd/main.go`).
- Every usecase/handler/infra method opens `defer newrelic.FromContext(ctx).StartSegment(...).End()`. When forking goroutines inside an errgroup, call `nrtx.NewGoroutine()` and rebind with `newrelic.NewContext`.
- `BlogAPIContext` is attached at the worker boundary (`blogapictx.New(messageID, "", "queue", nil, nil)`) so downstream logs carry the SQS message ID.
- DynamoDB access goes through `pqxd` (a `database/sql` driver for DynamoDB) wrapped in `sqlx` — the query string is a literal SQL-like statement, not the AWS SDK API. See `infra/dynamo/blogging_event.go`.

## Generated code (do not hand-edit)

- `internal/infra/rdb/sqlc/article/**`, `internal/infra/rdb/sqlc/tag/**` — regenerate with `sqlc generate` after editing the corresponding `schema.*.sql` / `query.*.sql`. `sqlc.yaml` pins `emit_prepared_queries`, `emit_db_tags`, `omit_unused_structs`, and type overrides (`timestamptz → types.UTCTime`, `thumbnail → string`).
- `internal/configs/di/wire_gen.go` — Wire output. Hand-edit `wire.go` (build tag `wireinject`) and providers in `provider.go`, then rerun Wire.

## Deployment

`.build/package/Dockerfile` targets `linux/arm64` (Go 1.25.1 alpine). CI/deploy runs on `ubuntu-24.04-arm`. Keep new code ARM64-compatible and avoid cgo-only deps.
