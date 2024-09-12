# Changelog

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