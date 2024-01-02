package internal

import (
	"gorm.io/gorm"
	"sync"
)

var (
	DialectorLock sync.RWMutex
	Dialector     *gorm.Dialector
	DBLock        sync.RWMutex
	DB            *gorm.DB
)
