package gorm

import (
	"github.com/cockroachdb/errors"
	"github.com/miyamo2/blogapi-core/log"
	"gorm.io/gorm"
	"sync"
)

var singletonDialector struct {
	mu        sync.RWMutex
	dialector *gorm.Dialector
}

var singletonDB struct {
	mu sync.RWMutex
	db *gorm.DB
}

var ErrDialectorNotInitialized = errors.New("gorm dialector is not initialized")

func Get() (*gorm.DB, error) {
	log.DefaultLogger().Info("get gorm connection")
	singletonDB.mu.Lock()
	defer singletonDB.mu.Unlock()
	db := singletonDB.db
	if db != nil {
		return db, nil
	}
	log.DefaultLogger().Info("gorm connection is not initialized")
	singletonDialector.mu.RLock()
	defer singletonDialector.mu.RUnlock()
	if singletonDialector.dialector == nil {
		return nil, ErrDialectorNotInitialized
	}
	dialector := *singletonDialector.dialector

	db, err := gorm.Open(dialector)
	if err != nil {
		log.DefaultLogger().Warn("failed to initialize gorm connection")
		return nil, err
	}
	singletonDB.db = db
	log.DefaultLogger().Info("completed gorm connection initialization")
	return db, nil
}

// InitializeDialector initializes gorm database dialector.
func InitializeDialector(dialector *gorm.Dialector) {
	log.DefaultLogger().Info("initialize gorm database dialector")
	singletonDialector.mu.Lock()
	defer singletonDialector.mu.Unlock()
	if singletonDialector.dialector != nil {
		log.DefaultLogger().Warn("gorm database dialector is already initialized")
		return
	}
	singletonDialector.dialector = dialector
	log.DefaultLogger().Info("completed gorm database dialector initialization")
}

// InvalidateDialector invalidates gorm database dialector.
func InvalidateDialector() {
	log.DefaultLogger().Info("invalidate gorm database dialector")
	singletonDialector.mu.Lock()
	defer singletonDialector.mu.Unlock()
	if singletonDialector.dialector == nil {
		log.DefaultLogger().Warn("gorm database dialector is not initialized")
		return
	}
	singletonDialector.dialector = nil
	log.DefaultLogger().Info("completed gorm database dialector invalidation")
}

// Invalidate invalidates gorm connection.
func Invalidate() {
	log.DefaultLogger().Info("invalidate gorm connection")
	singletonDB.mu.Lock()
	defer singletonDB.mu.Unlock()
	if singletonDB.db == nil {
		log.DefaultLogger().Warn("gorm connection is not initialized")
		return
	}
	singletonDB.db = nil
	log.DefaultLogger().Info("completed gorm connection invalidation")
}
