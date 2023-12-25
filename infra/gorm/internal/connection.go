package internal

import (
	"gorm.io/gorm"
	"sync"
)

var (
	Lock      sync.RWMutex
	Dialector *gorm.Dialector
)
