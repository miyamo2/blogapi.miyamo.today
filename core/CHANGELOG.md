# Changelog

## 0.24.0 - 2024-12-29

### 💥 Breaking Changes

- Package name is now `blogapi.miyamo.today/core` instead of `github.com/miyamo2/blogapi.miyamo.today/core`.

## 0.23.1 - 2024-12-24

## 0.23.0 - 2024-12-24

### ✨ New Features

- Added `util/url`

## 0.22.0 - 2024-12-14

### ✨ New Features

- Added utility functions for the tcp connection.

## 0.21.2 - 2024-12-09

🚑️ Fix Transaction Handling

## 0.21.1 - 2024-09-12

### ⬆️ Update dependencies

- `gorm.io/driver/postgres`
- `github.com/newrelic/go-agent/v3`
- `gorm.io/gorm`

## 0.21.0 - 2024-06-19

### 🚀️ New Features

- Added `db.GetAndStartWithDBSource`, the function specified the database source to start transaction.

## 0.20.0 - 2024-06-10

### 🚀️ New Features

- Added `gorm.Initialize`, the function initialize the GORM connection.

## 0.19.0 - 2024-06-10

### 🚀️ New Features

- `gorm.Statement` now runs without transactions by default.

### ⬆️ Update dependencies

- `github.com/cockroachdb/errors`

## 0.18.1 - 2024-05-04

⬆️ Update dependencies

- `github.com/DATA-DOG/go-sqlmock`
- `github.com/newrelic/go-agent/v3`
- `gorm.io/gorm`
- `gorm.io/driver/postgres`
- `github.com/miyamo2/altnrslog`

## 0.18.0 - 2024-05-04

💥 Breaking Changes

- Remove implementation of the `DBTransaction` for DynamoDB

## 0.17.1 - 2024-05-04

🚑️ Fix Broken Checksum

## 0.17.0 - 2024-04-25

- Remove instrumentation from the DB Transaction process