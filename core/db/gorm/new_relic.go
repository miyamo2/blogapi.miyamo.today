package gorm

import (
	"github.com/newrelic/go-agent/v3/newrelic"
	"gorm.io/gorm"
)

// TraceableScan is a wrapper around gorm.DB.Scan that instruments with New Relic.
func TraceableScan(tx *newrelic.Transaction, db *gorm.DB, dest interface{}) *gorm.DB {
	defer tx.StartSegment("gorm: Scan").End()
	return db.Scan(dest)
}
