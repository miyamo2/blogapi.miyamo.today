package conn

import (
	"sync"

	"gorm.io/gorm"
)

var (
	Mu       sync.RWMutex
	Instance *gorm.DB
)
