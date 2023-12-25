package gorm

import (
	"github.com/cockroachdb/errors"
	"github.com/miyamo2/blogapi-core/infra/gorm/internal"
	"github.com/miyamo2/blogapi-core/log"
	"gorm.io/gorm"
)

var ErrDialectorNotInitialized = errors.New("gorm dialector is not initialized")

func Get() (*gorm.DB, error) {
	internal.Lock.RLock()
	defer internal.Lock.RUnlock()
	if internal.Dialector == nil {
		return nil, ErrDialectorNotInitialized
	}
	db, err := gorm.Open(*internal.Dialector)
	if err != nil {
		return nil, err
	}
	return db, nil
}

// InitializeDialector initializes gorm database dialector.
func InitializeDialector(dialector *gorm.Dialector) {
	log.DefaultLogger().Info("initialize gorm database dialector")
	internal.Lock.Lock()
	defer internal.Lock.Unlock()
	if internal.Dialector != nil {
		log.DefaultLogger().Warn("gorm database dialector is already initialized")
		return
	}
	internal.Dialector = dialector
	log.DefaultLogger().Info("completed gorm database dialector initialization")
}

// InvalidateDialector invalidates gorm database dialector.
func InvalidateDialector() {
	log.DefaultLogger().Info("invalidate gorm database dialector")
	internal.Lock.Lock()
	defer internal.Lock.Unlock()
	if internal.Dialector == nil {
		log.DefaultLogger().Warn("gorm database dialector is not initialized")
		return
	}
	internal.Dialector = nil
	log.DefaultLogger().Info("completed gorm database dialector invalidation")
}
