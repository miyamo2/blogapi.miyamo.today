package gorm

import (
	"context"

	"blogapi.miyamo.today/core/db/gorm/internal/conn"
	"blogapi.miyamo.today/core/db/gorm/internal/dial"
	"blogapi.miyamo.today/core/log"
	"github.com/cockroachdb/errors"
	"github.com/miyamo2/altnrslog"
	"gorm.io/gorm"
)

var ErrDialectorNotInitialized = errors.New("gorm dialector is not initialized")

func Get(ctx context.Context) (*gorm.DB, error) {
	logger, err := altnrslog.FromContext(ctx)
	if err != nil {
		logger = log.DefaultLogger()
	}
	logger.Info("get gorm connection")
	conn.Mu.Lock()
	defer conn.Mu.Unlock()
	db := conn.Instance

	sc := gorm.Session{
		Context:                ctx,
		PrepareStmt:            true,
		SkipDefaultTransaction: true,
	}

	if db == nil {
		logger.Info("gorm connection is not initialized")
		dial.Mu.RLock()
		defer dial.Mu.RUnlock()
		if dial.Instance == nil {
			return nil, ErrDialectorNotInitialized
		}
		dialector := *dial.Instance

		db, err = gorm.Open(dialector, &gorm.Config{
			SkipDefaultTransaction: true,
			PrepareStmt:            true,
		})
		if err != nil {
			logger.Warn("failed to initialize gorm connection")
			return nil, err
		}
		conn.Instance = db
		logger.Info("completed gorm connection initialization")
	}
	return db.Session(&sc), nil
}

// Initialize initializes gorm database connection.
func Initialize(db *gorm.DB) {
	log.DefaultLogger().Info("initialize gorm database connection")
	conn.Mu.Lock()
	defer conn.Mu.Unlock()
	if conn.Instance != nil {
		log.DefaultLogger().Warn("gorm database connection is already initialized")
		return
	}
	conn.Instance = db
	log.DefaultLogger().Info("completed gorm database connection initialization")
}

// InitializeDialector initializes gorm database dialector.
func InitializeDialector(dialector *gorm.Dialector) {
	log.DefaultLogger().Info("initialize gorm database dialector")
	dial.Mu.Lock()
	defer dial.Mu.Unlock()
	if dial.Instance != nil {
		log.DefaultLogger().Warn("gorm database dialector is already initialized")
		return
	}
	dial.Instance = dialector
	log.DefaultLogger().Info("completed gorm database dialector initialization")
}

// InvalidateDialector invalidates gorm database dialector.
func InvalidateDialector() {
	log.DefaultLogger().Info("invalidate gorm database dialector")
	dial.Mu.Lock()
	defer dial.Mu.Unlock()
	if dial.Instance == nil {
		log.DefaultLogger().Warn("gorm database dialector is not initialized")
		return
	}
	dial.Instance = nil
	log.DefaultLogger().Info("completed gorm database dialector invalidation")
}

// Invalidate invalidates gorm connection.
func Invalidate() {
	log.DefaultLogger().Info("invalidate gorm connection")
	conn.Mu.Lock()
	defer conn.Mu.Unlock()
	if conn.Instance == nil {
		log.DefaultLogger().Warn("gorm connection is not initialized")
		return
	}
	conn.Instance = nil
	log.DefaultLogger().Info("completed gorm connection invalidation")
}
